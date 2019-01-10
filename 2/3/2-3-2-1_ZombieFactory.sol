// バージョンプラグマの指定
pragma solidity ^0.4.19;

// コントラクトの宣言
contract ZombieFactory {

    // イベントの宣言
    event NewZombie(uint zombieId, string name, uint dna);

    // 変数（uint型）の宣言
    uint dnaDigits = 16;
    uint dnaModulus = 10 ** dnaDigits;

    // 構造体の宣言
    struct Zombie {
        string name;
        uint dna;
    }

    // public変数（Zombie型配列）の宣言
    Zombie[] public zombies;

    // private関数の宣言
    function _createZombie(string _name, uint _dna) private {
        uint id = zombies.push(Zombie(_name, _dna)) - 1;
        NewZombie(id, _name, _dna); //※event発行について
    } 

    // private view関数の宣言
    function _generateRandomDna(string _str) private view returns (uint) {
        uint rand = uint(keccak256(_str));
        return rand % dnaModulus;
    }

    // public関数の宣言
    function createRandomZombie(string _name) public {
        uint randDna = _generateRandomDna(_name);
        _createZombie(_name, randDna);
    }

}