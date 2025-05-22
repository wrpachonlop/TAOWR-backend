package routes

import (
	"github.com/wrpachonlop/tool-inventory-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/employees", handlers.GetEmployees)
		api.POST("/employees", handlers.CreateEmployee)
		api.GET("/employees/:id", handlers.GetEmployeeByID)

		api.GET("/tools", handlers.GetTools)
		api.POST("/tools", handlers.CreateTool)
	}
}
