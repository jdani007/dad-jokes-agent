package tools

import (
	"encoding/json"
	"math/rand"
	"time"
	_ "embed"

	"google.golang.org/adk/v2/agent"
)

//go:embed data/jokes.json
var jokeData []byte

type jokesApi struct {
	Version     int         `json:"version"`
	GeneratedAt time.Time   `json:"generated_at"`
	Attribution attribution `json:"attribution"`
	Count       int         `json:"count"`
	Jokes       []jokes     `json:"jokes"`
}
type attribution struct {
	Source string `json:"source"`
	Notice string `json:"notice"`
}
type jokes struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}

type dadJokeInput struct{}

type dadJokeOutput jokes

func GetDadJoke(_ agent.Context, _ dadJokeInput) (dadJokeOutput, error) {

	joke, err := getRandomJoke()
	if err != nil {
		return dadJokeOutput{}, err
	}

	return joke, nil
}

func getRandomJoke() (dadJokeOutput, error) {
	var api jokesApi
	if err := json.Unmarshal(jokeData, &api); err != nil {
		return dadJokeOutput{}, err
	}

	randNum := rand.Intn(len(api.Jokes))
	return dadJokeOutput(api.Jokes[randNum]), nil
}