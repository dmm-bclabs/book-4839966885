pragma solidity >=0.4.24 <0.6.0;

contract SetOwnerContract {

    address owner; // 状態変数
    constructor() public { // コンストラクタの宣言
        owner = msg.sender; // 状態変数ownerにmsg.senderを代入
    }
}