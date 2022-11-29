package main

import "fmt"

//Creamos un struct(clase computadora)
type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "pong")
}

//Creamos una funcion que modifique el valor de un atributo del struct
func (myPc *pc) modificarRam() {
	myPc.ram = myPc.ram * 2
}
func (myPc *pc) modificarMarca() {
	myPc.brand = "Huawei"
}

/*func (myPc pc) String() string {
	return fmt.Sprintf("La computadora cuenta con %d Gb en Ram, %d Gb de disco duro y es de la marca %s", myPc.ram, myPc.disk, myPc.brand)
}*/
func main() {

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	//Instanciamos el struct
	myPc := pc{ram: 16, disk: 200, brand: "Lenovo"}
	fmt.Println(myPc)
	myPc.ping()

	//Invocamos la funcion que modifica el valor del atributo Ram
	fmt.Println(myPc)
	myPc.modificarRam()

	fmt.Println(myPc)
	myPc.modificarRam()

	fmt.Println(myPc)
	myPc.modificarRam()

	fmt.Println(myPc.brand, "Es la marca de la pc actual")
	myPc.modificarMarca()
	fmt.Println(myPc.brand, "Es la marca de la pc nueva")

}
