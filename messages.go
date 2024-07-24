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

// Users create Messages, Messages are a part of Threads.

func CreateUserMessage(thread_id string, prompt string) (string, error) {
	client := &http.Client{}
	var message_id string

	var data = strings.NewReader(fmt.Sprintf(`{
      "role": "user",
      "content": "%s"
    }`, prompt))

	fmt.Println(data)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", thread_id), data)
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

	var b interface{}
	if err := json.Unmarshal(bodyText, &b); err != nil {
		log.Fatal(err)
	}

	m := b.(map[string]interface{})

	for k, v := range m {
		if k == "id" {
			message_id = v.(string)
			fmt.Println(v)
			return message_id, nil
		}

		if k == "tool_resources" {
			fmt.Println(v)
		}
	}

	fmt.Printf("%s\n", bodyText)
  return "Message ID not found.", err
}

func ListMessages(thread_id string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", thread_id), nil)
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
