package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

type Repo struct {
	Name       string
	Created_at string
	Html_url   string
	Stars      int
}

func getUrl(username string) string {
	return "https://api.github.com/users/" + username + "/repos"
}

func fetchRepos(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprint(cli.ErrWriter, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var repos []Repo
	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Fprint(cli.ErrWriter, err)
	}

	for l := range repos {
		t, _ := time.Parse("2006-01-02T15:04:05Z0700",
			repos[l].Created_at)

		fmt.Printf("Name: %v \n", repos[l].Name)
		fmt.Printf("⭐ %v \n", repos[l].Stars)
		fmt.Printf("Created At: %v \n", t.Format(time.RFC822))
		fmt.Printf("Url: %v \n", repos[l].Html_url)
		fmt.Println("______________________________________________")
	}
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				Usage:   "GitHub username that you want to fetch the data of.",
			},
		},
		Action: func(c *cli.Context) error {
			username := c.String("username")
			url := getUrl(username)
			fetchRepos(url)
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
