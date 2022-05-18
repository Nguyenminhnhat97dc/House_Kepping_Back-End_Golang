package main

import (
	"BE_Golang/BE_Golang/controllers"
	"net/http"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/* func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.GET("/story/:id", controllers.Read)
				client.POST("/story/create", controllers.Create)
		   		client.PATCH("/story/update/:id", controllers.Update)
		   		client.DELETE("/story/:id", controllers.Delete)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
} */

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin, Authorization, Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 50 * time.Minute,
	}))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Data": "hello world"})
	})
	r.GET("/provider", controllers.FindProvider)
	r.POST("/provider/id", controllers.FindProviderID)
	r.GET("/services", controllers.FindServices)
	r.GET("/services/:count", controllers.LimitServices)
	r.POST("/requirement", controllers.UpdateRequirementCustomer)
	r.GET("/provider/services", controllers.ServiceProvider)
	r.POST("/servicesofprovider", controllers.AddServiceProvider)
	r.GET("/requirementcustomer", controllers.RequirementsCustomer)
	r.POST("/todolist", controllers.TodoList)
	r.POST("/loggin", controllers.Loggin)
	r.Run(":8080")
}