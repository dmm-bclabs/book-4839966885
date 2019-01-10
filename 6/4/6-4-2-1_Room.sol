pragma solidity ^0.4.24;

// ①
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

// ②
contract Room is Ownable {

    // ③
    constructor(address _creator) public payable {
        owner = _creator;
    }

    // ④
    function deposit() external payable {}

    // ⑤
    function sendReward(uint256 _reward, address _dest) external onlyOwner {
        _dest.transfer(_reward);
    }

    // ⑥
    function refundToOwner() external onlyOwner {
        owner.transfer(address(this).balance);
    }
}
