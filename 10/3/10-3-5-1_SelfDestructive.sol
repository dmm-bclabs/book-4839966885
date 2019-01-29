pragma solidity ^0.4.24;

import "github.com/OpenZeppelin/openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract SelfDestructive is Ownable {

    constructor() public payable { }

    function destroy(address _recipient) onlyOwner public {
        selfdestruct(_recipient);
    }
}