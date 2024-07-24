package openaiassistantsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

//TODO: Files endpoint for uploading files. (outputs file ids)
// TODO: Vector Storage for storing references to those files to be used with "file_search"

func DeleteFile(file_id string) error {

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.openai.com/v1/files/%s", file_id), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
    return err
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
    return err
	}

	fmt.Printf("%s\n", bodyText)

  return nil
}

func ListFiles() ([]string, error) {
	var file_ids []string

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.openai.com/v1/files", nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var b interface{}
	if err := json.Unmarshal(bodyText, &b); err != nil {
		log.Fatal(err)
	}

	m := b.(map[string]interface{})

	for k, v := range m {
		if k == "id" {
			file_ids = append(file_ids, v.(string))
			fmt.Println(v)
		}
	}

	fmt.Printf("%s\n", bodyText)

	return file_ids, nil
}

func UploadFile(filename string) (string, error) {
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	formField, err := writer.CreateFormField("purpose")
	if err != nil {
		log.Fatal(err)
	}
	_, err = formField.Write([]byte("assistants"))

	fw, err := writer.CreateFormFile("file", filepath.Base(filename))
	if err != nil {
		log.Fatal(err)
	}

	fd, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer fd.Close()

	_, err = io.Copy(fw, fd)
	if err != nil {
		log.Fatal(err)
	}

	writer.Close()

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/files", form)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var b interface{}

	if err := json.Unmarshal(bodyText, &b); err != nil {
		log.Fatal(err)
	}

	var file_id string
	m := b.(map[string]interface{})

	for k, v := range m {
		if k == "id" {
			file_id = v.(string)
			return file_id, nil
		}
	}
	fmt.Printf("%s\n", bodyText)

	return "File ID not found.", nil

}
