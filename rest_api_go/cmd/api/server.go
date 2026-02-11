package main

import (
	//"encoding/json"
	"crypto/tls"
	"fmt"
	"time"

	//"strings"
	//"io"
	"log"
	"net/http"
	"restapi/internal/api/middlewares"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		/*fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path,"/teachers/")
		userID := strings.TrimSuffix(path,"/")
		fmt.Println("User ID:",userID)
		fmt.Println("Query Params: ",r.URL.Query())
		queryParams := r.URL.Query()
		key:= queryParams.Get("key")
		sortBy:= queryParams.Get("sortby")
		sortOrder:= queryParams.Get("sortorder")

		if sortOrder == ""{
			sortOrder = "DESC"
		}
		fmt.Printf("SortBy: %v, SortOrder: %v, Key: %v ",sortBy,sortOrder,key)*/

		w.Write([]byte("Hello GET Method on Teachers Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Teachers Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
		return
	default:
		w.Write([]byte("Hello Teachers Route"))
		return
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Students Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Students Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Students Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Students Route"))
		return
	default:
		w.Write([]byte("Hello Students Route"))
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Execs Route"))
		return
	case http.MethodPost:
		fmt.Println("Query: ", r.URL.Query())
		fmt.Println("Name: ", r.URL.Query().Get("name"))

		//Parse form data
		err := r.ParseForm()
		if err != nil {
			return
		}
		fmt.Println("Form: ", r.Form)
		w.Write([]byte("Hello POST Method on Execs Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Execs Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Execs Route"))
		return
	default:
		w.Write([]byte("Hello Execs Route"))
		return
	}
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	rl := middlewares.NewRateLimiter(5, time.Minute)

	hppOptions := middlewares.HPPOptions{
		CheckBody:                   true,
		CheckQuery:                  true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		WhiteList:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}
	//Create custom server
	server := &http.Server{
		Addr:      port,
		Handler:   middlewares.HPP(hppOptions)(rl.Middleware(middlewares.Compression(middlewares.ResponseTimeMiddleware(middlewares.SecurityHeaders(middlewares.Cors(mux)))))),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}

}
