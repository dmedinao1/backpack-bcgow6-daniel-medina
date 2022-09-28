package main

import "fmt"

/*

Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

/*Product*/
type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(name string, price float64, quantity int) *Product {
	return &Product{Name: name, Price: price, Quantity: quantity}
}

/*END Product*/

/*User*/
type User struct {
	Name     string
	Lastname string
	Email    string
	Pass     string
	Products []Product
}

func (u *User) AddProduct(product *Product) {
	fmt.Printf("User en AddProduct: %p\n", u)
	u.Products = append(u.Products, *product)
}

func (u *User) DeleteProducts() {
	fmt.Printf("User en DeleteProducts: %p\n", u)
	u.Products = []Product{}
}

/*END User*/

func main() {
	myUser := User{Name: "Daniel"}

	fmt.Printf("User en main: %p\n", &myUser)

	chipProduct := NewProduct("Chips", 3.32, 2)
	appleProduct := NewProduct("Applers", 1.02, 10)

	myUser.AddProduct(chipProduct)
	myUser.AddProduct(appleProduct)

	fmt.Printf("%+v\n", myUser)

	myUser.DeleteProducts()

	fmt.Printf("%+v\n", myUser)
}
