package main

import (
	"net/http"

	"github.com/SanketPatil6869/Task_Management_Gin1/config"
	"github.com/SanketPatil6869/Task_Management_Gin1/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	//mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     w.Header().Set("Content-Type", "application/json")
	//     w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	// handler := cors.Default().Handler(mux)
	// http.ListenAndServe(":8080", handler)

	config.ConnectDatabase()

	/* routes */
	router.POST("/create", controller.CreateTask)
	router.GET("/tasks", controller.FindTasks)
	router.GET("/task/:id", controller.FindTask)
	router.PUT("/updateTask/:id", controller.UpdateTask)
	router.DELETE("/deleteTask/:id", controller.DeleteTask)
	router.GET("/taskStatus/:status", controller.FindTaskStatus)
	//router.GET("/GetTask2", controller.FindTasks_2)
	defer router.Run()
}
