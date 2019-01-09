pragma solidity ^0.4.24;

import "./Room.sol";

contract RoomFactory {

    function createRoom() external payable {
        address newRoom = (new Room).value(msg.value)(msg.sender);
    }
}
