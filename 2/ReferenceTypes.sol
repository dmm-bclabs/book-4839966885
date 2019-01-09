
pragma solidity ^0.4.24;

contract ReferenceTypes {

    uint[] public uintArray;

    struct Author {
        address addr;
        string name;
    }

    Author[] public authors;

    constructor () public {
        uintArray = [3, 65, 88, 1, 101010];

        address authorAddr = 0x0cE446255506E92DF41614C46F1d6df9Cc969183;  // CryptoZombiesのアドレス
        authors.push(Author(authorAddr, "blockchain_lab"));
    }
}