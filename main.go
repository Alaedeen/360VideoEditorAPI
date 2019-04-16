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
)





func main()  {
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
					&models.VideosLikes{})

	userRepo := repository.UserRepo{db}
	userHandler := handlers.UserHandler{&userRepo}
	videoRepo := repository.VideoRepo{db}
	videoHandler := handlers.VideoHandler{&videoRepo}
	// Init Router 
	r := mux.NewRouter()
	UserRouterHandler := router.UserRouterHandler{Router: r,Handler: userHandler}
	UserRouterHandler.HandleFunctions()
	VideoRouterHandler := router.VideoRouterHandler{Router: r,Handler: videoHandler}
	VideoRouterHandler.HandleFunctions()
	// start server
	port := ":" + strconv.Itoa(configuration.Server.Port) 
	log.Fatal(http.ListenAndServe(port,r))
}