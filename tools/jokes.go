package tools

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"sync"

	"google.golang.org/adk/v2/agent"
)

// jokeData stores the embedded jokes dataset that is parsed on first use.
//go:embed data/jokes.json
var jokeData []byte

// loadJokes lazily parses the embedded dataset once and caches the result.
var loadJokes = sync.OnceValues(func() (jokesApi, error) {
	var repo jokesApi
	if err := json.Unmarshal(jokeData, &repo); err != nil {
		return jokesApi{}, fmt.Errorf("failed to parse jokes.json: %w", err)
	}
	if len(repo.Jokes) == 0 {
		return jokesApi{}, fmt.Errorf("jokes.json contains zero jokes")
	}
	return repo, nil
})

// jokesApi mirrors the structure of the embedded joke dataset.
type jokesApi struct {
	Version     int         `json:"version"`
	GeneratedAt string      `json:"generated_at"`
	Attribution attribution `json:"attribution"`
	Count       int         `json:"count"`
	Jokes       []joke      `json:"jokes"`
}

// attribution captures source metadata for the embedded joke set.
type attribution struct {
	Source string `json:"source"`
	Notice string `json:"notice"`
}

// joke is a single dad-joke entry in the embedded dataset.
type joke struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}

// dadJokeInput is the empty input payload required by the tool handler.
type dadJokeInput struct{}

// dadJokeOutput is the tool result type returned for a single dad joke.
type dadJokeOutput joke

// GetDadJoke returns one randomly selected dad joke from the embedded dataset.
func GetDadJoke(_ agent.Context, _ dadJokeInput) (dadJokeOutput, error) {
	repo, err := loadJokes()
	if err != nil {
		return dadJokeOutput{}, err
	}

	randNum := rand.IntN(len(repo.Jokes))
	return dadJokeOutput(repo.Jokes[randNum]), nil
}
