
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"fmt"
		"github.com/gookit/color"
		"strconv"
		"strings"
	)

//--------------------------------------------------------Métodos-------------------------------------------------------

	func VerificarComandoUnMount() {

		//Variables
		var Substring string
		var NumeroString string
		var Id bool
		var ParametroExtra bool
		var AvisoError error
		var ArregloParametros []string

		//Asignación
		Id = true
		ParametroExtra = false
		Variables.MapComandos = make(map[string]string)

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos)-1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				Substring = ArregloParametros[0][0:2]
				NumeroString = ArregloParametros[0][2:len(ArregloParametros[0])]

				_, AvisoError = strconv.Atoi(NumeroString)

				if AvisoError != nil {

					Id = false

				} else {

					switch Substring {

					case "id":

						if len(ArregloParametros) > 1 {

							Variables.MapComandos[Metodos.Trim(ArregloParametros[0])] = Metodos.Trim(ArregloParametros[1])

						} else {

							Id = false

						}

					}

				}

			}

			if Id && !ParametroExtra {

				ComandoUnMount()

			} else {

				if ParametroExtra {

					color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
					color.HEX("#de4843", false).Println("Parametros Validos: ")
					color.HEX("#de4843", false).Println("1). -idn->    (Obligatorio)")
					fmt.Println("")

				}

				if !Id {

					color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -idn-> o")
					color.HEX("#de4843", false).Println("Existe Error En Sintaxis")
					fmt.Println("")

				}

			}
		}
	}

	func ComandoUnMount() {

		//Desmontar Particiones
		for Id := range Variables.MapComandos {

			for Contador := 0; Contador < len(Variables.ArregloParticionesMontadas); Contador++ {

				if strings.EqualFold(Metodos.Trim(Variables.MapComandos[Id]), Metodos.Trim(Variables.ArregloParticionesMontadas[Contador].IdentificadorMount)) {

					Variables.ArregloParticionesMontadas = append(Variables.ArregloParticionesMontadas[:Contador], Variables.ArregloParticionesMontadas[Contador+1:]...)

				}
			}

		}

		color.Success.Println("Particion/es Desmontadas Con Exito")
		fmt.Println("")

	}