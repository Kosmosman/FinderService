package wbserver

import (
	"github.com/Kosmosman/service/types"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
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

	enableCORS(&w)
	if _, err := w.Write([]byte("Hello Bro!\n")); err != nil {
		log.Fatal(err)
	}
}

func (s *ServerAPI) ordersHandler(w http.ResponseWriter, r *http.Request) {
	_, filename, _, _ := runtime.Caller(0)
	templatePath := filepath.Dir(filename) + "/templates/orders.html"
	tmpl := template.Must(template.ParseFiles(templatePath))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}

	uid := r.URL.Query().Get("uid")
	if uid == "" {
		return
	}
	if _, ok := s.Cache.Data[uid]; !ok {
		if _, err := w.Write([]byte("Order not found!\n")); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := w.Write(s.Cache.Data[uid]); err != nil {
			log.Fatal(err)
		}
	}
}
