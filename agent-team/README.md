# agent-team

This example shows how to build a Google ADK Go agent that can call a custom Go tool.

## What It Does

- Uses `gemini-2.5-flash` as the underlying LLM
- Registers a custom `get_weather` function tool
- Returns mock weather data for `newyork`, `london`, and `tokyo`
- Runs through the ADK full launcher, so you can use console, API, or Web UI modes

## Files

- `agent.go`: creates the Gemini model, registers the tool, and starts the launcher
- `tool.go`: defines the weather tool input/output types and mock tool logic
- `.example.env`: example environment variables for local setup

## Prerequisites

- Go `1.25.7`
- A valid `GOOGLE_API_KEY`
- Google ADK Go dependencies installed through `go mod`

## Setup

1. Copy the values from `.example.env` into your local shell or `.env`.
2. Set `GOOGLE_API_KEY` to a valid Gemini API key.
3. Install dependencies:

```bash
go mod tidy
```

## Run

From the `agent-team` directory:

```bash
go run . --help
```

Typical local run:

```bash
go run .
```

## Example Prompt

```text
What is the weather in Tokyo?
```

## Known Limits

- The weather tool uses mock data, not a real weather API.
- `taipei` is not included in the current mock dataset, so the tool returns an error response for that city.
- Gemini free-tier quotas can cause `429 RESOURCE_EXHAUSTED` errors even when the tool itself succeeds.

## Verify

```bash
GOCACHE=/tmp/go-build-agent-team go build ./...
```
