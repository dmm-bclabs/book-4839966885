const HDWalletProvider = require('truffle-hdwallet-provider')   // ①
const mnemonic = $MNEMONIC  // ②
const projectId = $INFURA_PROJECT_ID    // ③

module.exports = {
    networks: {
        development: {
            host: '127.0.0.1',
            port: 8545,
            network_id: 15,
            gas: 4700000
        },
        ganache: {
            host: '127.0.0.1',
            port: 7545,
            network_id: '*'
        },
        ropsten: {
            provider: () => {
                return new HDWalletProvider(
                    mnemonic,
                    `https://ropsten.infura.io/v3/${projectId}`
                );
            },
            network_id: 3,  // ④
            gas: 4700000
        }
    }
}
