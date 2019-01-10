await this.roomFactory.pause({ from: roomOwner1 })
    .should.be.rejectedWith(EVMRevert)
await this.roomFactory.pause({ from: factoryOwner })
    .should.be.fulfilled
