package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	// 1. Set up the model.
	// Note: Authentication is handled through the GOOGLE_API_KEY environment variable.
	model, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	weatherTool, err := weatherTool()
	if err != nil {
		log.Fatalf("Failed to create weather tool: %v", err)
	}

	// 2. Define the agent.
	a, err := llmagent.New(llmagent.Config{
		Name:        "multi_tool_agent",
		Model:       model,
		Description: "An agent that can answer questions using Custom tools",
		Instruction: "You are a helpful assistant. Use the available tools to answer questions.",
		Tools: []tool.Tool{
			weatherTool,
		},
	})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// 3. Configure the launcher and run it.
	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(a),
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}

}
