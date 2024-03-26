package main

import (
	"context"
	"flag"
	"go-rpc/db"
	"log"

	"github.com/joho/godotenv"
)

//go-rpc is the name of the project in go.mod
//db is the name of the package in the project
//when we import the package, we import it as go-rpc/db
//and it will have the functions and variables defined in the db package

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "Run the server locally")
	flag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}
	cfg := db.NewConfig()
	conn, err := db.NewMongoDBConn(ctx, cfg)
	if err != nil {
		log.Panicln(err)
	}
	if err != nil {
		log.Fatal("Can not connect mongodb", err)
	}

	defer func() {
		if err := conn.Disconnect((ctx)); err != nil {
			log.Fatal("Can not disconnect mongodb", err)
		}
	}()
	log.Println("Connected to mongodb")
}
