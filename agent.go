package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model/gemini"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/geminitool"
	"google.golang.org/genai"

	"dad-joke-agent/jokes"
	"dad-joke-agent/tools"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx, tools.MODEL, &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create Gemini model: %v", err)
	}

	dadJokeTool, err := tools.NewTool("getDadJoke", "Returns a random dad joke.", jokes.GetDadJoke)
	if err != nil {
		log.Fatalf("Failed to create tool: %v", err)
	}

	rootAgent, err := llmagent.New(llmagent.Config{
		Name:        "Dad Joke Agent",
		Model:       model,
		Description: "Dad joke agent",
		Instruction: `You are a helpful assistant that can tell dad jokes. 
					  Only tell jokes retrieved from the dadJokeTool unless requested by the user. 
					  You can also have a conversation on the jokes that you tell. One question or statement at a time.`,
		Tools: []tool.Tool{
			dadJokeTool,
			geminitool.GoogleSearch{},
		},
		GenerateContentConfig: &genai.GenerateContentConfig{
			ToolConfig: &genai.ToolConfig{
				IncludeServerSideToolInvocations: genai.Ptr(true),
			},
		},
	})

	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	if err := tools.LaunchAgent(ctx, rootAgent); err != nil {
		log.Fatalf("Run failed: %v", err)
	}
}
