pragma solidity ^0.4.24;

import "https://raw.githubusercontent.com/OpenZeppelin/openzeppelin-solidity/v1.11.0/contracts/ownership/rbac/RBAC.sol";

contract RoleSample is RBAC {

    string public constant ROLE_ADMIN = "admin";
    string public constant ROLE_MEMBER = "member";

    modifier onlyAdmin() {
        checkRole(msg.sender, ROLE_ADMIN);
        _;
    }
    
    constructor() public {
        addRole(msg.sender, ROLE_ADMIN);
    }
    
    function addMember(address addr) onlyAdmin public {
        addRole(addr, ROLE_MEMBER);
    }
    
    // ...
}