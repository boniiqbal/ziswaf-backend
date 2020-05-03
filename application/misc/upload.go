package misc

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type SavedFile struct {
	FileName string
	FilePath string
	FileURL  string
	MimeType string
	Concern  string
}

// UploadFileToS3 saves a file to aws bucket and returns the url to // the file and an error if there's any
func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	idAWS := os.Getenv("AWS_ID")
	secretAWS := os.Getenv("AWS_SECRET")
	tokenAWS := os.Getenv("AWS_TOKEN")
	regionAWS := os.Getenv("AWS_REGION")
	endpointAWS := os.Getenv("AWS_ENDPOINT")
	bucketAWS := os.Getenv("AWS_BUCKET")

	// create an AWS session which can be
	s, errSession := session.NewSession(&aws.Config{
		Region:      aws.String(regionAWS),
		Endpoint:    aws.String(endpointAWS),
		Credentials: credentials.NewStaticCredentials(idAWS, secretAWS, tokenAWS),
	})

	if errSession != nil {
		return "", errSession
	}
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// create a unique file name for the file
	tempFileName := "images/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucketAWS),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}

// UploadFile to save image get from request
func UploadFile(c *gin.Context, param string) (*SavedFile, error) {
	paramFile := param
	if param == "" {
		paramFile = "attachment_files"
	}
	file, errFiles := c.FormFile(paramFile)

	if file == nil {
		return &SavedFile{}, fmt.Errorf("Empty File: %s is empty", param)
	}

	if errFiles != nil {
		return &SavedFile{}, errFiles
	}
	mimeType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(mimeType, "image/") {
		return &SavedFile{}, errors.New("Only image file allowed")
	}

	folderPath := "./attachments/images/"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0755)
	}

	extfile := strings.Split(file.Filename, ".")
	randFileName := uuid.New()
	file.Filename = randFileName.String() + "image" + "." + extfile[len(extfile)-1]

	path := folderPath + file.Filename

	var host string
	if os.Getenv("GO_ENV") == "production" {
		host = "http://ziswaf-api.projects.refactory.id" //Development
	} else {
		host = "http://" + c.Request.Host
	}
	url := host + "/api/v1/images/" + file.Filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		return &SavedFile{}, err
	}

	return &SavedFile{
		FileName: file.Filename,
		FilePath: path,
		FileURL:  url,
		MimeType: file.Header.Get("Content-Type"),
	}, nil
}
