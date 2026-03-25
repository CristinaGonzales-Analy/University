package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Student struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
}

type Course struct {
	Code        string   `json:"code"`
	Title       string   `json:"title"`
	Credits     int      `json:"credits"`
	Enrolled    int      `json:"enrolled"`
	Instructors []string `json:"instructors"`
}

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, values := range headers {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := int64(1_048_576)
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	err := app.db.PingContext(ctx)
	dbStatus := "reachable"
	if err != nil {
		dbStatus = "unreachable: " + err.Error()
	}

	app.writeJSON(w, http.StatusOK, envelope{
		"status":    "available",
		"database":  dbStatus,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}, nil)
}

func (app *application) echoHeaders(w http.ResponseWriter, r *http.Request) {
	received := make(map[string]string)
	for name, values := range r.Header {
		received[name] = strings.Join(values, ", ")
	}
	app.writeJSON(w, http.StatusOK, envelope{"headers_received": received}, nil)
}