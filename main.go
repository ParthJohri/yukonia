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
			Title:       "Hugo Actually Explained (Websites, Themes, Layouts, and Intro to Scripting)",
			Thumbnail:   "https://i.ytimg.com/vi/ZFL09qhKi5I/hq720.jpg",
			PageURL:     "index.html",
			Description: "Hugo will allow you to create optimal website, but we shall overcome the \"Hugo Hump\" in this video. You gotta understand how templates and themes were before you get into the really impactful stuff.",
			Tags:        []string{"hugo"},
		},
		{
			Title:       "Based Cooking now using HUGO!",
			Thumbnail:   "https://i.ytimg.com/vi/jAXKSKb3etk/hqdefault.jpg",
			VideoURL:    "https://www.youtube.com/embed/jAXKSKb3etk",
			PageURL:     "index.html",
			Description: "(NOTE: I released this video on my PeerTube instance (https://videos.lukesmith.xyz) a while ago and realized I forgot to upload it to YouTube, so here it is, and still relevant).",
			Tags:        []string{"hugo"},
		},
		{
			Title:       "Simple Hugo Shortcodes absolutely MOG pathetic obese Wordpress!",
			Thumbnail:   "https://i.ytimg.com/vi/QTolhoxMyXg/hq720.jpg",
			VideoURL:    "https://www.youtube.com/embed/QTolhoxMyXg",
			PageURL:     "index.html",
			Description: "We cover Hugo shortcodes and the basics of using them and Hugo commands to create scriptable sections of your site. We cover range for for-loops, site variables and page variables, including titles, summaries, date formatting, word count, reading time, tags and much more.",
			Tags:        []string{"hugo"},
		},
	}

	videosTemplate := ""
	f, err := os.ReadFile("templates/video.html")
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

	err = ioutil.WriteFile(filepath.Join(buildDir, "videos.html"), videosHTML.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
