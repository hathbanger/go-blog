package server

import (
	"fmt"
	"io"
	"os"
	"net/http"

	"github.com/go-blog/models"
	"github.com/labstack/echo"
)

func CreatePost(c echo.Context) error {
	// title := c.FormValue("title")
	// body := c.FormValue("body")
	
	fmt.Println(c.FormValue("title"))
	fmt.Println(c.FormValue("body"))

	// Read form fields
	title := c.FormValue("title")
	body := c.FormValue("body")
	//-----------
	// Read file
	//-----------
	// get current destination
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(pwd)
	

	fmt.Println(c.FormFile("file"))
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}


	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(pwd + "/photos/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}


	imgurl := "https://api.catsrassholes.com/photos/" + file.Filename
	// return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, title, body))



	post := models.NewPost(title, body, imgurl )
	err = post.Save()
	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! There's was an issue..")
	}

	return c.JSON(http.StatusOK, post)
}




func GetAllPosts(c echo.Context) error {
	username := c.Param("username")
	posts, err := models.GetAllPosts(username)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}
