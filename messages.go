package openaiassistantsgo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Users create Messages, Messages are a part of Threads.

func CreateUserMessage(thread_id string, prompt string) {
	client := &http.Client{}

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

	fmt.Printf("%s\n", bodyText)
}

func ListMessages(thread_id string) {
  client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", thread_id), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + os.Getenv("OPENAI_API_KEY"))
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

