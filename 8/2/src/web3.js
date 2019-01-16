import Web3 from 'web3'

let web3

if (typeof window !== 'undefined' && typeof window.ethereum !== 'undefined') {
    web3 = new Web3(window.ethereum)
    window.ethereum.enable().catch(error => {
        // User denied account access
        console.log(error)
    })
} else if (typeof window !== 'undefined' && typeof window.web3 !== 'undefined') {
    web3 = new Web3(window.web3.currentProvider)
} else {
    const httpEndpoint = 'http://127.0.0.1:7545'
    const provider = new Web3.providers.HttpProvider(httpEndpoint)
    web3 = new Web3(provider)
}

export default web3
