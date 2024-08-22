package openaiassistantsgo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type VectorStorageFileResponse struct {
	ID               string  `json:"id"`
	Object           string  `json:"object"`
	UsageBytes       int     `json:"usage_bytes"`
	CreatedAt        int64   `json:"created_at"`
	VectorStoreID    string  `json:"vector_store_id"`
	Status           string  `json:"status"`
	LastError        *string `json:"last_error"`
	ChunkingStrategy struct {
		Type   string `json:"type"`
		Static struct {
			MaxChunkSizeTokens int `json:"max_chunk_size_tokens"`
			ChunkOverlapTokens int `json:"chunk_overlap_tokens"`
		} `json:"static"`
	} `json:"chunking_strategy"`
}

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
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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

func AttachFileToVectorStoreBak(file_id string, store_id string) error {
	client := &http.Client{}

	var data = strings.NewReader(fmt.Sprintf(`{
      "file_id": "%s"
    }`, file_id))

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/vector_stores/%s/files", store_id), data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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

func AttachFileToVectorStore(file_id string, store_id string) error {
	client := &http.Client{}

	// Initial POST request to attach the file
	var data = strings.NewReader(fmt.Sprintf(`{
      "file_id": "%s"
    }`, file_id))
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/vector_stores/%s/files", store_id), data)
	if err != nil {
		return fmt.Errorf("error creating POST request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending POST request: %w", err)
	}
	defer resp.Body.Close()

	// Check the initial response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("POST request failed with status code: %d", resp.StatusCode)
	}

	// Now, repeatedly check the status with GET requests
	for {
		time.Sleep(1 * time.Second) // Wait before checking status

		getReq, err := http.NewRequest("GET", fmt.Sprintf("https://api.openai.com/v1/vector_stores/%s/files/%s", store_id, file_id), nil)
		if err != nil {
			return fmt.Errorf("error creating GET request: %w", err)
		}
		getReq.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
		getReq.Header.Set("OpenAI-Beta", "assistants=v2")

		getResp, err := client.Do(getReq)
		if err != nil {
			return fmt.Errorf("error sending GET request: %w", err)
		}
		defer getResp.Body.Close()

		bodyText, err := io.ReadAll(getResp.Body)
		if err != nil {
			return fmt.Errorf("error reading GET response body: %w", err)
		}

		var fileResponse VectorStorageFileResponse
		err = json.Unmarshal(bodyText, &fileResponse)
		if err != nil {
			return fmt.Errorf("error unmarshaling GET response: %w", err)
		}

		fmt.Printf("File status: %s\n", fileResponse.Status)

		if fileResponse.Status == "completed" {
			return nil // Success, exit the function
		} else if fileResponse.Status == "failed" {
			return fmt.Errorf("file processing failed: %s", *fileResponse.LastError)
		}

		// If status is still processing, continue the loop
	}
}

func ListVectorStoreFiles(store_id string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openai.com/v1/vector_stores/%s/files", store_id), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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
