package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
	"google.golang.org/grpc"
	"log"
	"net"
	"test1/pkg/api"
	"test1/pkg/finder"
	"test1/pkg/storage"
)

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(8080:8080)/booksdb?multiStatements=true")
	if err != nil {
		panic(err)
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"test1/pkg/schema",
		"mysql",
		driver,
	)

	m.Steps(2)

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
