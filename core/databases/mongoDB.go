package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// 创建一个结构体
type Student struct {
	Name string
	Age  int
}

func initDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://10.249.213.177:27017")

	// 连接到mongoDB
	clientConnect, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err2 := clientConnect.Ping(context.TODO(), nil)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("连接mongoDB成功！")
	client = clientConnect
}

// 添加单个文档
func insertOne() {
	initDB()

	s := Student{
		Name: "zhangsan",
		Age:  18,
	}

	// 连接数据库Collection
	collection := client.Database("go_db").Collection("Student")

	insertOneResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted a single document: %v\n", insertOneResult.InsertedID)
}

// 添加多个文档
func insertMulti() {
	initDB()
	collection := client.Database("go_db").Collection("Student")

	s1 := Student{
		Name: "bob",
		Age:  44,
	}

	s2 := Student{
		Name: "Alice",
		Age:  34,
	}

	students := []interface{}{s1, s2}
	imr, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted multi document: %v\n", imr.InsertedIDs)
}

// 查询
func find() {
	initDB()

	ctx, cf := context.WithTimeout(context.Background(), time.Second*30)
	defer client.Disconnect(ctx)
	defer cf()

	collection := client.Database("go_db").Collection("Student")

	// 查询所有
	// findResults, err := collection.Find(ctx, bson.D{})

	// 条件查找
	findResults, err := collection.Find(ctx, bson.D{{"name", "abby"}})

	if err != nil {
		log.Fatal(err)
	}
	defer findResults.Close(ctx)

	for findResults.Next(ctx) {
		var result bson.D
		err := findResults.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result.Map()[\"name\"]: %v\n", result.Map()["name"])
	}
}

// 更新文档
func update() {
	initDB()
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	collection := client.Database("go_db").Collection("Student")
	update := bson.D{{"$set", bson.D{{"name", "11"}, {"age", 22}}}}

	ur, err := collection.UpdateMany(ctx, bson.D{{"name", "11"}}, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount)
}

// 删除文档
func delete() {
	initDB()

	ctx := context.TODO()
	defer client.Disconnect(ctx)

	collection := client.Database("go_db").Collection("Student")
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{"name", "11"}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleteResult.DeletedCount: %v\n", deleteResult.DeletedCount)
}

func main() {
	// initDB()

	// insertOne()
	// insertMulti()
	// find()
	// update()
	delete()
}
