package main

import (
	"backend/controllers"
	"backend/initializers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
	//helpers.SeedDB()

}

func main() {

	r := gin.Default()
	r.GET("/campground", middleware.CORSMiddleware, controllers.AllCampgrounds)
	r.GET("/campground/:id", middleware.CORSMiddleware, controllers.CampgroundByID)
	r.POST("/campground", controllers.CreateCampground)
	r.PUT("/campground/:id", controllers.UpdateCampground)
	r.Run()
}
