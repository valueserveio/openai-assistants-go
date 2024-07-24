package openaiassistantsgo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func CreateVectorStore(name string) (string, error) {

	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{
    "name": "%s"
  }`, name))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/vector_stores", data)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OpenAI-Beta", "assistants=v2")

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

	var store_id string

	var b interface{}

	if err := json.Unmarshal(bodyText, &b); err != nil {
		log.Fatal(err)
	}

	m := b.(map[string]interface{})

	for k, v := range m {
		if k == "id" {
			store_id = v.(string)
			fmt.Println(v)
			return store_id, nil
		}
	}

	fmt.Printf("%s\n", bodyText)

	return "No Vector Storage ID.", nil
}

func DeleteVectorStore(store_id string) error {
  
  client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.openai.com/v1/vector_stores/%s", store_id), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer " + os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", bodyText)
  return nil
}

func AttachFileToVectorStore(file_id string, store_id string) error {
  client := &http.Client{}

	var data = strings.NewReader(fmt.Sprintf(`{
      "file_id": "%s"
    }`, file_id))

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/vector_stores/%s/files", store_id), data)
	if err != nil {
		log.Fatal(err)
    return err
	}

	req.Header.Set("Authorization", "Bearer " + os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
    return err
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
    return err
	}

	fmt.Printf("%s\n", bodyText)
  return nil
}

func ListVectorStoreFiles(store_id string) error {
  client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openai.com/v1/vector_stores/%s/files", store_id), nil)
	if err != nil {
		log.Fatal(err)
    return err
	}

	req.Header.Set("Authorization", "Bearer " + os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
    return err
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
    return err
	}

	fmt.Printf("%s\n", bodyText)
  return nil
}
