package main

import (
	"testing"
	"time"
)

func BenchmarkVips(b *testing.B) {

	for n := 0; n < b.N; n++ {
		t0 := time.Now()
		generateThumbnailVips("./images/IMG_0901.JPG", "/tmp/thumbnailVips.jpg")
		b.Log("Took", time.Since(t0))
	}

}

func TestVips(t *testing.T) {

	generateThumbnailVips("./images/IMG_0901.JPG", "/tmp/thumbnailVips.jpg")
}

func BenchmarkTurbo(b *testing.B) {

	for n := 0; n < b.N; n++ {
		t0 := time.Now()
		generateThumbnailTurbo("./images/IMG_0901.JPG", "/tmp/thumbnailTurbo.jpg")
		b.Log("Took", time.Since(t0))
	}

}

func TestTurbo(t *testing.T) {

	generateThumbnailTurbo("./images/IMG_0901.JPG", "/tmp/thumbnailTurbo.jpg")
}

func BenchmarkNfnt(b *testing.B) {

	for n := 0; n < b.N; n++ {
		t0 := time.Now()
		generateThumbnailNfnt("./images/IMG_0901.JPG", "/tmp/thumbnailNfnt.jpg")
		b.Log("Took", time.Since(t0))
	}

}

func TestNfnt(t *testing.T) {

	generateThumbnailNfnt("./images/IMG_0901.JPG", "/tmp/thumbnailNfnt.jpg")
}
