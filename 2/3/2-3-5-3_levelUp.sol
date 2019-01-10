function levelUp (uint _zombieId) external payable {
    require(msg.value == 0.001 ether);

    zombies[_zombieId].level++;
}
