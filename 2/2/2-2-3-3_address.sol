pragma solidity ^0.4.24;

contract Address {

    constructor() payable public {
    }

    function transferToSender() payable public {
        address to = msg.sender;
        to.transfer(address(this).balance);
    }
    
    function getSenderBalance() view public returns(uint) {
        address to = msg.sender;
        return to.balance;
    }
}
