package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/mattes/migrate/source/file"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"test1/pkg/api"
	"test1/pkg/server"
	"test1/pkg/storage"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	s := grpc.NewServer()
	srv := &server.GRPCServer{St: storage.NewStorage(db)}
	api.RegisterLibraryServer(s, srv)

	l, err := net.Listen("tcp", os.Getenv("LISTENED_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
