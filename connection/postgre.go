package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DataBaseConnection() {
	var err error
	databaseUrl := "postgres://postgres:14121995@localhost:5432/personal_web_b43"
	Conn, err = pgx.Connect(context.Background(), databaseUrl) //koneksi database url. context background kosong
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err) //passing value
		os.Exit(1)
	}

	fmt.Println("Database connected,")
}

// postgres://{user}:{password}@{host}:{port}/{database}
// user = user postgre
// pasword = password postgrenya
// host = host postgrenya
// port = port postgrenya
// database = database postgrenya
