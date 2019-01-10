pragma solidity >=0.4.24 <0.6.0;

contract Deposit {
    uint depositThreshold;

    event LogDeposit(uint _time);

    constructor () public {
        depositThreshold = 1 ether;
    }

    function deposit () public payable {
        require(msg.value >= depositThreshold);

        emit LogDeposit(now);
    }

    function getBalance () public view returns(uint) {
        return address(this).balance;
    }
}