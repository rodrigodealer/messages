package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigodealer/users/mysql"
)

type HealthcheckServiceStatus struct {
	Working bool   `json:"working"`
	Service string `json:"service"`
}

type HealthcheckStatus struct {
	Status   string                      `json:"status"`
	Services []*HealthcheckServiceStatus `json:"services"`
}

func HealthcheckHandler(mysql mysql.MySQLConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		healthcheckStatus := &HealthcheckStatus{Status: "WORKING"}
		healthcheckStatus.Services = append(healthcheckStatus.Services, mysqlCheck(mysql))

		for _, service := range healthcheckStatus.Services {
			if !service.Working {
				healthcheckStatus.Status = "FAILED"
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		json.NewEncoder(w).Encode(healthcheckStatus)
	}
}

func mysqlCheck(mysql mysql.MySQLConn) *HealthcheckServiceStatus {
	var working, _ = mysql.Ping()
	return &HealthcheckServiceStatus{Working: working, Service: "MySQL"}
}
