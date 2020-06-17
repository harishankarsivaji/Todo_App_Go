package tests

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Test DB
const _connectionString = "mongodb://localhost:27017"

const _dbName = "testTodo"

const _collName = "testTodoList"

var _collection *mongo.Collection

func Test_Init(t *testing.T) {
	clientOptions := options.Client().ApplyURI(_connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	_collection = client.Database(_dbName).Collection(_collName)
	t.Log("MongoDB connected and collection created")
}

// func Test_GetAllTask(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/task", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(rr)

// 	GetAllTask(c)
// 	assert.Equal(t, 200, rr.Status)

// 	var got gin.H
// 	err := json.Unmarshal(&got, rr.Body().Bytes())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, want, got)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// }
