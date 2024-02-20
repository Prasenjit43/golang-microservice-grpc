package add

import (
	"fmt"
	"net/http"

	"github.com/Prasenjit43/golang-microservice-grpc/employee/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func InitHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp model.EmployeeModel
		err := c.BindJSON(&emp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Employee & in handler :", emp)
		fmt.Println("Employee & in handler :", &emp)

		insertionNum, insertErr := h.service.Add(&emp)
		if insertErr != nil {
			c.JSON(insertErr.ErrorStatus, gin.H{
				"error": insertErr.ErrorMessage,
			})
			return
		}
		c.JSON(http.StatusOK, insertionNum)
	}
}
