package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"status_page/internal/handlers"
	"status_page/internal/out"
)

type handler struct {
}

// NewHandler Создание хендлера
func NewHandler() handlers.Handler {
	return &handler{}
}

// Register Обработчики маршрута
func (h *handler) Register(r *mux.Router) {
	r.HandleFunc("/", h.handleConnection)
}

// handleConnection Возврвт ответа
func (h *handler) handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res []byte
	res, _ = json.Marshal(out.GetResultData())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)

	if err != nil {
		log.Fatal(err)
	}
}
