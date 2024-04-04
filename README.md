Saludos en Go

Este paquete prpociona una forma simple de obtener saludos personalizados y aleatorios

## Instalacion
Ejecuta el siguiente comando para instalar el paquete:

```bash 
go get -u github.com/mig-bn/greetins
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo:

``` go
package main

import (
    "fmt"
    "github.com/mig-bn/greetings"
)

func main() {
    message, err := greetings.Hello("Miguel")

    if err != nil {
        fmt.Println("Ocurrio un error: ", err)
        return
    }

    fmt.Println(message)
}

```

Este ejemplo importa el paquete github.com/mig-bn/greetins y llama la funcion Hello, que realiza un saludo personalizad y aletorio. Si ocurre un error, se imprime un mensaje de error. 
