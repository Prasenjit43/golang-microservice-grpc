package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Prasenjit43/golang-microservice-grpc/employee/add"
	"github.com/Prasenjit43/golang-microservice-grpc/employee/persistance"
	"github.com/Prasenjit43/golang-microservice-grpc/employee/update"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start of Employee microservices")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	group := router.Group("/api/employee")

	testFuncOutput := persistance.SetTest()
	fmt.Println("testFuncOutput : ", testFuncOutput)
	fmt.Println("testFuncOutput : ", testFuncOutput.Name)

	employeeDBInstance := persistance.DBInstance()

	fmt.Println("employeeDBInstance : ", employeeDBInstance)
	fmt.Println("employeeDBInstance : ", employeeDBInstance.Client)

	repo := getPersistanceObj()
	fmt.Println("repo : ", repo)

	registerAddRoutes(group, repo)
	registerUpdateRoutes(group, repo)

	router.Run(":" + port)

}

func getPersistanceObj() persistance.IEmployeeDBContext {
	return persistance.DBInstance()
}

func registerAddRoutes(group *gin.RouterGroup, repo persistance.IEmployeeDBContext) {
	service := add.InitService(repo)
	handler := add.InitHandler(service)
	router := add.InitRouter(handler)
	router.RegisterRoutes(group)
}

func registerUpdateRoutes(group *gin.RouterGroup, repo persistance.IEmployeeDBContext) {
	service := update.InitService(repo)
	handler := update.InitHandler(service)
	router := update.InitRouter(handler)
	router.RegisterRoutes(group)
}
