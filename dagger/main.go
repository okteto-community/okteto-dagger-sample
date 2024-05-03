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

import "context"

type OktetodoDagger struct{}

// Returns a container that has Okteto CLI with the correct context set
func (m *OktetodoDagger) SetContext(context string, token string) *Container {
	return dag.Container().
		From("okteto/okteto").
		WithEnvVariable("OKTETO_TOKEN", token).
		// WithEnvVariable("OKTETO_CONTEXT", token).
		WithExec([]string{"okteto", "ctx", "use", context})
}

func (m *OktetodoDagger) PreviewEnv(ctx context.Context,
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
	return endpointsOut, nil
}
