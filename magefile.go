//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	imageName    = "init-full-stack"
	imageVersion = "0.1.0"
)

// Run App Locally
func RunLocal() error {
	return sh.RunV("go", "run", "cmd/app/main.go")
}

// Remove generated files
func Clean() error {
	var err error

	if err = sh.RunV("sh", "-c", "find ./bin -mindepth 1 -delete"); err != nil {
		return err
	}
	fmt.Println("cleaned bin")

	return nil
}

// Build app binaries
func BuildBin() error {
	mg.Deps(Clean)

	var err error

	if err = sh.RunV("sh", "-c", "GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64/app cmd/app/main.go"); err != nil {
		return err
	}
	fmt.Println("built binary: ./bin/linux-amd64/app")

	if err = sh.RunV("sh", "-c", "GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin-arm64/app cmd/app/main.go"); err != nil {
		return err
	}
	fmt.Println("built binary: ./bin/darwin-arm64/app")

	return nil
}

// Build docker image
func BuildImage() error {
	mg.Deps(BuildBin)

	var (
		err error
		cmd string
	)

	cmd = fmt.Sprintf("docker build --platform linux/amd64 -t %s:latest -t %s:%s .", imageName, imageName, imageVersion)
	if err = sh.RunV("sh", "-c", cmd); err != nil {
		return err
	}
	fmt.Printf("built image: %s:%s\n", imageName, imageVersion)

	return nil
}

// Push Image
// 1. Login aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 12345.dkr.ecr.us-east-1.amazonaws.com
// 2. docker tag <image_id> <account_id>.dkr.ecr.<region_name>.amazonaws.com/<repository_name>:<tag>
// 3. docker push <account_id>.dkr.ecr.<region_name>.amazonaws.com/<repository_name>:<tag>
