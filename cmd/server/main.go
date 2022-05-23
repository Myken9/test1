package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"test1/pkg/api"
	"test1/pkg/finder"
	"test1/pkg/storage"
)

func main() {

	db, err := sql.Open("mysql", "root:123456@/booksdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := grpc.NewServer()
	srv := &finder.GRPCServer{St: storage.NewStorage(db)}
	api.RegisterLibraryServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
