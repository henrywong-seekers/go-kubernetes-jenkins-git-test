package main

import (
	"fmt"
	"log"
	"regexp"

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

	re, _ := regexp.Compile("^.*/(.*)")

	refs, _ := r.References()
	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && ref.Hash() == headRef.Hash() && ref.Name() != "HEAD" {
			fmt.Println(re.FindStringSubmatch(ref.Name().String())[1])
		}

		return nil
	})
}
