
//---------------------------------------------Paquetes E Imports-------------------------------------------------------

	package AnalisisYComandos

    import (
    	"../Variables"
	)

//-----------------------------------------------------Métodos----------------------------------------------------------

	func VerificarComandoExit() {

		ComandoExit()

	}

	func ComandoExit() {

	    Variables.Salir = true

	}
