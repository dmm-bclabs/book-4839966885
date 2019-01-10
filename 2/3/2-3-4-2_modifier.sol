modifier onlyOwner() {
    require(msg.sender == owner);
    _;
}

function _createZombie(string _name, uint _dna) public onlyOwner {
    // ゾンビを作成するための実装
}
