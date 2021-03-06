package server

import (
	"fmt"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
)



func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

// Restricted Access
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://catsrassholes.com", "https://www.catsrassholes.com"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))



// ROUTES
	e.Static("/photos", "photos")	
	e.GET("/", accessible)
	r.GET("", restricted)
	e.GET("/user/:username", GetUser)
	e.GET("/user/id/:user_id", GetUserById)
	e.POST("/user", CreateUser)
	e.GET("/users", GetAllUsers)
	e.POST("/post/create", CreatePost)
	e.POST("/post/delete/:post_id", RemovePost)
	// e.POST("/post/like/:post_id", LikePost)
	e.GET("/post/all", GetAllPosts)
	e.POST("/login", Login)

	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}

