pragma solidity >=0.4.24 <0.6.0;

contract SimpleStorage { // コントラクトの宣言
    uint storedData;

    function set(uint x) public {
        storedData = x;
    }

    function get() public view returns (uint) {
        return storedData;
    }
}