pragma solidity ^0.4.24;

// ①
import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "openzeppelin-solidity/contracts/lifecycle/Destructible.sol";
import "./Room.sol";

// ②
contract RoomFactory is Destructible, Pausable {

    // ③
    event RoomCreated(
        address indexed _creator,
        address _room,
        uint256 _depositedValue
    );

    // ④
    function createRoom() external payable whenNotPaused {
        address newRoom = (new Room).value(msg.value)(msg.sender);
        emit RoomCreated(msg.sender, newRoom, msg.value);
    }
}
