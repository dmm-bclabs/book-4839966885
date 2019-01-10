#!/bin/sh
geth --datadir /var/share/ethereum --nodiscover --maxpeers 0 \
init /var/share/ethereum/genesis.json \
&& \
geth --datadir /var/share/ethereum --networkid 15 \
--nodiscover --maxpeers 0 --mine --minerthreads 1 \
--rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" \
--rpcvhosts "*" --rpcapi "eth,web3,personal,net" \
--ipcpath /tmp/geth.ipc --ws --wsaddr "0.0.0.0" \
--wsapi "eth,web3,personal,net" --wsorigins "*" \
--unlock 0,1,2,3,4 --password /var/share/ethereum/password
