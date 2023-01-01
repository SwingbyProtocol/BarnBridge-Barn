import * as deploy from '../test/helpers/deploy';
import { diamondAsFacet } from '../test/helpers/diamond';
import { ChangeRewardsFacet, Rewards } from '../typechain';
import { BigNumber } from 'ethers';
import * as helpers from '../test/helpers/helpers';
require('dotenv').config();
import { ethers } from 'hardhat';


const _burn = "0xb4200c8c44b05a342A9f7FD0d27647C4bf9533e7"
const _rewards = "0x4A25F3815E159582E1E2E7805b78Db8e4cB12768" // new rewards contract
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
