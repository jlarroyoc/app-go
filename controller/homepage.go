package controller

import (
	"log"
	"sort"
	"bytes"
	"io"
	"os"
	"time"
	"net"
	"strings"
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
	log.Println("")
	log.Println("RemoteAddr: " + req.RemoteAddr)
	log.Println("RequestURI: "+req.RequestURI)
	log.Println("Host: "+ req.Host)

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

	log.Println("***************")
	writeRequest(c.Request)
	log.Println("***************")

	hostname, _ := os.Hostname()

	server :=  c.Request.Header.Get("X-Forwarded-Host") 
	port :=  c.Request.Header.Get("X-Forwarded-Port") 
	if len(server)==0 {
		if strings.Contains(c.Request.Host, ":") {
		log.Println("**dehh*")
			server, port, _ = net.SplitHostPort(c.Request.Host)
   	 	} else {
			server = c.Request.Host
		}
	}

	addrs, _ := net.LookupIP(hostname)
        var ipv4 net.IP = nil
	for _, addr := range addrs {
		ipv4 = addr.To4()
		if ipv4 != nil {
			log.Println("IPv4: ", ipv4)
			break
		}   
	}

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Geeksbeginner",
			"localName": hostname, 
			"localAddr": ipv4, 
			"serverName": server, 
			"serverPort": port, 
			"dateTime": time.Now().Format(time.RFC850), 
			"locale": c.Request.Header.Get("Accept-Language"), 
		},
	)
}

