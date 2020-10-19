package main

import (
    "fmt"
    
    "encoding/json"
	"github.com/streadway/amqp"
)


type info_envio struct{
    id int
    tipo string
    valor int
    intentos int
    fecha_entrega string
}

type balance struct{
	costos int
	perdidas int
	ganancias int
}


func balance_entregado(envio *info_envio, balance_global *balance){
	balance_global.ganancias=balance_global.ganancias+envio.valor
	balance_global.costos=balance_global.costos+(envio.intentos-1)*10
}

func balance_retail_NO_entregado(envio *info_envio, balance_global *balance){
	balance_global.ganancias=balance_global.ganancias+envio.valor
	balance_global.costos=balance_global.costos+(envio.intentos-1)*10
}

func balance_normal_NO_entregado(envio *info_envio, balance_global *balance){
	balance_global.ganancias=balance_global.ganancias+0
	balance_global.costos=balance_global.costos+(envio.intentos-1)*10
}


func balance_prioritario_NO_entregado(envio *info_envio, balance_global *balance){
	balance_global.ganancias=balance_global.ganancias+(30*envio.valor/100)
	balance_global.costos=balance_global.costos+(envio.intentos-1)*10
}



func inter_balance(envio *info_envio, balance_global *balance){
	if envio.fecha_entrega!="0"{
		balance_entregado(envio,balance_global)
	}else{
		if envio.tipo=="normal"{
			balance_normal_NO_entregado(envio,balance_global)
		}else if envio.tipo=="prioritario"{
			balance_prioritario_NO_entregado(envio,balance_global)
		}else{
			balance_retail_NO_entregado(envio,balance_global)
		}
	}
    fmt.Println("|",balance_global.costos,"|",balance_global.ganancias,"|", balance_global.perdidas)
}

func main() {
    
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"Finanzas",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
    
    
	balance_global:=&balance{
		costos: 0,
		perdidas: 0,
		ganancias: 0,
	}

    
    fmt.Println("Bienvenido al interfaz de Finanzas")

 	var info map[string]interface{}
    
	forever := make(chan bool)
	go func() {
		for d := range msgs {
            json.Unmarshal([]byte(d.Body),&info)
            
            msg := &info_envio{
                id: int(info["Id"].(float64)),//paquete.Get_id()
                tipo: fmt.Sprintf("%v", info["Tipo"]), //paquete.Get_tipo()
                valor: int(info["Valor"].(float64)),
                intentos: int(info["Intentos"].(float64)),
                fecha_entrega: fmt.Sprintf("%v", info["Fecha_entrega"]),
            }
            
                        
            fmt.Println("|COSTOS|GANANCIAS|PERDIDAS|")
            inter_balance(msg,balance_global)
		}
	}()

	fmt.Println("Esperando mensajes")
	<-forever
}
