import { ethers } from 'hardhat';

const _gov = '0x63d9f6A25ddD2c586F4441065Ce7C8412fbBB91e'; // todo: change address
const _barn = '0x009cc14ce70b2E667984C2276490d56ae3234c43'; // todo: change address

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
