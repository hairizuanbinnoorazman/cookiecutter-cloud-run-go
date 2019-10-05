package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type exampleHandler struct {
	logger *logrus.Logger
}

type PubsubMsg struct {
	Message      PubsubMsgDetails `json:"message"`
	Subscription string           `json:"subscription"`
}

type PubsubMsgDetails struct {
	Data        string `json:"data"`
	MessageID   string `json:"messageId"`
	PublishTime string `json:"publishTime"`
}

func (h exampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start Example Handler")
	defer h.logger.Info("End Example Handler")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Errorf("Error in reading body of request")
		w.WriteHeader(500)
		return
	}
	message := PubsubMsg{}
	json.Unmarshal(data, &message)
	h.logger.Infof("%+v", message)
	decodedMsg, _ := base64.StdEncoding.DecodeString(message.Message.Data)
	h.logger.Infof("%+v", string(decodedMsg))
	w.WriteHeader(200)
	w.Write([]byte("Example"))
}
