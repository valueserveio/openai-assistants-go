package main

import (
	"fmt"

	openaiassistantsgo "github.com/devhulk/openai-assistants-go"
)

func main() {
  assistant_instructions := `You are an AI Cybersecurity specialist who uses your knowledge base to help people make decisions based on the financial impact of cybersecurity solutions.`
  prompt := `I explain to me why someone would buy an AI security solution vs a regular security solution. Explain how much money it would cost a company depending on the decision they made.`

  fmt.Println(prompt, "\n -------------------------")
  fmt.Println(assistant_instructions, "\n -------------------------")

  // TODO: Each piece works seperately. Need to patch it all together.
  // TODO: Add proper error handling.
  // TODO: Files? Vector Stores?

  // assistant_id, err := openaiassistantsgo.CreateAssistant(assistant_instructions); if err != nil {
  //   fmt.Println(err)
  // }

  // thread_id := openaiassistantsgo.CreateThread()
  // fmt.Println(thread_id)

  //openaiassistantsgo.CreateUserMessage("thread_id", prompt)
  //openaiassistantsgo.CreateUserMessage("thread_id", prompt)
  //openaiassistantsgo.CreateRun("assistant_id", "thread_id")
  //openaiassistantsgo.ListMessages("thread_id")

  //openaiassistantsgo.DeleteAssistant("asst_8DfXObposm7OiUXStjnAQIAt")
  //openaiassistantsgo.ListAssistants()

}


