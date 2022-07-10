package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"google.golang.org/grpc"	
	sp "github.com/nickaxgit/seaports/protobuff"
	
)

var (
	port = flag.Int("port", 50051, "The server port")
	db = make(map[string]*sp.SeaPort)  //here is our 'database' a map, of ports codes to seaport structures
)

type server struct {
	sp.UnimplementedUpsertServer
}


func main() {
	
	file, _ := ioutil.ReadFile("../ports.json")

	err:= json.Unmarshal([]byte(file),&db) //load the json data into the map of string

	if err !=nil {
		log.Printf(err.Error())
	}

	log.Printf("Loaded " + strconv.FormatInt(int64(len(db)),10) + " ports")
	
	//yup, they loaded
	// for key, val := range db{
    //     fmt.Printf("Key: %s, Value: %s\n", key, val)
    // }


	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sp.RegisterUpsertServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}

func (s *server) Upsert(ctx context.Context, seaport *sp.SeaPort) (*sp.Status, error) {
	
	id:=seaport.GetUnlocs()[0]
	log.Printf("Received: %s", id)
	
	if db[id] !=nil { 
		log.Printf("Updating: %s", id)
	} else{
		log.Printf("Inserting: %s", id)
	}

	
	db[id]=seaport //store the (pointer to the) message in the map (database)
	
	log.Printf(seaport.String())
	
	return &sp.Status{Text: "OK"}, nil
}

