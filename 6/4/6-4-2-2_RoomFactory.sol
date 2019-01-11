pragma solidity ^0.4.24;

// ①
import "./Room.sol";

contract RoomFactory {

    // ②
    function createRoom() external payable {
        address newRoom = (new Room).value(msg.value)(msg.sender);
    }
}
