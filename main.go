package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func main() {
	err := call("http://localhost:3000/image", "POST")

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	log.Println("success")
}

func call(urlPath, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", "test.pdf")

	if err != nil {
		return err
	}

	file, err := os.Open("test.pdf")

	fmt.Println("file:", file)

	if err != nil {
		return err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}

	writer.Close()

	req, err := http.NewRequest(method, urlPath, bytes.NewReader(body.Bytes()))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)

	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}

	return nil
}
