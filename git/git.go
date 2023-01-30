package git

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type Repo struct {
	config GitArgs
	r      *git.Repository
	w      *git.Worktree
	path   string
}

func New(args GitArgs) (*Repo, error) {
	r := &Repo{
		config: args,
	}
	if err := r.init(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Repo) init() error {
	repoPath, err := os.MkdirTemp("", "CodeCallScanner")
	if err != nil {
		return err
	}
	r.path = repoPath
	return nil
}

func (r *Repo) Clone() error {
	var err error
	r.r, err = git.PlainClone(r.path, false, &git.CloneOptions{
		URL:  r.config.URL,
		Auth: r.config.Auth,
	})
	if err != nil {
		return err
	}
	r.w, err = r.r.Worktree()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) CheckOut() error {
	if err := r.w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(r.config.Ref),
	}); err != nil {
		return err
	}
	return nil
}

func (r *Repo) Destroy() error {
	return os.RemoveAll(r.path)
}

func (r *Repo) Path() string {
	return r.path
}

type GitArgs struct {
	Ref               string
	URL               string
	Auth              transport.AuthMethod
	RecurseSubmodules bool
}
