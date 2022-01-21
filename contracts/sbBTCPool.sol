// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.7.6;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/IBarn.sol";
import "./interfaces/ISwapContract.sol";

contract sbBTCPool is Ownable {
    using SafeMath for uint256;

    bool public disabled;
    uint256 public lastPullTs;

    uint256 public balanceBefore;
    uint256 public currentMultiplier;

    mapping(address => uint256) public userMultiplier;
    mapping(address => uint256) public owed;

    IBarn public immutable barn;
    IERC20 public immutable rewardToken;
    ISwapContract public swapContract;

    event Claim(address indexed user, uint256 amount);

    constructor(address _owner, address _sbBTC, address _barn) {
        require(_sbBTC != address(0), "reward token must not be 0x0");
        require(_barn != address(0), "barn address must not be 0x0");

        transferOwnership(_owner);

        rewardToken = IERC20(_sbBTC);
        barn = IBarn(_barn);
    }

    // claim calculates the currently owed reward and transfers the funds to the user
    function claim() public returns (uint256){
        require(swapContract.isNodeStaked(msg.sender), "msg.sender is not staker");

        _calculateOwed(msg.sender);

        uint256 amount = owed[msg.sender];
        require(amount > 0, "nothing to claim");

        owed[msg.sender] = 0;

        rewardToken.transfer(msg.sender, amount);

        // acknowledge the amount that was transferred to the user
        ackFunds();

        emit Claim(msg.sender, amount);

        return amount;
    }

    // ackFunds checks the difference between the last known balance of `token` and the current one
    // if it goes up, the multiplier is re-calculated
    // if it goes down, it only updates the known balance
    function ackFunds() public {
        uint256 balanceNow = rewardToken.balanceOf(address(this));

        if (balanceNow == 0 || balanceNow <= balanceBefore) {
            balanceBefore = balanceNow;
            return;
        }

        uint256 totalStakedBond = swapContract.totalStakedAmount();
        // if there's no bond staked, it doesn't make sense to ackFunds because there's nobody to distribute them to
        // and the calculation would fail anyways due to division by 0
        if (totalStakedBond == 0) {
            return;
        }

        uint256 diff = balanceNow.sub(balanceBefore);
        uint256 multiplier = currentMultiplier.add(diff.mul(decimals).div(totalStakedBond));

        balanceBefore = balanceNow;
        currentMultiplier = multiplier;
    }

    // setBarn sets the address of the BarnBridge Barn into the state variable
    function setSwap(address _swap) public {
        require(swapContract != address(0), 'barn address must not be 0x0');
        require(msg.sender == owner(), '!owner');

        swapContract = ISwapContract(_swap);
    }

    // _calculateOwed calculates and updates the total amount that is owed to an user and updates the user's multiplier
    // to the current value
    // it automatically attempts to pull the token from the source and acknowledge the funds
    function _calculateOwed(address user) internal {
        ackFunds();

        uint256 reward = _userPendingReward(user);

        owed[user] = owed[user].add(reward);
        userMultiplier[user] = currentMultiplier;
    }

    // _userPendingReward calculates the reward that should be based on the current multiplier / anything that's not included in the `owed[user]` value
    // it does not represent the entire reward that's due to the user unless added on top of `owed[user]`
    function _userPendingReward(address user) internal view returns (uint256) {

        uint256 multiplier = currentMultiplier.sub(userMultiplier[user]);

        return barn.balanceOf(user).mul(multiplier).div(decimals);
    }
}
