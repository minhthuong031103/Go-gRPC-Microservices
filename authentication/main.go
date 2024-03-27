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

//define a local variable//when run the program, can use command like -local=true or false to set the value of the local variable

func init() {
	// This is an initialization function named init().
	// In Go, init() functions are automatically called by Go's runtime before the main() function is executed.
	// 	Inside the init() function, it uses the flag package to define a boolean command-line flag -local using flag.BoolVar().
	// 	This flag determines whether to run the server locally.
	// flag.Parse() parses the command-line flags. This function must be called after all flags are defined and before accessing their values.
	flag.BoolVar(&local, "local", true, "Run the server locally")
	flag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//context.Background() returns a Context object representing the root of a context tree.
	//It's typically used as the starting point for creating other contexts.
	// In Go, a context is an object that provides a way to manage and propagate cancellation signals, deadlines, and other request-scoped values across API boundaries.
	//Think of it as a way to pass along important information about a request or operation as it travels through different parts of your program.
	//most use case is: manage cancellation, deadline, and request-scoped values across API boundaries
	//Contexts in Go form a hierarchy, where each context can have a parent context.
	//When you create a new context from an existing one, you're essentially creating a child context.
	//This hierarchical relationship allows for the propagation of deadlines, cancellation signals, and other request-scoped values from parent contexts to their children.
	//When you create a new context from an existing one using functions like context.WithCancel(), context.WithDeadline(), context.WithTimeout(), or context.WithValue(),
	//you're creating a derived context.
	//Derived contexts inherit certain properties from their parent context, such as cancellation signals or deadlines.
	// However, they can also have their own additional properties or values.

	defer cancel()
	// The cancel function is used to cancel the context and any contexts derived from it.
	//defer mean delay
	//The defer keyword is used to defer the execution of a function call until the surrounding function returns.
	//By deferring the invocation of cancel(), you ensure that the context is canceled when the main function returns, regardless of how it returns (e.g., normal execution, panic, or error).
	//In that case
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
		log.Fatal("Can not connect mongodb", err)

	}

	defer func() {
		if err := conn.Disconnect((ctx)); err != nil {
			log.Fatal("Can not disconnect mongodb", err)
		}
		// In this case, it's used to defer the disconnection of the MongoDB connection until the end of the main() function.
	}()
	log.Println("Connected to mongodb")
}
