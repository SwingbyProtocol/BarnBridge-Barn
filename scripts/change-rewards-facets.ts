import * as deploy from '../test/helpers/deploy';
import { diamondAsFacet } from '../test/helpers/diamond';
import { ChangeRewardsFacet, Rewards } from '../typechain';
import { BigNumber } from 'ethers';
import * as helpers from '../test/helpers/helpers';
require('dotenv').config();
import { ethers } from 'hardhat';


const _burn = "0x009cc14ce70b2E667984C2276490d56ae3234c43"
const _rewards = "0x313c379eC483678f0ae9F612e56eA983636dd627" // new rewards contract
// disabled because amount is changed dynamically.
// const rewardsAmount = BigNumber.from(process.env.REWARDSAMOUNT).mul(helpers.tenPow18);
async function main() {
    const [owner] = await ethers.getSigners();
    console.log(`signer = ${owner.address}`)

    const barn = await ethers.getContractAt('ChangeRewardsFacet', _burn) as ChangeRewardsFacet
    await barn.changeRewardsAddress(_rewards);
    console.log(`New Rewards attached at: ${_rewards}`);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
