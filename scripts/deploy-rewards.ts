import * as deploy from '../test/helpers/deploy';
import { ethers } from 'hardhat';

const _bond = '0xFCd51B56e65605C33024A9E98a7aaDfF2e1A15b9';
const _apr = 10;
const _cv = '0xb52E4F319D75918B03A3bdd02D20479bF4fa8399';

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

