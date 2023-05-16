package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	os.RemoveAll("tmp")
	_, err := git.PlainClone("tmp", false, &git.CloneOptions{
		URL:           "https://github.com/zerok/mre-go-git-win-permissions.git",
		ReferenceName: plumbing.NewBranchReferenceName("main"),
		Depth:         50,
	})
	if err != nil {
		log.Fatalf("Failed to clone repository into tmp folder: %s", err)
	}

	cmd := exec.Command("git", "status", "--short")
	cmd.Dir = "tmp"
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run git status in tmp folder: %s", err)
	}

	if len(output) == 0 {
		log.Printf("✅ Status after checkout UNCHANGED")
		return
	}
	log.Fatalf("🚨 Status after checkout CHANGED: \n\n%s", string(output))
}
