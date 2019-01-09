pragma solidity ^0.4.24;

contract Factory {
    
    address[] childList;

    function createChildContract() public payable {
      address newChild = new Child(msg.sender);            
      childList.push(newChild);   
    }

    function getDeployedChildContracts() public view returns (address[]) {
      return childList;
    }
}

contract Child {
    
    address owner;
    
    constructor(address _owner) public {
        owner = _owner;
    }
}
