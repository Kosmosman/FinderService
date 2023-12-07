package wbserver

import (
	"github.com/Kosmosman/service/types"
	"log"
	"net/http"
)

type ServerAPI struct {
	Cache *types.Cache
}

func (s *ServerAPI) StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.homepageHandler)
	mux.HandleFunc("/orders", s.ordersHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (s *ServerAPI) homepageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(&w)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	enableCORS(&w) // Set CORS headers for the actual request
	if _, err := w.Write([]byte("Hello Bro!\n")); err != nil {
		log.Fatal(err)
	}
}

func (s *ServerAPI) ordersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(&w)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Unfollowed method", http.StatusMethodNotAllowed)
		return
	}
	uid := r.URL.Query().Get("uid")
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)
	if _, ok := s.Cache.Data[uid]; !ok {
		http.Error(w, "Uid not found", http.StatusNotFound)
		return
	}
	if _, err := w.Write(s.Cache.Data[uid]); err != nil {
		log.Fatal(err)
	}
}
