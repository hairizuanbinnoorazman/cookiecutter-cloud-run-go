# cookiecutter-cloud-run-go
Cookiecutter template for Cloud Run project

The following template setups up dockerfiles as well as build scripts in order to automatically deploy this application to project. 

# Usage

In order to use the following, ensure that the following is installed.

- `gcloud`: Setup your gcloud cli properly by setting it to the right project
- `yq`: yq tool is meant to understand yaml files/content on cli more easily. 
- `cookiecutter`: Templating tool. Install it on your development machine and have it consume this repo

The main properties that needs to be set is:

- `golang_mod_name`: Golang module project name. `app_name` and `mod_name` would be auto generated from this
- `type`: Sets up Cloud Run project such that it consumes Google Pubsub messages (mainly for async tasks)