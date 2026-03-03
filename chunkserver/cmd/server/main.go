package main

import (
	"log"
	"net/http"
	"os"

	handler "mini-gfs/chunkserver/internal/http"
	"mini-gfs/chunkserver/internal/storage"
)

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	store := storage.New("./data_" + port)
	h := handler.NewHandler(store)

	http.HandleFunc("/chunk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			h.UploadChunk(w, r)
		case http.MethodGet:
			h.GetChunk(w, r)
		case http.MethodDelete:
			h.DeleteChunk(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})

	log.Println("Chunk server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}