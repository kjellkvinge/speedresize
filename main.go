package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/nfnt/resize"
	fjpeg "github.com/pixiv/go-libjpeg/jpeg"
)

func init() {
	vips.LoggingSettings(func(domain string, level vips.LogLevel, msg string) {
		fmt.Println(domain, level, msg)
	}, vips.LogLevelError)
	vips.Startup(nil)
}
func main() {
	generateThumbnailVips("./images/IMG_0901.JPG", "/tmp/thumbnail")
}

func generateThumbnailVips(original, thumbnail string) {

	//	defer vips.Shutdown()
	img, err := vips.NewThumbnailFromFile(original, 300, 300, vips.InterestingNone)
	if err != nil {
		fmt.Println(err)
		return
	}

	image1bytes, _, err := img.ExportNative()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(thumbnail, image1bytes, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func generateThumbnailTurbo(original, thumbnail string) {
	reader, err := os.Open(original)
	if err != nil {
		log.Println(err)
		return
	}
	defer reader.Close()
	originalImg, err := fjpeg.Decode(reader, &fjpeg.DecoderOptions{
		ScaleTarget: image.Rect(0, 0, 300, 300),
	},
	)
	if err != nil {
		log.Printf("error decoding image:%s", err)
		return
	}

	var outFile *os.File
	outFile, err = os.Create(thumbnail)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)

	err = fjpeg.Encode(b, originalImg, &fjpeg.EncoderOptions{Quality: 90})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func generateThumbnailNfnt(original, thumbnail string) {

	reader, err := os.Open(original)
	if err != nil {
		log.Println(err)
		return
	}
	defer reader.Close()

	originalImg, _, err := image.Decode(reader)
	if err != nil {
		//TODO: check file suffix for supported fileformats (jpg, png++)
		log.Printf("error decoding image:%s", err)
		return
	}

	// resize it
	resizedImg := resize.Resize(300, 0, originalImg, resize.Lanczos3)

	var outFile *os.File
	outFile, err = os.Create(thumbnail)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)

	err = jpeg.Encode(b, resizedImg, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
