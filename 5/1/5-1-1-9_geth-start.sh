geth --datadir ~/private-net --networkid 15 \
--nodiscover --maxpeers 0 --mine --minerthreads 1 \
--rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" \
--rpcvhosts "*" --rpcapi "eth,web3,personal,net" \
--ipcpath ~/private-net/geth.ipc --ws --wsaddr "0.0.0.0" \
--wsapi "eth,web3,personal,net" --wsorigins "*" \
--unlock 0,1,2,3,4 --password ~/private-net/password
--allow-insecure-unlock
