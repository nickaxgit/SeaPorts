package main


import (
	"context"
	"flag"
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	sp "github.com/nickaxgit/seaports/protobuff"
	

)


var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)



func main(){
	
	var port sp.SeaPort =sp.SeaPort{Name:"SpaceX Deep Water 1", Coordinates:[] float32 {57.92,38.21},Unlocs:[] string {"SPACEX1"}}
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sp.NewUpsertClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Upsert(ctx, &port)
	if err != nil {
		log.Fatalf("could not upsert: %v", err)
	}
	log.Printf("Upsert: %s", r.Text)

}







