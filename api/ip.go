package api

import (
	"context"
	"dnsupdater/db"
	"dnsupdater/models"
	"dnsupdater/views"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type PublicIpLogResponse struct {
	Result []models.PublicIpLog `json:"data"`
	Status int                  `json:"status"`
}

type PublicIpLogChartJsDatasetResponse struct {
	Data    []string `json:"data"`
	Label   string   `json:"label" default:"Public IP Address"`
	Fill    bool     `json:"fill" default:"false"`
	Tension float64  `json:"tension" default:"0.1"`
	Stepped bool     `json:"stepped" default:"false"`
}

type PublicIpLogChartJsDataResponse struct {
	Dataset []PublicIpLogChartJsDatasetResponse `json:"datasets"`
	Labels  []string                            `json:"labels"`
}

type PublicIpLogChartJsResponse struct {
	Data PublicIpLogChartJsDataResponse `json:"data"`
}

func uniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us

}

func ApiIpPublicGet(w http.ResponseWriter, r *http.Request, sql *gorm.DB) {
	query := r.URL.Query()
	format := strings.ToLower(query.Get("format"))

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}

	reverse, err := strconv.ParseBool(query.Get("reverse"))
	if err != nil {
		reverse = false
	}

	rows := db.PublicIpLogAll(sql, limit, reverse)
	res := PublicIpLogResponse{
		Result: rows,
		Status: 200,
	}

	w.WriteHeader(http.StatusOK)

	if format == "" || format == "json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

	if format == "chartjs" {
		values := []string{}
		labels := []string{}

		for _, row := range rows {
			labels = append(labels, row.CreatedAt.Format("01-Jan 15:04:05"))
			values = append(values, row.Ip)
		}

		resp := PublicIpLogChartJsResponse{
			Data: PublicIpLogChartJsDataResponse{
				Dataset: []PublicIpLogChartJsDatasetResponse{{
					Data: uniqueNonEmptyElementsOf(values),
				}},
				Labels: labels,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

	if format == "html" {
		w.Header().Set("Content-Type", "text/html")
		component := views.Api_ip_public_get(rows)
		component.Render(context.Background(), w)
	}

}

func ApiIpUpdateGet(w http.ResponseWriter, r *http.Request, sql *gorm.DB) {
	query := r.URL.Query()
	format := strings.ToLower(query.Get("format"))

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}

	reverse, err := strconv.ParseBool(query.Get("reverse"))
	if err != nil {
		reverse = false
	}

	rows := db.PublicIpLogUpdateAll(sql, limit, reverse)
	res := PublicIpLogResponse{
		Result: rows,
		Status: 200,
	}

	w.WriteHeader(http.StatusOK)

	if format == "" || format == "json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

	if format == "html" {
		w.Header().Set("Content-Type", "text/html")
		component := views.Api_ip_public_get(rows)
		component.Render(context.Background(), w)
	}
}
