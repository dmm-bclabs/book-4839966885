pragma solidity >=0.4.24 <0.6.0;

contract Variables {
    uint[] public stateVariable;

    constructor () public {
        stateVariable.push(100); // 0番目の値
        stateVariable.push(101); // 1番目の値

        setMemoryLocalVariable(stateVariable); //stateVariable[0]は100のまま
        setStorageLocalVariable(stateVariable); //stateVariable[1]は201に変更されている
    }

    function setMemoryLocalVariable (uint[] memory _stateVariable) pure private {
        _stateVariable[0] = 200;
    }

    function setStorageLocalVariable (uint[] storage _stateVariable) private {
        _stateVariable[1] = 201;
    }
}
