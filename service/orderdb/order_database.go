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
	"path/filepath"
	"runtime"
)

type OrderDB struct {
	Database *sql.DB
}

func (db *OrderDB) Connect() {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	if err := godotenv.Load(path + "/.env"); err != nil {
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
}

func (db *OrderDB) Get(uid *string) string {
	result, err := db.Database.Query(`SELECT "order" FROM orders WHERE $1 = order_id`, uid)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var data string
	for result.Next() {
		err = result.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
	}
	return data
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
	if _, err := db.Database.Exec(`TRUNCATE TABLE orders`); err != nil {
		log.Fatal(err)
	}
}
