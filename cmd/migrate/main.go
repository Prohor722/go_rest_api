package migrate

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'dowmn'")
	}

	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./data.db")

	if(err != nil) {
		log.Fatal(err)
	}
}