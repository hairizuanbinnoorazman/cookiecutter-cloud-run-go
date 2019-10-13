{% if cookiecutter.type == "msg" %}APP_NAME?="{{cookiecutter.app_name}}"
TOPIC_NAME?="{{cookiecutter.app_name}}"
SUB_NAME?="{{cookiecutter.app_name}}"
SERVICE_URL=$$(gcloud beta run services describe {{cookiecutter.app_name}} --region=asia-northeast1 --platform=managed | yq r - status.url)
DATA=$$(printf '{"sample": "data"}' | base64)

# The following commands are for handling when the application is to be paired with Google Pubsub
# Before running this command, ensure that PROJECT_ID environment variable is set
pubsub-prep:
	gcloud beta run services add-iam-policy-binding $(APP_NAME) \
   		--member=serviceAccount:cloud-run-pubsub-invoker@${PROJECT_ID}.iam.gserviceaccount.com \
		--region=asia-northeast1 \
   		--role=roles/run.invoker
	gcloud pubsub topics create $(TOPIC_NAME)
	gcloud beta pubsub subscriptions create $(SUB_NAME) --topic $(TOPIC_NAME) \
		--push-endpoint=$(SERVICE_URL) \
		--push-auth-service-account=cloud-run-pubsub-invoker@${PROJECT_ID}.iam.gserviceaccount.com

sendmsg:
	gcloud pubsub topics publish $(TOPIC_NAME) --message "sample message"

sampledata:
	echo $(DATA)

local-http:
	curl localhost:8080 -X POST -H "Content-Type: application/json" -d '{"message": {"data": "eyJzYW1wbGUiOiAiZGF0YSJ9"}}'
{% endif %}
setup:
	@echo "Run setup steps"
	{% if cookiecutter.use_gcs == "yes" %}~/go/bin/goimports main.go > main_temp.go && mv main_temp.go main.go{% endif %}
	{% if cookiecutter.type == "http" and cookiecutter.use_gcs == "yes" %}~/go/bin/goimports http_handlers.go > http_handlers_temp.go && mv http_handlers_temp.go http_handlers.go{% endif %}
	{% if cookiecutter.type == "msg" and cookiecutter.use_gcs == "yes" %}~/go/bin/goimports msg_handlers.go > msg_handlers_temp.go && mv msg_handlers_temp.go msg_handlers.go{% endif %}
	