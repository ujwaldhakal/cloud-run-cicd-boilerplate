### Cloud Run CI/CD
This is basic demonstration of making CI/CD in Cloud Run with some stacks. This can be used for different kinds of projects which is containerized

#### Stacks Used
* Container Registry -: It will store all docker images to be used in Cloud Run
* Google Storage -: It will save our state of Terraform.
* Cloud Run -: Serverless platform where our app will be hosted
* Terraform -: It will help us to spin up Cloud Run instance and to create multiple working environments like staging and production
* Go -: It will help us to trigger Terraform command whenever we want to with Github Actions
* Github Actions -: It will help us as an entry point to trigger those Go command and Go will trigger terraform

### Usage
* Copy `cicd` folder into your project
* Replace all `dev.tfvars`, `credentials/dev-cred.json` with your actual credentials
* trigger command `./cicd/deploy dev  master` where dev is environment name and master is the branch name
* Replace github repo url at `deploy`

Visit https://medium.com/p/528ff7964e7/edit this link for more