package main

import (
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
func queryProcessor() {
	hub := newHub()
	hub.startCounter()
	go hub.run()

	router := mux.NewRouter()

	router.HandleFunc("/", startPage)

	router.HandleFunc("/api/counter/start", func(w http.ResponseWriter, r *http.Request) {
		putStartCounter(hub, w, r)
	}).Methods("GET")

	router.HandleFunc("/api/counter/stop", func(w http.ResponseWriter, r *http.Request) {
		putStopCounter(hub, w, r)
	}).Methods("POST")

	router.HandleFunc("/api/counter/reset", func(w http.ResponseWriter, r *http.Request) {
		putResetCounter(hub, w, r)
	}).Methods("POST")

	router.HandleFunc("/api/counter/value", getCounterValue).Methods("GET")

	router.HandleFunc("/api/counter/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(hub, w, r)
	})

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         "192.168.88.246:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	queryProcessor()
}