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
        }
    }
}
