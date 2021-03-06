package main

import (
    "os"
    "log"
    "encoding/csv"
    "fmt"
	"net"

	"github.com/nchcl/sd/chat"
	"google.golang.org/grpc"
)

func main() {
    fmt.Println("Server on")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	writeCSV()
	
    //var seguimiento *int
    //seguimiento = new(int)
    //*seguimiento = 0
	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	
}

func writeCSV() {
    var data [][]string                                                                                                                                                       
    data = append(data, []string{"timestamp", "id-paquete", "tipo", "nombre","valor","origen","destino",
        "seguimiento"})   
    
    file, err := os.Create("result.csv")
    
    if err != nil {           
        fmt.Println("hello")
        log.Fatal(err)                                                                                                                                                        
    }     
    
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range data {
        err := writer.Write(value)
        if err != nil {           
        log.Fatal(err)                                                                                                                                                        
        }  
    }
    
}
