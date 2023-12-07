package orderdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Kosmosman/service/types"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type OrderDB struct {
	Database *sql.DB
}

func (db *OrderDB) Connect() {
	fp, _ := os.Getwd()
	if err := godotenv.Load(fp + "/orderdb/.env"); err != nil {
		log.Fatal(err)
	}
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s", os.Getenv("DB_USER"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PASSWORD"))
	var err error
	db.Database, err = sql.Open(os.Getenv("DB_CONNECTION"), connStr)
	if err != nil {
		log.Fatal()
	}
	if err := db.Database.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to orderdb!")
}

func (db *OrderDB) Add(uid *string, orderData *string) {
	if _, err := db.Database.Exec(`INSERT INTO orders(order_id, "order") VALUES($1, $2)`, uid, orderData); err != nil {
		log.Fatal(err)
	}
	println("Add new order!")
}

func (db *OrderDB) RestoreCache(cache *types.Cache) {
	rows, err := db.Database.Query(`SELECT * FROM orders`)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}
	var uid, data string
	cache.Data = make(map[string][]byte)
	for rows.Next() {
		if err = rows.Scan(&uid, &data); err != nil {
			log.Fatal(err)
		}
		dataBytes, err := json.Marshal(&data)
		if err != nil {
			log.Fatal(err)
		}
		cache.Data[uid] = dataBytes
	}
}

func (db *OrderDB) ClearDB() {
	if err, _ := db.Database.Exec(`DELETE FROM orders`); err != nil {
		log.Fatal(err)
	}
}
