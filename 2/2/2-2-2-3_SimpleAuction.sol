pragma solidity >=0.4.24 <0.6.0;

contract SimpleAuction {

    address public highestBidder;
    uint public highestBid;

    function bid() public payable { // 関数
        // ...
        highestBidder = msg.sender;
        highestBid = msg.value;
    }
}
