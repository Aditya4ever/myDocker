package  main

import (
	"fmt"

	"log"
	"net/http"
)


func main(){
	request1()
}

func homePage(response http.ResponseWriter,r *http.Request){
	fmt.Printf("welcome to Go Web API")
	fmt.Println("Endpoint hit:HOmepage")
}

func aboutMe(response http.ResponseWriter,r *http.Request){

	who := "AdityaCh"
	fmt.Fprintf(response, "A little more about me")
	fmt.Println("Endpoint hit: ", who)
	
}

func request1(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
