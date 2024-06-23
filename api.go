package main

import (
	"log"
	"net/http"
)

type APIServer struct{
	addr string 
}

func NewAPIServer(addr string) *APIServer{
	return&APIServer{
		addr: addr  ,
	}

}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/login", loginHandler)

	router.HandleFunc("/api/posts/{id}/comments", func(w http.ResponseWriter,r *http.Request){
		id := r.PathValue("id")
		w.Write([]byte("id: "+id))
	})
	// router.HandleFunc("/api/posts", )
	router.HandleFunc("/api/posts/{id}", func(w http.ResponseWriter,r *http.Request){
		id := r.PathValue("id")
		w.Write([]byte("id: "+id))
	} )
	router.HandleFunc("/api/posts/{id}/dislike", func(w http.ResponseWriter,r *http.Request){
		id := r.PathValue("id")
		w.Write([]byte("id: "+id))
	} )
	router.HandleFunc("/api/posts/{id}/like",  func(w http.ResponseWriter,r *http.Request){
		id := r.PathValue("id")
		w.Write([]byte("id: "+id))
	})
	// router.HandleFunc("/api/categories", )
	// router.HandleFunc("signup", )

	router.HandleFunc("/", homeHandler)
	server := http.Server{
		Addr : s.addr,
		Handler: router,
	}

	log.Printf("server has started %s", s.addr)
	return server.ListenAndServe()

} 


// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		http.ServeFile(w, r, filepath.Join("pages", "login.html"))
// 		return
// 	}

// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		http.ServeFile(w, r, filepath.Join("pages", "index.html"))
// 		return
// 	}

// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// }
