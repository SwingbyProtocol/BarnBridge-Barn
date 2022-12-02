import * as deploy from '../test/helpers/deploy';
import { diamondAsFacet } from '../test/helpers/diamond';
import { BarnFacet, OwnershipFacet, Rewards } from '../typechain';
import { BigNumber } from 'ethers';
import * as helpers from '../test/helpers/helpers';
require('dotenv').config();
import { ethers } from 'hardhat';


const _barn = "0x009cc14ce70b2E667984C2276490d56ae3234c43"
const _bond = "0xFCd51B56e65605C33024A9E98a7aaDfF2e1A15b9"
const _rewards = "0xE9039D90A2685116cFdeeC8A7Fe5af5ce679DBdC"
// disabled because amount is changed dynamically.
// const rewardsAmount = BigNumber.from(process.env.REWARDSAMOUNT).mul(helpers.tenPow18);
async function main() {
    const [owner] = await ethers.getSigners();
    console.log(`signer = ${owner.address}`)

    const barn = await ethers.getContractAt('BarnFacet', _barn) as BarnFacet
    await barn.initBarn(_bond, _rewards);
    //await rewards.setupPullToken(_cv, startTs, endTs, rewardsAmount);
    //await rewards.setupPullToken(_cv, startTs, endTs);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
