package server

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var dbURL = "postgres://YOUR_SUPABASE_USER:YOUR_SUPABASE_PASSWORD@YOUR_SUPABASE_HOST:5432/postgres"

func main() {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected to Supabase!")
}
