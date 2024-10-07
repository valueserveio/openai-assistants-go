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

func CreateAssistant(instructions string, store_id string) (string, error) {

	client := &http.Client{}
	var assistant_instructions string
	var assistant_id string

	if instructions != "" {
		assistant_instructions = instructions
	} else {
		assistant_instructions = "You are a business value assistant specializing in the value purchasing a new piece of software will or will not have on a business."
	}

	var data = strings.NewReader(fmt.Sprintf(`{
    "instructions": "%s",
    "name": "Test Assistant",
    "tools": [
      {"type": "file_search"}
    ],
    "model": "gpt-4o-mini",
    "tool_resources": {
      "file_search": {
        "vector_store_ids": ["%s"]
      }
    }
  }`, assistant_instructions, store_id))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/assistants", data)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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

	var b interface{}
	if err := json.Unmarshal(bodyText, &b); err != nil {
		log.Fatal(err)
		return "", err
	}

	m := b.(map[string]interface{})

	for k, v := range m {
		if k == "id" {
			fmt.Println(v)
			assistant_id = v.(string)
			return assistant_id, nil
		}

		if k == "tool_resources" {
			fmt.Println(v)
		}
	}

	fmt.Printf("%s\n", bodyText)
	return assistant_id, nil

}

func DeleteAssistant(id string) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.openai.com/v1/assistants/%s", id), nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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
}

func ListAssistants() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.openai.com/v1/assistants?order=desc&limit=20", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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
}

func UpdateAssistantModel(id string, model string) {
	client := &http.Client{}

	var data = strings.NewReader(fmt.Sprintf(`{
    "model": "%s"
  }`, model))

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/assistants/%s", id), data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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

}

func RetrieveAssistant(id string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openai.com/v1/assistants/%s", id), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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
}
