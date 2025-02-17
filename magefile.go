//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
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

	env, err := loadEnv()
	if err != nil {
		return err
	}

	var cmd string

	cmd = fmt.Sprintf("docker build --platform linux/amd64 -t %s:latest -t %s:%s .", env.ECRAppRepo, env.ECRAppRepo, env.AppVersion)
	if err = sh.RunV("sh", "-c", cmd); err != nil {
		return err
	}
	fmt.Printf("Built image: %s:%s\n", env.ECRAppRepo, env.AppVersion)

	return nil
}

// Push docker image to AWS ECR
func PushImage() error {
	mg.Deps(BuildImage)

	env, err := loadEnv()
	if err != nil {
		return err
	}

	fmt.Printf("Attempting to push image version %s to ECR...\n", env.AppVersion)

	localImageTag := fmt.Sprintf("%s:%s", env.ECRAppRepo, env.AppVersion)
	ecrInstance := fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", env.AWSAccountID, env.AWSRegion)
	newEcrImageTag := fmt.Sprintf("%s/%s:%s", ecrInstance, env.ECRAppRepo, env.AppVersion)

	var cmd string

	cmd = fmt.Sprintf("aws --profile %s ecr get-login-password --region %s | docker login --username AWS --password-stdin %s",
		env.AWSProfile,
		env.AWSRegion,
		ecrInstance,
	)
	if err = sh.RunV("sh", "-c", cmd); err != nil {
		return err
	}

	cmd = fmt.Sprintf("docker tag %s %s", localImageTag, newEcrImageTag)
	if err = sh.RunV("sh", "-c", cmd); err != nil {
		return err
	}

	cmd = fmt.Sprintf("docker push %s", newEcrImageTag)
	if err = sh.RunV("sh", "-c", cmd); err != nil {
		return err
	}

	fmt.Printf("Successfully pushed %q to ECR\n", newEcrImageTag)

	return nil
}

func loadEnv() (*MageEnv, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	env := &MageEnv{
		AWSAccountID: os.Getenv("APP_AWS_ACCOUNT_ID"),
		AWSProfile:   os.Getenv("APP_AWS_PROFILE"),
		AWSRegion:    os.Getenv("APP_AWS_REGION"),
		ECRAppRepo:   os.Getenv("APP_ECR_REPO"),
		AppVersion:   getAppVersion(),
	}

	// Check for empty strings in the MageEnv struct
	for _, value := range []string{
		env.AWSAccountID,
		env.AWSProfile,
		env.AWSRegion,
		env.ECRAppRepo,
		env.AppVersion,
	} {
		if value == "" {
			return nil, fmt.Errorf("one or more environment variables are empty")
		}
	}

	return env, nil
}

func getAppVersion() string {
	versionBytes, err := os.ReadFile("VERSION")
	if err != nil {
		return ""
	}
	return string(versionBytes)
}

type MageEnv struct {
	AWSAccountID string
	AWSProfile   string
	AWSRegion    string
	ECRAppRepo   string
	AppVersion   string
}
