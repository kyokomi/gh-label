package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	var token, labelName, owner, repo, branch string
	flag.StringVar(&labelName, "l", "", "labelName")
	flag.StringVar(&token, "t", "", "github access token")
	flag.StringVar(&owner, "o", "", "github owner name")
	flag.StringVar(&repo, "r", "", "github repository name")
	flag.StringVar(&branch, "b", "", "github branch name")
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pulls, _, err := client.PullRequests.List(ctx, owner, repo, nil)
	if err != nil {
		log.Fatalln(err)
	}

	for _, r := range pulls {
		if *r.Head.Ref != branch {
			continue
		}

		number := *r.Number
		if isAlreadyLabel(client, ctx, owner, repo, number, labelName) {
			fmt.Println("already labelName")
			return
		}

		_, _, err := client.Issues.AddLabelsToIssue(ctx, owner, repo, number, []string{labelName})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func isAlreadyLabel(client *github.Client, ctx context.Context, owner, repo string, number int, labelName string) bool {
	labels, _, err := client.Issues.ListLabelsByIssue(ctx, owner, repo, number, nil)
	if err != nil {
		log.Fatalln(err)
	}

	for i := range labels {
		fmt.Println(*labels[i].Name)
		if *labels[i].Name == labelName {
			return true
		}
	}

	return false
}
