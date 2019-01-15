pragma solidity >=0.4.24 <0.6.0;

contract Integer {

    constructor() public {
        uint a = 3;
        uint b = 2;
        uint c =  a / b * 2; // 3 / 2 は切り捨てで1になり、2倍してcは2になります。
        uint d =  3 / 2 * 2; // リテラルは切り捨てられず1.5になり、2倍してdは3になります。
    }
}

