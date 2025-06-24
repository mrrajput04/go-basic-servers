package handlers

import (
	"comic-store-apis/config"
	"comic-store-apis/model"
	"database/sql"
	"net/http"
	"strconv"
	"text/template"
)

// Parse HTML templates
var templates = template.Must(template.ParseFiles("view/appForm.html"))

func AppHandler(w http.ResponseWriter, r *http.Request) {
	db := config.DB

	if r.Method != http.MethodPost {
		templates.Execute(w, nil)
		return
	}

	//converting issueNumber and Price to int
	issueNumberConv, _ := strconv.Atoi(r.FormValue("issueNumber"))
	priceConv, _ := strconv.Atoi(r.FormValue("price"))

	// Extract form values and create a ComicInfo object
	comic := model.ComicInfo{
		Cid:         r.FormValue("cid"),
		Title:       r.FormValue("title"),
		IssueNumber: issueNumberConv,
		Collection:  r.FormValue("collection"),
		Publisher:   r.FormValue("publisher"),
		Description: r.FormValue("description"),
		Price:       priceConv,
		ImageURL:    r.FormValue("imageURL"),
	}

	// Perform different actions based on the submitted form button
	switch r.FormValue("submit") {
	case "Create":
		createComic(w, db, comic)
	case "Read":
		readComics(w, db)
	case "Update":
		updateComic(w, db, comic)
	case "Delete":
		deleteComic(w, db, comic)
	}
}

func readComics(w http.ResponseWriter, db *sql.DB) {
	// Query all records from the 'comicInfo' table
	rows, err := db.Query("select * from comicInfo")
	if err != nil {
		// Render an error message if the query fails
		templates.Execute(w, struct {
			Success bool
			Message string
		}{Success: false, Message: err.Error()})
		return
	}
	defer rows.Close()

	// Create bootstrap table data
	tableData := []struct {
		Cid         string
		Title       string
		IssueNumber int
		Collection  string
		Publisher   string
		Description string
		Price       int
		ImageURL    string
	}{}

	// Iterate through the query result and populate the tableData
	for rows.Next() {
		var c model.ComicInfo
		rows.Scan(&c.Cid, &c.Title, &c.IssueNumber, &c.Collection, &c.Publisher, &c.Description, &c.Price, &c.ImageURL)
		tableData = append(tableData, struct {
			Cid         string
			Title       string
			IssueNumber int
			Collection  string
			Publisher   string
			Description string
			Price       int
			ImageURL    string
		}{c.Cid, c.Title, c.IssueNumber, c.Collection, c.Publisher, c.Description, c.Price, c.ImageURL})
	}

	// Render the tableData in the template
	templates.Execute(w, struct {
		Success   bool
		Message   string
		TableData []struct {
			Cid         string
			Title       string
			IssueNumber int
			Collection  string
			Publisher   string
			Description string
			Price       int
			ImageURL    string
		}
	}{Success: true, Message: "Displaying all comics", TableData: tableData})
}

// createComic inserts a new comic record into the database
func createComic(w http.ResponseWriter, db *sql.DB, comic model.ComicInfo) {
	cid, _ := strconv.Atoi(comic.Cid)

	_, err := db.Exec("insert into comicInfo (cid, title, issueNumber, collection, publisher, description, price, imageURL) values(?,?,?,?,?,?,?,?)", cid, comic.Title, comic.IssueNumber, comic.Collection, comic.Publisher, comic.Description, comic.Price, comic.ImageURL)

	if err != nil {
		templates.Execute(w, struct {
			Success bool
			Message string
		}{Success: true, Message: err.Error()})
	} else {
		templates.Execute(w, struct {
			Success bool
			Message string
		}{Success: true, Message: "Comic Added to the database"})
	}
}

// updateComic updates an existing comic record in the database
func updateComic(w http.ResponseWriter, db *sql.DB, comic model.ComicInfo) {
	// Convert Cid to integer
	cid, _ := strconv.Atoi(comic.Cid)
	// Execute the SQL query to update the comic record
	result, err := db.Exec("update comicInfo set title=?, issueNumber=?, collection=?, publisher=?, description=?, price=?, imageURL=? where cid=?", comic.Title, comic.IssueNumber, comic.Collection, comic.Publisher, comic.Description, comic.Price, comic.ImageURL, cid)

	// Render success or error message in the template
	if err != nil {
		templates.Execute(w, struct {
			Success bool
			Message string
		}{Success: true, Message: err.Error()})
	} else {
		// Check the number of rows affected
		_, err := result.RowsAffected()
		if err != nil {
			templates.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: "comic not updated"})
		} else {
			templates.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: "comic  updated"})
		}
	}
}

// deleteComic deletes an existing comic record from the database
func deleteComic(w http.ResponseWriter, db *sql.DB, comic model.ComicInfo) {
	// Convert Cid to integer
	cid, _ := strconv.Atoi(comic.Cid)

	// Execute the SQL query to delete the comic record
	result, err := db.Exec("delete from comicInfo where cid=?", cid)

	if err != nil {
		templates.Execute(w, struct {
			Success bool
			Message string
		}{Success: true, Message: err.Error()})
	} else {
		// Check the number of rows affected
		_, err := result.RowsAffected()
		if err != nil {
			templates.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: "comic not deleted"})
		} else {
			templates.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: "comic  deleted successfully"})
		}
	}
}
