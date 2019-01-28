uint attackVictoryProbability = 70; // ① 攻撃が成功する確率

function attack(uint _zombieId, uint _targetId) external ownerOf(_zombieId) {
    Zombie storage myZombie = zombies[_zombieId]; // ② 自分のゾンビ
    Zombie storage enemyZombie = zombies[_targetId]; // ③ 敵のゾンビ
    uint rand = randMod(100);
    if(rand <= attackVictoryProbability) {
        myZombie.winCount++; // ④ 自分のゾンビが勝った回数
        myZombie.Level++;
        enemyZombie.lossCount++; // ⑤ 敵のゾンビが負けた回数
        feedAndMultiply(_zombieId, enemyZombie.dna, "zombie");
    }
}