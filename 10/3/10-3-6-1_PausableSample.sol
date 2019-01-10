pragma solidity ^0.4.24;

import "github.com/OpenZeppelin/openzeppelin-solidity/contracts/lifecycle/Pausable.sol";

contract PausableSample is Pausable {
    
    constructor() public payable { }
    
    function normalFunction() public whenNotPaused {
        // one can do only when not paused
    }
}