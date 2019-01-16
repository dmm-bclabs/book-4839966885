package main

// ①
import (
    "./bindings"
    "context"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "log"
    "math/big"
    "strings"
)

// ②
type RoomCreatedEvent struct {
    Creator common.Address
    Room common.Address
    DepositedValue *big.Int
}

func main() {
    // ③
    client, err := ethclient.Dial("ws://localhost:8546")
    if err != nil {
        log.Fatalf("err: %v\n", err)
    }

    // ④
    topics := map[string]common.Hash{
        "RoomCreated": crypto.Keccak256Hash([]byte("RoomCreated(address,address,uint256)")),
    }

    // ⑤
    query := ethereum.FilterQuery{
        Addresses: []common.Address{common.HexToAddress("0x7d3719f08bcb02e934eb2c8d5dce76952e79a130")},
        Topics: [][]common.Hash{{
            topics["RoomCreated"],
        }},
    }

    // ⑥
    event := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, event)
    if err != nil {
        log.Fatal(err)
    }

    roomFactoryAbi, err := abi.JSON(strings.NewReader(string(bindings.RoomFactoryABI)))
    if err != nil {
        log.Fatal(err)
    }

    // ⑦
    for {
        select {
            // ⑧
            case err := <-sub.Err():
                log.Println(err)
                close(event)
            // ⑨
            case vLog := <-event:
                if len(vLog.Topics) == 0 {
                    log.Println("Topicsが存在しません")
                }

                // ⑩
                switch vLog.Topics[0] {
                    case topics["RoomCreated"]:
                        var roomCreatedEvent RoomCreatedEvent
                        if err := roomFactoryAbi.Unpack(&roomCreatedEvent, "RoomCreated", vLog.Data); err != nil {
                            log.Println("unpackに失敗しました")
                            continue
                        }

                        // データベースに新しいRoomを登録する処理

                        break
                }
        }
    }
}
