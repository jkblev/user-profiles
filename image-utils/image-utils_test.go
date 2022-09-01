package image_utils

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"testing"
)

// TestResizeImageJPEGWide verifies that the wide JPEG test image is
// resized to be less than the MaxWidth variable (and the MaxHeight variable).
func TestResizeImageJPEGWide(t *testing.T) {
	openedFile, _ := os.Open("test-images/wide.jpeg")
	defer openedFile.Close()
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		panic(err)
	}

	decodedImage, err := jpeg.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		panic(err)
	}

	resized := ResizeImage(decodedImage)

	imageWidth := resized.Bounds().Max.X
	imageHeight := resized.Bounds().Max.Y

	if imageWidth > int(MaxWidth) {
		t.Errorf("expected image width to be <= %v, but it is %v", MaxWidth, imageWidth)
	}

	if imageHeight > int(MaxHeight) {
		t.Errorf("expected image height to be <= %v, but it is %v", MaxHeight, imageHeight)
	}
}

// TestResizeImageJPEGTall verifies that the tall JPEG test image is
// resized to be less than the MaxHeight variable (and the MaxWidth variable).
func TestResizeImageJPEGTall(t *testing.T) {
	openedFile, _ := os.Open("test-images/tall.jpeg")
	defer openedFile.Close()
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		panic(err)
	}

	decodedImage, err := jpeg.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		panic(err)
	}

	resized := ResizeImage(decodedImage)

	imageWidth := resized.Bounds().Max.X
	imageHeight := resized.Bounds().Max.Y

	if imageWidth > int(MaxWidth) {
		t.Errorf("expected image width to be <= %v, but it is %v", MaxWidth, imageWidth)
	}

	if imageHeight > int(MaxHeight) {
		t.Errorf("expected image height to be <= %v, but it is %v", MaxHeight, imageHeight)
	}
}

// TestResizeImagePNGWide verifies that the wide PNG test image is
// resized to be less than the MaxWidth variable (and the MaxHeight variable).
func TestResizeImagePNGWide(t *testing.T) {
	openedFile, _ := os.Open("test-images/wide.png")
	defer openedFile.Close()
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		panic(err)
	}

	decodedImage, err := png.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		panic(err)
	}

	resized := ResizeImage(decodedImage)

	imageWidth := resized.Bounds().Max.X
	imageHeight := resized.Bounds().Max.Y

	if imageWidth > int(MaxWidth) {
		t.Errorf("expected image width to be <= %v, but it is %v", MaxWidth, imageWidth)
	}

	if imageHeight > int(MaxHeight) {
		t.Errorf("expected image height to be <= %v, but it is %v", MaxHeight, imageHeight)
	}
}

// TestResizeImagePNGTall verifies that the tall PNG test image is
// resized to be less than the MaxHeight variable (and the MaxWidth variable).
func TestResizeImagePNGTall(t *testing.T) {
	openedFile, _ := os.Open("test-images/tall.png")
	defer openedFile.Close()
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		panic(err)
	}

	decodedImage, err := png.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		panic(err)
	}

	resized := ResizeImage(decodedImage)

	imageWidth := resized.Bounds().Max.X
	imageHeight := resized.Bounds().Max.Y

	if imageWidth > int(MaxWidth) {
		t.Errorf("expected image width to be <= %v, but it is %v", MaxWidth, imageWidth)
	}

	if imageHeight > int(MaxHeight) {
		t.Errorf("expected image height to be <= %v, but it is %v", MaxHeight, imageHeight)
	}
}

// TestGetNewPNGFileNameWithPNG verifies that a PNG file's
// new filename will take the existing filename and replace
// the .png suffix with _out.png
func TestGetNewPNGFileNameWithPNG(t *testing.T) {
	filename := "filename.png"

	received := GetNewPNGFileName(filename)

	if received != "filename_out.png" {
		t.Errorf("Expected new filename to be 'filename_out.png', received %v", received)
	}
}

// TestGetNewPNGFileNameWithPNG verifies that a JPEG file's
// new filename will take the existing filename and replace
// the .jpeg suffix with _out.png
func TestGetNewPNGFileNameWithJPEG(t *testing.T) {
	filename := "filename.jpeg"

	received := GetNewPNGFileName(filename)

	if received != "filename_out.png" {
		t.Errorf("Expected new filename to be 'filename_out.png', received %v", received)
	}
}

// TestGetNewPNGFileNameWithPNG verifies that a JPG file's
// new filename will take the existing filename and replace
// the .jpg suffix with _out.png
func TestGetNewPNGFileNameWithJPG(t *testing.T) {
	filename := "filename.jpg"

	received := GetNewPNGFileName(filename)

	if received != "filename_out.png" {
		t.Errorf("Expected new filename to be 'filename_out.png', received %v", received)
	}
}

// TestResizePNG verifies that a new file is created
// with the resized image of the raw PNG filebytes passed to ResizePNG
func TestResizePNG(t *testing.T) {
	openedFile, _ := os.Open("test-images/tall.png")
	defer openedFile.Close()

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		panic(err)
	}

	ResizePNG(fileBytes, "test-images/tall_resize.png")

	resizedFile, _ := os.Open("test-images/tall_resize.png")
	defer openedFile.Close()

	resizedFileBytes, err := io.ReadAll(resizedFile)
	if err != nil {
		panic(err)
	}
	decodedImage, err := png.Decode(bytes.NewReader(resizedFileBytes))
	if err != nil {
		panic(err)
	}

	imageWidth := decodedImage.Bounds().Max.X
	imageHeight := decodedImage.Bounds().Max.Y

	if imageWidth > int(MaxWidth) {
		t.Errorf("expected image width to be <= %v, but it is %v", MaxWidth, imageWidth)
	}

	if imageHeight > int(MaxHeight) {
		t.Errorf("expected image height to be <= %v, but it is %v", MaxHeight, imageHeight)
	}

	// Clean up the resized file
	removeErr := os.Remove("test-images/tall_resize.png")
	if removeErr != nil {
		panic(removeErr)
	}
}

// TestConvertJPEGToPNGAndResize verifies that a new file is created
// with the resized image of the raw JPEG filebytes passed to ConvertJPEGToPNGAndResize
func TestConvertJPEGToPNGAndResize(t *testing.T) {
	openedFile, _ := os.Open("test-images/tall.jpeg")
	defer openedFile.Close()

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		panic(err)
	}

	ConvertJPEGToPNGAndResize(fileBytes, "test-images/converted_test.png")

	resizedFile, _ := os.Open("test-images/converted_test.png")
	defer openedFile.Close()

	resizedFileBytes, err := io.ReadAll(resizedFile)
	if err != nil {
		panic(err)
	}
	decodedImage, err := png.Decode(bytes.NewReader(resizedFileBytes))
	if err != nil {
		panic(err)
	}

	imageWidth := decodedImage.Bounds().Max.X
	imageHeight := decodedImage.Bounds().Max.Y

	if imageWidth > int(MaxWidth) {
		t.Errorf("expected image width to be <= %v, but it is %v", MaxWidth, imageWidth)
	}

	if imageHeight > int(MaxHeight) {
		t.Errorf("expected image height to be <= %v, but it is %v", MaxHeight, imageHeight)
	}

	// Clean up the resized file
	removeErr := os.Remove("test-images/converted_test.png")
	if removeErr != nil {
		panic(removeErr)
	}
}
