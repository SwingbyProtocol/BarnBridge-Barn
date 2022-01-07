import * as deploy from '../test/helpers/deploy';
import { NodeRewards } from '../typechain';
require('dotenv').config();
import { ethers } from 'hardhat';

const _owner = process.env.OWNER;
const _bond = process.env.BOND;
const _barn = process.env.BARN ? process.env.BARN : "";
const _swap = process.env.SWAP
const _apr = 15;
// needed for rewards setup
const _cv = process.env.CV;
const startTs = process.env.STARTTS;
// ENDTS must be 1 year later from STARTTS (31536000 sec)
const endTs = process.env.ENDTS
// disabled because amount is changed dynamically.
// const rewardsAmount = BigNumber.from(process.env.REWARDSAMOUNT).mul(helpers.tenPow18);
console.log(startTs, endTs)
async function main() {
    const [owner] = await ethers.getSigners();
    console.log(owner.address)

    const barn = await ethers.getContractAt('Barn', _barn);

    const nodeRewards = (await deploy.deployContract('NodeRewards', [_owner, _bond, barn.address, _swap, _apr])) as NodeRewards;
    console.log(`nodeRewards deployed at: ${nodeRewards.address}`);

    await nodeRewards.setupPullToken(_cv, startTs, endTs);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
