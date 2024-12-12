const path = require('path');
const fs = require('fs-extra');
const process = require('process');

const getFiles = (dirPath, targetExtension) => {
    let filesArray = [];
    const files = fs.readdirSync(dirPath);
    files.forEach((file) => {
        const filePath = path.join(dirPath, file);
        const stat = fs.statSync(filePath);

        if (stat.isDirectory()) {
            filesArray = filesArray.concat(getFiles(filePath, targetExtension));
        } else {
            const fileExtension = path.extname(filePath);
            const fileName = path.basename(filePath);
            if (fileExtension == targetExtension && !filePath.includes('.dbg.') && fileName != 'Types.sol') {
                filesArray.push(filePath);
            }
        }
    });

    return filesArray;
};

const addressMap = {
    BankAddress: "0x0000000000000000000000000000000000001000",
    AuthAddress: "0x0000000000000000000000000000000000001001",
    GovAddress: "0x0000000000000000000000000000000000001002",
    StakingAddress: "0x0000000000000000000000000000000000001003",
    DistributionAddress: "0x0000000000000000000000000000000000001004",
    SlashingAddress: "0x0000000000000000000000000000000000001005",
    EvidenceAddress: "0x0000000000000000000000000000000000001006",
    EpochsAddress: "0x0000000000000000000000000000000000001007",
    AuthzAddress: "0x0000000000000000000000000000000000001008",
    FeemarketAddress: "0x0000000000000000000000000000000000001009",
    PaymentAddress: "0x000000000000000000000000000000000000100a",
    PermissionAddress: "0x000000000000000000000000000000000000100b",
    VirtualGroupAddress: "0x0000000000000000000000000000000000002000",
    StorageAddress: "0x0000000000000000000000000000000000002001",
    StorageProviderAddress: "0x0000000000000000000000000000000000002002"
};

const exportPrecompile = async () => {
    const dir = process.cwd();
    console.log(dir);
    const solidityPath = path.join(dir, 'contracts');
    const artifactPath = path.join(dir, 'artifacts/contracts');

    const solidityFiles = getFiles(solidityPath, '.sol');
    const artifactFiles = getFiles(artifactPath, '.json');

    console.log(solidityFiles, artifactFiles);

    const typesPath = path.join(dir, 'contracts/common/Types.sol');
    const typesContent = await fs
        .readFileSync(typesPath)
        .toString()
        .replace('// SPDX-License-Identifier: Apache-2.0', '')
        .replace('pragma solidity ^0.8.0;', '')
        .replace(/(\r\n|\n|\r){2,}/g, '\n');

    // console.log(typesContent);

    let precompiles = [];
    for (const filePath of solidityFiles) {
        const extname = path.extname(filePath);
        const name = path.basename(filePath).replace(extname, '').substring(1);
        const abiPath = artifactFiles.find((item) => item.includes(name));
        const abi = (await fs.readJSON(abiPath)).abi;

        const source = await fs.readFileSync(filePath).toString().replace('import "../common/Types.sol";', typesContent);

        precompiles.push({
            address: addressMap[name + 'Address'],
            name,
            bytecode: '0x00',
            compiler: 'v0.8.20+commit.a1b79de6',
            source,
            abi: JSON.stringify(abi, '', 2),
        });
    }

    await fs.writeFile('./precompiles.json', JSON.stringify(precompiles, '', 2));
};

exportPrecompile();
