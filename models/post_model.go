package models

import (
	"time"
	// "fmt"

	"labix.org/v2/mgo/bson"
	"github.com/go-blog/store"
)

type Post struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Title	string           `json:"title",bson:"title,omitempty"`
	Body	string           `json:"body",bson:"body,omitempty"`
	Imgurl	string           `json:"imgurl",bson:"imgurl,omitempty"`
}

func NewPost(title string, body string, imgurl string) *Post {
	p := new(Post)
	p.Id = bson.NewObjectId()
	p.Title = title
	p.Body = body
	p.Imgurl = imgurl

	return p
}

func (p *Post) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	// collection, err := store.ConnectToPostCollection(session, "posts")
	// if err != nil {
	// 	panic(err)
	// }

	collection := session.DB("test").C("posts")

	post := Post {
		Id:		p.Id,
		Timestamp:	p.Timestamp,
		Title:	p.Title,
		Body: 	p.Body,
		Imgurl: p.Imgurl,
	}


	err = collection.Insert(post)
	if err != nil {
		return err
	}

	return nil
}



func GetAllPosts(Title string) ([]*Post, error){

	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection := session.DB("test").C("posts")

	posts := []*Post{}

	err = collection.Find(bson.M{}).All(&posts)

	return posts, err
}