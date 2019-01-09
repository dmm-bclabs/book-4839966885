pragma solidity ^0.4.24;

contract SmartSpeaker {
    // publicなstate variableは自動的にgetterが生成される
    string public response;


    constructor () public {
        response = "";
    }

    function listen (uint _number) public {
        if (_number <= 100) {
            response = "small!";
        } else if (_number <= 1000) {
            response = "normal!!";
        } else if (_number <= 10000) {
            response = "large!!!";
        } else if (_number <= 100000) {
            response = "WOWOWO!!!!";
        } else {
            response = "unbelievable!!!!!";
        }
    }
}