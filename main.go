package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
)

type Commit struct {
	Message   string    `json:"message"`
	Committer Committer `json:"committer"`
	Content   string    `json:"content"`
	Path      string    `json:"path"`
}

type Committer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Config struct {
	Username    string `toml:"username"`
	Email       string `toml:"email"`
	Repository  string `toml:"repository"`
	Message     string `toml:"message"`
	Content     string `toml:"content"`
	AccessToken string `toml:"access_token"`
	Directory   string `toml:"directory"`
}

func main() {

	config := Config{}
	data, err := ioutil.ReadFile("config.conf")

	if err != nil {
		log.Fatal(err)
	}
	_, _err := toml.Decode(string(data), &config)

	if _err != nil {
		log.Fatal(_err)
	}

	filename := fmt.Sprintf("LOG-%s.log", time.Now().Format("20060102150405"))
	path := config.Directory + "/" + filename
	url := fmt.Sprintf("http://api.github.com/repos/%s/%s/contents/%s?access_token=%s", config.Username, config.Repository, path, config.AccessToken)

	commit := Commit{
		Message: config.Message,
		Committer: Committer{
			Name:  config.Username,
			Email: config.Email,
		},
		Content: b64.StdEncoding.EncodeToString([]byte(config.Content)),
		Path:    path,
	}

	client := http.Client{}

	b, err := json.Marshal(commit)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	req.Header.Add("Accept", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp)
}
