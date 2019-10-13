# cookiecutter-cloud-run-go
Cookiecutter template for Cloud Run project

The following template setups up dockerfiles as well as build scripts in order to automatically deploy this application to project. 

# Usage

In order to use the following, ensure that the following is installed.

- `gcloud`: Setup your gcloud cli properly by setting it to the right project
- `yq`: yq tool is meant to understand yaml files/content on cli more easily. 
- `cookiecutter`: Templating tool. Install it on your development machine and have it consume this repo

This template also assumes that you are using 

The main properties that needs to be set is:

- `golang_mod_name`: Golang module project name. `app_name` and `mod_name` would be auto generated from this
- `type`: Sets up Cloud Run project such that it consumes Google Pubsub messages (mainly for async tasks)
- `use_secret`: Provides user to enable deploy cloud run with secrets via cloud build. Configuration of key ring is not provided in the makefile - utilize the Google Cloud UI for that
- `use_gcs`: Provides some sample code to integrate gcs into 

# Features

The following is the features that this template supports:

- Able to create different template outputs: either http based or msg based.
- Users can get to choose whether to have projects that have secrets included or not
- Users can get to decide to include a basic GCS integration