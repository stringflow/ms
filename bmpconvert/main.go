package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
)

func convertFile(path string) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return
	}

	if contents[0] != 'B' || contents[1] != 'M' {
		return
	}

	png := path
	bmp := strings.TrimSuffix(path, filepath.Ext(path)) + ".bmp"

	err = os.Rename(png, bmp)
	if err != nil {
		return
	}

	cmd := exec.Command("ffmpeg", "-y", "-i", bmp, png)
	cmd.Run()
	fmt.Println("Converted", png)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s (path to screenshots folder)\n", os.Args[0])
		return
	}

	path := os.Args[1]

	r := regexp.MustCompile(".png$")

	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	err := w.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event := <-w.Event:
				convertFile(event.Path)
			case err := <-w.Error:
				log.Fatal(err)
			case <-w.Closed:
				return
			}
		}
	}()

	fmt.Println("Watching", path)
	err = w.Start(time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
