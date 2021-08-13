
//----------------------------------------------Paquetes E Imports------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Variables"
		"fmt"
		"github.com/gookit/color"
		"os"
		"os/exec"
	)

//----------------------------------------------------Métodos-----------------------------------------------------------

    func VerificarComandoCls() {

    	ComandoCls()

	}

	func ComandoCls() {

		//Variables
		var Command *exec.Cmd
		var AvisoError error


		//Verificar OS
		if Variables.SistemaOperativo == "windows" {

     		Command = exec.Command("cmd", "/c", "cls")
			Command.Stdout = os.Stdout
			AvisoError = Command.Run()

			if AvisoError != nil {

				color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando")
				fmt.Println("")

			}

		} else if Variables.SistemaOperativo == "linux" {

			Command = exec.Command("clear")
			Command.Stdout = os.Stdout
			Command.Stderr = os.Stderr
			AvisoError = Command.Run()

			if AvisoError != nil {

				color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
			fmt.Println("")

		}

		fmt.Println("")

	}