package middleware

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/harishankarsivaji/Todo_App_Go/server/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "todoApp"

// Collection name
const collName = "todolist"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
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

		log.Info("Logger has been initilized.")
	}()

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	log.Info("Collection instance created!")
}

// GetAllTask get all the task route
func GetAllTask(c *gin.Context) {
	c.Header("Context-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": payload,
	})
}

// CreateTask create task route
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

// TaskComplete update task route
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

// UndoTask undo the complete task route
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

// DeleteTask delete one task route
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

// DeleteAllTask delete all tasks route
func DeleteAllTask(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")

	count := deleteAllTask()
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": count,
	})
}

// get all task from the DB and return it
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

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Inserted a Single Record: ", insertResult.InsertedID)
}

// task complete method, update task's status to true
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

// task undo method, update task's status to false
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

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	del, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Deleted Document: ", del.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	del, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("No. of documents deleted: ", del.DeletedCount)
	return del.DeletedCount
}
