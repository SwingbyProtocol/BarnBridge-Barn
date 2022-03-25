import * as deploy from '../test/helpers/deploy';
import { ethers } from 'hardhat';

const _owner = '0x64496f51779e400C5E955228E56fA41563Fb4dd8';
const _sbBTC = '0x1B5B6dF2C72D7c406df1C30E640df8dBaE57d21d';
const _barn = '0x19cFBFd65021af353aB8A7126Caf51920163f0D2';

async function main () {
    const [owner] = await ethers.getSigners();
    console.log(owner.address)
    const _sbBTCPool = await deploy.deployContract('sbBTCPool');
    console.log(`sbBTCPool deployed at: ${_sbBTCPool.address}`);

    //await nodeRewards.setupPullToken(_cv, startTs, endTs);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });

