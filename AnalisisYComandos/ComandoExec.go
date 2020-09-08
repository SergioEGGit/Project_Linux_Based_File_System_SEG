
//----------------------------------------------Paquetes E Imports------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"fmt"
		"github.com/gookit/color"
		"strings"
	)

//---------------------------------------------------Métodos------------------------------------------------------------

	func VerificarComandoExec() {

		//Variables
		var Path bool
		var ParametroExtra bool
		var ComandoOk bool
		var ArregloParametros []string
		var ContadorPath int

		//Asignación
		Path = false
		ParametroExtra = false
		ComandoOk = false
		ContadorPath = 0
		Variables.MapComandos = make(map[string]string)

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos) - 1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				switch ArregloParametros[0] {

				case "path":

					if len(ArregloParametros) > 1 {


						ArregloParametros[1] = Metodos.QuitarComillas(ArregloParametros[1])
						ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])
						ComandoOk = Metodos.ExisteRuta(ArregloParametros[1])

						if ComandoOk {

							Variables.MapComandos["path"] = ArregloParametros[1]

						}

						Path = true
						ContadorPath++

					} else {

						Path = false

					}

				default:

					ParametroExtra = true

				}
			}
		}

		if Path && !ParametroExtra && ComandoOk && ContadorPath == 1 {

			ComandoExec()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println( "Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println( "Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				fmt.Println("")

			}

			if !ComandoOk {

				if Variables.NoExisteArchivo {

					color.HEX("#de4843", false).Println("El Archivo Indicado No Existe")
					fmt.Println("")

				} else {

					if !Path {

						color.HEX("#de4843", false).Println( "Debe Indicar Los Parametros Obligatorios")
						fmt.Println("")

					} else {

						color.HEX("#de4843", false).Println("Falta El Nombre Y/O Dirección Del Archivo")
						fmt.Println("")

					}
				}

			}

			if !Path {

				color.HEX("#de4843", false).Println("Error En Sintaxis En Parametro -path O No Se Encuentra")
				color.HEX("#de4843", false).Println("Sintaxis: -path->\"Ruta Archivo\"")
				fmt.Println("")

			}

			if ContadorPath > 1 {

				color.HEX("#de4843", false).Println("El Parametro -path Se Indico Mas De Una Una Vez")
                fmt.Println("")
			}

		}

	}

	func ComandoExec() {

		//Variables
		var Bandera bool
		var ArregloTemp []string

		//Asignacion
		ArregloTemp = make([]string, 0)

		//Leer Archivo
		Bandera, ArregloTemp = Metodos.LeerArchivoEntrada(Metodos.Trim(Variables.MapComandos["path"]))

		if Bandera {

			Metodos.RecuperarLDComando(ArregloTemp)

			for Contador :=0; Contador < len(Variables.ArregloArchivo); Contador++ {

				Variables.ArregloArchivo[Contador] = Metodos.Trim(Variables.ArregloArchivo[Contador])

				if Variables.ArregloArchivo[Contador] != "" {

					if strings.HasPrefix(Variables.ArregloArchivo[Contador], "#") {

						color.HEX("#26c941", false).Println(Variables.ArregloArchivo[Contador])
                        fmt.Println("")

					} else {

                        AnalisisComando(Variables.ArregloArchivo[Contador])

					}

				}

			}
		}
	}