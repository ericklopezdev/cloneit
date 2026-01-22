package github

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"cloneit/internal/types"
)

// CheckAuth checks if GitHub CLI is authenticated
func CheckAuth() error {
	cmd := exec.Command("gh", "auth", "status")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("GitHub CLI is not authenticated. Please run 'gh auth login' to authenticate")
	}
	return nil
}

// ListRepos fetches the list of repositories using GitHub CLI
func ListRepos() ([]types.Repo, error) {
	cmd := exec.Command("gh", "repo", "list", "--limit", "300", "--json", "name,sshUrl,description")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running gh repo list: %w", err)
	}

	var repos []types.Repo
	if err := json.Unmarshal(output, &repos); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return repos, nil
}
