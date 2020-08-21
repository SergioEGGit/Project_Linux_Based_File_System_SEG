
//-------------------------------------------------Paquetes E Imports---------------------------------------------------

    package AnalisisYComandos

	import (
		"fmt"
		"github.com/gookit/color"
	)

//------------------------------------------------------Métodos---------------------------------------------------------

	func VerificarComandoPause() {

		ComandoPause()

	}

	func ComandoPause() {

		color.HEX("#26abc9", false).Println("Presione Enter Para Continuar.....")
		fmt.Println("")
		_, _ = fmt.Scanln()

	}