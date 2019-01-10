pragma solidity >=0.4.21 <0.6.0;

contract SimpleAuction {
    event HighestBidIncreased(address bidder, uint amount); // イベントの宣言

    function bid() public payable {
        // ...
        emit HighestBidIncreased(msg.sender, msg.value); // イベントの発火
    }
}