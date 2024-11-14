package main

import (
    "fmt"
    "net"
    "net/http"
    "os"
)

func getPublicIP() string {
    conn, _ := net.Dial("udp", "8.8.8.8:80")
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP.String()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    publicIP := getPublicIP()
    privateIP, _ := net.InterfaceAddrs()
    hostname, _ := os.Hostname()
    port := "8080"
    
    fmt.Fprintf(w, "<html><head><title>Server Info</title></head><body>")
    fmt.Fprintf(w, "<h1>Server Information</h1>")
    fmt.Fprintf(w, "<p><strong>Public IP:</strong> %s</p>", publicIP)
    fmt.Fprintf(w, "<p><strong>Private IP:</strong> %s</p>", privateIP[1].String())
    fmt.Fprintf(w, "<p><strong>Running Port:</strong> %s</p>", port)
    fmt.Fprintf(w, "<p><strong>Machine Name:</strong> %s</p>", hostname)
    fmt.Fprintf(w, "<h2>Client Information</h2>")
    fmt.Fprintf(w, "<p><strong>Client IP:</strong> %s</p>", r.RemoteAddr)
    fmt.Fprintf(w, "<h3>Request Headers</h3><ul>")
    
    for name, values := range r.Header {
        fmt.Fprintf(w, "<li><strong>%s:</strong> %s</li>", name, values)
    }
    fmt.Fprintf(w, "</ul></body></html>")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "200 OK")
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/health", healthHandler)
	
    http.ListenAndServe(":8080", nil)
    fmt.Println("Server running on port 8080...")
}