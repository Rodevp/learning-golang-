package main

import (
	"course/course/utils"
	"fmt"
	"os"
)

func defer_y_mas() {

	utils.Example()

	//defer -> nos permite diferir algo, o sea, nos permite aplazar la ejecucion
	//de una funcion en donde este defer fue especificado, este se ejecutara
	//cuando la funcion haya retornado por completo y ahi es donde se ejecutará
	//la funcion que hemos diferido.

	//las funciones que diferamos se van a guardar en una pila donde la ultima
	//que entro es la que se ejecutará primero.

	defer fmt.Println(3) //este se ejecutara cuando la funcion se haya terminado.
	fmt.Println(1)
	fmt.Println(2)

	//defer en uso
	createFile()

	//panic
	division(2, 2)
	division(3, 0)
	division(5, 2)

}

//casos de uso de defer
//(limpiar recursos, cerrar archivos, con. a reds y controladores de bd)

func createFile() {
	file, err := os.Create("hola.txt")

	if err != nil {
		fmt.Print("error al abrir el archivo")
		return
	}

	//esta funcion se ejecutará cuando ya toda la funcion retorne y cierre
	//siempre se ejecutará file close. (cerrará este ya que es un puntero)
	defer file.Close()

	_, err2 := file.Write([]byte("Hola ghopers"))

	if err2 != nil {
		fmt.Print("error al abrir el archivo")
		return
	}

}

//panic -> nos permite entrar en panico, esto nos permite detener
//la ejecucion del programa cuando lo queramos y entre en panico.

//recovery -> nos ayuda a recuperarnos de un panico.
//la usamos junto con defer para que nos cachee ese error y solo se ejecute
//al terminar la funcion o retorne.

func division(divisor, dividendo int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperando el panico: ", r)
		}
	}()

	validaDividendo(dividendo)
	fmt.Println("division: ", (divisor / dividendo))

}

func validaDividendo(dividendo int) {
	if dividendo == 0 {
		panic("AAAAAAAAAAAAAAaaa, es cero")
	}
}
