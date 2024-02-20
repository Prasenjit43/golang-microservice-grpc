package update

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

func (h *Handler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateEmp model.EmployeeModel
		err := c.BindJSON(&updateEmp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Update Data :", updateEmp)

		updateResult, updateErr := h.service.Update(&updateEmp)
		if updateErr != nil {
			c.JSON(updateErr.ErrorStatus, gin.H{
				"error": updateErr.ErrorMessage,
			})
			return
		}

		c.JSON(http.StatusOK, updateResult)
	}
}
