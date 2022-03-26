# BarnBridge Barn
![](https://i.imgur.com/WhuS8Dv.png)

Implements continuous rewards for staking BOND in the DAO. Implements logic for determining DAO voting power based upon amount of BOND deposited (which becomes vBOND) plus a multiplier (up to 2x) awarded to those who lock their BOND in the DAO for a specified period (up to 1 year). Those that lock their vBOND for 1 year receive a 2x multiplier; those that lock their vBOND for 6 months receive a 1.5x multiplier, and so on. Also allows users to delegate their vBOND voting power to a secondary wallet address.
**NOTE:** The vBOND multiplier ONLY affects voting power; it does NOT affect rewards. All users who stake BOND receive the same reward rate regardless of the amount of time they have locked or not locked.

Any questions? Please contact us on [Discord](https://discord.gg/FfEhsVk) or read our [Developer Guides](https://integrations.barnbridge.com/) for more information.

##  Contracts
### Barn.sol
Allows users to deposit BOND into the DAO, withdraw it, lock for a time period to increase their voting power (does not affect rewards), and delegate their vBOND voting power to a secondary wallet address. Interacts with [Rewards.sol](https://github.com/BarnBridge/BarnBridge-Barn/blob/master/contracts/Rewards.sol) contract to check balances and update upon deposit/withdraw. Interacts with [Governance.sol](https://github.com/BarnBridge/BarnBridge-DAO/blob/master/contracts/Governance.sol) contract to specify how much voting power (vBOND) a wallet address has for use in voting on or creating DAO proposals.
#### Actions
- deposit
- withdraw
- lock
- delegate

### Rewards.sol
Rewards users who stake their BOND on a continuous basis. Allows users to Claim their rewards which are then Transfered to their wallet. Interacts with the [CommunityVault.sol](https://github.com/BarnBridge/BarnBridge-YieldFarming/blob/master/contracts/CommunityVault.sol) which is the source of the BOND rewards. The `Barn` contract calls the `registerUserAction` hook on each `deposit`/`withdraw` the user executes, and sends the results to the `Rewards` contract.
#### How it works
1. every time the `acKFunds` function detects a balance change, the multiplier is recalculated by the following formula:
```
newMultiplier = oldMultiplier + amountAdded / totalBondStaked
```
2. whenever a user action is registered (either by automatic calls from the hook or by user action (claim)), we calculate the amount owed to the user by the following formula:
```
newOwed = currentlyOwed + userBalance * (currentMultiplier - oldUserMultiplier)

where:
- oldUserMultiplier is the multiplier at the time of last user action
- userBalance = barn.balanceOf(user) -- the amount of $BOND staked into the Barn
```
3. update the oldUserMultiplier with the current multiplier -- signaling that we already calculated how much was owed to the user since his last action

## Smart Contract Architecture
Overview

![dao sc architecture](https://gblobscdn.gitbook.com/assets%2F-MIu3rMElIO-jG68zdaV%2F-MXHutr14sDo0hYi6gg3%2F-MXHwLegBZM5HWoEzudF%2Fdao.png?alt=media&token=51e3e2c7-4aab-4601-a3f1-46ae9e1b966f)


Check out more detailed smart contract Slither graphs with all the dependencies: [BarnBridge-Barn Slither Graphs](https://github.com/BarnBridge/sc-graphs/tree/main/BarnBridge-Barn).

### Specs
- user can stake BOND for vBOND
    - user can lock BOND for a period up to 1 year and he gets a bonus of vBOND
        - bonus is linear, max 1 year, max 2x multiplier
            - example:
                - lock 1000 BOND for 1 year → get back 2000 vBOND
                - lock 1000 BOND for 6 months → get back 1500 vBOND
        - bonus has a linear decay relative to locking duration
            - example: lock 1000 BOND for 1 year, get back 2000 vBOND at T0 → after 6 months, balance is 1500 vBOND → after 9 months, balance is 1250 vBOND
        - user can only withdraw their BOND balance after lock expires
    - user can keep BOND unlocked and no bonus is applied, vBOND balance = BOND balance
- user can stake more BOND
    - no lock → just get back the same amount of vBOND
    - lock
        - lock period stays the same
            - base balance is increased with the added BOND
            - multiplier stays the same
        - lock period is extended
            - base balance is increased with the added BOND
                - multiplier is recalculated relative to the new lock expiration date
- user can delegate vBOND to another user
    - there can be only one delegatee at a time
    - only actual balance can be delegated, not the bonus
    - delegated balance cannot be locked
    - user can take back the delegated vBONDs at any time



## Initial Setup
### Install NVM and the latest version of NodeJS 12.x
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
    # Restart terminal and/or run commands given at the end of the installation script
    nvm install 12
    nvm use 12
### Use Git to pull down the BarnBridge-SmartYieldBonds repository from GitHub
    git clone https://github.com/BarnBridge/BarnBridge-Barn.git
    cd BarnBridge-YieldFarming

## Updating the .env file
### Create an API key with a Provider such as Infura to deploy to Ethereum Public Testnet. In this guide, we are using Kovan.
1. Navigate to [Infura.io](https://infura.io/) and create an account
2. Log in and select "Get started and create your first project to access the Ethereum network"
3. Create a project and name it appropriately; switch the Network on the Settings page to Kovan
4. Then, copy the URL and paste it into the section named PROVIDER in the .env file

### Update .env with your wallet information
5. Insert the address of your testing wallet in the .env section labeled "OWNER"
6. Finally, insert the mnemonic phrase for your testing wallet into the .env file section named MNEMONIC. You can use ie a MetaMask instance, and switch the network to Kovan on the upper right. DO NOT USE YOUR PERSONAL METAMASK SEED PHRASE; USE A DIFFERENT BROWSER WITH AN INDEPENDENT METAMASK INSTALLATION
7. You'll need some Kovan-ETH (it is free) in order to pay the gas costs of deploying the contracts on the TestNet; you can use your GitHub account to authenticate to the [KovanFaucet](https://faucet.kovan.network/) and receive 2 Kovan-ETH for free every 24 hours

### Create an API key with Etherscan
8. Navigate to [EtherScan](https://etherscan.io/) and create an account
9. Log in and navigate to [MyAPIKey](https://etherscan.io/myapikey)
10. Use the Add button to create an API key, and paste it into the indicated section towards the section labeled ETHERSCAN in the .env file

### Update .env with the Community Vault contract address
11. Prior to running this deployment, deploy [BarnBridge-YieldFarming](https://github.com/BarnBridge/BarnBridge-YieldFarming) and then copy the contract address of the CommunityVault contract created into the section labeled CV in the .env file


## Installing
### Install NodeJS dependencies which include HardHat
    npm install

### Compile the contracts
    npm run compile

## Running Tests
    npm run test

**Note:** The result of tests is readily available [here](./test-results.md).
## Running Code Coverage Tests
    npm run coverage

## Deploying to Kovan

    npm run deploy-from-env

## Audits
- [QuantStamp](https://github.com/BarnBridge/BarnBridge-PM/blob/master/audits/BarnBridge%20DAO%20audit%20by%20Quanstamp.pdf)
- [Haechi](https://github.com/BarnBridge/BarnBridge-PM/blob/master/audits/BarnBridge%20DAO%20audit%20by%20Haechi.pdf)

## Deployed contracts For SWINGBY
### Mainnet
```
CommunityVault (for pre-staking)
https://etherscan.io/address/0x7afa809786cd034525336afdd45c7f759dcaa824

DiamondCutFacet
https://etherscan.io/address/0x8d5d176ef5a8448236c51372a982e710885b0e7f
DiamondLoupeFacet
https://etherscan.io/address/0x4010749aD0181250DEdb04D362F10188844E3aCa
OwnershipFacet
https://etherscan.io/address/0x06a69Af8008e80a6729636c9Fc5AFba2a25b541C
ChangeRewardsFacet
https://etherscan.io/address/0xBCF17C031Ea9C39261E345e65c8f60cAbdb1CD5A
BarnFact
https://etherscan.io/address/0x03e7081DB878ffeFE9a2800ccbe1F377208Bb546

Barn
https://etherscan.io/address/0xb4200c8c44b05a342a9f7fd0d27647c4bf9533e7
Rewards
https://etherscan.io/address/0xac01adc15878fae7d9b580d6fb695aa735738856

Gov (Governance.sol)
https://etherscan.io/address/0x6f1e586c62f7d8a1b7394f1a81a75aa68e109365
```
### Testnet
```
CommunityVault (for pre-staking)
https://ropsten.etherscan.io/address/0xd15c3473da91be7a49e27f3f7ece055e4ceecb0e
CommunityVault (for node-staking)
https://ropsten.etherscan.io/address/0xF9E552e3b21DB6638caaff4A583e3009cC579e68

DiamondCutFacet
https://ropsten.etherscan.io/address/0xD638356c72D08b0c55a6030A43eE28EAF00bDe7d
DiamondLoupeFacet
https://ropsten.etherscan.io/address/0x3d6aa943d23517857CF99C5F318c9c595bf0B925
OwnershipFacet
https://ropsten.etherscan.io/address/0x4f681746b8775a2c469355fdbb1ca937212a9a77
ChangeRewardsFacet
https://ropsten.etherscan.io/address/0xf92342cbbf265371740f8bac761d61ccaf6339f8
BarnFact
https://ropsten.etherscan.io/address/0x9cb78cf2a86ec22a9212ed96a8e61a16919ce2fe

Barn
https://ropsten.etherscan.io/address/0x9170f8d749dcf64467793325512a5e34b2b189eb
Rewards
https://ropsten.etherscan.io/address/0xbced010b27dc675c46f2526d21e4f1b01eac669f
NodeRewards
https://ropsten.etherscan.io/address/0xfA3ad7C76E3d4fD2113a5e3f476a1Ea0B9C07919

sbBTCPool
https://ropsten.etherscan.io/address/0x313c379ec483678f0ae9f612e56ea983636dd627

Gov (Governance.sol)
https://ropsten.etherscan.io/address/0xb7EAB16427009dae4e063cb723c6a1450C874996

```
## Discussion
For any concerns with the platform, open an issue on GitHub or visit us on [Discord](https://discord.gg/9TTQNUzg) to discuss.
For security concerns, please email info@barnbridge.com.

Copyright 2021 BarnBridge DAO
