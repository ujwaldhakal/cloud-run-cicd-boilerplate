package main

import (
	"os"
	"os/exec"

	"github.com/ujwaldhakal/go-gcp-docker-utils/docker"
	"github.com/ujwaldhakal/go-gcp-docker-utils/git"
	"github.com/ujwaldhakal/go-gcp-docker-utils/utils"
)

func main() {
	env := os.Args[1]
	credentialFilePath := getCredentialsFilePath(string(env))
	configJson := utils.ConvertTfConfigToJson(getTfVarFileName(string(env)))
	commitHash := git.GetCommitHash()
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentialFilePath)
	docker.Login(credentialFilePath)
	docker.Build(configJson.Project, configJson.Github_repo_name, commitHash)
	initTerraform()
	applyTerraform(git.GetCommitHash(), string(env))
}

func applyTerraform(commitHash string, env string) {
	cmd := exec.Command("terraform", "apply", "-var", "image_tag="+commitHash, "-var-file="+getTfVarFileName(env), "-auto-approve", "-lock=false")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic("error applying terraform")
	}
}

func initTerraform() {
	cmd := exec.Command("terraform",
		"init",
		"-backend-config",
		"bucket=tf-test-app")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic("error initializing terraform")
	}

}

func getTfVarFileName(env string) string {

	if env == "dev" {
		return "dev.tfvars"
	}

	if env == "production" {
		return "prof.tfvars"
	}
	panic("Please select correct environment only dev & production available at the moment")
}

func getCredentialsFilePath(env string) string {

	if env == "dev" {
		return "credentials/dev-cred.json"
	}

	if env == "production" {
		return "credentials/prod-cred.json"
	}

	panic("error on loading credentials")
}
