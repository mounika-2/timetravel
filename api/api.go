package api

import (
	"github.com/gorilla/mux"
	"github.com/rainbowmga/timetravel/service"
)

type API struct {
	records service.RecordService
	gemini  *service.GeminiService
}

func NewAPI(
	records service.RecordService,
	gemini *service.GeminiService,
) *API {
	return &API{
		records: records,
		gemini:  gemini,
	}
}

// generates all api routes
func (a *API) CreateRoutes(routes *mux.Router) {
	routes.Path("/records/{id}/analyze").HandlerFunc(a.AnalyzeRecordChanges).Methods("POST")
	routes.Path("/records/{id}").HandlerFunc(a.GetRecords).Methods("GET")
	routes.Path("/records/{id}/versions").HandlerFunc(a.ListRecordVersions).Methods("GET")
	routes.Path("/records/{id}").HandlerFunc(a.PostRecords).Methods("POST")
	routes.Path("/records/{id}/versions/{version}").HandlerFunc(a.GetRecordVersion).Methods("GET")
}
