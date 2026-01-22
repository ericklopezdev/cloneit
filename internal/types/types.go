package types

import (
	"github.com/charmbracelet/bubbles/list"
)

type Repo struct {
	Name        string  `json:"name"`
	SshUrl      string  `json:"sshUrl"`
	Description *string `json:"description"`
}

type RepoItem struct {
	repo Repo
}

func (r RepoItem) Title() string { return r.repo.Name }
func (r RepoItem) Description() string {
	if r.repo.Description != nil {
		return *r.repo.Description
	}
	return ""
}
func (r RepoItem) FilterValue() string { return r.repo.Name }

func (r RepoItem) GetRepo() Repo {
	return r.repo
}

func NewRepoItem(repo Repo) list.Item {
	return RepoItem{repo: repo}
}
