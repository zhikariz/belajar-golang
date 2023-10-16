package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	currentDirectory, _ = os.Getwd()
	tempPath            = filepath.Join(currentDirectory, "pipeline-pattern-fifo/pipeline-temp")
	totalFile           = 3000
	contentLength       = 5000
)

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

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

	// pipeline 1: job distribution
	chanFileIndex := generateFileIndexes()

	// pipeline2 : the main login (creating files)
	createFilesWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFilesWorker)

	// track and print output
	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}

	log.Printf("%d/%d files created", counterSuccess, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// wait group to control the workers
	wg := new(sync.WaitGroup)

	// allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		// dispatch N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				// listen to `chanIn` channel for incoming jobs
				for job := range chanIn {
					// do the jobs
					filePath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := os.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

					// construct the job's result, and send it to `chanOut`
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}

				// if `chanIn` is closed, and the remaining jobs are finished,
				// only then we mark the worker as complete.
				wg.Done()
			}(workerIndex)
		}
	}()

	// wait until `chanIn` closed and then all workers are done,
	// because right after that - we need to close the `chanOut` channel.
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

// 0.4503687
