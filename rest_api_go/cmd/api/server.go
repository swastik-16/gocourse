package main

import (
	//"encoding/json"
	"fmt"
	"strings"
	//"io"
	"log"
	"net/http"
)
type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request){
	//teachers/1 -> id 1 of teachers
	fmt.Println(r.Method)
	switch r.Method{
	case http.MethodGet :
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path,"/teachers/")
		userID := strings.TrimSuffix(path,"/")
		fmt.Println("User ID:",userID)
		w.Write([]byte("Hello GET Method on Teachers Route"))
		//fmt.Println("Hello GET Method on Teachers Route")
		return
	case http.MethodPost :
		w.Write([]byte("Hello POST Method on Teachers Route"))
		fmt.Println("Hello POST Method on Teachers Route")
		return
	case http.MethodPut :
			w.Write([]byte("Hello PUT Method on Teachers Route"))
			fmt.Println("Hello PUT Method on Teachers Route")
			return
		case http.MethodDelete :
			w.Write([]byte("Hello DELETE Method on Teachers Route"))
			fmt.Println("Hello DELETE Method on Teachers Route")
			return
		default:
			w.Write([]byte("Hello Teachers Route"))
			fmt.Println("Hello Teachers Route")
			return
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet :
		w.Write([]byte("Hello GET Method on Students Route"))
		fmt.Println("Hello GET Method on Students Route")
		return
	case http.MethodPost :
		w.Write([]byte("Hello POST Method on Students Route"))
		fmt.Println("Hello POST Method on Students Route")
		return
	case http.MethodPut :
		w.Write([]byte("Hello PUT Method on Students Route"))
		fmt.Println("Hello PUT Method on Students Route")
		return
	case http.MethodDelete :
		w.Write([]byte("Hello DELETE Method on Students Route"))
		fmt.Println("Hello DELETE Method on Students Route")
		return
	default:
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet :
		w.Write([]byte("Hello GET Method on Execs Route"))
		fmt.Println("Hello GET Method on Execs Route")
		return
	case http.MethodPost :
		w.Write([]byte("Hello POST Method on Execs Route"))
		fmt.Println("Hello POST Method on Execs Route")
		return
	case http.MethodPut :
		w.Write([]byte("Hello PUT Method on Execs Route"))
		fmt.Println("Hello PUT Method on Execs Route")
		return
	case http.MethodDelete :
		w.Write([]byte("Hello DELETE Method on Execs Route"))
		fmt.Println("Hello DELETE Method on Execs Route")
		return
	default:
		w.Write([]byte("Hello Execs Route"))
		fmt.Println("Hello Execs Route")
		return
	}
}

func main() {
	port := ":3000"

	http.HandleFunc("/",rootHandler)

	http.HandleFunc("/teachers/",teachersHandler)
	http.HandleFunc("/students/",studentsHandler)
	http.HandleFunc("/execs/",execsHandler)

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port,nil)
	if err!=nil{
		log.Fatalln("Error starting the server:",err)
	}

}