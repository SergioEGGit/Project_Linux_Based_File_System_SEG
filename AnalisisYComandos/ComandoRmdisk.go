
//----------------------------------------------Paquetes E Imports------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"bufio"
		"fmt"
		"github.com/gookit/color"
		"os"
		"strings"
	)

//----------------------------------------------------Métodos-----------------------------------------------------------

	func VerificarComandoRmdisk() {

		//Variables
	    var Path bool
		var ParametroExtra bool
		var ArregloParametros []string
		var ContadorPath int

		//Asignación
		Path = false
		ParametroExtra = false
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

					if ContadorPath == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Path = Metodos.ExisteRuta(ArregloParametros[1])

							if Path {

								Variables.MapComandos["path"] = Metodos.QuitarComillas(ArregloParametros[1])

							}

							ContadorPath++

						} else {

							Path = false

						}

					} else {

						ContadorPath++

					}

				default:

					ParametroExtra = true

				}
			}
		}


		if Path && !ParametroExtra && ContadorPath == 1 {

			ComandoRmdisk()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println( "Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				fmt.Println("")

			}

			if Variables.NoExisteArchivo {

				color.HEX("#de4843", false).Println("No Existe El Archivo Indicado")
				fmt.Println("")

			}

			if !Path {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -path-> O")
				color.HEX("#de4843", false).Println("Existe Error En Sintaxis")
				fmt.Println("")

			}

			if ContadorPath > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}

	func ComandoRmdisk() {

        //Variables
		var Lectura *bufio.Reader
		var Cadena string
		var AvisoError error

		//Asignaciones
		Lectura = bufio.NewReader(os.Stdin)

		//Ciclo Mensaje
		for {

			color.HEX("#c9265c", false).Print("Seguro Que Desea Eliminar El Disco?  s/n: ")
			Cadena, AvisoError = Lectura.ReadString('\n')
			_ = AvisoError

			Cadena = strings.ToLower(Cadena)
			Cadena = Metodos.Trim(Cadena)

			if Cadena == "s" {

				//Eliminar Archivo
				AvisoError = os.Remove(Metodos.Trim(Variables.MapComandos["path"]))

				//Catch Error
				if AvisoError != nil {

					color.HEX("#de4843", false).Println("Error: Disco No Eliminado")
					fmt.Println("")
					break

				} else {

                    color.Success.Println("Disco Eliminado Con Exito!")
                    fmt.Println("")
                    break

				}

			} else if Cadena == "n" {

				color.HEX("#c9265c", false).Println("Disco No Eliminado!")
			    fmt.Println("")
			    break

			} else {

				color.HEX("#de4843", false).Println("Debe De Ingresar s o n")
				fmt.Println("")

			}
		}

	}