package openaiassistantsgo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Runs take a Thread id which could contain Messages and runs it against an Assistant

func CreateRun(assistant_id string, thread_id string) {
	client := &http.Client{}

	var data = strings.NewReader(fmt.Sprintf(`{
    "assistant_id": "%s"
  }`, assistant_id))

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/threads/%s/runs", thread_id), data)
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
}
