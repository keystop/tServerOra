package handlers

import (
	"fmt"
	"io"
	"net/http"

	"tServerOra/internal/models"
)

var Repo models.Repository

func HandlerTCPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(models.UserKey).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// text, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	cTC := new(models.CardTC)
	cTC.DriverName = r.FormValue("driver_name")
	cTC.ModelTC = r.FormValue("model_tc")

	// err = json.Unmarshal(text, cTC)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	cTC.UserID = userID
	fmt.Println(cTC)
	Repo.SaveCard(ctx, cTC)

	//	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Zaebok")
}

func HandlerCheckDBConnect(w http.ResponseWriter, r *http.Request) {
	if err := Repo.CheckDBConnection(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Connection: OK")
}

func HandlerHi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hi my friend")
}

func NewHandlers(repo models.Repository) {
	Repo = repo
}
