package comments

import (
	"github.com/go-git/go-git/v5"
	git "github.com/go-git/go-git/v5"
)

// func GetRepo(path string) (*git.Repository, error) {
// 	repo, err := git.PlainOpen(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return repo, nil
// }

// func GetCommits(repo *git.Repository, opts *git.LogOptions) error {
// 	cIter, err := repo.Log(opts)
// 	if err != nil {
// 		return err
// 	}
// }
