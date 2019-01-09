pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract Room is Ownable {

    constructor(address _creator) public payable {
        owner = _creator;
    }

    function deposit() external payable {}

    function sendReward(uint256 _reward, address _dest) external onlyOwner {
        _dest.transfer(_reward);
    }

    function refundToOwner() external onlyOwner {
        owner.transfer(address(this).balance);
    }
}
