package main

import (
	"context"
	"net/http"
	"time"
	"github.com/lib/pq" 
)

func (app *application) listCourses(w http.ResponseWriter, r *http.Request) {
	query := `SELECT code, title, credits, enrolled, instructors 
			  FROM courses 
			  ORDER BY title`
	
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	rows, err := app.db.QueryContext(ctx, query)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var c Course
		err := rows.Scan(&c.Code, &c.Title, &c.Credits, &c.Enrolled, pq.Array(&c.Instructors))
		if err != nil {
			app.serverError(w, err)
			return
		}
		courses = append(courses, c)
	}

	if err = rows.Err(); err != nil {
		app.serverError(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"courses": courses}, nil)
}