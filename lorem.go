package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func ifError(err error) {
	// Nil instead of null. Neato!
	if err != nil {
		fmt.Printf("Hm, error occurred, exiting: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	// Hard coding these values but could be a CLI value instead
	url := "git@github.com:ossiesf/dummy_files.git"
	dir := "~/Downloads/dummy_files"
	ssh_key := os.Getenv("SSH_KEY_PATH")

	// Check if we can get the ssh key file
	_, err := os.Stat(ssh_key)
	ifError(err)

	// Create a public key from file
	publicKey, err := ssh.NewPublicKeysFromFile("git", ssh_key, "")
	ifError(err)

	fmt.Println("Downloading some very important stuff rn...")
	fmt.Printf("URL: %s,   dir: %s	key_file: %s\n", url, dir, ssh_key)

	// Per the examples in go-git repo this will give us an object for the git clone
	// Similar to Python we can use _ as a sort of placeholder 'throwaway' variable we won't use, since go dislikes that
	_, err = git.PlainClone(dir, false, &git.CloneOptions{URL: url, Auth: publicKey, Progress: os.Stdout})
	ifError(err)
}
