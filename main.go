package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	_ "github.com/denisenkom/go-mssqldb"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print(".env loaded")
	}
}

func main() {
	Info := New("sql")
	fmt.Printf(" password:%s\n", Info.Sqlconfig.User)

	TestType()
	TelegramBot()
	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
	// if *debug {
	// 	fmt.Printf(" connString:%s\n", connString)
	// }
	// conn, err := sql.Open("mssql", connString)
	// defer conn.Close()
	// if err != nil {
	// 	log.Fatal("Open connection failed:", err.Error())
	// }

	// //	rows, err := conn.Query("select appid from game_info")
	// rows, err := conn.Query("EXECUTE GetInfoFromBase")
	// if err != nil {
	// 	log.Fatal("EXECUTE failed:", err.Error())
	// }

	// for rows.Next() {
	// 	var appid int
	// 	var gamename string
	// 	var price_min int
	// 	var price_max int
	// 	var nextupdate int
	// 	err = rows.Scan(&appid, &gamename, &price_min, &price_max, &nextupdate)
	// 	fmt.Print(appid)
	// }
}
