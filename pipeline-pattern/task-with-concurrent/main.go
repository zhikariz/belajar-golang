package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	currentDirectory, _ = os.Getwd()
	tempPath            = filepath.Join(currentDirectory, "pipeline-pattern/pipeline-temp")
)

type FileInfo struct {
	FilePath  string // file location
	Content   []byte // file content
	Sum       string // md5 sum of content
	IsRenamed bool   // indicator whether the particular file is renamed already or not
}

func main() {
	log.Println("start")
	start := time.Now()

	// pipeline 1: loop all files and read it
	chanFileContent := readFiles()

	// pipeline 2: calculate md5sum
	numberOfWorkers := 100
	chanFileSum := getSum(chanFileContent, numberOfWorkers)

	// pipeline 3: rename files
	chanRename := rename(chanFileSum, numberOfWorkers)

	// print output
	counterRenamed := 0
	counterTotal := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	var wg = new(sync.WaitGroup)

	wg.Add(100)

	go func() {
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {

			// if there is an error, return immediatelly
			if err != nil {
				return err
			}

			// if it is a sub directory, return immediatelly
			if info.IsDir() {
				return nil
			}

			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}

			return nil
		})
		if err != nil {
			log.Println("ERROR:", err.Error())
		}

	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func getSum(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	wg := new(sync.WaitGroup)

	wg.Add(numberOfWorkers)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				for fileInfo := range chanIn {
					fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
					chanOut <- fileInfo
				}
				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func rename(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	wg := new(sync.WaitGroup)

	wg.Add(numberOfWorkers)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				for fileInfo := range chanIn {
					destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
					err := os.Rename(fileInfo.FilePath, destinationPath)
					fileInfo.IsRenamed = err == nil
					chanOut <- fileInfo
				}
				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan FileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
