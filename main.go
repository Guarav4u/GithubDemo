package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Package struct {
	FullName   string
	ForksCount int
	StarsCount int
}

// Fetch repository on github
func fetchRepository(ctx context.Context, client *github.Client) {

	repo, _, err := client.Repositories.Get(ctx, "Guarav4u", "foo")
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
	pack := Package{
		FullName:   *repo.FullName,
		ForksCount: *repo.ForksCount,
		StarsCount: *repo.StargazersCount,
	}
	fmt.Printf("%+v\n", pack)
}

// Create Repository on github
func createRepository(ctx context.Context, client *github.Client, repoName string, repoType bool) {

	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(repoType),
	}
	client.Repositories.Create(ctx, "", repo)

}

//Add Collaborator
func addCollaborator(ctx context.Context, client *github.Client, owner, repoName, collaborator string, option *github.RepositoryAddCollaboratorOptions) {
	repo, err := client.Repositories.AddCollaborator(ctx, owner, repoName, collaborator, option)
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
	fmt.Println(repo.Rate)
}

func main() {

	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "401e22c56b9515e10677fe285a0442effdbd6b12"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	// createRepository(context, client, "repoName", true)
	//fetchrepo
	//fetchRepository(context, client)

	// Add Collaborator

	addCollaborator(context, client, "Guarav4u", "foo", "PunchhGaurav", &github.RepositoryAddCollaboratorOptions{Permission: "admin"})
}
