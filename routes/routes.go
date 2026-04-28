package routes

import (
	"go-rest-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.ShowPageIndex)
	api := r.Group("/api")
	{
		api.GET("/students", controllers.AllStudents)
		api.GET("/students/:id", controllers.GetById)
		api.POST("/students", controllers.CreateStudents)
		api.DELETE("/students/:id", controllers.DeleteStudents)
		api.PATCH("/students/:id", controllers.EditStudent)
		api.GET("/students/cpf/:cpf", controllers.GetByCPF)
	}
	r.NoRoute(controllers.RouteNotFound)
	r.Run(":3001")
}
