package tools

import (
	_ "embed"
	"encoding/json"
	"math/rand/v2"

	"google.golang.org/adk/v2/agent"
)

//go:embed data/jokes.json
var jokeData []byte
var jokeRepo jokesApi

func init() {
	if err := json.Unmarshal(jokeData, &jokeRepo); err != nil {
		panic("failed to parse embedded jokes.json file: " + err.Error())
	}
	if len(jokeRepo.Jokes) == 0 {
		panic("file: jokes.json parsed successfully, but contains zero jokes")
	}
}

type jokesApi struct {
	Version     int         `json:"version"`
	GeneratedAt string      `json:"generated_at"`
	Attribution attribution `json:"attribution"`
	Count       int         `json:"count"`
	Jokes       []joke      `json:"jokes"`
}

type attribution struct {
	Source string `json:"source"`
	Notice string `json:"notice"`
}

type joke struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}

type dadJokeInput struct{}

type dadJokeOutput joke

func GetDadJoke(_ agent.Context, _ dadJokeInput) (dadJokeOutput, error) {
	randNum := rand.IntN(len(jokeRepo.Jokes))
	return dadJokeOutput(jokeRepo.Jokes[randNum]), nil
}
