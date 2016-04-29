package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

var (
	rootDir = "build"
)

func main() {

	dir = &rootDir

	/*
	 * modify the index.html to use the url prefix (basename) provided by environment variable
	 */

	//1. Get the environment variable
	baseUrl := os.Getenv("baseUrl")
	log.Println("baseUrl = ", baseUrl)
	if baseUrl != "" {

		indexPath := rootDir + "/index.html"
		oldIndexContent, err := ioutil.ReadFile(indexPath)
		if err != nil {
			log.Fatalf("unable to find index.html :%v", err)
		}

		newIndexContent := bytes.Replace(oldIndexContent, []byte("window.baseUrl=\"/\""), []byte("window.baseUrl=\"/"+baseUrl+"\""), -1)
		err = ioutil.WriteFile(indexPath, newIndexContent, 0644)
	}
	/**
	 * Start the HTTP server
	 */
	serveFiles()

}
