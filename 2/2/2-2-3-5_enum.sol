pragma solidity >=0.4.24 <0.6.0;

contract Enum {

    enum Colors {Red, Blue, Green}
    Colors color;

    function setColor() public {
        color = Colors.Blue;
    }
}
