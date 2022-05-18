package main

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/gookit/color.v1"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.bd.mongodb
// go get go.mongodb.org/mongo-driver/mongo
// go get gopkg.in/gookit/color.v1
// go mod tidy
// go run main.go

func main() {

	var ctx = context.TODO()
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	uri := "mongodb://localhost:27017/"
	uri = "mongodb://root:example@localhost:27017/"
	uri = "mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority"

	//credential := options.Credential{AuthSource: "TESTE", Username: "root", Password: "example"}
	//clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	clientOpts := options.Client().ApplyURI(uri)

	//client, err := mongo.NewClient(options.Client().ApplyURI(uri)) //  Connect to my cluster
	// if err != nil {
	// 	log.Fatal("mongo.NewClient:", err)
	// }

	client, err := mongo.Connect(ctx, clientOpts)
	//err = client.Connect(ctx)
	if err != nil {
		log.Fatal("client.Connect:", err)
	}

	//defer client.Disconnect(ctx)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal("client.Disconnect:", err)
		}
	}()

	err = client.Ping(ctx, nil)
	//err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("client.Ping:", err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{}) // List databases
	if err != nil {
		log.Fatal("client.ListDatabaseNames:", err)
	}
	log.Println("databases:", databases)

	// var collection *mongo.Collection
	collection := client.Database("teste").Collection("posts")

	// client.Database("<db>").Collection("<collection>").InsertOne(ctx, bson.D{{"x", 1}})

	docs := []interface{}{
		bson.D{{"title", "World"}, {"body", "Hello World"}},
		bson.D{{"title", "Mars"}, {"body", "Hello Mars"}},
		bson.D{{"title", "Pluto"}, {"body", "Hello Pluto"}},
	}

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal("collection.InsertMany:", insertErr)
	}
	log.Println(res)

	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		log.Fatal("collection.Find:", currErr)
	}
	defer cur.Close(ctx)

	var posts []Post

	if err = cur.All(ctx, &posts); err != nil {
		log.Fatal("cur.All:", err)
	}

	log.Println("posts")
	//log.Println(posts)
	for _, v := range posts {
		log.Println(v.Body, v.Title)
	}

	collection = client.Database("teste").Collection("task")

	log.Println("createTask")
	task := &Task{ID: primitive.NewObjectID(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Text: "Teste01", Completed: false}
	createTask(ctx, collection, task)
	task = &Task{ID: primitive.NewObjectID(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Text: "Teste02", Completed: false}
	createTask(ctx, collection, task)
	task = &Task{ID: primitive.NewObjectID(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Text: "Teste03", Completed: true}
	createTask(ctx, collection, task)
	task = &Task{ID: primitive.NewObjectID(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Text: "Teste04", Completed: false}
	createTask(ctx, collection, task)

	log.Println("getAll") // alistar todos
	tasks, err := getAll(ctx, collection)
	printTasks(tasks)

	log.Println("completeTask") // alterar Completed = true
	completeTask(ctx, collection, "Teste02")

	log.Println("getPending") // listar os Completed = false
	tasks, _ = getPending(ctx, collection)
	printTasks(tasks)

	log.Println("getFinished") // listar os Completed = true
	tasks, _ = getFinished(ctx, collection)
	printTasks(tasks)

	log.Println("deleteTask") // deletar
	deleteTask(ctx, collection, "Teste04")
	//deleteTask(ctx, collection, "cccccccccccc")

}

type Post struct {
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

func createTask(ctx context.Context, collection *mongo.Collection, task *Task) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}

func getAll(ctx context.Context, collection *mongo.Collection) ([]*Task, error) {
	filter := bson.D{{}}
	return filterTasks(ctx, collection, filter)
}

func filterTasks(ctx context.Context, collection *mongo.Collection, filter interface{}) ([]*Task, error) {
	var tasks []*Task

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t Task
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	cur.Close(ctx) // once exhausted, close the cursor

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func printTasks(tasks []*Task) {
	for i, v := range tasks {
		if v.Completed {
			color.Green.Printf("%d: %s\n", i+1, v.Text)
		} else {
			color.Yellow.Printf("%d: %s\n", i+1, v.Text)
		}
	}
}

func completeTask(ctx context.Context, collection *mongo.Collection, text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "completed", Value: true}}}}
	t := &Task{}
	return collection.FindOneAndUpdate(ctx, filter, update).Decode(t)
}

func getPending(ctx context.Context, collection *mongo.Collection) ([]*Task, error) {
	filter := bson.D{primitive.E{Key: "completed", Value: false}}
	return filterTasks(ctx, collection, filter)
}
func getFinished(ctx context.Context, collection *mongo.Collection) ([]*Task, error) {
	filter := bson.D{primitive.E{Key: "completed", Value: true}}
	return filterTasks(ctx, collection, filter)
}

func deleteTask(ctx context.Context, collection *mongo.Collection, text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("No tasks were deleted")
	}
	return nil
}
