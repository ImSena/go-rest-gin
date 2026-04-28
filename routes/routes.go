package routes

import (
	"go-rest-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/students", controllers.AllStudents)
		api.GET("/students/:id", controllers.GetById)
		api.POST("/students", controllers.CreateStudents)
		api.DELETE("/students/:id", controllers.DeleteStudents)
		api.PATCH("/students/:id", controllers.EditStudent)
		api.GET("/students/cpf/:cpf", controllers.GetByCPF)
	}
	r.Run(":3001")
}
