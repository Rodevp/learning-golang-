package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func funciones() {

	hello("rodrigo", "rodriguez")
	word := "nanda kure"

	change(&word)

	total := sum(2, 2)
	fmt.Println(total)

	lower, upper := convert("RoDrigO")
	fmt.Println(lower, upper)

	//errores en go
	//nos ayudaremos con el paquete errors para poder personalizar dicho error
	errorsInGo()

	//manejar un error
	err := boom()

	if err != nil { //se evalua que no sea nil, ya que si lo es no hay error
		fmt.Println(":O", err)
		//detiene la ejecucion, pero podemos hacer otra cosa
	}

	//otra forma resumida de manejar los errores
	if err2 := boom(); err2 != nil {
		fmt.Println("An error occurred:", err)
	}

	//variatic functions
	variaticFunctionSayHello() //puede no recibir el parametro
	variaticFunctionSayHello("Sammy")
	variaticFunctionSayHello("Sammy", "Jessica", "Drew", "Jamie")

	//funciones anonimas
	func(param string) {
		fmt.Println("funcion auto ejecuta", param)
	}("soy el parametro") //funcion anonima auto ejecutada

	//las podremos guardar en variables
	x := func(param string) {
		fmt.Println(param)
	}

	x("Hola mundo, soy una funcion anonima")

}

func hello(name string, lastName string) {
	fmt.Println("hola", name, lastName)
}

func change(value *string) {
	*value = "nami"
}

func sum(x int8, y int8) int8 {
	return x + y
}

// funciones que devuelven 2 valores
func convert(text string) (string, string) { //() van los tipos

	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may

}

func errorsInGo() {
	err := errors.New("soy un errorsito")
	fmt.Println(err)

	err = fmt.Errorf("ay! un error en este momento: %v", time.Now())
	fmt.Println(err)

}

//retornando un error

func boom() error {
	return errors.New("KABOOM!")
}

//las funciones variaticas son capaces de recibir muchos parametros
// parecido al operador rest en js.

func variaticFunctionSayHello(persons ...string) {
	fmt.Println("persons", persons)
	for _, n := range persons {
		fmt.Printf("Hello %s\n", n)
	}
}
