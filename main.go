package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	r, err := git.PlainOpen("./")
	if err != nil {
		log.Fatal(err)
	}

	headRef, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(headRef)

	refs, _ := r.References()
	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference {
			fmt.Println(ref)
		}

		return nil
	})
}
