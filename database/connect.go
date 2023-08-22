package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGNAME"),
	)
	var err error
	Pool, err = pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to the database!")
}

func DisconnectFromDB() {
	Pool.Close()
}
