package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gliderlabs/ssh"
)

type Tunnel struct{
	w 		io.Writer
	donech 	chan struct{}
}

var tunnels = map[int]chan Tunnel{

}

func main() {

	go func ()  {
		http.HandleFunc("/", handleRequest)
		log.Fatalln(http.ListenAndServe(":3000", nil))
	}()
	ssh.Handle(func(s ssh.Session) {
		id:= rand.Intn(math.MaxInt)
		tunnels[id] = make(chan Tunnel)
		fmt.Println("tunnel ID-> ", id)

		tunnel:=  <- tunnels[id]
		fmt.Println("tunnel is ready")

		//copying everyting form the cli to the browser 
		_, err :=io.Copy(tunnel.w, s )
		if err != nil {
			log.Fatalln(err)
		}
		close(tunnel.donech)
		s.Write([]byte("This is complete now"))
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}


func handleRequest( w http.ResponseWriter, r *http.Request){
	idstr:= r.URL.Query().Get("id")
	id,_ := strconv.Atoi(idstr)
	tunnel, ok := tunnels[id]
	if !ok{
		w.Write([]byte("Tunnel not found"))
		return
	}
	donech:= make(chan struct{})
	tunnel <- Tunnel{
		w:w,
		donech: donech,
	}

	<-donech
}