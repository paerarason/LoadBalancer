package main
import (
	"fmt"
    "net/http/httputil"
	"log"
	"net/url"
	"net/http"
)

type  server struct{
   IP string
   proxy *httputil.ReverseProxy 
}

type Server interface  {
	Address() string 
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

func NewServer(IP string ) *server{
   server_url,err:=url.Parse(IP)
   ERRhandler(err)
   return &server{
	IP:IP,
	proxy:httputil.NewSingleHostReverseProxy(server_url),
   }
}


type Loadbalancer struct{
	port string
	RoundrobinCount int16
	servers [] server
} 

func NewLoadbalancer(port string , servers [] server) *Loadbalancer {
	return &Loadbalancer{
		port:port,
		servers:servers,
		RoundrobinCount:0,}
}

func (lb *Loadbalancer) get_Available_server() server{
   	
}


func main(){
	fmt.Println("hello server")
}


func ERRhandler(err error){
	if err !=nil {
		log.Fatal(err)
	}
 }