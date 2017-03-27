package main

import (
	"io"
	"net/http"
	"os"
)

type Downloader interface {
	Get(url, filename string) error
}

func NewDownloader() Downloader {
	return &Down{}
}

type Down struct {
}

func (d *Down) Get(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}
