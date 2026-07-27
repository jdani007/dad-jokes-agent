# Dad Jokes Agent

A small Go agent that tells dad jokes using the Google ADK and a random dad joke API.

## Prerequisites

- Go 1.22+ installed
- A Google API key for Gemini

## Setup

1. Clone the repository and move into it:

```bash
git clone <your-repo-url>
cd dad-jokes
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set your API key in your shell:

```bash
export GOOGLE_API_KEY=YOUR-API-KEY
```

You can copy the format from `template.env`.

## Run the agent

Start the app from the project root:

```bash
go run agent.go
```

## Example interaction

```text
User -> tell me a new joke

Agent -> The punchline often arrives before the set-up. Do you know the problem with UDP jokes?

Pretty classic networking humor, right? Did you get it?
```

## Notes for command line usage

- Make sure `GOOGLE_API_KEY` is exported in the same terminal session where you run `go run agent.go`.
- If you open a new terminal tab/window, export the variable again unless you add it to your shell profile (for example `~/.zshrc` or `~/.bashrc`).

## Thanks

Special thanks to [Jonas Werner](https://github.com/jonas-werner) and the `dad-joke-api` project for providing the jokes API:

https://github.com/jonas-werner/dad-joke-api
