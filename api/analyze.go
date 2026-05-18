package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AnalyzeRequest struct {
	FromVersion int `json:"from_version"`
	ToVersion   int `json:"to_version"`
}

func (a *API) AnalyzeRecordChanges(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}

	var body AnalyzeRequest

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeError(w, "invalid request", http.StatusBadRequest)
		return
	}

	fromRecord, err := a.records.GetRecordVersion(
		ctx,
		idNumber,
		body.FromVersion,
	)

	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	toRecord, err := a.records.GetRecordVersion(
		ctx,
		idNumber,
		body.ToVersion,
	)

	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	analysis, err := a.gemini.AnalyzeChanges(
		fromRecord.Data,
		toRecord.Data,
	)

	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{
		"analysis": analysis,
	}, http.StatusOK)
}
