package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	Dsn() string
	DbName() string
}

// In Go, names starting with an uppercase letter are considered exported, meaning they are accessible from outside the package.
// Names starting with a lowercase letter are considered unexported, meaning they are only accessible from within the same package.

type config struct {
	// In the config struct, fields such as dbUser, dbPass, dbHost, dbPort, dbName, and dsn are all unexported because they start with a lowercase letter.
	//  This means they are only accessible within the db package.
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn    string
}

func NewConfig() Config {
	//in go, this is the way to define a constructor
	//this function is a constructor for the config struct
	//because the config struct is not exported, this function is the only way to create a new config struct
	var cfg config
	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbHost = os.Getenv("DATABASE_HOST")
	cfg.dbName = os.Getenv("DATABASE_NAME")
	var err error
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)

	return &cfg

}

func (c *config) Dsn() string {
	// c *config is a pointer receiver, this mean this function is a method of the config struct
	//this is the way go define methods for a struct
	// (c *config): This part specifies the receiver type.
	// It indicates that the Dsn() method is associated with a pointer to a config struct (*config).
	//  This means that the method operates on instances of the config struct by receiving a pointer to it.
	// This also allows the method to modify the fields of the config struct directly.
	// The (cfg *config) part is equivalent to "this" in other languages.
	//It allows the method to access and modify the fields of the config struct it's associated with.
	//By using a pointer receiver (*config), the method can modify the fields of the struct directly, rather than creating a copy of the struct.
	return c.dsn
}

func (c *config) DbName() string {
	// Go promotes separating data and behavior. Structs define the data, while methods define the behavior.
	// This separation keeps the code clean and organized.

	return c.dbName
}
