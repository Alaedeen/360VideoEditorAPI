package main

import (
	models "github.com/Alaedeen/360VideoEditorAPI/models"
	handlers "github.com/Alaedeen/360VideoEditorAPI/handlers" 
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/Alaedeen/360VideoEditorAPI/repository"
	"github.com/spf13/viper"
	"github.com/Alaedeen/360VideoEditorAPI/config"
	router "github.com/Alaedeen/360VideoEditorAPI/router"
	"github.com/rs/cors"
)





func main()  {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost", "http://localhost:8082"},
		AllowCredentials: true,
		AllowedHeaders: []string{"Authorization","Content-Type"},
		AllowedMethods: []string{"GET","POST","DELETE","PUT","HEAD"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	//connect to the data base
	UserName :=  configuration.Database.UserName
	Password :=  configuration.Database.Password
	DataBase :=  configuration.Database.DataBase
	Charset :=  configuration.Database.Charset
	ParseTime :=  configuration.Database.ParseTime
	db, err := gorm.Open("mysql", UserName+":"+Password+"@/"+DataBase+"?charset=" +
										Charset+"&parseTime="+ParseTime+"&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	
	defer db.Close() 
	db.AutoMigrate(&models.AddedShapes{},&models.AddedTags{},
					&models.Comment{},&models.CommentsDislikes{},
					&models.CommentsLikes{},&models.Font{},
					&models.Picture{},&models.Project{},
					&models.RepliesDislikes{},&models.RepliesLikes{},
					&models.Reply{},&models.Shape{},
					&models.Subscriptions{},&models.TagElements{},
					&models.User{},&models.Video{},
					&models.Video2D{},&models.VideosDislikes{},
					&models.VideosLikes{},&models.Test{})

	userRepo := repository.UserRepo{db}
	userHandler := handlers.UserHandler{&userRepo}
	videoRepo := repository.VideoRepo{db}
	videoHandler := handlers.VideoHandler{&videoRepo}
	projectRepo := repository.ProjectRepo{db}
	projectHandler := handlers.ProjectHandler{&projectRepo}
	// Init Router 
	r := mux.NewRouter()
	// serve static files
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",http.FileServer(http.Dir("./assets/"))))
	UserRouterHandler := router.UserRouterHandler{Router: r,Handler: userHandler}
	UserRouterHandler.HandleFunctions()
	VideoRouterHandler := router.VideoRouterHandler{Router: r,Handler: videoHandler}
	VideoRouterHandler.HandleFunctions()
	ProjectRouterHandler := router.ProjectRouterHandler{Router: r,Handler: projectHandler}
	ProjectRouterHandler.HandleFunctions()
	// start server
	port := ":" + strconv.Itoa(configuration.Server.Port) 
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(port,handler))
}