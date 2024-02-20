package persistance

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Prasenjit43/golang-microservice-grpc/employee/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmployeeMongoDB struct {
	Client *mongo.Client
}

type Test struct {
	Name string
}

func SetTest() Test {
	return Test{
		Name: "Monu",
	}
}

func DBInstance() EmployeeMongoDB {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongoadmin:password@localhost:27018"))

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected To Mongodb")

	return EmployeeMongoDB{
		Client: client,
	}

}

func openConnection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	userCollection := client.Database(databaseName).Collection(collectionName)
	return userCollection
}

func (e EmployeeMongoDB) AddEmployee(emp *model.EmployeeModel) (*mongo.InsertOneResult, *model.ErrorDetail) {

	fmt.Println("Employee  in database :", emp)
	fmt.Println("Employee & in database :", &emp)
	fmt.Println("Employee * in database :", *emp)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userCollection := openConnection(e.Client, "empdata", "employee")

	emp.ID = primitive.NewObjectID()
	emp.UserId = emp.ID.Hex()

	ressultInsertionum, insertErr := userCollection.InsertOne(ctx, emp)
	if insertErr != nil {
		return nil, &model.ErrorDetail{
			ErrorStatus:  http.StatusInternalServerError,
			ErrorMessage: insertErr.Error(),
		}
	}

	fmt.Println("ressultInsertionum  in AddEmployee : ", ressultInsertionum)
	return ressultInsertionum, nil
}

func (e EmployeeMongoDB) UpdateEmployee(emp *model.EmployeeModel) (*mongo.UpdateResult, *model.ErrorDetail) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userConnection := openConnection(e.Client, "empdata", "employee")
	filter := bson.D{{"userid", emp.UserId}}

	update := bson.D{{
		"$set", bson.D{
			{"department", emp.Department},
		},
	}}

	updateResult, updateErr := userConnection.UpdateOne(ctx, filter, update)

	if updateErr != nil {
		return nil, &model.ErrorDetail{
			ErrorStatus:  http.StatusBadRequest,
			ErrorMessage: updateErr.Error(),
		}
	}
	return updateResult, nil
}
