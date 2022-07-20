package main

import "fmt"

func main() {
	// fmt.Println("Hola Mundo")
	//Comentarios

	//Constantes

	const pi float64 = 3.1415
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//variables
	var altura int = 14
	base := 12
	var area int
	fmt.Println(altura, base, area)

	//Zero value

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado:", areaCuadrado)

	x := 50
	y := 10

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = x / y
	fmt.Println("División:", result)
	result = x % y

	// Modulo
	fmt.Println("Modulo:", result)

	// Incremetal
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Reto Área
	// Rectángulo Trapecio y Círculo

	// Rectángulo
	baseRectangulo := 10
	alturaRectangulo := 20
	areaRectagulo := baseRectangulo * alturaRectangulo

	fmt.Println("Área del Rectángulo:", areaRectagulo)

	// Trapecio
	baseMenor := 3.5
	baseMayor := 9.5
	alturaTrapecio := 4
	areaTrapecio := (baseMayor + baseMenor) * float64(alturaTrapecio) / 2.0

	fmt.Println("Área del Trapecio:", areaTrapecio)

	// Círculo

	radioCirculo := 2.5
	areaCirculo := pi * float64(radioCirculo)  * float64(radioCirculo)

	fmt.Println("Área del Circulo:", areaCirculo)
}
