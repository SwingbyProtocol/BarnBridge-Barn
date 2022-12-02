import * as deploy from '../test/helpers/deploy';
import { NodeRewards } from '../typechain';
require('dotenv').config();
import { ethers } from 'hardhat';

const _swingby = process.env.SWINGBY;
const _apr = 15;
// needed for rewards setup
const _source = process.env.CV;
// const startTs = process.env.STARTTS;
// disabled because amount is changed dynamically.
// const rewardsAmount = BigNumber.from(process.env.REWARDSAMOUNT).mul(helpers.tenPow18);
console.log(_swingby, _apr, _source)
async function main() {
    const [owner] = await ethers.getSigners();
    console.log(`signer = ${owner.address}`)
    const nodeRewards = (await deploy.deployContract('NodeRewards', [owner.address, _swingby, _apr, _source])) as NodeRewards;
    console.log(`nodeRewards deployed at: ${nodeRewards.address}`);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
