package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Query struct {
	Id            uint32
	Category      string
	Name          string
	Price         float32
	Savings       float32
	Compare_price string
	Compare_unit  string
	Brand_weight  string
	SortOrder     uint8
}

type List struct {
	Id          uint32
	Name        string
	Price       string
	BrandWeight string
	SortOrder   uint8
}

type QueryList struct {
	Query []Query
	List  []List
}

func handleIndex(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	list, err := db.Query("select * from list")
	if err != nil {
		log.Fatal(err)
	}
	defer list.Close()

	data := []List{}

	for list.Next() {
		listQuery := List{}
		list.Scan(&listQuery.Id, &listQuery.Name, &listQuery.Price, &listQuery.BrandWeight, &listQuery.SortOrder)
		data = append(data, listQuery)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmplpath, _ := filepath.Abs("./views/index.html")
		tmpl, err := template.ParseFiles(tmplpath)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
		}

		log.Printf("data: %v\n", data)

		tmpl.Execute(w, data)
	}
}

func handleSearchItem(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item := r.URL.Query().Get("item")
		rows, err := db.Query(fmt.Sprintf("SELECT id, name, price, savings FROM test WHERE name LIKE '%%%v%%'", item))
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var queryData []Query
		i := 0
		for rows.Next() {
			if i >= 5 {
				break
			}
			var row Query
			err := rows.Scan(&row.Id, &row.Name, &row.Price, &row.Savings)
			if err != nil {
				log.Fatal(err)
			}

			row.Price = row.Price - row.Savings
			queryData = append(queryData, row)
			i++
		}

		tmplpath, _ := filepath.Abs("./views/searchhits.html")
		tmpl := template.Must(template.ParseFiles(tmplpath))
		tmpl.ExecuteTemplate(w, "options", queryData)
	}
}

func handleNewItem(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	rows, err := db.Query("SELECT * FROM list")
	if err != nil {
		log.Fatal(fmt.Sprintf("Select * from list %v", err))
	}

	var data []List

	for rows.Next() {
		list := List{}
		rows.Scan(&list.Name, &list.Price, &list.BrandWeight, &list.SortOrder)
		data = append(data, list)

	}

	rows.Close()

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PostFormValue("add")
		newRow := List{}

		row, err := db.Query(fmt.Sprintf("SELECT name, price, brand_weight, SortOrder FROM test WHERE id=%v", id))
		if err != nil {
			log.Fatal(err)
		}

		if row.Next() {
			err := row.Scan(&newRow.Name, &newRow.Price, &newRow.BrandWeight, &newRow.SortOrder)
			if err != nil {
				log.Fatalf("Scan didn't work: %v\n", err)
			}
		}

		row.Close()

		_, err = db.Exec(`INSERT INTO list (
		name,
		price,
		BrandWeight,
		SortOrder
		)
		VALUES (?, ?, ?, ?)`,
			newRow.Name, newRow.Price, newRow.BrandWeight, newRow.SortOrder)

		if err != nil {
			log.Fatal(err)
		}

		data = append(data, newRow)
		tmplpath, _ := filepath.Abs("./views/index.html")
		tmpl := template.Must(template.ParseFiles(tmplpath))
		tmpl.ExecuteTemplate(w, "shopping-list-element", data)
	}
}

func main() {
	dbpath, _ := filepath.Abs("./db/willys.db")
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s", dbpath))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS
	list(
		id INTEGER PRIMARY KEY,
		name TEXT,
		price DECIMAL(4,2),
		BrandWeight TEXT,
		SortOrder INTEGER
	)`)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not prepare DB: %v\n", err))
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not create 'list': %v\n", err))
	}

	http.HandleFunc("/", handleIndex(db))
	http.HandleFunc("/new-item/", handleNewItem(db))
	http.HandleFunc("/search-item/", handleSearchItem(db))

	log.Fatal(http.ListenAndServe(":3333", nil))
}
