package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	lg = "_lg"
	md = "_md"
	sm = "_sm"
	th = "_th"
)

type Image struct {
	Original  string `json:"orig"`
	Large     string `json:"lg"`
	Medium    string `json:"md"`
	Small     string `json:"sm"`
	Thumbnail string `json:"th"`
}

func printJSON(root string) {
	// Walkthrough the folder and compress images
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info.Name())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	var data = map[string][]Image{}
	var page = 1

	// var paginatedArray = [][]Image{}
	var images = []Image{}

	for _, file := range files {
		if strings.Contains(file, lg) ||
			strings.Contains(file, md) ||
			strings.Contains(file, sm) ||
			strings.Contains(file, th) {
		} else {
			images = append(images, Image{
				Original:  file,
				Large:     imageNameWithSize(file, lg),
				Medium:    imageNameWithSize(file, md),
				Small:     imageNameWithSize(file, sm),
				Thumbnail: imageNameWithSize(file, th),
			})
		}

		if len(images) == 20 {
			key := strconv.Itoa(page)

			data[key] = images
			images = []Image{}
			page++
		}
	}

	if len(images) > 0 {
		key := strconv.Itoa(page)
		data[key] = images
	}

	// Convert the array to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON data to a file
	err = ioutil.WriteFile("images.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func imageNameWithSize(filename string, size string) string {
	lastIndexOfDot := strings.LastIndex(filename, ".")
	filenameWithoutExt := filename[0:lastIndexOfDot]
	ext := filename[lastIndexOfDot+1:]

	return fmt.Sprintf("%s%s.%s", filenameWithoutExt, size, ext)
}
