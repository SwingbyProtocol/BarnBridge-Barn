
const { BigNumber, Ethers } = require('ethers')
require('dotenv').config();


const _owner = process.env.OWNER;
const _bond = process.env.BOND

//rewards
const diamond = '0x56d3cf28c094A508E5031e331Fb67F1C78182F11' //barn

module.exports = [
    _owner,
    _bond,
    diamond
]
/**
 //barn
//npx hardhat verify --constructor-args ./scripts/args.js 0x56d3cf28c094A508E5031e331Fb67F1C78182F11 --network ropsten
module.exports = [
    [['0xE9aCeF220EA60CfC4f9bBf35C27D7A0832C53130',
        0,
        ['0x1f931c1c']],
    ['0xe94fB173E2D30635252021DfD4bd289E1cDBEa09',
        0,
        ['0xcdffacc6',
            '0x52ef6b2c',
            '0xadfca15e',
            '0x7a0ed627',
            '0x01ffc9a7']],
    ['0x46024e861d8eEc2498418Ec405ec13231Ef7FBcA',
        0,
        ['0x8da5cb5b', '0xf2fde38b']],
    ['0xa2B18457011877dB95eD388E5f1b861bc4bcD741',
        0,
        ['0x8d240d8b']],
    ['0xBab53e08f1D999D2528A088ef2a3b42cD9f1aE4a',
        0,
        ['0x65a5d5f0',
            '0x20e77a84',
            '0x417edd4d',
            '0x70a08231',
            '0xc2077e81',
            '0xf77f962f',
            '0x697eb324',
            '0x5c19a95c',
            '0x169df064',
            '0xd265a115',
            '0xb6b55f25',
            '0xbfc10279',
            '0x5107a3ae',
            '0xdd467064',
            '0x7a141096',
            '0x8e4a5248',
            '0x18ab6a3c',
            '0x6f121578',
            '0xbef624d8',
            '0xbf0ae48c',
            '0xc07473f6',
            '0xcbf8eda9',
            '0x2e1a7d4d']]],
    '0xbD8332654deFf42Ee2E7ec3E927e58b4e9c6CCF4' //_owner
]
 */