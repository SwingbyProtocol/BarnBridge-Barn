import * as deploy from '../test/helpers/deploy';
import { ethers } from 'hardhat';

const _bond = '0x8287C7b963b405b7B8D467DB9d79eEC40625b13A';
const _apr = 5;
const _cv = '0x7Afa809786Cd034525336aFDD45c7f759Dcaa824';

async function main() {
    const [owner] = await ethers.getSigners();
    console.log(`signer = ${owner.address}`)
    const rewards = await deploy.deployContract('Rewards', [owner.address, _bond, _apr, _cv]);
    console.log(`New Rewards deployed at: ${rewards.address}`);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });

