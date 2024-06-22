package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Video struct {
	Title       string   `json:"title"`
	Thumbnail   string   `json:"thumbnail"`
	VideoURL    string   `json:"video_url,omitempty"`
	Date        string   `json:"date"`
	PageURL     string   `json:"page_url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func main() {
	// Read the contents of the file
	data, err := ioutil.ReadFile("videos.json")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the JSON data into a slice of Video structs
	var videos []Video
	err = json.Unmarshal(data, &videos)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the HTML
	err = generateHTML(videos)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTML file created successfully.")
}

func generateHTML(videos []Video) error {
	// Read the template file
	t, err := template.ParseFiles("templates/videos.template")
	if err != nil {
		return err
	}

	// Create the build directory if it doesn't exist
	buildDir := filepath.Join("build")
	err = os.MkdirAll(buildDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Execute the template and write the HTML to a file
	outputFile := filepath.Join(buildDir, "index.html")
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, videos)
	if err != nil {
		return err
	}

	return nil
}
