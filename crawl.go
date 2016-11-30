package main
import (
	"fmt"
	"crypto/tls"
	"flag"
	"os"
	"collectlinks" 
	"net/http"
)

func main() {
	flag.Parse()
	args := flag.Args()


	if len(args) < 1 {
		fmt.Println("Please give me a start page")
		os.Exit(1)
	}

	tlsConfig := &tls.Config{                 
                 InsecureSkipVerify: true,  
               }  

    transport := &http.Transport{    
    TLSClientConfig: tlsConfig,    
	}                                

	client := http.Client{Transport: transport}


	resp, err := http.Get(args[0])
	if err != nil {
		return
	}
	defer resp.Body.Close() 
	links := collectlinks.All(resp.Body) 

	for _, link := range(links) {  
    	fmt.Println(link)            
  	} 


}

