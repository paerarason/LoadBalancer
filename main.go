package main
import (
	"fmt"
    "net/http/httputil"
	"log"
	"net/url"
)

type  server struct{
   IP string
   proxy *httputil.ReverseProxy 
}

func NewServer(IP string ) *server{
   server_url,err:=url.Parse(IP)
   ERRhandler(err)
   return &server{
	IP:IP,
	proxy:httputil.NewSingleHostReverseProxy(server_url),
   }
}
 

func main(){
	fmt.Println("hello server")
}


func ERRhandler(err error){
	if err !=nil {
		log.Fatal(err)
	}
 }