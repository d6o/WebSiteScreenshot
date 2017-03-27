package main

import (
	"fmt"
	"strings"
	"time"
)

const ScreenShotURL = "https://api.thumbnail.ws/api/%s/thumbnail/get?url=%s&width=1280"

func (a *App) takeScreenShot() error {
	url := fmt.Sprintf(ScreenShotURL, a.Config.API(), a.Config.URL())
	filename := a.createFilename(a.Config.URL())

	down := NewDownloader()
	err := down.Get(url, filename)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) createFilename(url string) string {
	filename := strings.Replace(url, "/", "", -1)
	filename = strings.Replace(filename, ":", "", -1)
	filename = strings.Replace(filename, ".", "", -1)

	t := time.Now()
	filename += "_" + t.Format("2006_01_02_15_04_05") + ".jpg"

	return filename
}
