package main

import (
	"github.com/alameddinc/temporal-workflow-golang-example/signals"
	"github.com/alameddinc/temporal-workflow-golang-example/starters"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Client Starting...")
	r := mux.NewRouter()
	r.HandleFunc("/about", GetAbout)
	r.HandleFunc("/start-workflow", StartWorkflow)
	r.HandleFunc("/send-signal/{workflowID}", SendSignal)
	log.Fatal(http.ListenAndServe(":3310", r))
}

func SendSignal(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	workflowId := vars["workflowID"]
	if err := signals.SendPaymentSignal(workflowId, true); err != nil {
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write([]byte("Sent it"))
	return
}

func StartWorkflow(writer http.ResponseWriter, request *http.Request) {
	workflowUUID, err := uuid.NewUUID()
	if err != nil {
		writer.Write([]byte(err.Error()))
		return
	}
	starters.StartWorkflowFunc(workflowUUID.String(), "alameddin")
	writer.Write([]byte("OK"))
}

func GetAbout(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("It's running..."))
}
