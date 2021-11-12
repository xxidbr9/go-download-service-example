package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	fileName := ""
	if len(args) == 1 {
		fileName = "sample"
	} else {
		fileName = args[1]
	}
	URL := "http://192.168.10.113:3000/_next/image?url=https%3A%2F%2Fstorage.googleapis.com%2Fellaskincare-backend-production%2Fcatalog%2Fwhats_on_article%2Fc5babta23akg0093cnkg_WhatsApp%20Image%202021-09-29%20at%2015.45.53.jpeg_1633068533&w=1920&q=75"
	downloadedFile, ext, err := downloadFile(URL)

	if err != nil {
		log.Fatal(err)
	}

	err = saveToLocal(fileName, ext, downloadedFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %s downlaod in current working directory", fileName)
}

func downloadFile(URL string) (savedFile io.ReadCloser, ext string, err error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return nil, "", err
	}

	if response.StatusCode != 200 {
		return nil, "", errors.New("received non 200 response code")
	}

	ext = response.Header.Get("Content-Type")
	ext2, err := mime.ExtensionsByType(ext)
	ext = ext2[1]
	if err != nil {
		return nil, "", err
	}

	return response.Body, ext, nil
}

func saveToLocal(fileName, fileExt string, file io.ReadCloser) error {
	//Create a empty file

	savedFile, err := os.Create(fmt.Sprintf("%s%s", fileName, fileExt))
	if err != nil {
		return err
	}
	defer file.Close()
	
	_, err = io.Copy(savedFile, file)
	if err != nil {
		return err
	}

	return nil
}
