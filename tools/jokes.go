package tools

import (
	_ "embed"
	"encoding/json"
	"math/rand/v2"

	"google.golang.org/adk/v2/agent"
)

// jokeData holds the embedded jokes file so it can be loaded at startup.
//go:embed data/jokes.json
var jokeData []byte

// jokeRepo stores the parsed joke dataset used by the tool.
var jokeRepo jokesApi

// init loads the embedded joke content and validates that it contains usable data.
func init() {
	if err := json.Unmarshal(jokeData, &jokeRepo); err != nil {
		panic("failed to parse embedded jokes.json file: " + err.Error())
	}
	if len(jokeRepo.Jokes) == 0 {
		panic("file: jokes.json parsed successfully, but contains zero jokes")
	}
}

// jokesApi represents the structure of the embedded joke dataset.
type jokesApi struct {
	Version     int         `json:"version"`
	GeneratedAt string      `json:"generated_at"`
	Attribution attribution `json:"attribution"`
	Count       int         `json:"count"`
	Jokes       []joke      `json:"jokes"`
}

// attribution captures metadata about where the jokes came from.
type attribution struct {
	Source string `json:"source"`
	Notice string `json:"notice"`
}

// joke is a single dad joke entry in the embedded dataset.
type joke struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}

type dadJokeInput struct{}

type dadJokeOutput joke

// GetDadJoke returns a single random dad joke from the embedded dataset.
func GetDadJoke(_ agent.Context, _ dadJokeInput) (dadJokeOutput, error) {
	randNum := rand.IntN(len(jokeRepo.Jokes))
	return dadJokeOutput(jokeRepo.Jokes[randNum]), nil
}
