// Imgur images handler.
// Author:
//    Andrea Cervesato <andrea.cervesato@mailbox.org>

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Imgur subreddit image.
type SubredditImage struct {
	ID           string
	Title        string
	Description  string
	Datetime     int
	Type         string
	Animated     bool
	Width        int
	Height       int
	Size         int
	Views        int
	Bandwidth    int
	Vote         int
	Favorite     bool
	NSFW         bool
	Section      string
	AccountUrl   string
	AccountId    string
	IsAd         bool
	Tags         []string
	InMostViral  bool
	InGallery    bool
	Link         string
	CommentCount int
	Ups          int
	Downs        int
	Points       int
	Score        int
	IsAlbum      bool
}

// Imgur subreddit gallery reply.
type SubredditGallery struct {
	Data    []SubredditImage
	Success bool
	Status  int
}

// Imgur communication class.
type Imgur struct {
	UrlApi   string
	ClientID string
}

// Create a new Imgur communication object.
func NewImgur(client_id string) *Imgur {
	obj := Imgur{
		UrlApi:   "https://api.imgur.com/3",
		ClientID: client_id,
	}

	return &obj
}

// Make a GET request to imgur and return body response as bytes.
func (obj *Imgur) Request(url string) []byte {
	// make a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// use no-password authorization
	req.Header.Set("Authorization", "Client-ID "+obj.ClientID)
	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// close at return
	defer resp.Body.Close()

	// read response data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := []byte(body)
	return data
}

// Select a random image from a subreddit gallery.
func (obj *Imgur) RandSubredditImage(gallery string) string {
	// fetch gallery data
	url := fmt.Sprintf("%s/gallery/r/%s/time/week/0", obj.UrlApi, gallery)
	data := obj.Request(url)

	// str -> JSON
	var result SubredditGallery
	json.Unmarshal(data, &result)

	// pick a random image
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(len(result.Data))
	image := result.Data[value]
	image_url := image.Link

	return image_url
}
