package add

import (
	"fmt"

	"github.com/Prasenjit43/golang-microservice-grpc/employee/model"
	"github.com/Prasenjit43/golang-microservice-grpc/employee/persistance"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	repository persistance.IEmployeeDBContext
}

func InitService(repo persistance.IEmployeeDBContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Add(emp *model.EmployeeModel) (*mongo.InsertOneResult, *model.ErrorDetail) {
	fmt.Println("Employee  in service :", emp)
	fmt.Println("Employee & in service :", &emp)
	fmt.Println("Employee * in service :", *emp)

	insertResult, errDetails := s.repository.AddEmployee(emp)
	return insertResult, errDetails
}
