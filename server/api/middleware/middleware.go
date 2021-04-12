package middleware

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harishankarsivaji/Todo_App_Go/server/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = os.Getenv("MONGODB_CONN_STRING")

var dbName = os.Getenv("DB_NAME")

var collName = os.Getenv("COLLECTION_NAME")

var collection *mongo.Collection

func init() {

	func() {
		// 	var filename string = "logfile.log"
		// 	// Create the log file if doesn't exist. And append to it if it already exists.
		// 	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// Formatter := new(log.JSONFormatter)
		// Formatter.TimestampFormat = "02-01-2006 15:04:05"
		// log.SetFormatter(Formatter)

		// 	wrt := io.MultiWriter(os.Stdout, file)

		// log.SetOutput()

		// 	// Calling method as a field - Logs the func and file path
		// log.SetReportCaller(true)

		log.Info("Logger has been initialized.")
	}()

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	log.Info("Collection instance created!")
}

func GetAllTask(c *gin.Context) {
	c.Header("Context-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": payload,
	})
}

func CreateTask(c *gin.Context) {
	c.Header("Context-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList
	c.BindJSON(&task)
	insertOneTask(task)
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": task,
	})
}

func TaskComplete(c *gin.Context) {

	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	id := c.Params.ByName("id")
	taskComplete(id)
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": id,
	})
}

func UndoTask(c *gin.Context) {

	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	id := c.Params.ByName("id")
	undoTask(id)
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": id,
	})
}

func DeleteTask(c *gin.Context) {

	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	id := c.Param("id")
	deleteOneTask(id)
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": id,
	})

}

func DeleteAllTask(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")

	count := deleteAllTask()
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": count,
	})
}

func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.Background())
	return results
}

func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Inserted a Single Record: ", insertResult.InsertedID)
}

func taskComplete(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Modified count: ", result.ModifiedCount)
}

func undoTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Modified count: ", result.ModifiedCount)
}

func deleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	del, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Deleted Document: ", del.DeletedCount)
}

func deleteAllTask() int64 {
	del, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("No. of documents deleted: ", del.DeletedCount)
	return del.DeletedCount
}
