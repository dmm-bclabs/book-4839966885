pragma solidity ^0.4.24;
//コンパイラのバージョンが0.4.24,25,26に限定されていることに注意してください

contract Address {

    constructor() payable public {
    }

    function transferToSender() payable public {
        address to = msg.sender;
        to.transfer(address(this).balance);
    }

    function getSenderBalance() view public returns (uint) {
        address to = msg.sender;
        return to.balance;
    }
}
