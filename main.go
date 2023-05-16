package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	repoURL := "https://github.com/zerok/mre-go-git-win-permissions.git"

	os.RemoveAll("tmp-gogit")
	os.RemoveAll("tmp-git")

	_, err := git.PlainClone("tmp-gogit", false, &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.NewBranchReferenceName("main"),
		Depth:         50,
	})
	if err != nil {
		log.Fatalf("Failed to clone repository into tmp folder: %s", err)
	}

	if err := exec.Command("git", "clone", repoURL, "tmp-git").Run(); err != nil {
		log.Fatalf("Failed to clone repository using native client: %s", err)
	}

	log.SetPrefix("go-git: ")
	cmd := exec.Command("git", "status", "--short")
	cmd.Dir = "tmp-gogit"
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run git status in tmp folder: %s", err)
	}

	if len(output) == 0 {
		log.Printf("Status after checkout UNCHANGED")
	} else {
		log.Printf("Status after checkout CHANGED:\n%s", string(output))
	}

	log.SetPrefix("git:    ")
	cmd = exec.Command("git", "status", "--short")
	cmd.Dir = "tmp-git"
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run git status in tmp folder: %s", err)
	}

	if len(output) == 0 {
		log.Printf("Status after checkout UNCHANGED")
	} else {
		log.Printf("Status after checkout CHANGED:\n%s", string(output))
	}
}
