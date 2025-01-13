package main

/*
TODO:
  - custom screenshots folder path
  - error handling and logging
  - handling of >9 digit prices (rare)
*/

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/oliamb/cutter"
	"github.com/otiai10/gosseract/v2"
	"github.com/sergeymakinen/go-bmp"
)

const (
	shops_per_page = 8
)

type store struct {
	id       string
	name     string
	quantity int
	price    int
}

type field struct {
	xoffset int
	width   int
}

func FindNewestScreenshots(limit time.Duration) []string {
	home, _ := os.UserHomeDir()
	screenshot_dir := home + "/MapleLegends/Screenshots/"

	files, _ := os.ReadDir(screenshot_dir)
	newest_time := time.Time{}

	for _, file := range files {
		info, _ := file.Info()
		modtime := info.ModTime()
		if modtime.After(newest_time) {
			newest_time = modtime
		}
	}

	newest_screenshots := make([]string, 0)
	for _, file := range files {
		info, _ := file.Info()
		modtime := info.ModTime()
		if newest_time.Sub(modtime) <= limit {
			newest_screenshots = append(newest_screenshots, screenshot_dir+file.Name())
		}
	}

	return newest_screenshots
}

func ExtractString(client *gosseract.Client, img image.Image, field field, row int) string {
	croppedImg, _ := cutter.Crop(img, cutter.Config{
		Width:  field.width,
		Height: 22,
		Anchor: image.Point{field.xoffset, 258 + row*22},
	})

	buffer := new(bytes.Buffer)
	png.Encode(buffer, croppedImg)

	client.SetImageFromBytes(buffer.Bytes())
	text, _ := client.Text()
	return text
}

func SanatizeID(s string) string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	s = regex.ReplaceAllString(s, "${1}")
	return s
}

func CompareStore(a store, b store) bool {
	return a.price < b.price
}

func FormatCommas(num int) string {
	regex := regexp.MustCompile(`(\d+)(\d{3})`)

	str := strconv.Itoa(num)
	for n := ""; n != str; {
		n = str
		str = regex.ReplaceAllString(str, "$1,$2")
	}

	return str
}

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("eng")

	newest_screenshots := FindNewestScreenshots(time.Minute)

	stores := make([]store, 0)
	errors := 0
	for _, filepath := range newest_screenshots {
		file, _ := os.Open(filepath)
		img, _ := bmp.Decode(file)

		for i := 0; i < shops_per_page; i++ {
			var err error

			store := store{}

			client.SetWhitelist("")
			store.id = SanatizeID(ExtractString(client, img, field{192, 70}, i))
			store.name = strings.TrimSpace(ExtractString(client, img, field{262, 70}, i))

			client.SetWhitelist("0123456789")
			store.quantity, err = strconv.Atoi(ExtractString(client, img, field{332, 30}, i))
			store.price, err = strconv.Atoi(ExtractString(client, img, field{361, 75}, i))

			if store.id == "" && store.name == "" && store.price == 0 && store.quantity == 0 {
				continue
			}
			if err != nil {
				errors += 1
				continue
			}

			stores = append(stores, store)
		}
	}

	sort.Slice(stores, func(i int, j int) bool {
		return CompareStore(stores[i], stores[j])
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Store Name", "Price", "Quantity"})
	for _, store := range stores {
		table.Append([]string{store.id, store.name, FormatCommas(store.price), FormatCommas(store.quantity)})
	}
	table.Render()

	if errors != 0 {
		fmt.Printf("Couldn't parse %d shops\n", errors)
	}
}
