
package main

import "io/ioutil"
import "net/http"
import "fmt"
import "os"
import "github.com/pusher/pusher-http-go/v5"

func pusherAuth(res http.ResponseWriter, req *http.Request) {
	pusherClient := pusher.Client {
		AppID: os.Getenv("APP_ID"),
		Key: os.Getenv("APP_KEY"),
		Secret: os.Getenv("APP_SECRET"),
		Cluster: os.Getenv("APP_CLUSTER"),
	}

	params, _ := ioutil.ReadAll(req.Body)
	// This authenticates every user. Don't do this in production!
	response, err := pusherClient.AuthorizePrivateChannel(params)
  
	if err != nil {
	  panic(err)
	}
  
	fmt.Fprintf(res, string(response))
  }
  
  func main() {
	http.HandleFunc("/pusher/auth", pusherAuth)
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	http.ListenAndServe(":" + port, nil)
	fmt.Sprintf("Server start listening on port: %s", port);
  }