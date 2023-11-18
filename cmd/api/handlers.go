package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status: "active",
		Message: "Maps service up and running",
		Version: "0.0.0",
	}

	out, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error json marshalling: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) GetCoordinates(w http.ResponseWriter, r *http.Request) {
	var coordinates []models.Coordinate

	stockholm := models.Coordinate {
		Lng: 18.055175,
		Lat: 59.327490,
	}

	getryggen := models.Coordinate {
		Lng: 12.311031,
		Lat: 63.183692,
	}

	coordinates = append(coordinates, getryggen, stockholm)

	out, err := json.Marshal(coordinates)

	if err != nil {
		fmt.Println("Error parsing coordinates: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}