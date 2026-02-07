package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"crypto/x509"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool{
	clientCAs := x509.NewCertPool()
	caCert,err :=os.ReadFile("cert.pem")
	if err!=nil {
		log.Fatalln("Could not load client CA: ", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
	
}
func main() {
	http.HandleFunc("/orders",func(w http.ResponseWriter, r *http.Request){
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling orders")
	})
	http.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request){
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users")
	})
	port := 3000

	//load the TLS certificate and key
	cert := "cert.pem"
	key := "key.pem"

	//config TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert, //enforce mTLS
		ClientCAs: loadClientCAs(),
	}

	//Create a custom server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	//enabling http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port: ", port)
	
	err:= server.ListenAndServeTLS(cert,key)
	if err!= nil{
		log.Fatal("Error starting server: ", err)
	}
	
	
	//HTTP 1.1 Server without TLS
	//err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	//if err!= nil{
	//	log.Fatal("Error starting server: ", err)
	//}


}

func logRequestDetails(r *http.Request){
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP Version: ", httpVersion)

	if r.TLS != nil{
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS version: ",tlsVersion)
	}else{
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16)string{
	switch version{
		case tls.VersionTLS10: return "TLS 1.0"
		case tls.VersionTLS11: return "TLS 1.1"
		case tls.VersionTLS12: return "TLS 1.2"
		case tls.VersionTLS13: return "TLS 1.3"
		default: return fmt.Sprintf("Unknown TLS Version: %d", version)
	}
}