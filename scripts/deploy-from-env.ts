import * as deploy from '../test/helpers/deploy';
import { diamondAsFacet } from '../test/helpers/diamond';
import { BarnFacet, Rewards } from '../typechain';
import { BigNumber } from 'ethers';
import * as helpers from '../test/helpers/helpers';
require('dotenv').config();
import { ethers } from 'hardhat';


const _bond = process.env.BOND;
const _apr = 15;
// needed for rewards setup
const _cv = process.env.CV;
//const startTs = process.env.STARTTS;
// ENDTS must be 1 year later from STARTTS (31536000 sec)
//const endTs = process.env.ENDTS
// disabled because amount is changed dynamically.
// const rewardsAmount = BigNumber.from(process.env.REWARDSAMOUNT).mul(helpers.tenPow18);
//console.log(startTs, endTs)
async function main() {
    const [owner] = await ethers.getSigners();
    console.log(`signer = ${owner.address}`)
    const cutFacet = await deploy.deployContract('DiamondCutFacet');
    console.log(`DiamondCutFacet deployed to: ${cutFacet.address}`);

    const loupeFacet = await deploy.deployContract('DiamondLoupeFacet');
    console.log(`DiamondLoupeFacet deployed to: ${loupeFacet.address}`);

    const ownershipFacet = await deploy.deployContract('OwnershipFacet');
    console.log(`OwnershipFacet deployed to: ${ownershipFacet.address}`);

    const crf = await deploy.deployContract('ChangeRewardsFacet');
    console.log(`ChangeRewardsFacet deployed to: ${crf.address}`);

    const barnFacet = await deploy.deployContract('BarnFacet');
    console.log(`BarnFacet deployed at: ${barnFacet.address}`);

    const diamond = await deploy.deployDiamond(
        'Barn',
        [cutFacet, loupeFacet, ownershipFacet, crf, barnFacet],
        owner.address,
    );
    console.log(`Barn deployed at: ${diamond.address}`);

    const rewards = (await deploy.deployContract('Rewards', [owner.address, _bond, _apr, _cv])) as Rewards;
    console.log(`Rewards deployed at: ${rewards.address}`);

    console.log('Calling initBarn');
    const barn = (await diamondAsFacet(diamond, 'BarnFacet')) as BarnFacet;
    await barn.initBarn(_bond, rewards.address);

    //await rewards.setupPullToken(_cv, startTs, endTs, rewardsAmount);
    //await rewards.setupPullToken(_cv, startTs, endTs);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
