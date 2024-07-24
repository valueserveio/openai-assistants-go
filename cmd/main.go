package main

import (
	"fmt"
	"log"

	openaiassistantsgo "github.com/devhulk/openai-assistants-go"
)

func main() {
	// assistant_instructions := `You are an AI Cybersecurity specialist who uses your knowledge base to help people make decisions based on the financial impact of cybersecurity solutions.`
	// prompt := `Explain to me why someone would buy an AI security solution vs a regular security solution. Explain how much money it would cost a company depending on the decision they made.`
	//
	// fmt.Println(prompt, "\n -------------------------")
	// fmt.Println(assistant_instructions, "\n -------------------------")
	//
	//  assistant_id, err := openaiassistantsgo.CreateAssistant(assistant_instructions); if err != nil {
	//    log.Fatal(err)
	//  }
	//
	//  thread_id, err := openaiassistantsgo.CreateThread(); if err != nil {
	//    log.Fatal(err)
	//  }
	//  fmt.Println(thread_id)
	//
	// openaiassistantsgo.CreateUserMessage(thread_id, prompt)
	//
	// _, err2 := openaiassistantsgo.CreateUserMessage(thread_id, prompt); if err != nil {
	//   log.Fatal(err2)
	// }
	//
	// run_id, err := openaiassistantsgo.CreateRun(assistant_id, thread_id); if err != nil {
	//   log.Fatal(err)
	// }
	//
	// for {
	//   time.Sleep(time.Second)
	//
	//   run_status, err := openaiassistantsgo.GetRunStatus(run_id, thread_id); if err != nil {
	//     log.Fatal(err)
	//   }
	//
	//   if run_status == "completed" {
	//     break
	//   }
	//
	// }
	//
	// openaiassistantsgo.ListMessages(thread_id)
	// openaiassistantsgo.DeleteAssistant(assistant_id)

	// file_id, err := openaiassistantsgo.UploadFile("test_files/test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(file_id)

  file_ids, err := openaiassistantsgo.ListFiles(); if err != nil {
    log.Fatal(err)
  }

  fmt.Println(file_ids)
}
