pragma solidity >=0.4.24 <0.6.0;

contract Bytes {

    constructor() public {
        bytes2 b = 'ba'; // 2バイトの固定バイト配列 
        bytes memory a = 'baaaaaaaaaa'; // 動的バイト配列
        string memory name = 'Miguel'; // stringの宣言
    }
}

