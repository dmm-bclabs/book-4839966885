contract SetOwnerContract {

    address owner; // 状態変数
    constructor () public { // コンストラクタの宣言
        owner = msg.sender; // 状態変数ownerにmsg.senderを代入
    }
}