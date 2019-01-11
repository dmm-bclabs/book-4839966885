const RoomFactory = artifacts.require('./RoomFactory.sol');

module.exports = deployer => {
    deployer.deploy(RoomFactory).then(instance => {
        console.log('ABI:', JSON.stringify(instance.abi));
    });
}
