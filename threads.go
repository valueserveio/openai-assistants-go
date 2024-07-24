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

// Assistants interact with Threads. Threads contain messages.

func CreateThread() (string, error) {
	client := &http.Client{}

  var thread_id string
	var data = strings.NewReader(``)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/threads", data)
	if err != nil {
		log.Fatal(err)
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
	}

	m := b.(map[string]interface{})

	for k, v := range m {
		if k == "id" {
      thread_id = v.(string)
			fmt.Println(v)
      return thread_id, nil
		}

		if k == "tool_resources" {
			fmt.Println(v)
		}
	}

	fmt.Printf("%s\n", bodyText)
  return "ID was not found", nil

}

func DeleteThread(id string) error {

  client := &http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.openai.com/v1/threads/%s", id), nil)
	if err != nil {
		log.Fatal(err)
    return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + os.Getenv("OPENAI_API_KEY"))
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
