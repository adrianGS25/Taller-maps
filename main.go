package main

import "fmt"

func main() {
	fmt.Println("Ingrese una frase o texto:")

	var palabras [100]string
	var destinos []interface{}
	for i := range palabras {
		destinos = append(destinos, &palabras[i])
	}
	fmt.Scanln(destinos...)
	conteo := make(map[string]int)
	for _, palabra := range palabras {
		if palabra != "" {
			conteo[palabra]++
		}
	}

	if len(conteo) == 0 {
		fmt.Println("No ingresó ningún texto.")
		return
	}
	fmt.Println("\nConteo de palabras:")
	for palabra, veces := range conteo {
		fmt.Printf("%s: %d\n", palabra, veces)
	}
	masRepetida := ""
	maximo := 0
	for palabra, veces := range conteo {
		if veces > maximo {
			maximo = veces
			masRepetida = palabra
		}
	}

	fmt.Printf("\nLa palabra que más se repite es \"%s\" con %d repeticiones.\n", masRepetida, maximo)
}


package main

import "fmt"

func valorTotal(precio float64, cantidad int) float64 {
	return precio * float64(cantidad)
}

func tieneStock(cantidad int) bool {
	return cantidad > 0
}

func mostrarProducto(nombre string, precio float64, cantidad int) {
	fmt.Println("Nombre:", nombre)
	fmt.Println("Precio:", precio)
	fmt.Println("Cantidad:", cantidad)
	fmt.Println("Valor total:", valorTotal(precio, cantidad))
	if tieneStock(cantidad) {
		fmt.Println("Estado: disponible")
	} else {
		fmt.Println("Estado: sin stock")
	}
	fmt.Println()
}

func main() {
	p1 := struct {
		Nombre   string
		Precio   float64
		Cantidad int
	}{"Laptop", 850.50, 5}

	p2 := struct {
		Nombre   string
		Precio   float64
		Cantidad int
	}{"Mouse", 12.99, 0}

	p3 := struct {
		Nombre   string
		Precio   float64
		Cantidad int
	}{"Teclado", 25.00, 10}

	mostrarProducto(p1.Nombre, p1.Precio, p1.Cantidad)
	mostrarProducto(p2.Nombre, p2.Precio, p2.Cantidad)
	mostrarProducto(p3.Nombre, p3.Precio, p3.Cantidad)
}