// A generated module for OktetodoDagger functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"
)

type OktetodoDagger struct{}

// Define a struct to match the JSON structure
type Endpoint struct {
	URL     string `json:"url"`
	Private bool   `json:"private"`
}

// Returns a container that has Okteto CLI with the correct context set
// example usage:
// dagger call set-context --context=arsh.okteto.me --token=$OKTETO_TOKEN
func (m *OktetodoDagger) SetContext(context string, token string) *Container {
	return dag.Container().
		From("okteto/okteto").
		WithEnvVariable("OKTETO_TOKEN", token).
		// WithEnvVariable("OKTETO_CONTEXT", token).
		WithExec([]string{"okteto", "ctx", "use", context})
}

// Deploys a preview environment in the specified okteto context
// example usage:
// dagger call preview-deploy --repo=https://github.com/RinkiyaKeDad/okteto-dagger-sample --branch=name-change --pr=https://github.com/RinkiyaKeDad/okteto-dagger-sample/pull/1 --context=arsh.okteto.me --token=$OKTETO_TOKEN
func (m *OktetodoDagger) PreviewDeploy(ctx context.Context,
	// Repo to deploy
	repo string,
	// Branch to deploy
	branch string,
	// URL of the pull request to attach in the Okteto Dashboard
	pr string,
	// Okteto context to be used for deployment
	context string,
	// Token to be used to authenticate with the Okteto context
	token string) (string, error) {
	c := m.SetContext(context, token).WithExec([]string{
		"okteto", "preview", "deploy", "--branch", branch, "--sourceUrl", pr, "--repository", repo, "--wait", branch,
	}).WithExec([]string{
		"okteto", "preview", "endpoints", branch, "--output=json",
	})

	endpointsOut, err := c.Stdout(ctx)
	if err != nil {
		return "", err
	}

	// Variable to hold the parsed data
	var endpoints []Endpoint

	// Parse the JSON data into the slice of Endpoint structs
	err = json.Unmarshal([]byte(endpointsOut), &endpoints)
	if err != nil {
		log.Fatal(err)
	}

	// StringBuilder to hold all URLs
	var urlsBuilder strings.Builder

	// Iterate through the parsed data and append each URL to the StringBuilder
	for _, endpoint := range endpoints {
		urlsBuilder.WriteString(endpoint.URL + "\n")
	}

	// Get the string with all URLs
	allURLs := urlsBuilder.String()

	return allURLs, nil
}

// Destorys a preview environment at the specified okteto context
// example usage:
// dagger call preview-destroy --branch=name-change --context=arsh.okteto.me --token=$OKTETO_TOKEN
func (m *OktetodoDagger) PreviewDestroy(ctx context.Context,
	// Branch to deploy (to be used as the name for the preview env)
	branch string,
	// Okteto context to be used for deployment
	context string,
	// Token to be used to authenticate with the Okteto context
	token string) (string, error) {
	c := m.SetContext(context, token).WithExec([]string{
		"okteto", "preview", "destroy", branch,
	})
	destoryOut, err := c.Stdout(ctx)
	if err != nil {
		return "", err
	}
	return destoryOut, nil
}
