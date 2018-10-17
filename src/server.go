package anachromeDb

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	// "os"
)

func main() {
	tempFilename := "test.db"
	// defer os.Remove(tempFilename)
	db, err := sql.Open("sqlite3", tempFilename)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	_, err = db.Exec("drop table foo")
	_, err = db.Exec("create table foo (id INTEGER PRIMARY KEY,name text)")
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	for i := 0; i < 10; i++ {
		res, err := db.Exec(fmt.Sprintf("insert into foo(id,name) values(%d,%s)", i, fmt.Sprint(i)))
		if err != nil {
			log.Fatal("Failed to insert record:", err)
		}
		affected, _ := res.RowsAffected()
		if affected != 1 {
			log.Fatalf("Expected %d for affected rows, but %d:", 1, affected)
		}
	}
	rows, err := db.Query("select id,name from foo")
	if err != nil {
		log.Fatal("Failed to select records:", err)
	}
	defer rows.Close()

	rows.Next()

	var result int
	rows.Scan(&result)
	if result != 123 {
		log.Fatalf("Expected %d for fetched result, but %d :", 123, result)
	}

	log.Printf("Success")
}
