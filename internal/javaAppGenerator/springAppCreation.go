package javaAppGenerator

import (
	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/utils"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	BaseUrl        = "https://start.spring.io"
	InitializerUrl = "/starter.zip"
	ZipFormat      = ".zip"
)

type SpringAppGenerator struct {
	TargetDirectory string
	ProjectName     string
}

func (sag SpringAppGenerator) CreateProjectStructure() {
	requestUrl := BaseUrl + InitializerUrl
	filePath := sag.TargetDirectory + "/" + sag.ProjectName

	res, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %v\n", res.Status)

	sag.saveProject(res, filePath)
}

func (sag SpringAppGenerator) saveProject(res *http.Response, filePath string) {
	file, err := os.Create(filePath + ZipFormat)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := res.Body.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := file.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	utils.Unzip(filePath+ZipFormat, filePath)
}
