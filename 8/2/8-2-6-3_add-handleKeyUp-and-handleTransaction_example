handleKeyUP = e => {
    this.setState({ deposit: e.target.value })
}

handleCreateRoom = async () => {
    const [ethaddress] = await web3.eth.getAccounts()

    try { // ①
        const receipt = await roomFactory.methods.createRoom().send({
            from: ethaddress,
            value: web3.utils.toWei(this.state.deposit.toString(), 'ether')
        })
    } catch (error) {
        console.log(error)
    }
}
