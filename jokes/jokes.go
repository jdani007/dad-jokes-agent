package jokes

import (
	"encoding/json"
	"io"
	"net/http"

	"google.golang.org/adk/v2/agent"
)

const url = "https://dadjokes.bamboozledaardvark.com/api/jokes/random"

type dadJokeInput struct {}

type dadJokeOutput struct {
	Joke string `json:"joke"`
}

func GetDadJoke(_ agent.Context, _ dadJokeInput) (dadJokeOutput, error){
	
	resp, err := http.Get(url)
	if err != nil {
		return dadJokeOutput{}, err
	}
	defer resp.Body.Close()

	body,  err := io.ReadAll(resp.Body)
	if err != nil {
		return dadJokeOutput{}, err
	}

	var joke dadJokeOutput
	json.Unmarshal(body, &joke)

	return joke, nil
}