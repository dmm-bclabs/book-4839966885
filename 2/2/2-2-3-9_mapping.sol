mapping(address => uint) public balances;

function update(uint newBalance) public {
    balances[msg.sender] = newBalance;
}
