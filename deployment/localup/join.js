import { exec } from 'child_process';
import { fileURLToPath } from 'node:url';
import { dirname } from 'node:path';
import fs from 'fs-extra';
import path from 'path';
import util from 'util';

const execPromis = util.promisify(exec);

const validatorCount = 4 + 1; // 4 validator + 1 common node(for test become validator)
const workPath = path.join(dirname(fileURLToPath(import.meta.url)), '../../');

export const sleep = (time) => {
  return new Promise((resolve) => setTimeout(resolve, time));
};

const main = async function () {
  console.log({ workPath });
  try {
    console.log('stop all mechain');
    const stopCmd = `bash ./deployment/localup/localup.sh stop ${validatorCount}`;
    await execPromis(stopCmd, { cwd: workPath });
    await sleep(3000);

    console.log('init mechain');
    const initCmd = `bash ./deployment/localup/localup.sh init ${validatorCount} 1`;
    await execPromis(initCmd, { cwd: workPath });
    await sleep(3000);

    console.log('generate mechain');
    const generateCmd = `bash ./deployment/localup/localup.sh generate ${validatorCount} 1`;
    await execPromis(generateCmd, { cwd: workPath });
    await sleep(3000);

    console.log('remove latest validator from genesis file');
    let message;
    for (let i = 0; i < validatorCount; i++) {
      const genesisPath = path.join(workPath, `deployment/localup/.local/validator${i}/config/genesis.json`);
      const genesis = await fs.readJSON(genesisPath);
      const tx = genesis.app_state.genutil.gen_txs.pop();
      message = tx.body.messages[0];
      message.from = '0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2';
      genesis.app_state.gov.params.quorum = '0.010000000000000000';
      genesis.app_state.gov.params.voting_period = '6s';
      await fs.outputJson(genesisPath, genesis, { spaces: 2 });
    }

    console.log('start all mechain node');
    const startCmd = `bash ./deployment/localup/localup.sh start ${validatorCount}`;
    execPromis(startCmd, { cwd: workPath }); // We got stuck here, so we continued without waiting for it to succeed.
    await sleep(6000);

    console.log('authz grant cosmos.staking.v1beta1.MsgDelegate');
    const grantCmd = `./build/mechaind tx authz grant 0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2 generic --msg-type=/cosmos.staking.v1beta1.MsgDelegate --gas="600000" --gas-prices="10000000000azkme"  --from=validator_delegator${
      validatorCount - 1
    } --home=./deployment/localup/.local/validator${validatorCount - 1} --keyring-backend=test --broadcast-mode sync -y`;
    execPromis(grantCmd, { cwd: workPath });
    await sleep(3000);

    console.log('generate proposal.json');
    const proposal = {
      title: `create validator${validatorCount - 1}`,
      summary: `use proposal create validator${validatorCount - 1}`,
      messages: [message],
      metadata: 'ipfs://CID',
      deposit: '1000000000000000000azkme',
    };
    const proposalPath = path.join(workPath, 'build/proposal.json');
    await fs.outputJson(proposalPath, proposal, { spaces: 2 });

    console.log('submint proposal');
    const proposalCmd = `./build/mechaind tx gov submit-proposal ./build/proposal.json --gas="600000" --gas-prices="10000000000azkme" --from=validator0 --home=./deployment/localup/.local/validator0 --keyring-backend=test --broadcast-mode sync -y`;
    execPromis(proposalCmd, { cwd: workPath });
    await sleep(3000);

    console.log('vote proposal');
    const voteCmd = `./build/mechaind tx gov vote 1 yes --gas="600000" --gas-prices="10000000000azkme" --from=validator0 --home=./deployment/localup/.local/validator0 --keyring-backend=test --broadcast-mode sync -y`;
    execPromis(voteCmd, { cwd: workPath });
    await sleep(9000);

    let rsp;

    rsp = await fetch('http://127.0.0.1:1317/cosmos/gov/v1/proposals/1');
    console.log('proposal rsp', await rsp.json());

    await sleep(6000);
    rsp = await fetch('http://127.0.0.1:26657/block');
    const signatures = (await rsp.json()).result.block.last_commit.signatures;
    console.log('block rsp');
    console.log('signatures length', signatures.length, signatures);
    if (signatures.filter((signature) => signature.block_id_flag == 2).length == validatorCount) {
      console.log(`validator${validatorCount - 1} join success`);
    } else {
      console.log(`validator${validatorCount - 1} join fail`);
    }

    console.log('end........');

    process.exit(0);
  } catch (error) {
    console.log('error', error);
  }
};

main();
