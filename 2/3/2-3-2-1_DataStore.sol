pragma solidity ^0.4.24;

contract DataStore {
    uint public storedData;

    constructor() public {
        storedData = 100;
    }
}