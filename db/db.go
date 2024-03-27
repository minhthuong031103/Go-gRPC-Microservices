package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

//in go when a method return a pointer, it typically indicates
//that the method will modify the value of the receiver

// In Go, when you declare a struct field with a type like *SomeType,
// it means that field holds a pointer to an instance of SomeType
// Storing pointers to mgo.Session and mgo.Database instances
// allows you to directly interact with the actual sessions and databases when working with conn objects.
// This can be more efficient and allows you to easily modify or access the session and database data as needed.

func NewMongoDBConn(ctx context.Context, cfg Config) (*mongo.Client, error) {
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(cfg.Dsn()).SetConnectTimeout(connectTimeout).
			SetMaxConnIdleTime(maxConnIdleTime).
			SetMinPoolSize(minPoolSize).
			SetMaxPoolSize(maxPoolSize),
	)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil

}

//if we not define the c *conn as the pointer receiver
//when we use the c variable, we will not see the methods Close() and DB() defined on the conn struct

//With (c conn) as the receiver, these methods will operate on a copy of the conn struct.
// Any modifications made to c within the methods will not affect the original conn struct.
//so we use a pointer receiver to modify the original struct
//and because of that, the NewCOnnection function has to return a pointer to the conn struct
//because the methods Close() and DB() are defined on the pointer receiver, not the value receiver

//Now, the conn struct in your code has methods Close() and DB()
//with the exact signatures required by the Connection interface:

//Even though there is no explicit declaration that conn implements COnnection interface
//Go will automatically infer that conn implements Connection interface
// Because it has methods that match the signature defined in the interface
// and also because in the func NewConnection(cfg Config) (Connection, error) function
// we are returning a pointer to conn struct which has the methods Close() and DB() defined
//so, the conn struct implements the Connection interface
//and the Connection interface is satisfied by the conn struct
