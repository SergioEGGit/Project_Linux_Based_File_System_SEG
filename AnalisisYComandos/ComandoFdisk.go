
//---------------------------------------------Paquetes E Imports-------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"fmt"
		"github.com/gookit/color"
		"strconv"
		"strings"
		"unsafe"
	)

//-----------------------------------------------------Métodos----------------------------------------------------------

    func VerificarComandoFdisk() {

		//Variables
    	var CrearParticion bool
    	var ArregloParametros []string

    	//Asignación
    	CrearParticion = true
    	
        //Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos) - 1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				if ArregloParametros[0] == "delete" {

				    VerificarDeleteFdisk()
				    CrearParticion = false
				    break

				} else if ArregloParametros[0] == "add" {

                	VerificarAddFdisk()
                	CrearParticion = false
                	break

				}

			}
			
			if CrearParticion {

				VerificarCrearFdisk()
				
			}
		} else {

			color.HEX("#de4843", false).Println("Debe De Colocar Todos Los Parametros Obligatorios")
			fmt.Println("")

		}
    }

    func VerificarDeleteFdisk() {

		//Variables
		var Delete bool
		var Path bool
		var Name bool
		var ParametroExtra bool
		var ArregloParametros []string
		var ArregloArchivo []string
		var ContadorDelete int
		var ContadorPath int
		var ContadorName int

		//Asignación
		Delete = false
		Path = false
		Name = false
		ParametroExtra = false
		ContadorDelete = 0
		ContadorPath = 0
		ContadorName = 0
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

				case "delete":

					if ContadorDelete == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = strings.ToLower(Metodos.Trim(ArregloParametros[1]))

							if ArregloParametros[1] == "fast" {

								Variables.MapComandos["delete"] = "fast"
								Delete = true

							} else if ArregloParametros[1] == "full" {

								Variables.MapComandos["delete"] = "full"
								Delete = true

							} else {

								color.HEX("#de4843", false).Println("En El Parametro delete Debe De Ingresar La palabra full O fast")
								fmt.Println("")
								Delete = false

							}

							ContadorDelete++

						} else {

							Delete = false

						}

					} else {

						ContadorDelete++

					}

				case "path":

					if ContadorPath == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Path = Metodos.ExisteRuta(ArregloParametros[1])

							if Path {

								ArregloArchivo = Metodos.SplitArchivo(ArregloParametros[1])

								if len(ArregloArchivo) > 1 {

									if ArregloArchivo[1] == "dsk" {

										Variables.MapComandos["path"] = Metodos.QuitarComillas(ArregloParametros[1])
										Path = true

									} else {

										color.HEX("#de4843", false).Println("La Extension Del Archivo Debe De Ser .dsk")
										fmt.Println("")
										Path = false

									}

								} else {

									color.HEX("#de4843", false).Println("Debe Indicar La Extension Del Archivo")
									fmt.Println("")
									Path = false

								}

							}

							ContadorPath++

						} else {

							Path = false

						}

					} else {

						ContadorPath++

					}

				case "name":

					if ContadorName == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Variables.MapComandos["name"] = Metodos.Trim(ArregloParametros[1])
							Name = true

							ContadorName++

						} else {

							Name = false

						}

					} else {

						ContadorName++

					}

				default:

					ParametroExtra = true

				}
			}
		}


		if Path && Name && Delete && !ParametroExtra && ContadorPath == 1 && ContadorDelete == 1 && ContadorName == 1  {

			VerificarMBR()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "2). -delete->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "3). -name->    (Obligatorio)")
				fmt.Println("")

			}

			if !Path {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -path-> o")
				color.HEX("#de4843", false).Println("El Archivo No Existe")
				fmt.Println("")

			}

			if !Delete {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -Delete-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")
			}

			if !Name {

				color.HEX("#de4843", false).Println("No Se Encuentra el Parametro -name-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if ContadorDelete > 1 || ContadorPath > 1 || ContadorName > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}

	func VerificarAddFdisk() {

		//Variables
		var Add bool
		var Path bool
		var Name bool
		var Unit bool
		var ParametroExtra bool
		var ArregloParametros []string
		var ArregloArchivo []string
		var ContadorAdd int
		var ContadorPath int
		var ContadorName int
		var ContadorUnit int

		//Asignación
		Add = false
		Path = false
		Name = false
		Unit = true
		ParametroExtra = false
		ContadorAdd = 0
		ContadorPath = 0
		ContadorName = 0
		ContadorUnit = 0
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

				case "path":

					if ContadorPath == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Path = Metodos.ExisteRuta(ArregloParametros[1])

							if Path {

								ArregloArchivo = Metodos.SplitArchivo(ArregloParametros[1])

								if len(ArregloArchivo) > 1 {

									if ArregloArchivo[1] == "dsk" {

										Variables.MapComandos["path"] = Metodos.QuitarComillas(ArregloParametros[1])
										Path = true

									} else {

										color.HEX("#de4843", false).Println("La Extension Del Archivo Debe De Ser .dsk")
										fmt.Println("")
										Path = false

									}

								} else {

									color.HEX("#de4843", false).Println("Debe Indicar La Extension Del Archivo")
									fmt.Println("")
									Path = false

								}

							}

							ContadorPath++

						} else {

							Path = false

						}

					} else {

						ContadorPath++

					}

				case "name":

					if ContadorName == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Variables.MapComandos["name"] = Metodos.Trim(ArregloParametros[1])
							Name = true

							ContadorName++

						} else {

							Name = false

						}

					} else {

						ContadorName++

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


		if Path && Name && Add && Unit && !ParametroExtra && ContadorPath == 1 && ContadorAdd == 1 && ContadorName == 1 && (ContadorUnit == 1 || ContadorUnit == 0) {

			VerificarMBR()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "2). -add->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "3). -name->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "4). -unit->    (Opcional)")
				fmt.Println("")

			}

			if !Path {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -path-> o")
				color.HEX("#de4843", false).Println("El Archivo No Existe")
				fmt.Println("")

			}

			if !Add {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -add-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")
			}

			if !Name {

				color.HEX("#de4843", false).Println("No Se Encuentra el Parametro -name-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if !Unit {

				color.HEX("#de4843", false).Println("No Se Encuentra el Parametro -unit-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if ContadorAdd > 1 || ContadorPath > 1 || ContadorName > 1 || ContadorUnit > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}

	func VerificarCrearFdisk() {

		//Variables
		var Size bool
		var Path bool
		var Name bool
		var Unit bool
		var Type bool
		var Fit bool
		var ParametroExtra bool
		var ArregloParametros []string
		var ArregloArchivo []string
		var ContadorSize int
		var ContadorPath int
		var ContadorName int
		var ContadorUnit int
		var ContadorType int
		var ContadorFit int

		//Asignación
		Size = false
		Path = false
		Name = false
		Unit = true
		Type = true
		Fit = true
		ParametroExtra = false
		ContadorSize = 0
		ContadorPath = 0
		ContadorName = 0
		ContadorUnit = 0
		ContadorType = 0
		ContadorFit = 0
		Variables.MapComandos = make(map[string]string)
		Variables.MapComandos["unit"] = "1024"
		Variables.MapComandos["type"] = "p"
		Variables.MapComandos["fit"] = "wf"

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos) - 1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				switch ArregloParametros[0] {

				case "size":

					if ContadorSize == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							var Tamanio int
							var ErrorEntero error

							Tamanio, ErrorEntero = strconv.Atoi(ArregloParametros[1])

							if ErrorEntero != nil {

								color.HEX("#de4843", false).Println("El Parametro Size Debe Ser Un Número")
								fmt.Println("")

							} else {

								if Tamanio > 0 {

									Variables.MapComandos["size"] = ArregloParametros[1]
									Size = true

								} else {

									Size = false
									color.HEX("#de4843", false).Println("El Parametro Size Debe Ser Un Número Mayor A 0")
									fmt.Println("")

								}
								ContadorSize++

							}

						} else {

							Size = false

						}
					} else {

						ContadorSize++

					}

				case "path":

					if ContadorPath == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Path = Metodos.ExisteRuta(ArregloParametros[1])

							if Path {

								ArregloArchivo = Metodos.SplitArchivo(ArregloParametros[1])

								if len(ArregloArchivo) > 1 {

									if ArregloArchivo[1] == "dsk" {

										Variables.MapComandos["path"] = Metodos.QuitarComillas(ArregloParametros[1])
										Path = true

									} else {

										color.HEX("#de4843", false).Println("La Extension Del Archivo Debe De Ser .dsk")
										fmt.Println("")
										Path = false

									}

								} else {

									color.HEX("#de4843", false).Println("Debe Indicar La Extension Del Archivo")
									fmt.Println("")
									Path = false

								}

							}

							ContadorPath++

						} else {

							Path = false

						}

					} else {

						ContadorPath++

					}

				case "name":

					if ContadorName == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Variables.MapComandos["name"] = Metodos.Trim(ArregloParametros[1])
							Name = true

							ContadorName++

						} else {

							Name = false

						}

					} else {

						ContadorName++

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

				case "type":

					if ContadorType == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = strings.ToLower(Metodos.Trim(ArregloParametros[1]))

							if ArregloParametros[1] == "p" {

								Variables.MapComandos["type"] = "p"
								Type = true

							} else if ArregloParametros[1] == "e" {

								Variables.MapComandos["type"] = "e"
								Type = true

							} else if ArregloParametros[1] == "l" {

								Variables.MapComandos["type"] = "l"
								Type = true

							} else {

								color.HEX("#de4843", false).Println("En El Parametro Type Debe De Ingresar La Letra p (Primaria) O La Letra e (Extendida) O La Letra l (Logica)")
								fmt.Println("")
								Type = false

							}

							ContadorType++

						} else {

							Type = false

						}

					} else {

						ContadorType++

					}

				case "fit":

					if ContadorFit == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = strings.ToLower(Metodos.Trim(ArregloParametros[1]))

							if ArregloParametros[1] == "bf" {

								Variables.MapComandos["fit"] = "bf"
								Fit = true

							} else if ArregloParametros[1] == "ff" {

								Variables.MapComandos["fit"] = "ff"
								Fit = true

							} else if ArregloParametros[1] == "wf" {

								Variables.MapComandos["fit"] = "wf"
								Fit = true

							} else {

								color.HEX("#de4843", false).Println("En El Parametro Fit Debe De Ingresar Las Letras bf (Best Fit) O Las Letras ff (First Fit) O Las Letras wf (Worst Fit)")
								fmt.Println("")
								Fit = false

							}

							ContadorFit++

						} else {

							Fit = false

						}

					} else {

						ContadorFit++

					}


				default:

					ParametroExtra = true

				}
			}
		}


		if Path && Size && Name && Unit && Fit && Type && !ParametroExtra && ContadorPath == 1 && ContadorSize == 1 && ContadorName == 1 && (ContadorUnit == 1 || ContadorUnit == 0) && (ContadorFit == 1 || ContadorFit == 0) && (ContadorType == 1 || ContadorType == 0){

			VerificarMBR()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "2). -size->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "3). -name->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "4). -unit->    (Opcional)")
				color.HEX("#de4843", false).Println( "5). -fit->    (Opcional)")
				color.HEX("#de4843", false).Println( "6). -type->    (Opcional)")
				fmt.Println("")

			}

			if !Path {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -path-> o")
				color.HEX("#de4843", false).Println("No Existe El Archivo")
				fmt.Println("")

			}

			if !Size {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -size-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")
			}

			if !Name {

				color.HEX("#de4843", false).Println("No Se Encuentra el Parametro -name-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if !Unit {

				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis En El Paremtro -unit->")
				fmt.Println("")

			}

			if !Fit {

				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis En El Paremtro -fit->")
				fmt.Println("")

			}

			if !Type {

				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis En El Paremtro -type->")
				fmt.Println("")

			}

			if ContadorSize > 1 || ContadorPath > 1 || ContadorName > 1 || ContadorUnit > 1 || ContadorFit > 1 || ContadorType > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}

	func VerificarMBR() {

		//Variables
		var MBRAuxiliar Variables.MBREstructura
		var Bandera bool

		//Asignacion
		MBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioArraglo(Variables.MapComandos["path"])

		if Bandera {

			if !VerificarNombreParticion(MBRAuxiliar) {

				if Variables.MapComandos["type"] == "l" {


				} else if Variables.MapComandos["type"] == "p" || Variables.MapComandos["type"] == "e" {

					if !VerificarSizeParticion(MBRAuxiliar) {

						if !VerificarTipoParticion(MBRAuxiliar) {



						}

					} else {

						color.HEX("#de4843", false).Println("Error No Hay Espacio Suficiente Para Crear La Partición")

					}

				}

			} else {

				color.HEX("#de4843", false).Println("Error Ya Existe Una Particion Con El Nombre Indicado")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando fdisk")
			color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
			fmt.Println("")

		}

	}

	func VerificarNombreParticion(MBRAuxiliar Variables.MBREstructura) bool {

		//Variables
		var Nombre string
		var Bandera bool

		//Asignación
		Nombre = Metodos.Trim(Variables.MapComandos["name"])
	    Bandera = false


		if MBRAuxiliar.Particion1MBR.SizePart != 0 {

			if Nombre == string(MBRAuxiliar.Particion1MBR.NamePart[:]) {

				Bandera = true

			}

		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			if Nombre == string(MBRAuxiliar.Particion2MBR.NamePart[:]) {

				Bandera = true

			}

		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			if Nombre == string(MBRAuxiliar.Particion3MBR.NamePart[:]) {

				Bandera = true

			}

		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			if Nombre == string(MBRAuxiliar.Particion4MBR.NamePart[:]) {

				Bandera = true

			}

		}

		return Bandera

	}

	func VerificarSizeParticion(MBRAuxiliar Variables.MBREstructura) bool {

	    //Variables
		var Bandera bool
		var SizeUsado int64
		var SizeDisco int64

		//Asignacion
		Bandera = false
		SizeUsado = 0
		SizeDisco = MBRAuxiliar.SizeMbr

		if MBRAuxiliar.Particion1MBR.SizePart != 0 {

			SizeUsado += MBRAuxiliar.Particion1MBR.SizePart

		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			SizeUsado += MBRAuxiliar.Particion2MBR.SizePart

		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			SizeUsado += MBRAuxiliar.Particion3MBR.SizePart

		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			SizeUsado += 89

		}

		if SizeUsado >= SizeDisco {

			Bandera = true

		} else {

			Bandera = false

		}

		return Bandera

	}

    func VerificarTipoParticion(MBRAuxiliar Variables.MBREstructura) bool {

    	//Variables
    	var Bandera bool
    	var Extendidas int
    	var Primarias int
    	var Total int

    	//Asignación
    	Bandera = false
    	Extendidas = 0
    	Primarias = 0
    	Total = 0

    	if MBRAuxiliar.Particion1MBR.SizePart != 0 {

    		if string(MBRAuxiliar.Particion1MBR.TipoPart) == "p" {

    			Primarias++

			} else if string(MBRAuxiliar.Particion1MBR.TipoPart) == "e" {

				Extendidas++

			}
		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			if string(MBRAuxiliar.Particion2MBR.TipoPart) == "p" {

				Primarias++

			} else if string(MBRAuxiliar.Particion2MBR.TipoPart) == "e" {

				Extendidas++

			}
		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			if string(MBRAuxiliar.Particion3MBR.TipoPart) == "p" {

				Primarias++

			} else if string(MBRAuxiliar.Particion3MBR.TipoPart) == "e" {

				Extendidas++

			}
		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			if string(MBRAuxiliar.Particion4MBR.TipoPart) == "p" {

				Primarias++

			} else if string(MBRAuxiliar.Particion4MBR.TipoPart) == "e" {

				Extendidas++

			}
		}

		Total = Primarias + Extendidas

		if Total < 4 {

			if Variables.MapComandos["type"] == "p" {

				Bandera = false

			} else if Variables.MapComandos["type"] == "e" {

				if Extendidas == 1 {

					color.HEX("#de4843", false).Println("Unicamente Se Pueden Crear Una Particion Extendida En El Disco")
					Bandera = true

				} else {

					Bandera = false

				}

			}

		} else if Total == 4 {

			color.HEX("#de4843", false).Println("Ya No Se Pueden Crear Mas Particiones Ya Que Existen 4")
			Bandera = true

		}

		return Bandera
	}


	func ComandoFdiskCrearParticion(MBRAuxiliar Variables.MBREstructura) {

        //Variables
		var ParticionAuxiliar Variables.ParticionEstructura
		var Bandera bool
		var Size int
		var Unit int
		var SizeTotal int

		//Asignación
		ParticionAuxiliar = Variables.ParticionEstructura{}
		Bandera = false
		Size = 0
		Unit = 0
		SizeTotal = 0

		//Crear Nueva Particion
		//Estado De La Particion
		ParticionAuxiliar.StatusPart = 'n'

		//Tipo De Particion
		if Variables.MapComandos["type"] == "p" {

			ParticionAuxiliar.TipoPart = 'p'
			Bandera = true

		} else if Variables.MapComandos["type"] == "e" {

			ParticionAuxiliar.TipoPart = 'e'
			Bandera = true

		} else if Variables.MapComandos["type"] == "l" {

			color.HEX("#de4843", false).Println("No Puede Crear Particiones Logicas Ya Que No Existe Un Partición Extendida")
            fmt.Println("")
			Bandera = false
		}

		//Fit De Particion
		if Variables.MapComandos["fit"] == "bf" {

			ParticionAuxiliar.FitPart = 'b'

		} else if Variables.MapComandos["fit"] == "ff" {

			ParticionAuxiliar.FitPart = 'f'

		} else if Variables.MapComandos["fit"] == "wf" {

			ParticionAuxiliar.FitPart = 'w'

		}

		//Inicio De La Partición
		ParticionAuxiliar.InicioPart = int64(unsafe.Sizeof(Variables.MBREstructura{}) + 1)

		//Tamaño De La Partición
		Size, _ = strconv.Atoi(Variables.MapComandos["size"])
		Unit, _ = strconv.Atoi(Variables.MapComandos["unit"])
		SizeTotal = Size * Unit
		ParticionAuxiliar.SizePart = int64(SizeTotal)

		//Nombre De La Partición
		copy(ParticionAuxiliar.NamePart[:], Variables.MapComandos["name"])

		if Bandera {

			MBRAuxiliar.Particion1MBR = ParticionAuxiliar
			Metodos.EscribirArchivoBinarioArreglo(MBRAuxiliar)
			color.Success.Println("Particion Creada Con Exito")

		}

	}

	func ComandoFdisk(MBRAuxiliar Variables.MBREstructura) {

		//Variables
		var Primaria int
		var Extendida int

		//Asignación
		Primaria = 0
		Extendida = 0

		//Verificar Si No Existe Particiones

		//Particion1
		if string(MBRAuxiliar.Particion1MBR.TipoPart) != ""  {

		    if string(MBRAuxiliar.Particion1MBR.TipoPart) == "p" {

		    	Primaria++

			} else if string(MBRAuxiliar.Particion1MBR.TipoPart) == "e" {

				Extendida++

			}

		}

		//Particion2
		if string(MBRAuxiliar.Particion2MBR.TipoPart) != ""  {

			if string(MBRAuxiliar.Particion2MBR.TipoPart) == "p" {

				Primaria++

			} else if string(MBRAuxiliar.Particion2MBR.TipoPart) == "e" {

				Extendida++

			}

		}

		//Particion3
		if string(MBRAuxiliar.Particion3MBR.TipoPart) != ""  {

			if string(MBRAuxiliar.Particion3MBR.TipoPart) == "p" {

				Primaria++

			} else if string(MBRAuxiliar.Particion3MBR.TipoPart) == "e" {

				Extendida++

			}

		}
		//Particion1
		if string(MBRAuxiliar.Particion4MBR.TipoPart) != ""  {

			if string(MBRAuxiliar.Particion4MBR.TipoPart) == "p" {


			} else if string(MBRAuxiliar.Particion4MBR.TipoPart) == "e" {


			}

		}


	}