package jobserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/concourse/atc/api/present"
	"github.com/concourse/atc/db"
	"github.com/google/jsonapi"
)

func (s *Server) CreateJobBuild(pipeline db.Pipeline) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.Session("create-job-build")

		jobName := r.FormValue(":job_name")

		job, found, err := pipeline.Job(jobName)
		if err != nil {
			logger.Error("failed-to-get-resource-types", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

			if !found {
			w.WriteHeader(http.StatusNotFound)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Job Not Found Error",
				Detail: fmt.Sprintf("Job with name '%s' not found.", jobName),
				Status: "404",
			}})
			return
		}

		if job.Config().DisableManualTrigger {
			w.WriteHeader(http.StatusConflict)
			return
		}

		scheduler := s.schedulerFactory.BuildScheduler(pipeline, s.externalURL, s.variablesFactory.NewVariables(pipeline.TeamName(), pipeline.Name()))

		resourceTypes, err := pipeline.ResourceTypes()
		if err != nil {
			logger.Error("failed-to-get-resource-types", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		versionedResourceTypes := resourceTypes.Deserialize()

		resources, err := pipeline.Resources()
		if err != nil {
			logger.Error("failed-to-get-resources", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		build, _, err := scheduler.TriggerImmediately(logger, job, resources, versionedResourceTypes)
		if err != nil {
			logger.Error("failed-to-trigger", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "failed to trigger: %s", err)
			return
		}

		json.NewEncoder(w).Encode(present.Build(build))
	})
}
