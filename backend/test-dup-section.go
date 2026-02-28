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

	_, err = db.Exec(
		"INSERT INTO sections (id, portfolio_id, type, position, data, is_visible, column_span) VALUES (?, ?, ?, ?, ?, ?, ?)",
		"TEST_SECTION_ID", "TEST_D2", "hero", 0, "{}", true, "full",
	)
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}
	fmt.Println("Success")
}
