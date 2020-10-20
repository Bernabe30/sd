package main

import (
	"context"
	"log"
    "encoding/csv"
    "os"
    "time"
    "fmt"
    "math/rand"
    "strconv"
	"github.com/nchcl/sd/chat"
	"google.golang.org/grpc"

)

type Camion struct{
    id int
    tipo int32
    paquete_1 Paquete
    paquete_2 Paquete
    cantidad_paquetes int
}


type Paquete struct{
    id string
    seguimiento string
    tipo string
    valor int
    origen string
    destino string
    intentos int
    fecha_entrega string
}

type message_paquete struct{
    idpaquete string
    seguimiento string
    tipo string
    valor string
    intentos string
    estado string
    origen string
    destino string
}

func writeCSV() {
    var paquetes_camion_retail_1 [][]string
    var paquetes_camion_retail_2 [][]string
    var paquetes_camion_normal_1 [][]string                                                                                                                                                   
    paquetes_camion_retail_1 = append(paquetes_camion_retail_1, []string{"id-paquete", "tipo", "valor", "origen", "destino", "intentos", "fecha_entrega"}) 
    paquetes_camion_retail_2 = append(paquetes_camion_retail_2, []string{"id-paquete", "tipo", "valor", "origen", "destino", "intentos", "fecha_entrega"}) 
    paquetes_camion_normal_1 = append(paquetes_camion_normal_1, []string{"id-paquete", "tipo", "valor", "origen", "destino", "intentos", "fecha_entrega"}) 
    file, err := os.Create("Actividad_camion_retail_1.csv")
    if err != nil {           
        fmt.Println("hello")
        log.Fatal(err)                                                                                                                                                        
    }     
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range paquetes_camion_retail_1 {
        err := writer.Write(value)
        if err != nil {           
        log.Fatal(err)                                                                                                                                                        
        }  
    }
    
    file, err = os.Create("Actividad_camion_retail_2.csv")
    if err != nil {           
        fmt.Println("hello")
        log.Fatal(err)                                                                                                                                                        
    }     
    defer file.Close()

    writer = csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range paquetes_camion_retail_2 {
        err = writer.Write(value)
        if err != nil {           
        log.Fatal(err)                                                                                                                                                        
        }  
    }
    file, err = os.Create("Actividad_camion_normal_1.csv")
    if err != nil {           
        fmt.Println("hello")
        log.Fatal(err)                                                                                                                                                        
    }     
    defer file.Close()

    writer = csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range paquetes_camion_normal_1 {
        err = writer.Write(value)
        if err != nil {           
        log.Fatal(err)                                                                                                                                                        
        }  
    }
}
func appendCSV(camion *Camion, nro_paquete int) { 
    var path string
    var data [][]string
    if camion.id==1{
        path = "Actividad_camion_retail_1.csv" 
        if nro_paquete==1{
            data=append(data, []string{camion.paquete_1.id, camion.paquete_1.tipo, strconv.Itoa(camion.paquete_1.valor), camion.paquete_1.origen, camion.paquete_1.destino, strconv.Itoa(camion.paquete_1.intentos), camion.paquete_1.fecha_entrega})
        }else{
            data=append(data, []string{camion.paquete_2.id, camion.paquete_2.tipo, strconv.Itoa(camion.paquete_2.valor), camion.paquete_2.origen, camion.paquete_2.destino, strconv.Itoa(camion.paquete_2.intentos), camion.paquete_2.fecha_entrega})
        }
    }else if camion.id==2{
        path = "Actividad_camion_retail_2.csv"
        if nro_paquete==1{
            data=append(data, []string{camion.paquete_1.id, camion.paquete_1.tipo, strconv.Itoa(camion.paquete_1.valor), camion.paquete_1.origen, camion.paquete_1.destino, strconv.Itoa(camion.paquete_1.intentos), camion.paquete_1.fecha_entrega})
        }else{
            data=append(data, []string{camion.paquete_2.id, camion.paquete_2.tipo, strconv.Itoa(camion.paquete_2.valor), camion.paquete_2.origen, camion.paquete_2.destino, strconv.Itoa(camion.paquete_2.intentos), camion.paquete_2.fecha_entrega})
        }
    }else{
        path = "Actividad_camion_normal_1.csv" 
        if nro_paquete==1{
            data=append(data, []string{camion.paquete_1.id, camion.paquete_1.tipo, strconv.Itoa(camion.paquete_1.valor), camion.paquete_1.origen, camion.paquete_1.destino, strconv.Itoa(camion.paquete_1.intentos), camion.paquete_1.fecha_entrega})
        }else{
            data=append(data, []string{camion.paquete_2.id, camion.paquete_2.tipo, strconv.Itoa(camion.paquete_2.valor), camion.paquete_2.origen, camion.paquete_2.destino, strconv.Itoa(camion.paquete_2.intentos), camion.paquete_2.fecha_entrega})
        }
    }     

    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)     
    
    if err != nil {           
        fmt.Println("hello")
        log.Fatal(err)                                                                                                                                                        
    }  
    
    defer f.Close()                                                                                                                                                           
                                                                                                                                                                                                                                                                                                                                                                                                                        
    w := csv.NewWriter(f)                                                                                                                                                     
    w.WriteAll(data)                                                                                                                                                          
                                                                                                                                                                                
    if err := w.Error(); err != nil {                                                                                                                                         
        log.Fatal(err)                                                                                                                                                        
    }
}


func probabilidad_entrega() bool{
    rand.Seed(time.Now().UTC().UnixNano())
    porcentaje := 1 + rand.Intn(99)
    fmt.Println("La probabilidad de Entregar es:", porcentaje)
    if porcentaje<=80{
        return true
    } else {
        return false
    }

}


func actualizar_paquete(camion *Camion, estado string, nro_paquete int){
    var mensaje_server message_paquete
    if nro_paquete==1{
        fmt.Println(camion.paquete_1.id, camion.paquete_1.tipo, camion.paquete_1.valor, camion.paquete_1.origen, camion.paquete_1.destino, camion.paquete_1.intentos, camion.paquete_1.fecha_entrega)
        mensaje_server=message_paquete{
            idpaquete: camion.paquete_1.id,
            seguimiento: camion.paquete_1.seguimiento,
            tipo: camion.paquete_1.tipo,
            valor: strconv.Itoa(camion.paquete_1.valor),
            intentos: strconv.Itoa(camion.paquete_1.intentos),
            estado: estado,
            origen: camion.paquete_1.origen,
            destino: camion.paquete_1.destino,
        }
        //fmt.Println(mensaje_server) //Este es el que debes enviar y leer para actualizar
        c := chat.NewChatServiceClient(*ptr)
        response, err := c.EnvioTerminado(context.Background(), &chat.Paquete{Idpaquete: mensaje_server.idpaquete,Seguimiento: mensaje_server.seguimiento,Tipo: mensaje_server.tipo,Valor: mensaje_server.valor,Intentos: mensaje_server.intentos,Estado: mensaje_server.estado,Origen: mensaje_server.origen, Destino: mensaje_server.destino})
        if err != nil {
            log.Fatalf("Error when calling SayHello: %s", err)
        }
        log.Printf(response.Body)
        
    }else{
        fmt.Println(camion.paquete_2.id, camion.paquete_2.tipo, camion.paquete_2.valor, camion.paquete_2.origen, camion.paquete_2.destino, camion.paquete_2.intentos, camion.paquete_2.fecha_entrega)
        mensaje_server=message_paquete{
            idpaquete: camion.paquete_2.id,
            seguimiento: camion.paquete_2.seguimiento,
            tipo: camion.paquete_2.tipo,
            valor: strconv.Itoa(camion.paquete_2.valor),
            intentos: strconv.Itoa(camion.paquete_2.intentos),
            estado: estado,
            origen: camion.paquete_2.origen,
            destino: camion.paquete_2.destino,
        }
        //fmt.Println(mensaje_server) //Este es el que debes enviar y leer para actualizar
        c := chat.NewChatServiceClient(*ptr)
        response, err := c.EnvioTerminado(context.Background(), &chat.Paquete{Idpaquete: mensaje_server.idpaquete,Seguimiento: mensaje_server.seguimiento,Tipo: mensaje_server.tipo,Valor: mensaje_server.valor,Intentos: mensaje_server.intentos,Estado: mensaje_server.estado,Origen: mensaje_server.origen, Destino: mensaje_server.destino})
        if err != nil {
            log.Fatalf("Error when calling SayHello: %s", err)
        }
        log.Printf(response.Body)
    }
}

func entregar_paquete_1(camion *Camion, tiempo_envio int){
    camion.paquete_1.intentos=camion.paquete_1.intentos+1
    //actualizar_paquete(camion, "En Camino", 1)
    time.Sleep(time.Duration(tiempo_envio) * time.Second)
    if probabilidad_entrega(){
        t := time.Now()
        camion.paquete_1.fecha_entrega=t.Format("02-01-2006 15:04")//Day()+"-"+t.Month()+"-"+t.Year()+" "+t.Hour()+":"t.Minute()
        actualizar_paquete(camion, "Recibido", 1)
        // ENTREGADO PAQUETE 1
        if camion.paquete_2.fecha_entrega=="0"{
            // REVISO SI EL PAQUETE FUE ENVIADO
            if camion.paquete_2.tipo=="retail"{
                if camion.paquete_2.intentos<3{
                    entregar_paquete_2(camion, tiempo_envio)
                    //VOY A ENTREGAR EL OTRO PAQUETE
                }else{
                    return
                    //RETORNO A BODEGA.
                }
            }else{
                if camion.paquete_2.intentos*10<=camion.paquete_2.valor && camion.paquete_2.intentos<3{
                    entregar_paquete_2(camion, tiempo_envio)
                    //VOY A ENTREGAR EL OTRO PAQUETE
                }else{
                    return
                    //RETORNO A BODEGA
                }
            }
        }else{
            return
            //RETORNO A BODEGA
        }
    }else{
        actualizar_paquete(camion, "No Recibido", 1)
        // NO FUE ENTREGADO SE SIGUE CON EL SIGUIENTE
        if camion.paquete_2.fecha_entrega=="0"{
            //REVISION DE SI SE ENTREGO
            if camion.paquete_2.tipo=="retail"{
                if camion.paquete_2.intentos<3{
                    entregar_paquete_2(camion, tiempo_envio)
                    //VOY A ENTREGAR EL SIGUIENTE PAQUETE
                }else{
                    return
                    //RETORNO A BODEGA
                }
            }else{
                if camion.paquete_2.intentos*10<=camion.paquete_2.valor && camion.paquete_2.intentos<3{
                    entregar_paquete_2(camion, tiempo_envio)
                    //A ENTREGAR EL SIGUIENTE PAQUETE
                }else{
                    return
                    //RETORNO BODEGA
                }
            }
        }else if camion.paquete_1.fecha_entrega=="0"{
            if camion.paquete_1.tipo=="retail"{
                if camion.paquete_1.intentos<3{
                    entregar_paquete_1(camion, tiempo_envio)
                    //VOY A ENTREGAR EL SIGUIENTE
                }else{
                    return
                    //RETORNO BODEGA
                }
            }else{
                if camion.paquete_1.intentos*10<=camion.paquete_1.valor && camion.paquete_1.intentos<3{
                    entregar_paquete_1(camion, tiempo_envio)
                    //VOY A ENTREGAR EL SIGUIENTE
                }else{
                    return
                    //RETORNO A BODEGA
                }
            }
        }else{
            return 

            //RETORNO BODEGA
        }
    }
}


func entregar_paquete_2(camion *Camion, tiempo_envio int){
    camion.paquete_2.intentos=camion.paquete_2.intentos+1
    //actualizar_paquete(camion, "En Camino", 2)
    time.Sleep(time.Duration(tiempo_envio) * time.Second)
    if probabilidad_entrega(){
        t := time.Now()
        camion.paquete_2.fecha_entrega=t.Format("02-01-2006 15:04")//t.Day()+"-"+t.Month()+"-"+t.Year()+" "+t.Hour()+":"t.Minute()
        actualizar_paquete(camion, "Recibido", 2)
        if camion.paquete_1.fecha_entrega=="0"{
            if camion.paquete_1.tipo=="retail"{
                if camion.paquete_1.intentos<3{
                    entregar_paquete_1(camion, tiempo_envio)
                }else{
                    return
                }
            }else{
                if camion.paquete_1.intentos*10<=camion.paquete_1.valor && camion.paquete_1.intentos<3{
                    entregar_paquete_1(camion, tiempo_envio)
                }else{
                    return
                }
            }
        }else{
            return
        }
    }else{
        actualizar_paquete(camion, "No Recibido", 2)
        if camion.paquete_1.fecha_entrega=="0"{
            if camion.paquete_1.tipo=="retail"{
                if camion.paquete_1.intentos<3{
                    entregar_paquete_1(camion, tiempo_envio)
                }else{
                    return
                }
            }else{
                if camion.paquete_1.intentos*10<=camion.paquete_1.valor && camion.paquete_1.intentos<3{
                    entregar_paquete_1(camion, tiempo_envio)
                }else{
                    return
                }
            }
        }else if camion.paquete_2.fecha_entrega=="0"{
            if camion.paquete_2.tipo=="retail"{
                if camion.paquete_2.intentos<3{
                    entregar_paquete_2(camion, tiempo_envio)
                }else{
                    return
                }
            }else{
                if camion.paquete_2.intentos*10<=camion.paquete_2.valor && camion.paquete_2.intentos<3{
                    entregar_paquete_2(camion, tiempo_envio)
                }else{
                    return
                }
            }
        }else{
            return 
        }
    }
}



func inter_entrega(camion *Camion, tiempo_envio int) {
    //Debo enviar aviso de EN CAMINO.
    //for camion.cantidad_paquetes!=0{
    if camion.paquete_1.id!=""{
        if camion.paquete_1.valor>=camion.paquete_2.valor {
            entregar_paquete_1(camion, tiempo_envio)
            appendCSV(camion, 1)
            if camion.paquete_2.id!=""{
                appendCSV(camion, 2)
            }
        } else {
            entregar_paquete_2(camion, tiempo_envio)
            appendCSV(camion, 2)
            appendCSV(camion, 1)
        }
    }
    //}
}

func solicitar_paquete(conexion *grpc.ClientConn, camion *Camion) Paquete{
    c := chat.NewChatServiceClient(conexion)
    response, err := c.Camion(context.Background(), &chat.Tipo{Tipo: camion.tipo})
    if err != nil {
        //log.Fatalf("Error when calling SayHello: %s", err)
        return Paquete{}
    }else{
        if (response.Idpaquete=="-1"){
            fmt.Println("No PAQUETE")
            return Paquete{}
        }
        valor, err:=strconv.Atoi(response.Valor)
        if err != nil {
            //log.Fatalf("Error when calling SayHello: %s", err)
            return Paquete{}
        }
        return Paquete{
            id: response.Idpaquete ,//paquete.Get_id()
            seguimiento: response.Seguimiento,
            tipo: response.Tipo, //paquete.Get_tipo()
            valor: valor,//response.Valor, //paquete.Get_valor()
            origen: response.Origen, //paquete.Get_origen()
            destino: response.Destino, //paquete.Get_destino()
            intentos: 0,
            fecha_entrega: "0",
        }
    }
}

func espera_paquetes(camion *Camion ,tiempo_espera int, tiempo_envio int, conexion *grpc.ClientConn) {
    //Primero espera la llegada de los paquetes
    for{
        camion.paquete_1 = solicitar_paquete(conexion, camion) 
        time.Sleep(time.Duration(tiempo_espera) * time.Second)
        camion.paquete_2 = solicitar_paquete(conexion, camion) 
        inter_entrega(camion, tiempo_envio)
    }

}

var ptr **grpc.ClientConn

func main() {
    writeCSV()
    var conn *grpc.ClientConn
    conn, err := grpc.Dial("dist118:9000", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    ptr = &conn
    

    var tiempo int
    var tiempo_envio int
    camion_retail_1 := &Camion{
        id: 1 ,//paquete.Get_id()
        tipo: 1, //paquete.Get_tipo()
        paquete_1: Paquete{}, //paquete.Get_valor()
        paquete_2: Paquete{}, //paquete.Get_origen()
        cantidad_paquetes: 0,
    }
    camion_retail_2 := &Camion{
        id: 2 ,//paquete.Get_id()
        tipo: 1, //paquete.Get_tipo()
        paquete_1: Paquete{}, //paquete.Get_valor()
        paquete_2: Paquete{}, //paquete.Get_origen()
        cantidad_paquetes: 0,
    }
    camion_normal_1 := &Camion{
        id: 3 ,//paquete.Get_id()
        tipo: 2, //paquete.Get_tipo()
        paquete_1: Paquete{}, //paquete.Get_valor()
        paquete_2: Paquete{}, //paquete.Get_origen()
        cantidad_paquetes: 0,
    }

    fmt.Println("Bienvenido al interfaz de Camiones")
    fmt.Println("El tiempo de espera entre paquetes sera:")
    fmt.Scan(&tiempo)
    fmt.Println("Tiempo que demora el envio: ")
    fmt.Scan(&tiempo_envio)
    
    fmt.Println("| id-paquete | tipo | valor | origen | destino | intentos | fecha-entrega |")
        
    go espera_paquetes(camion_retail_1, tiempo, tiempo_envio, conn)
    go espera_paquetes(camion_retail_2, tiempo, tiempo_envio, conn)
    espera_paquetes(camion_normal_1, tiempo, tiempo_envio, conn)
    /*for i := 0; i < 3; i++ {
        if i==0{
            go espera_paquetes(camion_retail_1, tiempo, tiempo_envio, conn)
        }else if i==1{
            go espera_paquetes(camion_retail_2, tiempo, tiempo_envio, conn)
        }else{
            go espera_paquetes(camion_normal_1, tiempo, tiempo_envio, conn)
        }
    }
    /*go func(){
        for{
            espera_paquetes(camion_normal_1, tiempo, conn2)
        }
    }()
    /*
    switch tipo {
        case 1:
            retail(tiempo)
        case 2:
            pyme(tiempo)
    }
*/	
}
