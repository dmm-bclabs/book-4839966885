pragma solidity >=0.4.24 <0.6.0;

contract Address {

    constructor() payable public {
    }

    function transferToSender() payable public {
        address payable to = msg.sender;
        to.transfer(address(this).balance);
    }
    
    function getSenderBalance() view public returns(uint) {
        address payable to = msg.sender;
        return to.balance;
    }
}
