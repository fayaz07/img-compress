package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	smFolder = "small"
	mdFolder = "medium"
	lgFolder = "large"
	thFolder = "thumbnail"
)

var smallSaveLoc string
var mdSaveLoc string
var lgSaveLoc string
var thSaveLoc string

var originalLoc string

func main() {
	root := "/home/fayaz/Pictures/edit"
	// compress(root)
	printJSON(root)
}

func compress(root string) {
	originalLoc = root
	// replacePath := root
	smallSaveLoc = path.Join(root, smFolder)
	mdSaveLoc = path.Join(root, mdFolder)
	lgSaveLoc = path.Join(root, lgFolder)
	thSaveLoc = path.Join(root, thFolder)

	// if runtime.GOOS == "windows" {
	// 	replacePath = flipSlashes(replacePath)
	// 	// fmt.Println(replacePath)
	// }

	fmt.Println("Creating directories required for saving compressed images")
	err := os.Mkdir(smallSaveLoc, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(mdSaveLoc, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(lgSaveLoc, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(thSaveLoc, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Walkthrough the folder and compress images
	var files []string
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info.Name())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Started compressing images without loosing quality")

	for index, file := range files {
		fmt.Println("Processing: " + file)
		resizeThisPhoto(file)
		fmt.Printf("Processed %d of %d\n\n", index+1, len(files))
	}

	fmt.Println("Done")
}

func flipSlashes(loc string) string {
	return strings.ReplaceAll(loc, "/", "\\")
}

func resizeThisPhoto(filename string) {
	lastIndexOfDot := strings.LastIndex(filename, ".")
	filenameWithoutExt := filename[0:lastIndexOfDot]
	ext := filename[lastIndexOfDot+1:]
	// fmt.Println(ext)
	// fmt.Println(filenameWithoutExt)

	img, err := imaging.Open(path.Join(originalLoc, filename))
	if err != nil {
		// panic(err)
		fmt.Println("Error processing the file: " + filename)
		return
	}

	large := imaging.Resize(img, largeSize(img.Bounds().Size().X), largeSize(img.Bounds().Size().Y), imaging.Lanczos)
	imaging.Save(large, path.Join(lgSaveLoc, fmt.Sprintf("%s_%s.%s", filenameWithoutExt, "lg", ext)))

	medium := imaging.Resize(img, mediumSize(img.Bounds().Size().X), mediumSize(img.Bounds().Size().Y), imaging.Lanczos)
	imaging.Save(medium, path.Join(mdSaveLoc, fmt.Sprintf("%s_%s.%s", filenameWithoutExt, "md", ext)))

	small := imaging.Resize(img, smallSize(img.Bounds().Size().X), smallSize(img.Bounds().Size().Y), imaging.Lanczos)
	imaging.Save(small, path.Join(smallSaveLoc, fmt.Sprintf("%s_%s.%s", filenameWithoutExt, "sm", ext)))

	x, y := thumbnail(img.Bounds().Size().X, img.Bounds().Size().Y)

	th := imaging.Resize(img, x, y, imaging.Lanczos)
	imaging.Save(th, path.Join(thSaveLoc, fmt.Sprintf("%s_%s.%s", filenameWithoutExt, "th", ext)))

}
