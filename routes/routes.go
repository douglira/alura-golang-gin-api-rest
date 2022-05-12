package routes

import (
	"github.com/douglira/alura-golang-gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine) {
	apiRouter := r.Group("/api")
	apiRouter.GET("/students/:id", controllers.FindStudentById)
	apiRouter.GET("/students/search/identity-number", controllers.FindStudentByIdentityNumber)
	apiRouter.GET("/students", controllers.AllStudents)
	apiRouter.POST("/students", controllers.RegisterStudent)
	apiRouter.DELETE("/students/:id", controllers.DeleteStudent)
	apiRouter.PUT("/students/:id", controllers.UpdateStudent)
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	router(r)
	return r
}
