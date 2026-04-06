# my_agent

A minimal Google ADK agent project written in Go.

## What this project does

- Creates a Gemini model client.
- Builds a single `llmagent`.
- Enables the built-in `google_search` tool.
- Starts the ADK launcher (console/web modes).

## Prerequisites

- Go `1.25.7` or compatible
- A valid `GOOGLE_API_KEY`

## Setup

```bash
go mod tidy
export GOOGLE_API_KEY="<your_api_key>"
```

## Run

Run in default mode:

```bash
go run .
```

Run in console mode:

```bash
go run . console
```

Run in web mode:

```bash
go run . web
```

If startup fails, check:

- `GOOGLE_API_KEY` is set in your shell
- network access to Google APIs is available

## Project files

- `agent.go`: app entrypoint and agent configuration
- `go.mod` / `go.sum`: dependencies
- `.gitignore`: basic Go and env ignore rules
