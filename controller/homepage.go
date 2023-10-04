package controller

import (
	"log"
	"sort"
	"bytes"
	"io"
	"os"
	"time"
	"net"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
)

func printHeaders(h http.Header) {
	sortedKeys := make([]string, 0, len(h))

	for key := range h {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		for _, value := range h[key] {
			log.Printf("%s: %s\n", key, value)
		}
	}
}

func writeRequest( req *http.Request) {
	log.Printf("%s %s %s\n", req.Method, req.URL, req.Proto)
	log.Println("")

	log.Printf("Host: %s\n", req.Host)
	printHeaders(req.Header)

	var body bytes.Buffer
	io.Copy(&body, req.Body) // nolint:errcheck

	if body.Len() > 0 {
		log.Println("")
		log.Println(body)
	}
}

func Router(r *gin.Engine) {
	r.GET("/", index)
}

func index(c *gin.Context) {

	res1, _ := json.MarshalIndent(c.Request.Header,"", "\t")
	log.Println("***************")
	log.Println(string(res1))
	log.Println("***************")
	writeRequest(c.Request)
	log.Println("***************")
	log.Println(c.Request.RemoteAddr)
	log.Println(c.Request.RequestURI)
	log.Println(os.Hostname())
	hostname, _ := os.Hostname()
	server, port, _ := net.SplitHostPort(c.Request.Host)

	addrs, _ := net.LookupIP(hostname)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			log.Println("IPv4: ", ipv4)
		}   
	}

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Geeksbeginner",
			"localName": hostname, 
			"serverName": server, 
			"serverPort": port, 
			"dateTime": time.Now().Format(time.RFC850), 
			"locale": c.Request.Header.Get("Accept-Language"), 
		},
	)
}

