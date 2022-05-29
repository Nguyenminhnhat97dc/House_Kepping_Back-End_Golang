package main

import (
	"BE_Golang/BE_Golang/controllers"
	//"log"
	"net/http"
	//"time"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/websocket"
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
/* var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket returns json format
func jsonApi(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	var data struct {
		A string `json:"Id"`
		B int    `json:"b"`
	}
	//Read data in ws
	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}

	//Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 1000 {
			break
		}

		err = ws.WriteJSON(struct {
			A string `json:"a"`
			B int    `json:"b"`
			C int    `json:"c"`
		}{
			A: data.A,
			B: data.B,
			C: count,
		})
		if err != nil {
			log.Println("error write json: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
*/

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.New()
	//r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin, Authorization, Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
	}))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Data": "hello world"})
	})
	//r.GET("/provider", controllers.FindProvider)
	r.POST("/provider/id", controllers.FindProviderID)
	r.GET("/services", controllers.FindServices)
	r.GET("/services/:count", controllers.LimitServices)
	r.POST("/requirement", controllers.AddRequirementCustomer)
	r.GET("/provider/services", controllers.ServiceProvider)
	r.GET("/servicesofprovider", controllers.AddServiceProvider)
	r.GET("/requirementcustomer", controllers.RequirementsCustomer)
	r.GET("/todolist", controllers.TodoList)
	r.GET("history", controllers.HistoryList)
	r.POST("/loggin", controllers.Loggin)
	r.POST("/priceservices", controllers.FindPriceOfServices)
	r.POST("/addprice", controllers.AddPrice)
	r.POST("/addtodolist", controllers.AddTodoList)
	r.POST("/paginationrequirement", controllers.CountPaginationRequirement)
	r.POST("/paginationtodolist", controllers.CountPaginationToDoList)
	r.POST("/update_todolist", controllers.UpdateTodoList)
	r.POST("/deleteservices", controllers.DeleteServicesProvider)
	r.POST("/updateprovider", controllers.UpdateInformationProvider)

	r.Run(":" + port)
}
