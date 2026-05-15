package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *API) GetRecordVersion(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	params := mux.Vars(r)

	idNumber, err := strconv.ParseInt(
		params["id"],
		10,
		32,
	)

	if err != nil || idNumber <= 0 {
		err := writeError(
			w,
			"invalid id; id must be a positive number",
			http.StatusBadRequest,
		)
		logError(err)
		return
	}

	versionNumber, err := strconv.ParseInt(
		params["version"],
		10,
		32,
	)

	if err != nil || versionNumber <= 0 {
		err := writeError(
			w,
			"invalid version",
			http.StatusBadRequest,
		)
		logError(err)
		return
	}

	record, err := a.records.GetRecordVersion(
		ctx,
		int(idNumber),
		int(versionNumber),
	)

	if err != nil {
		err := writeError(
			w,
			fmt.Sprintf(
				"record %v version %v does not exist",
				idNumber,
				versionNumber,
			),
			http.StatusBadRequest,
		)

		logError(err)
		return
	}

	err = writeJSON(w, record, http.StatusOK)
	logError(err)
}

func (a *API) ListRecordVersions(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	params := mux.Vars(r)

	idNumber, err := strconv.ParseInt(
		params["id"],
		10,
		32,
	)

	if err != nil || idNumber <= 0 {
		err := writeError(
			w,
			"invalid id; id must be a positive number",
			http.StatusBadRequest,
		)

		logError(err)
		return
	}

	versions, err := a.records.ListRecordVersions(
		ctx,
		int(idNumber),
	)

	if err != nil {
		err := writeError(
			w,
			"could not retrieve versions",
			http.StatusInternalServerError,
		)

		logError(err)
		return
	}

	err = writeJSON(w, versions, http.StatusOK)
	logError(err)
}
