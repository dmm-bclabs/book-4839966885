pragma solidity ^0.4.24;

// ①
import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "openzeppelin-solidity/contracts/lifecycle/Destructible.sol";

contract Room is Destructible, Pausable {

    // ②
    mapping (uint256 => bool) public rewardSent;

    // ③
    event Deposited(
        address indexed _depositor,
        uint256 _depositedValue
    );

    // ④
    event RewardSent(
        address indexed _dest,
        uint256 _reward,
        uint256 _id
    );

    // ⑤
    event RefundedToOwner(
        address indexed _dest,
        uint256 _refundedBalance
    );

    constructor(address _creator) public payable {
        owner = _creator;
    }

    // ⑥
    function deposit() external payable whenNotPaused {
        require(msg.value > 0); // ⑦
        emit Deposited(msg.sender, msg.value);
    }

    // ⑧
    function sendReward(uint256 _reward, address _dest, uint256 _id) external onlyOwner {
        require(!rewardSent[_id]); // ⑨
        require(_reward > 0); // ⑩
        require(address(this).balance >= _reward); // ⑪
        require(_dest != address(0)); // ⑫
        require(_dest != owner); // ⑬

        rewardSent[_id] = true; // ⑨
        _dest.transfer(_reward);
        emit RewardSent(_dest, _reward, _id);
    }

    function refundToOwner() external onlyOwner {
        require(address(this).balance > 0); // ⑭

        uint256 refundedBalance = address(this).balance;
        owner.transfer(refundedBalance);
        emit RefundedToOwner(msg.sender, refundedBalance);
    }
}
