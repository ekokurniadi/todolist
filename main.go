package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ekokurniadi/tokopedia-go-submittion/entity"
	"github.com/ekokurniadi/tokopedia-go-submittion/handler"
	"github.com/ekokurniadi/tokopedia-go-submittion/repository"
	"github.com/ekokurniadi/tokopedia-go-submittion/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// uncomment this part for running on local computer
	// env := godotenv.Load()
	// if env != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	host := os.Getenv("DB_HOST")
	userHost := os.Getenv("DB_USER")
	userPass := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")
	databasePort := os.Getenv("DB_PORT")

	//if you want to running on local computer, please change sslmode to disable, i using require for running on heroku
	dsn := "host=" + host + " user=" + userHost + " password=" + userPass + " dbname=" + databaseName + " port=" + databasePort + " sslmode=require TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&entity.Todo{})
	fmt.Println("Database Connected")
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("website/templates/**/*")
	router.Static("css", "./website/assets/css")
	router.Static("js", "./website/assets/js")
	router.Static("images", "./website/assets/images")
	api := router.Group("/api/v1")

	router.GET("/", todoHandler.Index)
	api.POST("/todos", todoHandler.CreateTodo)
	api.GET("/todos", todoHandler.GetTodos)
	api.GET("/todos/incomplete", todoHandler.GetTodosInComplete)
	api.GET("/todos/:id", todoHandler.GetTodo)
	api.PUT("/todos/:id", todoHandler.UpdateTodo)
	api.DELETE("/todos/:id", todoHandler.DeleteTodo)

	router.Run()
}
