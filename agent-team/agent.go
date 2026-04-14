package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	const (
		appName = "weather_team_go"
		userID  = "vincent"
	)
	// 0. Set up the agent.
	rootAgent, err := setupAgent(ctx)
	if err != nil {
		log.Fatalf("Failed to set up agent: %v", err)
	}

	// 1. Create a session service.
	sessionService := session.InMemoryService()

	// 2. Create a session for this user.
	sessionResp, err := sessionService.Create(ctx, &session.CreateRequest{
		AppName: appName,
		UserID:  userID,
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	sess := sessionResp.Session
	sessionID := sess.ID()

	// 3. Set up the runner.
	r, err := runner.New(runner.Config{
		AppName:        appName,
		Agent:          rootAgent,
		SessionService: sessionService,
	})
	if err != nil {
		log.Fatalf("Failed to create runner: %v", err)
	}

	// 4. Prepare the user message.
	userMsg := &genai.Content{
		Role: string(genai.RoleUser),
		Parts: []*genai.Part{
			genai.NewPartFromText("What's the weather in Tokyo?"),
		},
	}

	// 5. Execute the runner and print streamed text responses.
	streamingMode := agent.StreamingModeSSE
	fmt.Print("\nAgent -> ")

	for event, err := range r.Run(
		ctx,
		userID,
		sessionID,
		userMsg,
		agent.RunConfig{
			StreamingMode: streamingMode,
		},
	) {
		if err != nil {
			log.Printf("run error: %v", err)
			continue
		}

		if event.LLMResponse.Content == nil {
			continue
		}

		for _, p := range event.LLMResponse.Content.Parts {
			if streamingMode != agent.StreamingModeSSE || event.LLMResponse.Partial {
				fmt.Print(p.Text)
			}
		}
	}
	fmt.Println()
}

func setupAgent(ctx context.Context) (agent.Agent, error) {
	// Set up the model.
	// Note: Authentication is handled through the GOOGLE_API_KEY environment variable.
	model, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		return nil, fmt.Errorf("create model: %w", err)
	}

	weatherTool, err := weatherTool()
	if err != nil {
		return nil, fmt.Errorf("create weather tool: %w", err)
	}

	// Define the agent.
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
		return nil, fmt.Errorf("create agent: %w", err)
	}

	return a, nil
}
