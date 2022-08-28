package main

import "fmt"

func main() {

	//los arrays en go tienen un tamaño unico y se declaran de la sig, forma

	var fruits [3]string

	fruits[0] = "manzana"
	fruits[1] = "pera"
	fruits[2] = "piña"

	//tenemos los arrays literales
	//fast_food := [3]string{"hot dog", "hamburguer", "pizza"}

	//slices -> no contienen valores, son apuntadores a arrays

	set := [7]int{1, 2, 3, 4, 5, 6, 7}
	//[init:end] -> el indice final es excluyente
	five := set[0:5] //puede ser [:5]

	fmt.Println(five)

	//si modificamos algun valor del slice, se modificara en el arr original
	//ya que este guarda su referencia de memoria.

	//MAPAS

	//son datos clave-valor los cuales.

	//[type-key]type-value

	//con funcion make

	animals := make(map[string]string)

	animals["cat"] = "gato"
	animals["dog"] = "perro"

	//mapas literales

	sports := map[string]string{
		"soccer": "futbol",
		"boxing": "boxeo",
	}

	fmt.Println(sports, animals)

	//metodos - mapas

	delete(sports, "soccer") //eliminamos la llave y por ende el valor

	fmt.Println(animals["cat"]) //obtener valor

	//los mapas devuelven 2 valores cuando accedemos a este

	//value -> que es el valor y si la llave existe o no
	//son variables pero por nomenclatura a una la llamaremos ok

	//sino usaremos la variable usaremos la nomenclatura de _ ya que go las variables
	//que no usemos nos va dar error al compilar

	_, ok := sports["american football"]

	fmt.Println(ok)

}
