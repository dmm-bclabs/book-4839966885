pragma solidity >=0.4.24 <0.6.0;

contract Mapping {

    mapping(address => uint) public balances;

    function update(uint newBalance) public {
        balances[msg.sender] = newBalance;
    }
}
