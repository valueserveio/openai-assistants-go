package main

import (
	"fmt"
	"log"
	"time"

	openaiassistantsgo "github.com/devhulk/openai-assistants-go"
)

// TODO: Test Integrations

func main() {
	assistant_instructions := `You are an AI Cybersecurity specialist who uses your knowledge base to help people make decisions based on the financial impact of cybersecurity solutions.`
	prompt := `What is upload file endpoint used for? Use the file I uploaded.`

	fmt.Println(prompt, "\n -------------------------")
	fmt.Println(assistant_instructions, "\n -------------------------")

	//TODO: Implement Store Creation and Attach File
	store_id, err := openaiassistantsgo.CreateVectorStore("test_store")
	if err != nil {
		log.Fatal(err)
	}
	assistant_id, err := openaiassistantsgo.CreateAssistant(assistant_instructions, store_id)
	if err != nil {
		log.Fatal(err)
	}

	// openaiassistantsgo.DeleteAssistant(assistant_id)

	file_id, err := openaiassistantsgo.UploadFile("test_files/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file_id)

  if err := openaiassistantsgo.AttachFileToVectorStore(file_id, store_id); err != nil {
    log.Fatal(err)
  }
  
	thread_id, err := openaiassistantsgo.CreateThread()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(thread_id)

	openaiassistantsgo.CreateUserMessage(thread_id, prompt)

	_, err2 := openaiassistantsgo.CreateUserMessage(thread_id, prompt)
	if err != nil {
		log.Fatal(err2)
	}

	run_id, err := openaiassistantsgo.CreateRun(assistant_id, thread_id)
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second)

		run_status, err := openaiassistantsgo.GetRunStatus(run_id, thread_id)
		if err != nil {
			log.Fatal(err)
		}

		if run_status == "completed" {
			break
		}

	}

	openaiassistantsgo.ListMessages(thread_id)

}
