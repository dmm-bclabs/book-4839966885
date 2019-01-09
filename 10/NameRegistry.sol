pragma solidity ^0.4.24;

contract NameRegistry {

   mapping(string => address) registry;

   function registerName(string name, address addr) public returns (bool) {
       registry[name] = addr;
       return true;
   }

    function getContractAddress(string name) public view returns(address) {
      return registry[name];
   }

}
