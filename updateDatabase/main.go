package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"willys/source"
	"willys/structs"
)

func getProducts(url string) structs.WillysJson {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonResp := structs.WillysJson{}

	json.Unmarshal(body, &jsonResp)

	return jsonResp
}

func addToDatabase(db *sql.DB, answer structs.WillysJson, sortOrder int) {
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS 
		test(
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		price DECIMAL(4, 2), 
		savings DECIMAL(3,2), 
		compare_price CHAR(10), 
		compare_unit CHAR(10), 
		brand_weight TEXT, 
		sortOrder INTEGER)`)
	if err != nil {
		panic(err)
	}

	statement.Exec()
	statement, err = db.Prepare(`INSERT INTO test (
		name, 
		price, 
		savings, 
		compare_price, 
		compare_unit, 
		brand_weight, 
		sortOrder
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(answer.Results); i++ {
		statement.Exec(
			answer.Results[i].Name,
			answer.Results[i].PriceValue,
			answer.Results[i].SavingsAmount,
			answer.Results[i].ComparePrice,
			answer.Results[i].ComparePriceUnit,
			answer.Results[i].ProductLine2,
			sortOrder,
		)
	}
}

func main() {
	database, err := sql.Open("sqlite3", "willys.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	for i := 0; i < len(source.Alla); i++ {
		resp := getProducts(source.Alla[i])

		addToDatabase(database, resp, i)
	}
}
