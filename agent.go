package main

import (
	"context"
	"log"
	"os"

	"dad-joke-agent/helpers"
	"dad-joke-agent/tools"

	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model/gemini"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/geminitool"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	// Create the Gemini model using the configured API key.
	model, err := gemini.NewModel(ctx, helpers.MODEL, &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create Gemini model: %v", err)
	}

	// Register the custom dad-joke tool so the agent can fetch jokes on demand.
	dadJokeTool, err := helpers.NewTool("getDadJoke", "Returns a random dad joke.", tools.GetDadJoke)
	if err != nil {
		log.Fatalf("Failed to create tool: %v", err)
	}

	// Configure the root agent with instructions and the tools it can use.
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

	// Launch the interactive CLI experience for the configured agent.
	if err := helpers.LaunchAgent(ctx, rootAgent); err != nil {
		log.Fatalf("Run failed: %v", err)
	}
}
