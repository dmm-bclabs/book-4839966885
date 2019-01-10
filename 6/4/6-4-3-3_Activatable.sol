pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract Activatable is Ownable {

    event Deactivate(address indexed _sender);
    event Activate(address indexed _sender);

    // ①
    bool public active = false;

    // ②
    modifier whenActive() {
        require(active);
        _;
    }

    // ③
    modifier whenNotActive() {
        require(!active);
        _;
    }

    // ④
    function deactivate() public whenActive onlyOwner {
        active = false;
        emit Deactivate(msg.sender);
    }

    // ⑤
    function activate() public whenNotActive onlyOwner {
        active = true;
        emit Activate(msg.sender);
    }
}
