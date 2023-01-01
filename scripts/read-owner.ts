import * as deploy from '../test/helpers/deploy';
import { diamondAsFacet } from '../test/helpers/diamond';
import { BarnFacet, OwnershipFacet, Rewards } from '../typechain';
import { BigNumber } from 'ethers';
import * as helpers from '../test/helpers/helpers';
require('dotenv').config();
import { ethers } from 'hardhat';

// tesnet
// const _burn = "0x009cc14ce70b2E667984C2276490d56ae3234c43"
// mainnet
const _burn = "0xb4200c8c44b05a342A9f7FD0d27647C4bf9533e7"
// disabled because amount is changed dynamically.
// const rewardsAmount = BigNumber.from(process.env.REWARDSAMOUNT).mul(helpers.tenPow18);
async function main() {
    const [owner] = await ethers.getSigners();
    console.log(`signer = ${owner.address}`)

    const barn = await ethers.getContractAt('OwnershipFacet', _burn) as OwnershipFacet

    const ow = await barn.owner()
    console.log(`contract owner = ${ow}`)
    //await rewards.setupPullToken(_cv, startTs, endTs, rewardsAmount);
    //await rewards.setupPullToken(_cv, startTs, endTs);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
