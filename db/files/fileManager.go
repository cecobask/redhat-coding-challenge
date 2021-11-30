package files

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cecobask/redhat-coding-challenge/model"
)

// FileManager ...
type FileManager interface {
	CreateFile(multipart.File, model.Object) error
	RetrieveFile(model.Object) (string, error)
	DeleteFile(model.Object) error
}

// FileManager ...
type fileManager struct{}

// NewFileManager ...
func NewFileManager() FileManager {
	log.Println("Initialized file manager")
	return fileManager{}
}

func (fm fileManager) CreateFile(mpFile multipart.File, object model.Object) error {
	log.Println("FileManager.CreateFile() invoked")
	fileBytes, err := io.ReadAll(mpFile)
	if err != nil {
		log.Println(err)
		return err
	}
	// Create a file within the uploads directory
	uploadsDir := filepath.Join("uploads", object.BucketName)
	err = os.MkdirAll(uploadsDir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}
	fileName := fmt.Sprintf("%s.%s", object.ObjectName, object.ObjectExtension)
	file, err := os.Create(filepath.Join(uploadsDir, filepath.Base(fileName)))
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	file.Write(fileBytes)
	log.Println("Created new file in directory", uploadsDir, "called", fileName)
	return nil
}

func (fm fileManager) RetrieveFile(object model.Object) (string, error) {
	log.Println("FileManager.RetrieveFile() invoked")
	_, err := os.Stat(object.ObjectPath)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return object.ObjectPath, nil
}

func (fm fileManager) DeleteFile(object model.Object) error {
	log.Println("FileManager.DeleteFile() invoked")
	err := os.Remove(object.ObjectPath)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Removed file %s.%s from bucket %s", object.ObjectName, object.ObjectExtension, object.BucketName)
	return nil
}
