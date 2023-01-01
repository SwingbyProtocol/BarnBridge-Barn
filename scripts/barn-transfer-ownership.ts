import { ethers } from 'hardhat';

const _gov = '0x63d9f6A25ddD2c586F4441065Ce7C8412fbBB91e'; // todo: change address
const _barn = '0xb4200c8c44b05a342A9f7FD0d27647C4bf9533e7'; // todo: change address

async function main () {
    const barnOwnership=  await ethers.getContractAt('OwnershipFacet', _barn);
    await barnOwnership.transferOwnership(_gov);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
