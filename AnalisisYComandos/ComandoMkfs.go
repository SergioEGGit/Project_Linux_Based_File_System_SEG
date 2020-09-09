
//---------------------------------------------Paquetes E Imports-------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"fmt"
		"github.com/asaskevich/govalidator"
		"github.com/gookit/color"
		"strconv"
		"strings"
	)

//-----------------------------------------------------Métodos----------------------------------------------------------

	func VerificarComandoMkfs() {

		//Variables
		var FormatearParticion bool
		var ArregloParametros []string

		//Asignación
		FormatearParticion = true

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos) - 1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				if ArregloParametros[0] == "add" {

					VerificarMkfsAdd()
					FormatearParticion = false
					break

				}

			}

			if FormatearParticion {

				VerificarMkfsFormateo()

			}

		} else {

			color.HEX("#de4843", false).Println("Debe De Colocar Todos Los Parametros Obligatorios")
			fmt.Println("")

		}
	}

	func VerificarMkfsFormateo() {

		//Variables
		var Id bool
		var Tipo bool
		var ParametroExtra bool
		var ArregloParametros []string
		var ContadorId int
		var ContadorTipo int

		//Asignación
		Id = false
		Tipo = true
		ParametroExtra = false
		ContadorId = 0
		ContadorTipo = 0
		Variables.MapComandos = make(map[string]string)
		Variables.MapComandos["tipo"] = "full"

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos) - 1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				switch ArregloParametros[0] {

				case "id":

					if ContadorId == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

								Variables.MapComandos["id"] = ArregloParametros[1]
								Id = true
								ContadorId++

						} else {

							Id = false

						}
					} else {

						ContadorId++

					}

				case "tipo":

					if ContadorTipo == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = strings.ToLower(Metodos.Trim(ArregloParametros[1]))

							if ArregloParametros[1] == "full" {

								Variables.MapComandos["type"] = "full"
								Tipo = true

							} else if ArregloParametros[1] == "fast" {

								Variables.MapComandos["type"] = "fast"
								Tipo = true

							} else {

								color.HEX("#de4843", false).Println("En El Parametro Type Debe De Ingresar La palabra full O La palabra fast")
								fmt.Println("")
								Tipo = false

							}

							ContadorTipo++

						} else {

							Tipo = false

						}

					} else {

						ContadorTipo++

					}

				default:

					ParametroExtra = true

				}
			}
		}


		if Id && Tipo && !ParametroExtra && ContadorId == 1 && (ContadorTipo == 1 || ContadorTipo == 0) {

			ComandoMkfs()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -id->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "4). -tipo->    (Opcional)")
				fmt.Println("")

			}

			if !Id {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -id-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if !Tipo {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -tipo-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")
			}

			if ContadorId > 1 || ContadorTipo > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}

	func VerificarMkfsAdd() {

		//Variables
		var Add bool
		var Unit bool
		var ParametroExtra bool
		var ArregloParametros []string
		var ContadorAdd int
		var ContadorUnit int
		var ContadorAuxiliar int

		//Asignación
		Add = false
		Unit = true
		ParametroExtra = false
		ContadorAdd = 0
		ContadorUnit = 0
		ContadorAuxiliar = 0
		Variables.MapComandos = make(map[string]string)
		Variables.MapComandos["unit"] = "1024"

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos) - 1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				switch ArregloParametros[0] {

				case "add":

					ContadorAuxiliar = Contador + 1

					if ContadorAuxiliar < len(Variables.ArregloComandos) {

						//Obtener Parametro
						Variables.ArregloComandos[ContadorAuxiliar] = Metodos.Trim(Variables.ArregloComandos[ContadorAuxiliar])
						ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[ContadorAuxiliar])

						//Verificar Si ES Digito
						if govalidator.IsInt(ArregloParametros[0]) {

							ArregloParametros = append(ArregloParametros, "-" + ArregloParametros[0])
							Contador += 1

						} else {

							//Obtener Parametro
							Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
							ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

						}

					}

					if ContadorAdd == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							var Tamanio int
							var ErrorEntero error

							Tamanio, ErrorEntero = strconv.Atoi(ArregloParametros[1])

							if ErrorEntero != nil {

								color.HEX("#de4843", false).Println("El Parametro add Debe Ser Un Número")
								fmt.Println("")

							} else {

								if Tamanio > 0 || Tamanio < 0 {

									Variables.MapComandos["add"] = ArregloParametros[1]
									Add = true

								} else {

									Add = false
									color.HEX("#de4843", false).Println("El Parametro Add No Puede Ser 0")
									fmt.Println("")

								}
								ContadorAdd++
							}

						} else {

							Add = false

						}

					} else {

						ContadorAdd++

					}

				case "unit":

					if ContadorUnit == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = strings.ToLower(Metodos.Trim(ArregloParametros[1]))

							if ArregloParametros[1] == "k" {

								Variables.MapComandos["unit"] = "1024"
								Unit = true

							} else if ArregloParametros[1] == "m" {

								Variables.MapComandos["unit"] = "1048576"
								Unit = true

							} else if ArregloParametros[1] == "b" {

								Variables.MapComandos["unit"] = "1"
								Unit = true

							} else {

								color.HEX("#de4843", false).Println("En El Parametro Unit Debe De Ingresar La Letra m (Megabytes) O La Letra k (Kylobytes) O La Letra b (Bytes)")
								fmt.Println("")
								Unit = false

							}

							ContadorUnit++

						} else {

							Unit = false

						}

					} else {

						ContadorUnit++

					}

				default:

					ParametroExtra = true

				}
			}
		}


		if Add && Unit && !ParametroExtra && ContadorAdd == 1 && (ContadorUnit == 1 || ContadorUnit == 0) {

			ComandoMkfsAdd()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println( "2). -add->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "4). -unit->    (Opcional)")
				fmt.Println("")

			}

			if !Add {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -add-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")
			}

			if !Unit {

				color.HEX("#de4843", false).Println("No Se Encuentra el Parametro -unit-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if ContadorAdd > 1 || ContadorUnit > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}
	
	func ComandoMkfs() {
		
		//Variables
		var Bandera bool

		//Asignacion
		Bandera = false

		//Verificar Particon Montada
		for Contador := 0; Contador < len(Variables.ArregloParticionesMontadas); Contador++ {

			if strings.EqualFold(Variables.MapComandos["id"], Variables.ArregloParticionesMontadas[Contador].IdentificadorMount) {

				Bandera = true

			}

		}

		if Bandera {

			println("Hago Formateo")

		} else {

			color.HEX("#de4843", false).Println("No Existe El Id Indicado")
			fmt.Println("")

		}

		
	}

	func ComandoMkfsAdd() {

		print("Comando Mkfs")

	}