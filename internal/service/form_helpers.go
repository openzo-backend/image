package service
// package service

// import (
// 	"mime/multipart"
// 	"strconv"
// 	"time"
// )

// func StringToInt(s string) int {
// 	i, err := strconv.Atoi(s)

// 	if err != nil {
// 		return 0
// 	}
// 	return i
// }

// func StringToTime(s string) time.Time {
// 	t, err := time.Parse(time.DateTime, s)

// 	if err != nil {
// 		return time.Now().Add(24 * time.Hour)
// 	}
// 	return t
// }

// func UploadFile(file *multipart.FileHeader) (string, error) {
// 	fileHeader := file
// 	f, err := fileHeader.Open()
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()
// 	// f, err := imageProcessing(file, 20)

// 	uploaderURL, err := SaveFile(f, fileHeader)
// 	if err != nil {
// 		return "", err
// 	}
// 	return uploaderURL, nil

// }

// func UploadMultipleFiles(files []*multipart.FileHeader) ([]string, error) {
// 	var imageURLs []string
// 	for _, file := range files {
// 		fileHeader := file
// 		f, err := fileHeader.Open()
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer f.Close()

// 		uploaderURL, err := SaveFile(f, fileHeader)
// 		if err != nil {
// 			return nil, err
// 		}
// 		imageURLs = append(imageURLs, uploaderURL)

// 	}

// 	return imageURLs, nil
// }
