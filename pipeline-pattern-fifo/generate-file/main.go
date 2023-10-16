package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var (
	currentDirectory, _ = os.Getwd()
	tempPath            = filepath.Join(currentDirectory, "pipeline-pattern-fifo/pipeline-temp")
	totalFile           = 3000
	contentLength       = 5000
)

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(contentLength)
		err := os.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", filename)
		}
		log.Println(i, "files created")
	}
	log.Printf("%d of total files created", totalFile)
}

// 0.8349291
