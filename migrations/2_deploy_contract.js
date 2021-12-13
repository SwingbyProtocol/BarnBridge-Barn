const Barn = artifacts.require("Barn")
const Rewards = artifacts.require("Rewards")

module.exports = async function (deployer, net) {

    if (net == "ropsten") {
        console.log("DEPLOYING TO ROPSTEN")

        const arguments = []

        await deployer.deploy(Barn, arguments)
        const barn = await Barn.deployed()
    }


};