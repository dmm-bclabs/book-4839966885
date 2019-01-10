package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	resp "github.com/nicklaw5/go-respond"
)

type (
	room struct {
		ID           uint64         `json:"id"`
		OwnerID      uint64         `json:"owner_id"`
		OwnerAddress string         `json:"owner_address"`
		Title        string         `json:"title"`
		Description  string         `json:"description"`
		EventCode    string         `json:"event_code"`
		Address      string         `json:"address"`
		CreateTxHash string         `json:"create_tx_hash"`
		IsPrivate    bool           `json:"is_private"`
		WeiBalance   uint64         `json:"wei_balance"`
		WeiPrize     uint64         `json:"wei_prize"`
		StartTime    mysql.NullTime `json:"start_time"`
		EndTime      mysql.NullTime `json:"end_time"`
		Active       bool           `json:"active"`
	}

	rooms []room
)

func main() {
	log.Println("Starting API Server!!")

	db := connectDB()
	defer db.Close()

	router := chi.NewRouter()

	// @Tags Room
	// @Accept json
	// @Summary ルーム情報を取得します
	// @Success 200 {array} []room
	// @Router /rooms [get]
	router.Get("/rooms", func(w http.ResponseWriter, r *http.Request) {
		rooms := []room{}
		db.Find(&rooms)

		resp.NewResponse(w).Ok(rooms)
	})

	http.ListenAndServe(":8080", router)
}

// Gormを利用してDBに接続する
func connectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql)"
	DBNAME := "test_db"
	option := "?charset=utf8&parseTime=True"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + option

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
