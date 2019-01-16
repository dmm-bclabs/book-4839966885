pragma solidity 0.4.24;

contract RoomFactory is Destructible, Pausable {
    // イベントの宣言
    event RoomCreated(
        address indexed _creator,
        address _room,
        uint256 _depositedValue
    );

    // 〜(略)〜

    function createRoom() external payable whenNotPaused {
        address newRoom = (new Room).value(msg.value)(msg.sender);
        emit RoomCreated(msg.sender, newRoom, msg.value); // イベントの発火
    }

    // 〜(略)〜

}
