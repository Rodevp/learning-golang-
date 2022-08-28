package main

import "fmt"

func punteros() {

	/*
		los punteros en golang son variables que almacenan la direccion de
		memoria de un valor, es decir, al obtener esta direcion tenemos su
		referencia.

	*/

	//& -> operador para obtener la direccion de memoria de una variable

	fruit := "apple"

	fmt.Println("direccion de memoria fruit", &fruit)

	//recordemos, los punteros almacenan direcciones de memoria y obtienes el valor
	//de las mismas

	//*desreferenciacion y crear punteros

	var p *string = &fruit

	fmt.Println("direccion de memoria fruit", &fruit, "direccion del puntero", p)

	//si modificamos fruit, el valor del puntero cambiará al tener su ref
	fruit = "banana"

	fmt.Println("valor fruit", fruit, "valor puntero", *p)

	//si cambiamos el valor del puntero cambiaremos el valor de fruit ya que
	//nuestro puntero tiene la referencia a su direccion de memoria.

	*p = "pera"

	fmt.Println("valor fruit", fruit, "valor puntero", *p)

}
