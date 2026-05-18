package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rainbowmga/timetravel/entity"
)

// POST /records/{id}
// if the record exists, the record is updated.
// if the record doesn't exist, the record is created.
func (a *API) PostRecords(
	w http.ResponseWriter,
	r *http.Request,
) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload map[string]string

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// check if record exists
	_, err = a.records.GetRecord(ctx, id)

	// RECORD DOES NOT EXIST -> CREATE
	if err != nil {

		record := entity.Record{
			ID:   id,
			Data: payload,
		}

		err = a.records.CreateRecord(ctx, record)
		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		json.NewEncoder(w).Encode(record)
		return
	}

	// RECORD EXISTS -> UPDATE
	updates := map[string]*string{}

	for key, value := range payload {
		v := value
		updates[key] = &v
	}

	record, err := a.records.UpdateRecord(
		ctx,
		id,
		updates,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(record)
}
