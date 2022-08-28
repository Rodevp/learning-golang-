package main

import "fmt"

type Point struct {
	x int16
	y int16
}

//para cambiar los valores de nuestra "instancia" de nuestra estructura
//debemos pasar un puntero, de no hacerlo asi, solo será una copia.

func changePointWithReference(p *Point) {
	p.x = 250
	fmt.Println("cambiamos el original", p)
}

func changePointNotReference(p Point) {
	p.x = 100
	fmt.Println("crea una copia", p)
}

//podemos componer estructuras con otras a traves de un puntero mismo o una copia.

type Circle struct {
	radius float32
	center *Point
}

//mientras quiera modificar el valor de las estructura original, debemos pasarle
//su referencia, de no ser asi solo actuará como un tipo, funcion pura.

func estructura() {

	/*
		Las estructuras son tipos que podemos crear nosotros y este a su vez puede
		Componenerse de otros tipos. Los structs nos ayudan a emular las clases
		pero a diferencia de este no incluyen metodos, tenderemos que hacer
		referencia a nuestra "instancia" de nuestra estructura por medio de los
		Punteros para poder manipular.

		Esto nos hace pensar en composicion lo cual puede ser más practico y nos
		olvidamos del concepto de clases.

		Otra cosa es que go, esta compuesto de estructuras y punteros.

	*/

	//podemos llamarlos posicionales o nombrados  a las propiedades
	var p1 Point = Point{2, 5}

	//podemos hacerlo literal
	p2 := Point{y: 7, x: 6}

	fmt.Println("sin modificar", p1, p2)

	changePointWithReference(&p1)
	changePointNotReference(p2)

	fmt.Println("con modificar", p1, p2)

	c1 := Circle{radius: 3.14, center: &p1}

	fmt.Println("circle", c1)
	fmt.Println("circle point", c1.center)

}
