
//----------------------------------------------Paquetes E Imports------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"github.com/gookit/color"
		"strings"
	)

//-----------------------------------------------------Métodos----------------------------------------------------------

	func AnalisisComando(CadenaComando string) {

		//Asignaciones
		CadenaComando = Metodos.Trim(CadenaComando)
		Variables.ArregloComandos = Metodos.SplitComando(CadenaComando)
		Variables.ArregloComandos[0] = strings.ToLower(Variables.ArregloComandos[0])
		Variables.ArregloComandos[0] = Metodos.Trim(Variables.ArregloComandos[0])
		Variables.NoExisteArchivo = false


		//Revisión Existencia Comando
		switch Variables.ArregloComandos[0] {

		case "exec":

			VerificarComandoExec()

		case "pause":

			VerificarComandoPause()

		case "mkdisk":

			VerificarComandoMkdisk()

		case "rmdisk":

			VerificarComandoRmdisk()

		case "fdisk":

			VerificarComandoFdisk()

		case "cls":

			VerificarComandoCls()

		case "exit":

			VerificarComandoExit()

		default:

			color.HEX("#de4843", false).Println("Comando No Soportado!")

		}

	}