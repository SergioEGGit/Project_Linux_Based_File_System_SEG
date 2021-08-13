
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package Metodos

	import (
		"../Variables"
		"fmt"
	)

//---------------------------------------------------Variables----------------------------------------------------------

	var miDisco int
	var discoBinario  []bool
	var EspaciosDisponibles [200]Espacio

//--------------------------------------------------Estructura----------------------------------------------------------

	type Espacio struct {
		P1 int
		Tamano int
		Disponible bool
	}

//----------------------------------------------------Métodos-----------------------------------------------------------

    func GeneraEspacios() {
		elemento := 0
		limpiaVacios()
		for i:=0; i<= miDisco -1 ; i++ {
			if i <= 0 && !discoBinario[i]  {
				EspaciosDisponibles[elemento].P1 = 0
				EspaciosDisponibles[elemento].Disponible = true
			}
			if i > 0 &&  i <= miDisco - 1 && discoBinario[i -1] && !discoBinario[i]  {
				elemento++
				EspaciosDisponibles[elemento].P1  = i
				EspaciosDisponibles[elemento].Disponible  = true
			}
			sumaVacios(i, elemento)
		}
    }

	func sumaVacios(inicio int, elemento int) {
		if   !discoBinario[inicio]  {
			EspaciosDisponibles[elemento].Tamano++
		}
	}

	func limpiaVacios(){
		for i:= 0; i<= 200 - 1; i++{
			EspaciosDisponibles[i].Tamano = 0
			EspaciosDisponibles[i].P1 = 0
			EspaciosDisponibles[i].Disponible = false
		}
	}

	func LimpiaDisco(){
		 discoBinario = make([]bool,0)
	}

	func LlenaDisco(inicio int, tamano int){
		for i:= inicio; i<=inicio + tamano - 1; i++ {
			discoBinario[i] = true
		}
	}

	func CreaDisco( Disco int){
		miDisco = Disco
		for i :=0; i<= miDisco - 1; i++ {
			discoBinario = append(discoBinario, false)
		}
	}

	func TamanoLibreTotal() int {
		var total int
		for i := 0; i<= 200 - 1 ;i++ {
			if EspaciosDisponibles[i].Disponible {
				total = total + EspaciosDisponibles[i].Tamano
			}
		}
		return total
	}

	func TamanoOcupadoTotal() int {
		var total int
		total = 0
		for i := 0; i<= miDisco - 1 ;i++ {
			 if discoBinario[i] {
				total++
			 }
		}
		return total
	}

	func MostrarEspacios() {
		siExiste := true
		for i := 0; i<= 200 - 1; i++ {
			if EspaciosDisponibles[i].Disponible {
				fmt.Println("espacios vacios: inicio ", EspaciosDisponibles[i].P1," tamaño ", EspaciosDisponibles[i].Tamano)
				siExiste = false
			}
		}
		if  siExiste {
			fmt.Println("espacios vacios  no hay disponibilidad")
		}
	}

	func LLenarParticiones(MBRAuxiliar Variables.MBREstructura) {

		//Verificar Particiones

		if MBRAuxiliar.Particion1MBR.SizePart != 0 {

			LlenaDisco(int(MBRAuxiliar.Particion1MBR.InicioPart), int(MBRAuxiliar.Particion1MBR.SizePart))

		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			LlenaDisco(int(MBRAuxiliar.Particion2MBR.InicioPart), int(MBRAuxiliar.Particion2MBR.SizePart))

		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			LlenaDisco(int(MBRAuxiliar.Particion3MBR.InicioPart), int(MBRAuxiliar.Particion3MBR.SizePart))

		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			LlenaDisco(int(MBRAuxiliar.Particion4MBR.InicioPart), int(MBRAuxiliar.Particion4MBR.SizePart))

		}

	}

	func LLenarParticionesAdd(MBRAuxiliar Variables.MBREstructura, NumeroParticion int) {

		//Verificar Particiones

		if MBRAuxiliar.Particion1MBR.SizePart != 0 && NumeroParticion != 1 {

			LlenaDisco(int(MBRAuxiliar.Particion1MBR.InicioPart), int(MBRAuxiliar.Particion1MBR.SizePart))

		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 && NumeroParticion != 2 {

			LlenaDisco(int(MBRAuxiliar.Particion2MBR.InicioPart), int(MBRAuxiliar.Particion2MBR.SizePart))

		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 && NumeroParticion != 3 {

			LlenaDisco(int(MBRAuxiliar.Particion3MBR.InicioPart), int(MBRAuxiliar.Particion3MBR.SizePart))

		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 && NumeroParticion != 4 {

			LlenaDisco(int(MBRAuxiliar.Particion4MBR.InicioPart), int(MBRAuxiliar.Particion4MBR.SizePart))

		}

	}
