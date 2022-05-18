package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseMongo struct {
	Context  context.Context
	Client   *mongo.Client
	Database *mongo.Database
}

var (
	usr      = "jesus"
	pwd      = "123456"
	host     = "localhost"
	port     = 27017
	database = "tutorial"
)

func (c *DatabaseMongo) StartDB() {

	//uri := "mongodb://localhost:27017/"
	//uri := "mongodb://root:example@localhost:27017/"
	uri := "mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority"
	// uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)

	//client, err := mongo.NewClient(options.Client().ApplyURI(uri)) //  Connect to my cluster
	// if err != nil {
	// 	log.Fatal("mongo.NewClient:", err)
	// }

	//credential := options.Credential{AuthSource: "TESTE", Username: "root", Password: "example"}
	//clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	clientOpts := options.Client().ApplyURI(uri)

	var ctx = context.TODO()
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	//err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to mongo: error=%v", err)
	}

	// defer client.Disconnect(ctx)
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		log.Fatal("client.Disconnect:", err)
	// 	}
	// }()

	err = client.Ping(ctx, nil)
	//err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error ping to mongo: error=%v", err)
	}

	c.Context = ctx
	c.Client = client
	c.Database = c.Client.Database("cms-api-fiber")
}

func (c *DatabaseMongo) GetDatabase() *mongo.Database {
	return c.Database
}

func (c *DatabaseMongo) Close() {
	c.Client.Disconnect(c.Context)
}
