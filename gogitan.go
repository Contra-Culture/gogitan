package gogitan

import (
	"fmt"

	git "github.com/go-git/go-git/v5"
)

// dumb
func Check() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	_, err = fmt.Printf("\n\trepo: %#v", repo)
	if err != nil {
		panic(err)
	}
}
