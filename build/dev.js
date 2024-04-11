import { HDNodeWallet, Wallet } from "ethers";
import crypto from "crypto";
import { ethToBech32 } from "@quarix/address-converter";
import { exec } from "child_process";
import fs from "fs-extra";
import path from "path";
import util from "util";
import yargs from "yargs";
import { hideBin } from "yargs/helpers";

export const sleep = (time) => {
  return new Promise((resolve) => setTimeout(resolve, time));
};

export const privKeyToBurrowAddres = (privKey, isBase64 = true) => {
  if (isBase64) {
    privKey = Buffer.from(privKey, "base64").toString("hex");
  }
  const publicKey = privKey.substring(64, 128);
  const digest = crypto.createHash("sha256").update(Buffer.from(publicKey, "hex")).digest("hex");
  return digest.toLowerCase().substring(0, 40);
};

let argv = yargs(hideBin(process.argv))
  .option("n", {
    alias: "nohup",
    demandOption: false,
    default: true,
    describe: "Whether the startup script is nohup",
    type: "bool",
  })
  .option("init", {
    alias: "init",
    demandOption: false,
    default: true,
    describe: "Init the chain data",
    type: "bool",
  })
  .option("c", {
    alias: "compile",
    demandOption: false,
    default: false,
    describe: "Whether compile code",
    type: "bool",
  })
  .option("k", {
    alias: "keep",
    demandOption: false,
    default: false,
    describe: "Whether keep the data",
    type: "bool",
  })
  .option("f", {
    alias: "fixed",
    demandOption: false,
    default: true,
    describe: "Whether fixed the first validator",
    type: "bool",
  })
  .option("v", {
    alias: "validators",
    demandOption: false,
    default: 1,
    describe: "Number of validators to initialize the testnet with (default 1)",
    type: "number",
  })
  .option("cn", {
    alias: "commonNode",
    demandOption: false,
    default: 0,
    describe: "Number of common node to initialize the testnet with (default 0)",
    type: "number",
  })
  .option("s", {
    alias: "start",
    demandOption: false,
    default: "all",
    describe: "Whether after initialize immediate start",
    type: "string",
  })
  .option("stop", {
    alias: "stop",
    demandOption: false,
    default: "",
    describe: "stop nodes",
    type: "string",
  })
  .number(["v"])
  .number(["cn"])
  .boolean(["n", "c", "k", "init", "f"]).argv;

const isNohup = argv.nohup;
const start = argv.start;
const stop = argv.stop;
const isCompile = argv.compile;
const isKeep = argv.keep;
const fixedFirstValidator = argv.fixed;
const init = argv.init;
const commonNode = argv.commonNode;
const validators = argv.validators;
const nodesCount = validators + commonNode;
const platform = process.platform;
const execPromis = util.promisify(exec);
const curDir = process.cwd();
const nodesDir = path.join(curDir, "nodes");
let chainId = "chain_88888888-1";
let clientCfg = `
# The network chain ID
chain-id = "${chainId}"
# The keyring's backend, where the keys are stored (os|file|kwallet|pass|test|memory)
keyring-backend = "test"
# CLI output format (text|json)
output = "text"
# <host>:<port> to Tendermint RPC interface for this chain
node = "tcp://localhost:26657"
# Transaction broadcasting mode (sync|async)
broadcast-mode = "sync"
`;
const scriptStop = path.join(nodesDir, platform == "win32" ? "stopAll.vbs" : "stopAll.sh");
const scriptStart = path.join(nodesDir, platform == "win32" ? "startAll.vbs" : "startAll.sh");

const updatePorts = (data, ports, index) => {
  let lines = data.split(/\r?\n/);
  for (const key in ports) {
    let [k1, k2] = key.split("."); // key for example "api.address"
    let port = ports[key];
    let find = false;
    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];
      //  for example: [json-rpc]
      if (line.startsWith(`[${k1}]`)) {
        find = true;
      }
      //for example: "tcp://0.0.0.0:1317"
      if (find && line.startsWith(`${k2} = `)) {
        const oldPort = line.split(":").pop().split(`"`)[0];
        const newPort = String(port + index);
        // console.log(line, oldPort, newPort);
        lines[i] = line.replace(oldPort, newPort).replace("localhost", "0.0.0.0").replace("127.0.0.1", "0.0.0.0");
        break;
      }
    }
  }
  return lines.join("\n");
};

const updateCfg = (data, cfg) => {
  let lines = data.split(/\r?\n/);
  for (const key in cfg) {
    let find = true;
    let k1;
    let k2 = key;
    if (key.indexOf(".") > 0) {
      [k1, k2] = key.split(".");
      find = false;
    }
    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];
      if (!find && line.startsWith(`[${k1}]`)) {
        find = true;
      }
      if (find && line.startsWith(`${k2} = `)) {
        lines[i] = `${k2} = ${cfg[key]}`;
        break;
      }
    }
  }
  return lines.join("\n");
};

const main = async function () {
  console.log(`validators:${argv.validators}, commonNode:${argv.commonNode}, init:${argv.init}, fixed:${argv.fixed}, compile:${argv.compile}, start:${argv.start}, stop:${argv.stop}, keep:${argv.keep}, nohup:${argv.nohup}\n`);

  try {
    if (!fs.existsSync("./config.json")) {
      await fs.copyFile("./config.default.json", "./config.json");
    }

    let config = await fs.readJson("./config.json");
    const { app, tendermint, preMinePerAccount, preMineAccounts, privateKeys } = config;
    if (app.chain_id) {
      clientCfg = clientCfg.replaceAll(chainId, app.chain_id);
      chainId = app.chain_id;
    }

    const chainName = app.chain_name;
    const daemon = `${app.chain_name}d`;
    const daemonApp = platform == "win32" ? `${app.chain_name}d.exe` : `${app.chain_name}d`;

    if (stop) {
      if (stop === "all") {
        console.log(`Stop all ${chainName} node under the folder nodes`);
        await execPromis(scriptStop, { cwd: nodesDir });
      } else {
        const argv = stop.split(",");
        for (let i = 0; i < argv.length; i++) {
          const script = path.join(nodesDir, process.platform == "win32" ? `stop${argv[i]}.bat` : `stop${argv[i]}.sh`);
          const { stdout, stderr } = await execPromis(script, { cwd: nodesDir });
          console.log(`${stdout}${stderr}`);
        }
      }
      return;
    }

    if (await fs.pathExists(scriptStop)) {
      console.log(`Try to stop the ${chainName} under the nodes directory`);
      await execPromis(scriptStop, { cwd: nodesDir }); // Anyway, stop it first
      await sleep(platform == "win32" ? 600 : 300);
    }

    if (init || isKeep) {
      // These three sets of data(nodeKey, privValidatorConsensusKey, createValidator) need to be replaced.
      // for go update here testutil.GenerateSaveCoinKey
      const nodeKey = { priv_key: { type: "tendermint/PrivKeyEd25519", value: "A+L4//cF93rKhDoKF1oTNmCUACTzjBE9ni/3EWn2qrNdgPO/h325EwWs4Xkxvt6hxF4UmXO2kChXhulLIYw2lA==" } };
      const privValidatorConsensusKey = { address: "A66A07BDB89B792D81ACF20E52E4A37F3BAF2D61", pub_key: { type: "tendermint/PubKeyEd25519", value: "JDMdH9KXBxbWNIuqMxsDjL9rwQiM5rc48PXwzce4sUY=" }, priv_key: { type: "tendermint/PrivKeyEd25519", value: "+hrWS+TbGqX86JfKhL9+AG60hPtRctQW0hHft42UFvQkMx0f0pcHFtY0i6ozGwOMv2vBCIzmtzjw9fDNx7ixRg==" } };
      const createValidator = {
        body: {
          messages: [
            {
              "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
              description: { moniker: "node0", identity: "", website: "", security_contact: "", details: "" },
              commission: { rate: "0.100000000000000000", max_rate: "1.000000000000000000", max_change_rate: "1.000000000000000000" },
              min_self_delegation: "1",
              delegator_address: "mc1hajh6rhhkjqkwet6wqld3lgx8ur4y3khscnzem",
              validator_address: "mcvaloper1hajh6rhhkjqkwet6wqld3lgx8ur4y3khf8q9ma",
              pubkey: { "@type": "/cosmos.crypto.ed25519.PubKey", key: "JDMdH9KXBxbWNIuqMxsDjL9rwQiM5rc48PXwzce4sUY=" },
              value: { denom: "azkme", amount: "100000000000000000000" },
            },
          ],
          memo: "e2795df5deec6ded2bbfb9636485d8ac27eef6b8@192.168.0.1:26656",
          timeout_height: "0",
          extension_options: [],
          non_critical_extension_options: [],
        },
        auth_info: { signer_infos: [{ public_key: { "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey", key: "A50rbJg3TMPACbzE5Ujg0clx+d4udBAtggqEQiB7v9Sc" }, mode_info: { single: { mode: "SIGN_MODE_DIRECT" } }, sequence: "0" }], fee: { amount: [], gas_limit: "0", payer: "", granter: "" }, tip: null },
        signatures: ["m/CGLpFEjUACsw38f7zgVYFjowcglGfk/6dJ117lZd9Hbyu9lPTiHWEJZV4bJb9m/Pqs1jJBo64+Cty+RhXIdwE="],
      };

      const validatorSecret = "october pride genuine harvest reunion sight become tuna kingdom punch girl lizard cat crater fee emotion seat test output safe volume caught design soft";
      const validatorWallet = HDNodeWallet.fromPhrase(validatorSecret);
      const keySeed = {
        secret: validatorSecret,
        privateKey: validatorWallet.privateKey.replace("0x", ""),
        publicKey: validatorWallet.publicKey.replace("0x", ""),
        address: validatorWallet.address,
        bip39Address: ethToBech32(validatorWallet.address, app.prefix),
      };

      if (!fs.existsSync(daemonApp) || isCompile) {
        console.log(`Start recompiling ${daemonApp}...`);
        let make = await execPromis(`go build -o ${daemonApp} ../cmd/${daemon}`, { cwd: curDir });
        console.log(`${daemonApp} compile finished`, make);
      }

      if (!fs.existsSync(daemonApp)) {
        console.log(`${daemonApp} Executable file does not exist`);
        return;
      }

      if (validators < 1) {
        console.log("validators >= 1");
        return;
      }
      if (!isKeep) {
        console.log("Start cleaning up folder nodes");
        await fs.emptyDir(nodesDir);
        await fs.ensureDir(nodesDir);
        console.log("Folder nodes has been cleaned up");

        {
          const initFiles = `${platform !== "win32" ? "./" : ""}${daemonApp} testnet init-files --v ${nodesCount} --output-dir ./nodes --chain-id ${chainId} --keyring-backend test`;
          const initFilesValidator = `${platform !== "win32" ? "./" : ""}${daemonApp} testnet init-files --v ${validators} --output-dir ./nodes --chain-id ${chainId} --keyring-backend test`;
          console.log(`Exec cmd: ${initFiles}`);
          const { stdout, stderr } = await execPromis(initFiles, { cwd: curDir });
          console.log(`${stdout}${stderr}\n`);

          if (commonNode > 0) {
            for (let i = 0; i < validators; i++) {
              await fs.remove(path.join(nodesDir, `node${i}`));
            }
            await fs.remove(path.join(nodesDir, `gentxs`));

            // re init validator, and turn a validator node into a common node
            await execPromis(initFilesValidator, { cwd: curDir });
            const genesisPath = path.join(nodesDir, `node0/${daemon}/config/genesis.json`);
            for (let i = validators; i < nodesCount; i++) {
              await fs.copy(genesisPath, path.join(nodesDir, `node${i}/${daemon}/config/genesis.json`));
            }
          }

          if (fixedFirstValidator) {
            await fs.writeJSON(path.join(nodesDir, `node0/${daemon}/config/node_key.json`), nodeKey);
            await fs.writeJSON(path.join(nodesDir, `node0/${daemon}/config/priv_validator_key.json`), privValidatorConsensusKey);
            await fs.outputJSON(path.join(nodesDir, `node0/${daemon}/key_seed.json`), keySeed);
            const keyringPath = path.join(nodesDir, `node0/${daemon}/keyring-test`);
            await fs.emptyDir(keyringPath);
            await fs.writeFile(
              path.join(keyringPath, `bf657d0ef7b48167657a703ed8fd063f075246d7.address`),
              "eyJhbGciOiJQQkVTMi1IUzI1NitBMTI4S1ciLCJjcmVhdGVkIjoiMjAyMi0wOC0yNCAxODowOTowNC43NjQ4NTEgKzA4MDAgQ1NUIG09KzAuMjI4NTE5MjUxIiwiZW5jIjoiQTI1NkdDTSIsInAyYyI6ODE5MiwicDJzIjoiVHM3QXhNRmV4MlZtMTZpeiJ9.OrWluGLeod9SjmLDqvXTcA63z9P1VZ-D0l5LFzwVOhJG67vl3b0HXQ.BrINO_FqPHviDFff.yk2tJKWkWIo-OXZfxr7INBATtLws_mHvT5s4kSfwDkbpp2JJVyoEwFcozQHp5hh9owc3bPG7HRa_QHQarB5_Oz-fXJkuPlTxR955P6azI1C8vuWqBcZ7nfZkAhoFHgSZzQAPuFp6sPTWoDampAqocmtWu2lYPSiRnDHRZ6gEmP1slwsRwJTlASEwpmzjBeDsqrwCn9cT_jNrI7ilWB4LBUUXAkkKVu-p1X9bkqo8yZ_UrFFR2rI.6rVArcxnth5pzzgbEtuHSQ"
            );
            await fs.writeFile(
              path.join(keyringPath, `node0.info`),
              "eyJhbGciOiJQQkVTMi1IUzI1NitBMTI4S1ciLCJjcmVhdGVkIjoiMjAyMi0wOC0yNCAxODowOTowNC43NTg1NjYgKzA4MDAgQ1NUIG09KzAuMjIyMjM0MDQzIiwiZW5jIjoiQTI1NkdDTSIsInAyYyI6ODE5MiwicDJzIjoicmk3MzV2Y3Fid2VkUF9JcCJ9.ht-BieDMdmkOBfb1saBx2nvBDaD9anNxP5RTirHIk-tHUXJr6HbeKA.FvpzGpaY6il86ngO.WwHd6HTneYvxg3KkEhsXx1_F_XkmzHqVJwSmQrnX9ZSg2L8ZCAxV6rvliuRwt30816o8tElb06qpp1krFGwGL_LvP1FtnOiX4GdJJxAyX1lgBgJQrhZuqKc6EEE78ArwUR1Mb6b3ax_6oV7IB42izg1ci2PP5bgXN-510EM9RrSi9fnVl3UMoAanoBL8NfJGYHo2Cusn_Y14yEnPDHxS96vTl7wZx_pZrjtapyQ9ktnDQHVBfsupIKmIYXSwpQ16FQ9G4eclfKGhit4uUFofdT0UMG1g_aQEGHt1nPG08w66w8PxmW8ma_D8yCQp0TW6m9pTLWODiCztorLucEr9RFW9mJLofi4pFdCuqHrGm_o.X06PXwtrfTMDgiQDIpPS0g"
            );
          }
        }

        await fs.copy(daemonApp, `./nodes/${daemonApp}`);

        let nodeIds = [];
        for (let i = 0; i < nodesCount; i++) {
          const nodeKey = await fs.readJSON(path.join(nodesDir, `node${i}/${daemon}/config/node_key.json`));
          const nodeId = privKeyToBurrowAddres(nodeKey.priv_key.value);
          nodeIds.push(nodeId);

          const keySeedPath = path.join(nodesDir, `node${i}/${daemon}/key_seed.json`);
          let curKeySeed = await fs.readJSON(keySeedPath);
          const wallet = HDNodeWallet.fromPhrase(curKeySeed.secret);
          curKeySeed.privateKey = wallet.privateKey.replace("0x", "");
          curKeySeed.publicKey = wallet.publicKey.replace("0x", "");
          curKeySeed.address = wallet.address;
          curKeySeed.bip39Address = ethToBech32(wallet.address, app.prefix);
          await fs.outputJson(keySeedPath, curKeySeed, { spaces: 2 });
        }

        const account = { "@type": "/ethermint.types.v1.EthAccount", base_account: { address: "", pub_key: null, account_number: "0", sequence: "0" }, code_hash: "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470" };
        const balance = { address: "", coins: [] };
        for (let i = 0; i < nodesCount; i++) {
          let accounts = [];
          let balances = [];
          if (Array.isArray(preMineAccounts)) {
            let duplicate = {};
            for (const ac of preMineAccounts) {
              let address = ac;
              if (ac.length == 64) {
                const wallet = new Wallet(ac);
                address = ethToBech32(wallet.address, app.prefix);
              } else if (address.startsWith("0x")) {
                address = ethToBech32(ac, app.prefix);
              }
              if (duplicate[address]) {
                continue;
              }
              duplicate[address] = true;

              accounts.push(Object.assign(JSON.parse(JSON.stringify(account)), { base_account: { address } }));
              balances.push(Object.assign(JSON.parse(JSON.stringify(balance)), { address }));
            }
          }

          const genesisPath = path.join(nodesDir, `node${i}/${daemon}/config/genesis.json`);
          let genesis = await fs.readJSON(genesisPath);
          let appState = genesis.app_state;
          appState.auth.accounts.push(...accounts);
          appState.bank.balances.push(...balances);
          if (commonNode > 0) {
            for (let i = nodesCount - commonNode; i < nodesCount; i++) {
              const keySeedPath = path.join(nodesDir, `node${i}/${daemon}/key_seed.json`);
              const curKeySeed = await fs.readJSON(keySeedPath);
              const address = curKeySeed.bip39Address;
              appState.auth.accounts.push(Object.assign(JSON.parse(JSON.stringify(account)), { base_account: { address } }));
              appState.bank.balances.push(Object.assign(JSON.parse(JSON.stringify(balance)), { address }));
            }
          }

          for (let balances of appState.bank.balances) {
            balances.coins = app.denoms.sort().map((denom) => ({ denom, amount: "0" }));
            for (let coin of balances.coins) {
              coin.amount = preMinePerAccount;
            }
          }

          if (fixedFirstValidator) {
            appState.auth.accounts[0].base_account.address = keySeed.bip39Address;
            appState.bank.balances[0].address = keySeed.bip39Address;
            appState.genutil.gen_txs[0] = createValidator;
          }

          const genesisCfg = config.genesisCfg;
          if (Array.isArray(genesisCfg)) {
            for (const cfg of genesisCfg) {
              eval("genesis." + cfg);
            }
          }

          // Use zero address to occupy the first account, Because of account_ Accounts with number 0 cannot send Cosmos transactions
          appState.auth.accounts.unshift(Object.assign(JSON.parse(JSON.stringify(account)), { base_account: { address: ethToBech32("0x0000000000000000000000000000000000000000", app.prefix) } }));

          await fs.outputJson(genesisPath, genesis, { spaces: 2 });
        }

        // update app.toml and config.toml
        for (let i = 0; i < nodesCount; i++) {
          let data;
          const appConfigPath = path.join(nodesDir, `node${i}/${daemon}/config/app.toml`);
          data = await fs.readFile(appConfigPath, "utf8");
          data = updatePorts(data, app.port, i);
          data = updateCfg(data, app.cfg);
          await fs.writeFile(appConfigPath, data);

          const configPath = path.join(nodesDir, `node${i}/${daemon}/config/config.toml`);
          data = await fs.readFile(configPath, "utf8");
          data = updatePorts(data, tendermint.port, i);
          // replace persistent_peers
          let peers = [];
          const p2pPort = tendermint.port["p2p.laddr"];
          for (let j = 0; j < nodesCount && nodesCount > 1; j++) {
            if (i != j) {
              peers.push(`${nodeIds[j]}@127.0.0.1:${p2pPort + j}`);
            }
          }
          tendermint.cfg["p2p.persistent_peers"] = `"${peers.join()}"`;
          data = updateCfg(data, tendermint.cfg);
          await fs.writeFile(configPath, data);

          const clientConfigPath = path.join(nodesDir, `node${i}/${daemon}/config/client.toml`);
          data = clientCfg;
          data = data.replace("26657", tendermint.port["rpc.laddr"] + i + "");
          await fs.writeFile(clientConfigPath, data);
        }

        if (Array.isArray(privateKeys)) {
          for (const privateKey of privateKeys) {
            const cmd = `echo -n "your-password" | ./${daemonApp} keys unsafe-import-eth-key ${privateKey.name} ${privateKey.key} --home ./nodes/node0/${daemon} --keyring-backend test`;
            await execPromis(cmd, { cwd: curDir });
          }
        }

        // 生成启动命令脚本
        let vbsStart = platform == "win32" ? `set ws=WScript.CreateObject("WScript.Shell")\n` : `#!/bin/bash\n`;
        let vbsStop = platform == "win32" ? `set ws=WScript.CreateObject("WScript.Shell")\n` : `#!/bin/bash\n`;
        for (let i = 0; i < nodesCount; i++) {
          let p2pPort = tendermint.port["p2p.laddr"] + i;
          let start = (platform == "win32" ? "" : "#!/bin/bash\n") + (isNohup && platform !== "win32" ? "nohup " : "") + (platform !== "win32" ? "./" : "") + `${daemonApp} start --keyring-backend test --home ./node${i}/${daemon}/` + (isNohup && platform !== "win32" ? ` >./${daemon}${i}.log 2>&1 &` : "");
          let stop =
            platform == "win32"
              ? `@echo off
  for /f "tokens=5" %%i in ('netstat -ano ^| findstr 0.0.0.0:${p2pPort}') do set PID=%%i
  taskkill /F /PID %PID%`
              : `pid=\`lsof -iTCP:${p2pPort} -sTCP:LISTEN -t\`;
  if [[ -n $pid ]]; then kill -15 $pid; fi`;
          let startPath = path.join(nodesDir, `start${i}.` + (platform == "win32" ? "bat" : "sh"));
          let stopPath = path.join(nodesDir, `stop${i}.` + (platform == "win32" ? "bat" : "sh"));
          await fs.writeFile(startPath, start);
          await fs.writeFile(stopPath, stop);

          if (platform == "win32") {
            vbsStart += `ws.Run ".\\start${i}.bat",0\n`;
            vbsStop += `ws.Run ".\\stop${i}.bat",0\n`;
          } else {
            vbsStart += `./start${i}.sh\n`;
            vbsStop += `./stop${i}.sh\n`;
            await fs.chmod(startPath, 0o777);
            await fs.chmod(stopPath, 0o777);
          }
        }

        // 生成总的启动脚本
        let startAllPath = path.join(nodesDir, `startAll.` + (platform == "win32" ? "vbs" : "sh"));
        let stopAllPath = path.join(nodesDir, `stopAll.` + (platform == "win32" ? "vbs" : "sh"));
        await fs.writeFile(startAllPath, vbsStart);
        await fs.writeFile(stopAllPath, vbsStop);
        if (!(platform == "win32")) {
          await fs.chmod(startAllPath, 0o777);
          await fs.chmod(stopAllPath, 0o777);
        }
      } else {
        await fs.copy(daemonApp, `./nodes/${daemonApp}`, { overwrite: true });
      }
    }

    if (start && start.toLowerCase() !== "no") {
      if (start === "all") {
        console.log(`Start all ${chainName} node under the folder nodes`);
        await execPromis(scriptStart, { cwd: nodesDir });
      } else {
        let vbsStart = platform == "win32" ? `set ws=WScript.CreateObject("WScript.Shell")\n` : `#!/bin/bash\n`;
        const argv = start.split(",");
        for (let i = 0; i < argv.length; i++) {
          if (platform == "win32") {
            vbsStart += `ws.Run ".\\start${argv[i]}.bat",0\n`;
          } else {
            vbsStart += `./start${argv[i]}.sh\n`;
          }
        }

        const script = path.join(nodesDir, `startTemp.` + (platform == "win32" ? "vbs" : "sh"));
        await fs.writeFile(script, vbsStart);
        if (platform != "win32") {
          await fs.chmod(script, 0o777);
        }
        const { stdout, stderr } = await execPromis(script, { cwd: nodesDir });
        console.log(`${stdout}${stderr}`);
        await fs.remove(script);
      }
    }
  } catch (error) {
    console.log("error", error);
  }
};

main();
