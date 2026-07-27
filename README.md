# Dad Jokes Agent

A small Go agent that tells dad jokes using the Google ADK and a random dad joke API.

## Source of jokes

This agent can provide jokes in three ways:

- Pull a random joke from [`dad-joke-api`](https://github.com/jonas-werner/dad-joke-api)
- Create an original dad joke when you ask it to make one.
- Use web search tools to find new jokes from other places online.

## Prerequisites

- Go 1.22+ installed
- A Google API key for Gemini

## Setup

1. Clone the repository and move into it:

```bash
git clone https://github.com/jdani007/dad-jokes-agent.git
cd dad-jokes-agent
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set your API key in your shell (or use the `.env` file):

> ```bash
> export GOOGLE_API_KEY=YOUR-API-KEY
> ```
>
> You can copy the format from `template.env`.
>
> ```bash
> source .env
> ```

## Run the agent

Start the app from the project root:

```bash
go run agent.go
```

## Build the agent

You can build the agent into an executable binary with:

```bash
go build -o dad-jokes-agent agent.go
./dad-jokes-agent
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
