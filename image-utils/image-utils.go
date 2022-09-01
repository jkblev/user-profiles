package image_utils

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
)

/* Global variables */

// MaxWidth is the max number of pixels that an image can be downsized to, by width
var MaxWidth uint = 256

// MaxHeight is the max number of pixels that an image can be downsized to, by height
var MaxHeight uint = 256

// ResizeImage will resize a given image.Image to fit the max height and
// max width as provided in the global variables. Uses nearest-neighbor interpolation.
func ResizeImage(image image.Image) image.Image {
	return resize.Thumbnail(MaxWidth, MaxHeight, image, resize.NearestNeighbor)
}

// GetFileContentType determines the Content-Type of the given file
// that was passed in as part of a multipart request
func GetFileContentType(file *multipart.FileHeader) (string, error) {
	openedFile, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer openedFile.Close()

	// Only the first 512 bytes are used to detect content type
	buf := make([]byte, 512)
	_, err = openedFile.Read(buf)
	if err != nil {
		return "", err
	}

	// the function that actually does the trick
	contentType := http.DetectContentType(buf)

	return contentType, nil
}

// GetNewPNGFileName accepts a file name, strips the suffix, and
// replaces the suffix with _out.png from the previous filename.
func GetNewPNGFileName(fileName string) string {
	regex := regexp.MustCompile(`(\.png)|(\.jp(e)?g)`)
	newFileName := regex.ReplaceAllString(fileName, "_out.png")

	return newFileName
}

// ResizePNG will create a new file and resize the given raw PNG file bytes,
// encode the resized bytes, and put them in the new file location
func ResizePNG(pngBytes []byte, newFileName string) {
	decodedImage, err := png.Decode(bytes.NewReader(pngBytes))
	if err != nil {
		panic(err)
	}

	resized := ResizeImage(decodedImage)

	newFile, err := os.Create(newFileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	if err := png.Encode(newFile, resized); err != nil {
		panic(err)
	}

	return
}

// ConvertJPEGToPNGAndResize will create a new PNG file, resize the
// given raw JPEG file bytes, encode the resized bytes into PNG, and put them
// in the new file location
func ConvertJPEGToPNGAndResize(jpegBytes []byte, newFileName string) {
	decodedImage, err := jpeg.Decode(bytes.NewReader(jpegBytes))
	if err != nil {
		panic(err)
	}

	resized := ResizeImage(decodedImage)

	newFile, err := os.Create(newFileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	if err := png.Encode(newFile, resized); err != nil {
		panic(err)
	}

	return
}
