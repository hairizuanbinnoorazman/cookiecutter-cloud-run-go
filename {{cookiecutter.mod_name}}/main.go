package main

import (
	"net/http"
	"os"
	"time"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ServiceName denotes name of service. In order to reduce confusion, try to make it similar to name on Cloud Run UI
var ServiceName = "{{cookiecutter.app_name}}"

// Version denotes version no of service. Change it as necessary
var Version = "v0.1.0"

func main() {
	logger := logrus.New()
	logger.Formatter = stackdriver.NewFormatter(
		stackdriver.WithService(ServiceName),
		stackdriver.WithVersion(Version),
	)
	logger.Level = logrus.InfoLevel
	logger.Info("Application Start Up")
	defer logger.Info("Application Ended")

	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "LOCAL"
	}
	logger.Infof("Application Mode: %v", mode)

	{% if cookiecutter.use_gcs == "yes" %}credJSON, err := ioutil.ReadFile("/example-cred.json")
	if err != nil {
		logger.Error("Unable to load example-cred.json file")
	}
	xClient, err := storage.NewClient(context.Background(), option.WithCredentialsJSON(credJSON))
	if err != nil {
		logger.Error("Unable to create storage client")
	}

	{% endif %}r := mux.NewRouter()
	{% if cookiecutter.use_gcs == "yes" %}r.Handle("/", exampleHandler{logger: logger, client: xClient})
	{% else %}r.Handle("/", exampleHandler{logger: logger}){% endif %}

	srv := http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(srv.ListenAndServe())
}
