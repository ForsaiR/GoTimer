package controller

import (
	"awesomeProject1/restapiserver/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.staticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

//Обработчик запросов
func QueryProcessor() {

	router := mux.NewRouter()
	router.HandleFunc("/", startPage)
	router.HandleFunc("/api/counter/value", getCounterValue).Methods("GET")
	router.HandleFunc("/api/counter/reset", putCounterValue).Methods("PUT")

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         "192.168.88.246:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

//Стартовая страница
func startPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

//GET запрос для отправки значения счетчика
func getCounterValue(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]int{"count": service.GetCount()})
}

//PUT запрос для установки значения счетчика 0
func putCounterValue(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"status": service.ResetTimer()})
}
