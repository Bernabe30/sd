package main

import (
	"context"
	"log"
    "encoding/csv"
    "os"
    "time"
    "fmt"

	"github.com/nchcl/sd/chat"
	"google.golang.org/grpc"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func retail(tiempo int) {
    
    var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

    records := readCsvFile("retail.csv")
	
    for i := 1; i < len(records); i++ {
        c := chat.NewChatServiceClient(conn)

        response, err := c.OrderRetail(context.Background(), &chat.Retail{Id: records[i][0], Producto: records[i][1], Valor: records[i][2], Tienda: records[i][3], Destino: records[i][4]})
        if err != nil {
            log.Fatalf("Error when calling SayHello: %s", err)
        }
        log.Printf("Response from server: %s", response.Body)
        
        time.Sleep(time.Duration(tiempo) * time.Second)
    }
    
}

func pyme(tiempo int) {
    
    var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

    records := readCsvFile("pymes.csv")
	
    for i := 1; i < len(records); i++ {
        c := chat.NewChatServiceClient(conn)

        response, err := c.OrderPyme(context.Background(), &chat.Pyme{Id: records[i][0], Producto: records[i][1], Valor: records[i][2], Tienda: records[i][3], Destino: records[i][4], Prioritario:records[i][5]})
        if err != nil {
            log.Fatalf("Error when calling SayHello: %s", err)
        }
        log.Printf("Response from server: %s", response.Body)
        
        time.Sleep(time.Duration(tiempo) * time.Second)
    }
    
}

func main() {
    
    var tipo int
    var tiempo int
    
    fmt.Println("1. Retail")
    fmt.Println("2. Pyme")
    fmt.Scan(&tipo)
    
    fmt.Println("Tiempo entre ordenes")
    fmt.Scan(&tiempo)

    
    switch tipo {
        case 1:
            retail(tiempo)
        case 2:
            pyme(tiempo)
    }
	
}
