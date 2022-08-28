package main

import (
	"fmt"
)

func main() {

	//control de flujo if

	name := "tim"

	fmt.Println("antes del if")

	if name == "tim" {
		fmt.Println("dentro del if", name)
	} else {
		fmt.Println("dentro del else", name)
	}

	fmt.Println("después del if")

	//switch - automaticamente go ejecuta el codigo y rompe el flujo

	age := 17

	switch age {
	case 17:
		fmt.Println("te falta edad")
	case 18, 19, 20: //nos deja unir casos con ,
		fmt.Println("puedes pasar primer piso")
	default:
		fmt.Println("pasa a sona mayor")

	}

	typeAdmin := "regional"

	switch {
	case typeAdmin == "reginal":
		fmt.Println("2.000.000")
	case typeAdmin == "departamental":
		fmt.Println("6.000.000")
	default:
		fmt.Println("1.500.000")

	}

	//ciclos
	//for solo tenemos una instrucion for, con sus variantes

	//for clasico
	for i := 0; i < 10; i++ {
		fmt.Println("for clasico", i)
	}

	//for continuo - similar a un while
	n := 1
	for n <= 5 {
		fmt.Println("for continuo", n)
		n++
	}

	//for ever - se ejecuta para siempre
	/*
		for {
			fmt.Println("forever", n)
			n++
		}
		//para conecciones que necesiten ser escuchadas siempre (sockets) entre otros

	*/

	//for range

	nums := []uint8{2, 3, 4, 5}

	for range nums {
		fmt.Println("range sin parametros")
	}

	for i, v := range nums {
		fmt.Println("range obteniendo indice y valor", i, v)
	}

	for i := range nums {
		fmt.Println("range solo index", i)
	}

	/*
		for _, v := range nums {
			fmt.Println("range solo valor", v)
		}
	*/

	//el for range devuelve una copia sobre el valor, entonces debemos usar la posicion
	//para hacer referencia al mismo

	for i := range nums {
		nums[i] *= 2
	}

	fmt.Println("modificado", nums)

	//iterar sobre mapas
	sports := map[string]string{
		"soccer": "futbol",
		"boxing": "boxeo",
	}

	for key, v := range sports {
		fmt.Println("key", key, "value", v)
	}

	//iterar strings
	hello := "hello"

	for _, v := range hello {
		fmt.Println(v, "sin castear") //devuelve su numero de bits
		fmt.Println(string(v), "casteado")
	}

}
