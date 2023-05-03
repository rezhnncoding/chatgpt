package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("missing api key")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apikey)
	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"first gpt working"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{","},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}
