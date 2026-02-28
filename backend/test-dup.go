package main

import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "ploykong:warisa2708@tcp(127.0.0.1:3306)/ploykong?parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`INSERT INTO portfolios (id, user_id, slug, title, description, theme, seo_title, seo_desc,
								 is_published, password_hash, expires_at, view_count)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		"TEST_D2", "01J5EXAMPLE000000000000001", "test-d2", "Test D2", "Desc", "{}",
		"", "", false, nil, nil, 0,
	)
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}
	fmt.Println("Success")
}
