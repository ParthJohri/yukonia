package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Video struct {
	Title       string
	Thumbnail   string
	VideoURL    string
	PageURL     string
	Description string
	Tags        []string
}

func main() {
	videos := []Video{
		{
			Title:       "Me at the zoo",
			Thumbnail:   "../thumbnails/jawed.jpeg",
			VideoURL:    "https://www.youtube.com/embed/jNQXAC9IVRw",
			PageURL:     "index.html",
			Description: "Me talking about elephants.",
			Tags:        []string{"jawed", "youtube", "elephants"},
		},
		{
			Title:       "Gardening in the Sun: How to Grow Beautiful Flowers",
			Thumbnail:   "../thumbnails/jawed.jpeg",
			VideoURL:    "https://www.youtube.com/embed/jNQXAC9IVRw",
			PageURL:     "index.html",
			Description: "Learn how to create a beautiful garden with these easy-to-follow steps.",
			Tags:        []string{"gardening", "flowers", "sun"},
		},
		{
			Title:       "The Secrets to Growing Healthy Herbs: Tips and Tricks",
			Thumbnail:   "../thumbnails/jawed.jpeg",
			VideoURL:    "https://www.youtube.com/embed/jNQXAC9IVRw",
			PageURL:     "index.html",
			Description: "Discover the best ways to grow healthy herbs for your garden.",
			Tags:        []string{"gardening", "herbs", "health"},
		},
		{
			Title:       "The Art of Gardening: Tips for Creating a Stunning Garden Bed",
			Thumbnail:   "../thumbnails/jawed.jpeg",
			VideoURL:    "https://www.youtube.com/embed/jNQXAC9IVRw",
			PageURL:     "index.html",
			Description: "Learn how to create a beautiful garden bed with these easy-to-follow steps.",
			Tags:        []string{"gardening", "garden bed", "art"},
		},
		{
			Title:       "The Science of Gardening: How Plants Grow and Thrive",
			Thumbnail:   "../thumbnails/jawed.jpeg",
			VideoURL:    "https://www.youtube.com/embed/jNQXAC9IVRw",
			PageURL:     "index.html",
			Description: "Discover the science behind how plants grow and thrive in your garden.",
			Tags:        []string{"gardening", "science", "plants"},
		},
		{
			Title:       "The Magic of Gardening: How to Create a Whimsical Garden",
			Thumbnail:   "../thumbnails/jawed.jpeg",
			VideoURL:    "https://www.youtube.com/embed/jNQXAC9IVRw",
			PageURL:     "/",
			Description: "Learn how to create a whimsical garden with these easy-to-follow steps.",
			Tags:        []string{"gardening", "whimsical", "magic"},
		},
	}
	videosTemplate := ""
	f, err := os.ReadFile("video.templ")
	if err != nil {
		log.Fatal(err)
	}
	videosTemplate = string(f)

	buildDir := filepath.Join("build")
	err = os.MkdirAll(buildDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = writeHTML(videosTemplate, videos, buildDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTML file created successfully.")
}

func writeHTML(templateString, videos interface{}, buildDir string) error {
	t, err := template.New("videos").Parse(templateString.(string))
	if err != nil {
		return err
	}

	videosHTML := new(bytes.Buffer)
	err = t.Execute(videosHTML, videos)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(buildDir, "index.html"), videosHTML.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
