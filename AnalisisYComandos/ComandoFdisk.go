
//---------------------------------------------Paquetes E Imports-------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"bufio"
		"bytes"
		"fmt"
		"github.com/asaskevich/govalidator"
		"github.com/gookit/color"
		"os"
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

							ArregloParametros[1] = Metodos.QuitarComillas(ArregloParametros[1])
							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Path = Metodos.ExisteRuta(ArregloParametros[1])

							if Path {

								ArregloArchivo = Metodos.SplitArchivo(ArregloParametros[1])

								if len(ArregloArchivo) > 1 {

									if ArregloArchivo[1] == "dsk" {

										Variables.MapComandos["path"] = ArregloParametros[1]
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

			VerificarMBRDelete()

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
		var ContadorAuxiliar int

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

				case "path":

					if ContadorPath == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.QuitarComillas(ArregloParametros[1])
							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Path = Metodos.ExisteRuta(ArregloParametros[1])

							if Path {

								ArregloArchivo = Metodos.SplitArchivo(ArregloParametros[1])

								if len(ArregloArchivo) > 1 {

									if ArregloArchivo[1] == "dsk" {

										Variables.MapComandos["path"] = ArregloParametros[1]
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

			VerificarMBRAdd()

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

							ArregloParametros[1] = Metodos.QuitarComillas(ArregloParametros[1])
							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Path = Metodos.ExisteRuta(ArregloParametros[1])

							if Path {

								ArregloArchivo = Metodos.SplitArchivo(ArregloParametros[1])

								if len(ArregloArchivo) > 1 {

									if ArregloArchivo[1] == "dsk" {

										Variables.MapComandos["path"] = ArregloParametros[1]
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

			VerificarMBRCrear()

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

	func VerificarMBRCrear() {

		//Variables
		var MBRAuxiliar Variables.MBREstructura
		var Bandera bool
		var BanderaExtendida bool
		var InicioListaLogica int64
		var SizeExtendida int64

		//Asignacion
		MBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioArraglo(Variables.MapComandos["path"])

		if Bandera {

			if !VerificarNombreParticionCrear(MBRAuxiliar) {

				if Variables.MapComandos["type"] == "l" {

					BanderaExtendida, InicioListaLogica, SizeExtendida = VerficiarExisteParticionExtendida(MBRAuxiliar)

				    if BanderaExtendida {

				    	if VerificarSizeParticionExtendida(InicioListaLogica, SizeExtendida) {

				    		ComandoFdiskCrearParticionLogica(InicioListaLogica, SizeExtendida)

						} else {

							color.HEX("#de4843", false).Println("Error No Hay Espacio Suficiente Para Crear La Partición Logica")
							fmt.Println("")

						}

					} else {

						color.HEX("#de4843", false).Println("Error Debe Existir Un Particion Extendida Para Crear Particiones Logicas")
						fmt.Println("")

					}

				} else if Variables.MapComandos["type"] == "p" || Variables.MapComandos["type"] == "e" {

					if !VerificarSizeParticionCrear(MBRAuxiliar) {

						if !VerificarTipoParticionCrear(MBRAuxiliar) {

							ComandoFdiskCrearParticion(MBRAuxiliar)

						}

					} else {

						color.HEX("#de4843", false).Println("Error No Hay Espacio Suficiente Para Crear La Partición")
						fmt.Println("")

					}

				}

			} else {

				color.HEX("#de4843", false).Println("Error Ya Existe Una Particion Con El Nombre Indicado")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando fdisk")
			color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
			fmt.Println("")

		}

	}

	func VerificarMBRDelete() {

		//Variables
		var MBRAuxiliar Variables.MBREstructura
		var Bandera bool
		var ExisteNombre bool
		var PrimariaExtendida bool
		var NumeroParticion int
		var InicioExtendida int64

		//Asignacion
		MBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioArraglo(Variables.MapComandos["path"])

		if Bandera {

			ExisteNombre, PrimariaExtendida, NumeroParticion, InicioExtendida = VerificarNombreParticionDelete(MBRAuxiliar)

			if ExisteNombre {

				if PrimariaExtendida {

					ComandoFdiskDeleteParticion(MBRAuxiliar, NumeroParticion)

				} else {

					ComandoFdiskDeleteParticionLogica(InicioExtendida)

				}

			} else {

				color.HEX("#de4843", false).Println("No Existe La Particion Indicada En El Disco")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando fdisk")
			color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
			fmt.Println("")

		}

	}

	func VerificarMBRAdd() {

		//Variables
		var MBRAuxiliar Variables.MBREstructura
		var Bandera bool
		var ExisteNombre bool
		var PrimariaExtendida bool
		var NumeroParticion int
		var InicioExtendida int64
		var SizeExtendida int64

		//Asignacion
		MBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioArraglo(Variables.MapComandos["path"])

		if Bandera {

			ExisteNombre, PrimariaExtendida, NumeroParticion, InicioExtendida, SizeExtendida = VerificarNombreParticionAdd(MBRAuxiliar)

			if ExisteNombre {

				if PrimariaExtendida {

					ComandoFdiskAddParticion(MBRAuxiliar, NumeroParticion)

				} else {

					ComandoFdiskAddParticionLogica(InicioExtendida, SizeExtendida)

				}

			} else {

				color.HEX("#de4843", false).Println("No Existe La Particion Indicada En El Disco")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando fdisk")
			color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
			fmt.Println("")

		}

	}

	func VerificarNombreParticionCrear(MBRAuxiliar Variables.MBREstructura) bool {

		//Variables
		var Nombre string
		var NombreArray1 string
		var NombreArray2 string
		var NombreArray3 string
		var NombreArray4 string
		var NombreExtendida string
		var Bandera bool
		var InicioExtendida int64
		var ArregloEBR []Variables.EBREstructura

		//Asignación
		Nombre = Metodos.Trim(strings.ToLower(Variables.MapComandos["name"]))
		Bandera = false
		InicioExtendida = 0
		NombreArray1 = string(bytes.Trim(MBRAuxiliar.Particion1MBR.NamePart[:], "\x00"))
		NombreArray2 = string(bytes.Trim(MBRAuxiliar.Particion2MBR.NamePart[:], "\x00"))
		NombreArray3 = string(bytes.Trim(MBRAuxiliar.Particion3MBR.NamePart[:], "\x00"))
		NombreArray4 = string(bytes.Trim(MBRAuxiliar.Particion4MBR.NamePart[:], "\x00"))

		if MBRAuxiliar.Particion1MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray1) {

				Bandera = true

				if MBRAuxiliar.Particion1MBR.TipoPart == 'e' {

					InicioExtendida = MBRAuxiliar.Particion1MBR.InicioPart

				}

			}

		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray2)  {

				Bandera = true

			}

			if MBRAuxiliar.Particion2MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion2MBR.InicioPart

			}

		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray3)   {

				Bandera = true

			}

			if MBRAuxiliar.Particion3MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion3MBR.InicioPart

			}

		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray4)   {

				Bandera = true

			}

			if MBRAuxiliar.Particion4MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion4MBR.InicioPart

			}

		}

		if InicioExtendida != 0 {

			ArregloEBR = ObtenerEBR(InicioExtendida)

			for Contador := 0; Contador < len(ArregloEBR); Contador++ {

				NombreExtendida = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

				if strings.EqualFold(Variables.MapComandos["name"], NombreExtendida) {

					Bandera = true

				}

			}

		}

		return Bandera

	}

	func VerificarSizeParticionCrear(MBRAuxiliar Variables.MBREstructura) bool {

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

			SizeUsado += MBRAuxiliar.Particion4MBR.SizePart

		}

		if SizeUsado >= SizeDisco {

			Bandera = true

		} else {

			Bandera = false

		}

		return Bandera

	}

    func VerificarTipoParticionCrear(MBRAuxiliar Variables.MBREstructura) bool {

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
					fmt.Println("")
					Bandera = true

				} else {

					Bandera = false

				}

			}

		} else if Total == 4 {

			color.HEX("#de4843", false).Println("Ya No Se Pueden Crear Mas Particiones Ya Que Existen 4")
			fmt.Println("")
			Bandera = true

		}

		return Bandera
	}

	func VerificarNombreParticionDelete(MBRAuxiliar Variables.MBREstructura) (bool, bool, int, int64) {

		//Variables
		var Nombre string
		var NombreArray1 string
		var NombreArray2 string
		var NombreArray3 string
		var NombreArray4 string
		var NombreExtendida string
		var Bandera bool
		var PrimariaExtendida bool
		var NumeroParticion int
		var InicioExtendida int64
		var ArregloEBR []Variables.EBREstructura

		//Asignación
		Nombre = Metodos.Trim(strings.ToLower(Variables.MapComandos["name"]))
		Bandera = false
		NumeroParticion = 0
		NombreArray1 = string(bytes.Trim(MBRAuxiliar.Particion1MBR.NamePart[:], "\x00"))
		NombreArray2 = string(bytes.Trim(MBRAuxiliar.Particion2MBR.NamePart[:], "\x00"))
		NombreArray3 = string(bytes.Trim(MBRAuxiliar.Particion3MBR.NamePart[:], "\x00"))
		NombreArray4 = string(bytes.Trim(MBRAuxiliar.Particion4MBR.NamePart[:], "\x00"))

		if MBRAuxiliar.Particion1MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray1) {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 1

				if MBRAuxiliar.Particion1MBR.TipoPart == 'e' {

					InicioExtendida = MBRAuxiliar.Particion1MBR.InicioPart

				}

			}

		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray2)  {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 2

			}

			if MBRAuxiliar.Particion2MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion2MBR.InicioPart

			}

		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray3)   {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 3
			}

			if MBRAuxiliar.Particion3MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion3MBR.InicioPart

			}

		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray4)   {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 4

			}

			if MBRAuxiliar.Particion4MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion4MBR.InicioPart

			}

		}

		if InicioExtendida != 0 {

			ArregloEBR = ObtenerEBR(InicioExtendida)

			for Contador := 0; Contador < len(ArregloEBR); Contador++ {

				NombreExtendida = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

				if strings.EqualFold(Variables.MapComandos["name"], NombreExtendida) {

					Bandera = true

				}

			}

		}

		return Bandera, PrimariaExtendida, NumeroParticion, InicioExtendida


	}

	func VerficiarExisteParticionExtendida(MBRAuxiliar Variables.MBREstructura) (bool, int64, int64) {

		//Variables
		var Bandera bool
		var Extendidas int
		var Primarias int
		var InicioParticion int64
		var SizeParticion int64

		//Asignación
		Bandera = false
		Extendidas = 0
		Primarias = 0

		if MBRAuxiliar.Particion1MBR.SizePart != 0 {

			if string(MBRAuxiliar.Particion1MBR.TipoPart) == "p" {

				Primarias++

			} else if string(MBRAuxiliar.Particion1MBR.TipoPart) == "e" {

				InicioParticion = MBRAuxiliar.Particion1MBR.InicioPart
				SizeParticion = MBRAuxiliar.Particion1MBR.SizePart
				Extendidas++

			}
		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			if string(MBRAuxiliar.Particion2MBR.TipoPart) == "p" {

				Primarias++

			} else if string(MBRAuxiliar.Particion2MBR.TipoPart) == "e" {

				InicioParticion = MBRAuxiliar.Particion2MBR.InicioPart
				SizeParticion = MBRAuxiliar.Particion2MBR.SizePart
				Extendidas++

			}
		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			if string(MBRAuxiliar.Particion3MBR.TipoPart) == "p" {

				Primarias++

			} else if string(MBRAuxiliar.Particion3MBR.TipoPart) == "e" {

				InicioParticion = MBRAuxiliar.Particion3MBR.InicioPart
				SizeParticion = MBRAuxiliar.Particion3MBR.SizePart
				Extendidas++

			}
		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			if string(MBRAuxiliar.Particion4MBR.TipoPart) == "p" {

				Primarias++

			} else if string(MBRAuxiliar.Particion4MBR.TipoPart) == "e" {

				InicioParticion = MBRAuxiliar.Particion4MBR.InicioPart
				SizeParticion = MBRAuxiliar.Particion4MBR.SizePart
				Extendidas++

			}
		}

		if Extendidas == 1 {

			Bandera = true

		}

		return Bandera, InicioParticion, SizeParticion

	}

	func VerificarSizeParticionExtendida(InicioListaExtendida int64, SizeExtendida int64) bool {

		//Variables
		var Bandera bool
		var SizeUsado int
		var SizeEBR int
		var EBRAuxiliar Variables.EBREstructura

		//Asignación
		Bandera = true
		SizeEBR = int(unsafe.Sizeof(Variables.EBREstructura{}))
		SizeUsado = 0

		for {

			//Leer EBR
			EBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioEBR(Variables.MapComandos["path"], InicioListaExtendida)

			//Lista Corrupta
			if !Bandera {

				return false

			}

			SizeUsado += int(EBRAuxiliar.SizeEBR) + SizeEBR
			InicioListaExtendida = EBRAuxiliar.SiguienteEBR

			if EBRAuxiliar.SiguienteEBR == -1 {

				break

			}

		}

		if SizeUsado >= int(SizeExtendida) {

			Bandera = false

		}

		return Bandera

	}

	func ObtenerEBR(InicioListaExtendida int64) []Variables.EBREstructura {

		//Variables
		var Contador int
		var Bandera bool
		var EBRAuxiliar Variables.EBREstructura
		var ArregloEBR []Variables.EBREstructura

		//Asignación
		Contador = 0

		for {

			//Leer EBR
			EBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioEBR(Variables.MapComandos["path"], InicioListaExtendida)

			//Lista Corrupta
			if !Bandera {

				return ArregloEBR

			}

			//fmt.Println("Size: ", EBRAuxiliar.SizeEBR, "Inicio: ", EBRAuxiliar.InicioEBR, "Siguiente: ", EBRAuxiliar.SiguienteEBR, "Nombre: ", string(EBRAuxiliar.NameEBR[:]))

			ArregloEBR = append(ArregloEBR, EBRAuxiliar)
			InicioListaExtendida = ArregloEBR[Contador].SiguienteEBR
			Contador++

			if EBRAuxiliar.SiguienteEBR == -1 {

				break

			}

		}

		return ArregloEBR

	}

	func VerificarNombreParticionAdd(MBRAuxiliar Variables.MBREstructura) (bool, bool, int, int64, int64) {

		//Variables
		var Nombre string
		var NombreArray1 string
		var NombreArray2 string
		var NombreArray3 string
		var NombreArray4 string
		var NombreExtendida string
		var Bandera bool
		var PrimariaExtendida bool
		var NumeroParticion int
		var InicioExtendida int64
		var SizeExtendida int64
		var ArregloEBR []Variables.EBREstructura

		//Asignación
		Nombre = Metodos.Trim(strings.ToLower(Variables.MapComandos["name"]))
		Bandera = false
		NumeroParticion = 0
		SizeExtendida = 0
		NombreArray1 = string(bytes.Trim(MBRAuxiliar.Particion1MBR.NamePart[:], "\x00"))
		NombreArray2 = string(bytes.Trim(MBRAuxiliar.Particion2MBR.NamePart[:], "\x00"))
		NombreArray3 = string(bytes.Trim(MBRAuxiliar.Particion3MBR.NamePart[:], "\x00"))
		NombreArray4 = string(bytes.Trim(MBRAuxiliar.Particion4MBR.NamePart[:], "\x00"))

		if MBRAuxiliar.Particion1MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray1) {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 1

				if MBRAuxiliar.Particion1MBR.TipoPart == 'e' {

					InicioExtendida = MBRAuxiliar.Particion1MBR.InicioPart
					SizeExtendida = MBRAuxiliar.Particion1MBR.SizePart

				}

			}

		}

		if MBRAuxiliar.Particion2MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray2)  {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 2

			}

			if MBRAuxiliar.Particion2MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion2MBR.InicioPart
				SizeExtendida = MBRAuxiliar.Particion2MBR.SizePart

			}

		}

		if MBRAuxiliar.Particion3MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray3)   {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 3
			}

			if MBRAuxiliar.Particion3MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion3MBR.InicioPart
				SizeExtendida = MBRAuxiliar.Particion3MBR.SizePart

			}

		}

		if MBRAuxiliar.Particion4MBR.SizePart != 0 {

			if strings.EqualFold(Nombre, NombreArray4)   {

				Bandera = true
				PrimariaExtendida = true
				NumeroParticion = 4

			}

			if MBRAuxiliar.Particion4MBR.TipoPart == 'e' {

				InicioExtendida = MBRAuxiliar.Particion4MBR.InicioPart
				SizeExtendida = MBRAuxiliar.Particion4MBR.SizePart

			}

		}

		if InicioExtendida != 0 {

			ArregloEBR = ObtenerEBR(InicioExtendida)

			for Contador := 0; Contador < len(ArregloEBR); Contador++ {

				NombreExtendida = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

				if strings.EqualFold(Variables.MapComandos["name"], NombreExtendida) {

					Bandera = true

				}

			}

		}

		return Bandera, PrimariaExtendida, NumeroParticion, InicioExtendida, SizeExtendida


	}

	func ComandoFdiskCrearParticion(MBRAuxiliar Variables.MBREstructura) {

        //Variables
		var ParticionAuxiliar Variables.ParticionEstructura
		var EBRAuxiliar Variables.EBREstructura
		var Bandera bool
		var Size int
		var Unit int
		var SizeTotal int

		//Asignación
		ParticionAuxiliar = Variables.ParticionEstructura{}
		EBRAuxiliar = Variables.EBREstructura{}
		Bandera = false
		Size = 0
		Unit = 0
		SizeTotal = 0

		//Crear Nuevo Disco
		Metodos.LimpiaDisco()
		Metodos.CreaDisco(int(MBRAuxiliar.SizeMbr) + 201)
		Metodos.LlenaDisco(0, int(unsafe.Sizeof(Variables.MBREstructura{})))
		Metodos.LLenarParticiones(MBRAuxiliar)
		Metodos.GeneraEspacios()

		//Crear Nueva Particion

		//Estado De La Particion
		ParticionAuxiliar.StatusPart = 'n'

		//Tipo De Particion
		if Variables.MapComandos["type"] == "p" {

			ParticionAuxiliar.TipoPart = 'p'

		} else if Variables.MapComandos["type"] == "e" {

			ParticionAuxiliar.TipoPart = 'e'

		} else if Variables.MapComandos["type"] == "l" {

			ParticionAuxiliar.TipoPart = 'l'

		}

		//Fit De Particion
		if Variables.MapComandos["fit"] == "bf" {

			ParticionAuxiliar.FitPart = 'b'

		} else if Variables.MapComandos["fit"] == "ff" {

			ParticionAuxiliar.FitPart = 'f'

		} else if Variables.MapComandos["fit"] == "wf" {

			ParticionAuxiliar.FitPart = 'w'

		}

		//Tamaño De La Partición
		Size, _ = strconv.Atoi(Variables.MapComandos["size"])
		Unit, _ = strconv.Atoi(Variables.MapComandos["unit"])
		SizeTotal = Size * Unit
		ParticionAuxiliar.SizePart = int64(SizeTotal)

		for Contador := 0; Contador <= 200 - 1; Contador++ {

			if Metodos.EspaciosDisponibles[Contador].Disponible {

				if SizeTotal <= Metodos.EspaciosDisponibles[Contador].Tamano {

					//Inicio Particion
					ParticionAuxiliar.InicioPart = int64(Metodos.EspaciosDisponibles[Contador].P1)
					Bandera = true
					break

				}

			}

		}

		//Nombre De La Partición
		copy(ParticionAuxiliar.NamePart[:], Variables.MapComandos["name"])

		if Bandera {

			//Buscar Particion Vacia

			if ParticionAuxiliar.TipoPart == 'e' {

				//Crear EBR Particion Extendida
				EBRAuxiliar.StatusEBR = 'n'
				EBRAuxiliar.FitEBR = 'w'
				EBRAuxiliar.InicioEBR = ParticionAuxiliar.InicioPart
				EBRAuxiliar.SizeEBR = 0
				EBRAuxiliar.SiguienteEBR = -1
				copy(EBRAuxiliar.NameEBR[:], "none")

				//Escribir EBR
				Metodos.EscribirArchivoBinarioEBR(EBRAuxiliar, ParticionAuxiliar.InicioPart)

			}

			if MBRAuxiliar.Particion1MBR.SizePart == 0 {

				MBRAuxiliar.Particion1MBR = ParticionAuxiliar
				Metodos.EscribirArchivoBinarioArreglo(MBRAuxiliar)
				color.Success.Println("Particion Creada Con Exito")
				fmt.Println("")

			} else if MBRAuxiliar.Particion2MBR.SizePart == 0 {

				MBRAuxiliar.Particion2MBR = ParticionAuxiliar
				Metodos.EscribirArchivoBinarioArreglo(MBRAuxiliar)
				color.Success.Println("Particion Creada Con Exito")
				fmt.Println("")

			} else if MBRAuxiliar.Particion3MBR.SizePart == 0 {

				MBRAuxiliar.Particion3MBR = ParticionAuxiliar
				Metodos.EscribirArchivoBinarioArreglo(MBRAuxiliar)
				color.Success.Println("Particion Creada Con Exito")
				fmt.Println("")

			} else if MBRAuxiliar.Particion4MBR.SizePart == 0 {

				MBRAuxiliar.Particion4MBR = ParticionAuxiliar
				Metodos.EscribirArchivoBinarioArreglo(MBRAuxiliar)
				color.Success.Println("Particion Creada Con Exito")
				fmt.Println("")

			}

   		} else {

			color.HEX("#de4843", false).Println("No Existe Espacio Disponible Para La Particion Indicada")
			fmt.Println("")
		}

	}

	func ComandoFdiskCrearParticionLogica(InicioListaLogica int64, SizeExtendida int64) {

		//Variables
		var Size int
		var Unit int
		var SizeTotal int
		var ContadorAuxiliar int
		var InicioPart int64
		var InicioReal int64
		var SizePart int64
		var Bandera bool
		var SiCambio int
		var ParticionAuxiliar Variables.EBREstructura
		var ArregloEBR []Variables.EBREstructura
		var ArregloAntes []Variables.EBREstructura
		var ArregloDespues []Variables.EBREstructura

		//Asignación
		Bandera = false
		ParticionAuxiliar = Variables.EBREstructura{}
		ArregloEBR = make([]Variables.EBREstructura, 0)
		ArregloEBR = ObtenerEBR(InicioListaLogica)

		// Crear Nueva Particion

		// Estado De La Particion
		ParticionAuxiliar.StatusEBR = 'n'

		//Fit De Particion
		if Variables.MapComandos["fit"] == "bf" {

			ParticionAuxiliar.FitEBR = 'b'

		} else if Variables.MapComandos["fit"] == "ff" {

			ParticionAuxiliar.FitEBR = 'f'

		} else if Variables.MapComandos["fit"] == "wf" {

			ParticionAuxiliar.FitEBR = 'w'

		}

		// Siguiente Particion
		ParticionAuxiliar.SiguienteEBR = -1

		//Nombre De La Partición
		copy(ParticionAuxiliar.NameEBR[:], Variables.MapComandos["name"])

		// Crear Disco Virtual
		Metodos.LimpiaDisco()
		Metodos.CreaDisco(int(SizeExtendida))

		// Rellenar Particiones Existentes
		for Contador := 0; Contador < len(ArregloEBR); Contador++ {

			InicioPart = ArregloEBR[Contador].InicioEBR
			InicioReal = InicioPart - InicioListaLogica
			SizePart = ArregloEBR[Contador].SizeEBR
			SizeTotal = 0
			if SizePart != 0 {

				SizeTotal = int(SizePart + int64(unsafe.Sizeof(Variables.EBREstructura{})))

			}

			Metodos.LlenaDisco(int(InicioReal), SizeTotal)
		}

		Metodos.GeneraEspacios()

		//Tamaño De La Partición
		Size, _ = strconv.Atoi(Variables.MapComandos["size"])
		Unit, _ = strconv.Atoi(Variables.MapComandos["unit"])
		SizeTotal = Size * Unit
		ParticionAuxiliar.SizeEBR = int64(SizeTotal)
		SizeTotal += int(unsafe.Sizeof(Variables.EBREstructura{}))

		for Contador := 0; Contador <= 200 - 1; Contador++ {

			if Metodos.EspaciosDisponibles[Contador].Disponible {

				if SizeTotal <= Metodos.EspaciosDisponibles[Contador].Tamano {

					//Inicio Particion

					ParticionAuxiliar.InicioEBR = int64(Metodos.EspaciosDisponibles[Contador].P1) + InicioListaLogica
					Bandera = true
					break

				}

			}

		}

		if Bandera {

			//Reorganizar EBRS
			for Contador := 0; Contador < len(ArregloEBR); Contador++ {

				if ArregloEBR[Contador].InicioEBR < ParticionAuxiliar.InicioEBR {

					SiCambio = 0

				} else if ArregloEBR[Contador].InicioEBR > ParticionAuxiliar.InicioEBR {

					SiCambio = 1

				}

				if SiCambio == 0 {

					ArregloAntes = append(ArregloAntes, ArregloEBR[Contador])

				} else if SiCambio == 1 {

					ArregloDespues = append(ArregloDespues, ArregloEBR[Contador])

				}

			}

			// Escribir Nuevo Arreglo
			ArregloEBR = make([]Variables.EBREstructura, 0)

			if len(ArregloAntes) != 0 {

				for Contador := 0; Contador < len(ArregloAntes); Contador++ {

					ArregloEBR = append(ArregloEBR, ArregloAntes[Contador])
					ContadorAuxiliar = Contador

				}

				ArregloEBR[ContadorAuxiliar].SiguienteEBR = ParticionAuxiliar.InicioEBR

			}

			if len(ArregloDespues) != 0 {

				// Cambiar Puntero
				ParticionAuxiliar.SiguienteEBR = ArregloDespues[0].InicioEBR
				ArregloEBR = append(ArregloEBR, ParticionAuxiliar)

				for Contador := 0; Contador < len(ArregloDespues); Contador++ {

					ArregloEBR = append(ArregloEBR, ArregloDespues[Contador])

				}

			} else {

				// Agregar Ultima Particion
				ArregloEBR = append(ArregloEBR, ParticionAuxiliar)

			}

			// Escribir EBR

			for Contador := 0; Contador < len(ArregloEBR); Contador++ {

				//Escribir EBR
				Metodos.EscribirArchivoBinarioEBR(ArregloEBR[Contador], ArregloEBR[Contador].InicioEBR)

			}

			color.Success.Println("Particion Creada Con Exito!")



			ArregloEBR = make([]Variables.EBREstructura, 0)
			ArregloEBR = ObtenerEBR(InicioListaLogica)

			// Crear Disco Virtual
			Metodos.LimpiaDisco()
			Metodos.CreaDisco(int(SizeExtendida))

			// Rellenar Particiones Existentes
			for Contador := 0; Contador < len(ArregloEBR); Contador++ {

				InicioPart = ArregloEBR[Contador].InicioEBR
				InicioReal = InicioPart - InicioListaLogica
				SizePart = ArregloEBR[Contador].SizeEBR
				SizeTotal = 0
				if SizePart != 0 {

					SizeTotal = int(SizePart + int64(unsafe.Sizeof(Variables.EBREstructura{})))

				}

				Metodos.LlenaDisco(int(InicioReal), SizeTotal)
			}

			Metodos.GeneraEspacios()
			fmt.Println("")
			fmt.Println("Espacios Vacios:")
			Metodos.MostrarEspacios()
			fmt.Println("")

		} else {

			color.HEX("#de4843", false).Println("No Existe Espacio Disponible Para La Particion Indicada")
			fmt.Println("")

		}

	}

	func ComandoFdiskDeleteParticion(MBRAuxiliar Variables.MBREstructura, NumeroParticion int) {

		//Variables
		var Lectura *bufio.Reader
		var Cadena string
		var AvisoError error
		var PosicionFinal int64

		//Asignaciones
		Lectura = bufio.NewReader(os.Stdin)

		//Ciclo Mensaje
		for {

			color.HEX("#c9265c", false).Print("Seguro Que Desea Eliminar La Particion?  s/n: ")
			fmt.Print("")
			Cadena, AvisoError = Lectura.ReadString('\n')
			_ = AvisoError

			Cadena = strings.ToLower(Cadena)
			Cadena = Metodos.Trim(Cadena)

			if Cadena == "s" {

				if NumeroParticion == 1 {

					if Variables.MapComandos["delete"] == "full" {

						PosicionFinal = MBRAuxiliar.Particion1MBR.InicioPart + MBRAuxiliar.Particion1MBR.SizePart
						Metodos.EscribirArchivoBinarioArregloDelete(MBRAuxiliar, MBRAuxiliar.Particion1MBR.InicioPart, PosicionFinal)
						MBRAuxiliar.Particion1MBR = Variables.ParticionEstructura{}

					} else if Variables.MapComandos["delete"] == "fast" {

						MBRAuxiliar.Particion1MBR = Variables.ParticionEstructura{}

					}


				} else if NumeroParticion == 2 {

					if Variables.MapComandos["delete"] == "full" {

						PosicionFinal = MBRAuxiliar.Particion2MBR.InicioPart + MBRAuxiliar.Particion2MBR.SizePart
						Metodos.EscribirArchivoBinarioArregloDelete(MBRAuxiliar, MBRAuxiliar.Particion2MBR.InicioPart, PosicionFinal)
						MBRAuxiliar.Particion2MBR = Variables.ParticionEstructura{}

					} else if Variables.MapComandos["delete"] == "fast" {

						MBRAuxiliar.Particion2MBR = Variables.ParticionEstructura{}

					}

				} else if NumeroParticion == 3 {

					if Variables.MapComandos["delete"] == "full" {

						PosicionFinal = MBRAuxiliar.Particion3MBR.InicioPart + MBRAuxiliar.Particion3MBR.SizePart
						Metodos.EscribirArchivoBinarioArregloDelete(MBRAuxiliar, MBRAuxiliar.Particion3MBR.InicioPart, PosicionFinal)
						MBRAuxiliar.Particion3MBR = Variables.ParticionEstructura{}

					} else if Variables.MapComandos["delete"] == "fast" {

						MBRAuxiliar.Particion3MBR = Variables.ParticionEstructura{}

					}

				} else if NumeroParticion == 4 {

					if Variables.MapComandos["delete"] == "full" {

						PosicionFinal = MBRAuxiliar.Particion4MBR.InicioPart + MBRAuxiliar.Particion4MBR.SizePart
						Metodos.EscribirArchivoBinarioArregloDelete(MBRAuxiliar, MBRAuxiliar.Particion4MBR.InicioPart, PosicionFinal)
						MBRAuxiliar.Particion4MBR = Variables.ParticionEstructura{}

					} else if Variables.MapComandos["delete"] == "fast" {

						MBRAuxiliar.Particion4MBR = Variables.ParticionEstructura{}

					}

				}

				Metodos.EscribirArchivoBinarioArreglo(MBRAuxiliar)
				color.Success.Println("Particion Eliminada Con Exito")
				fmt.Println("")
				break

			} else if Cadena == "n" {

				color.HEX("#c9265c", false).Println("Particion No Eliminada!")
				fmt.Println("")
				break

			} else {

				color.HEX("#de4843", false).Println("Debe De Ingresar s o n")
				fmt.Println("")

			}
		}

	}

	func ComandoFdiskDeleteParticionLogica(InicioExtendida int64) {

		//Variables
		var Lectura *bufio.Reader
		var Cadena string
		var NombreExtendida string
		var ContadorAuxiliar int
		var SiCambio bool
		var EBREliminado bool
		var AvisoError error
		var EBRAuxiliar Variables.EBREstructura
		var EBREliminar Variables.EBREstructura
		var ArregloEBR []Variables.EBREstructura
		var ArregloAntes []Variables.EBREstructura
		var ArregloDespues []Variables.EBREstructura

		//Asignaciones
		Lectura = bufio.NewReader(os.Stdin)
		ArregloEBR = ObtenerEBR(InicioExtendida)
		SiCambio = true
		EBREliminado = false

		//Ciclo Mensaje
		for {

			color.HEX("#c9265c", false).Print("Seguro Que Desea Eliminar La Particion?  s/n: ")
			fmt.Print("")
			Cadena, AvisoError = Lectura.ReadString('\n')
			_ = AvisoError

			Cadena = strings.ToLower(Cadena)
			Cadena = Metodos.Trim(Cadena)

			if Cadena == "s" {

				//Reorganizar EBRS
				for Contador := 0; Contador < len(ArregloEBR); Contador++ {

					NombreExtendida = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

					if strings.EqualFold(NombreExtendida, Variables.MapComandos["name"])  {

						EBREliminar = ArregloEBR[Contador]
						SiCambio = false
						EBREliminado = true

					}

					if SiCambio {

						if !EBREliminado {

							ArregloAntes = append(ArregloDespues, ArregloEBR[Contador])

						}

					} else {

						if !EBREliminado {

							ArregloDespues = append(ArregloDespues, ArregloEBR[Contador])

						}


					}

					EBREliminado = false

				}

				// Escribir Nuevo Arreglo
				ArregloEBR = make([]Variables.EBREstructura, 0)

				if len(ArregloAntes) >= 1 {

					for Contador := 0; Contador < len(ArregloAntes); Contador++ {

						ArregloEBR = append(ArregloEBR, ArregloAntes[Contador])
						ContadorAuxiliar = Contador

					}

					if len(ArregloDespues) != 0 {

						ArregloEBR[ContadorAuxiliar].SiguienteEBR = ArregloDespues[0].InicioEBR

					} else {

						ArregloEBR[ContadorAuxiliar].SiguienteEBR = -1

					}

				} else {

					//Crear EBR Particion Extendida
					EBRAuxiliar.StatusEBR = 'n'
					EBRAuxiliar.FitEBR = 'w'
					EBRAuxiliar.InicioEBR = InicioExtendida
					EBRAuxiliar.SizeEBR = 0

					if len(ArregloDespues) > 0 {

						EBRAuxiliar.SiguienteEBR = ArregloDespues[0].InicioEBR

					} else {

						EBRAuxiliar.SiguienteEBR = -1

					}

					copy(EBRAuxiliar.NameEBR[:], "none")

					ArregloEBR = append(ArregloEBR, EBRAuxiliar)

				}

				if len(ArregloDespues) != 0 {

			  		for Contador := 0; Contador < len(ArregloDespues); Contador++ {

						ArregloEBR = append(ArregloEBR, ArregloDespues[Contador])

					}

				}

				//Tipo De Borrado

				if Variables.MapComandos["delete"] == "full" {

					if len(ArregloEBR) > 1 {

						Metodos.EscribirArchivoBinarioEBRDelete(EBREliminar, EBREliminar.InicioEBR)

					}

				} else if Variables.MapComandos["delete"] == "fast" {

					// Nada

				}

				// Escribir EBR

				for Contador := 0; Contador < len(ArregloEBR); Contador++ {

					//Escribir EBR
					Metodos.EscribirArchivoBinarioEBR(ArregloEBR[Contador], ArregloEBR[Contador].InicioEBR)

				}

				color.Success.Println("Particion Eliminada Con Exito!")

				var InicioPart int64
				var InicioReal int64
				var SizePart int64
				var SizeTotal int

				// Crear Disco Virtual
				Metodos.LimpiaDisco()
				Metodos.CreaDisco(25500)
				ArregloEBR = make([]Variables.EBREstructura, 0)

				ArregloEBR = ObtenerEBR(InicioExtendida)

				// Rellenar Particiones Existentes
				for Contador := 0; Contador < len(ArregloEBR); Contador++ {

					InicioPart = ArregloEBR[Contador].InicioEBR
					InicioReal = InicioPart - InicioExtendida
					SizePart = ArregloEBR[Contador].SizeEBR
					SizeTotal = 0
					if SizePart != 0 {

						SizeTotal = int(SizePart + int64(unsafe.Sizeof(Variables.EBREstructura{})))

					}

					Metodos.LlenaDisco(int(InicioReal), SizeTotal)

				}

				Metodos.GeneraEspacios()
				fmt.Println("")
				fmt.Println("Espacios Vacios:")
				Metodos.MostrarEspacios()
				fmt.Println("")

				break

			} else if Cadena == "n" {

				color.HEX("#c9265c", false).Println("Particion No Eliminada!")
				fmt.Println("")
				break

			} else {

				color.HEX("#de4843", false).Println("Debe De Ingresar s o n")
				fmt.Println("")

			}
		}

	}

	func ComandoFdiskAddParticion(MBRAuxiliar Variables.MBREstructura, NumeroParticion int) {

		//Variables
		var CantidadAdd int
		var UnitAdd int
		var NuevoSize int
		var Bandera bool
		var MenosEspacio bool
		var MBRModificar Variables.ParticionEstructura

		//Asignaciones
		CantidadAdd, _ = strconv.Atoi(Variables.MapComandos["add"])
		UnitAdd, _ = strconv.Atoi(Variables.MapComandos["unit"])
		CantidadAdd = CantidadAdd * UnitAdd
		MBRModificar = Variables.ParticionEstructura{}
		Bandera = false
		MenosEspacio = false

		if NumeroParticion == 1 {

			NuevoSize = int(MBRAuxiliar.Particion1MBR.SizePart) + CantidadAdd
			MBRModificar = MBRAuxiliar.Particion1MBR

			if CantidadAdd > 0 {

				//Crear Disco
				Metodos.LimpiaDisco()
				Metodos.CreaDisco(int(MBRAuxiliar.SizeMbr) + 200)
				Metodos.LlenaDisco(0, 200)
				Metodos.LLenarParticionesAdd(MBRAuxiliar, 1)
				Metodos.LlenaDisco(int(MBRAuxiliar.Particion1MBR.InicioPart), 1)
				Metodos.GeneraEspacios()

				for Contador := 0; Contador <= 200 - 1; Contador++ {

					if Metodos.EspaciosDisponibles[Contador].Disponible {

						if Metodos.EspaciosDisponibles[Contador].P1 == int(MBRAuxiliar.Particion1MBR.InicioPart) + 1 {

							if NuevoSize <= Metodos.EspaciosDisponibles[Contador].Tamano + 1 {

								MBRModificar.SizePart = int64(NuevoSize)
								Bandera = true
								break

							}

						}

					}

				}

				MBRAuxiliar.Particion1MBR = MBRModificar

			} else if CantidadAdd < 0 {

				if NuevoSize > 0 {

					MBRAuxiliar.Particion1MBR.SizePart = int64(NuevoSize)
					Bandera = true

				} else {

					Bandera = false
					MenosEspacio = true

				}

			}


		} else if NumeroParticion == 2 {

			NuevoSize = int(MBRAuxiliar.Particion2MBR.SizePart) + CantidadAdd
			MBRModificar = MBRAuxiliar.Particion2MBR

			if CantidadAdd > 0 {

				//Crear Disco
				Metodos.LimpiaDisco()
				Metodos.CreaDisco(int(MBRAuxiliar.SizeMbr) + 200)
				Metodos.LlenaDisco(0, 200)
				Metodos.LLenarParticionesAdd(MBRAuxiliar, 2)
				Metodos.LlenaDisco(int(MBRAuxiliar.Particion2MBR.InicioPart), 1)
				Metodos.GeneraEspacios()

				for Contador := 0; Contador <= 200 - 1; Contador++ {

					if Metodos.EspaciosDisponibles[Contador].Disponible {

						if Metodos.EspaciosDisponibles[Contador].P1 == int(MBRAuxiliar.Particion2MBR.InicioPart) + 1 {

							if NuevoSize <= Metodos.EspaciosDisponibles[Contador].Tamano + 1 {

								MBRModificar.SizePart = int64(NuevoSize)
								Bandera = true
								break

							}

						}

					}

				}

				MBRAuxiliar.Particion2MBR = MBRModificar

			} else if CantidadAdd < 0 {

				if NuevoSize > 0 {

					MBRAuxiliar.Particion2MBR.SizePart = int64(NuevoSize)
					Bandera = true

				} else {

					Bandera = false
					MenosEspacio = true

				}

			}

		} else if NumeroParticion == 3 {

			NuevoSize = int(MBRAuxiliar.Particion3MBR.SizePart) + CantidadAdd
			MBRModificar = MBRAuxiliar.Particion3MBR

			if CantidadAdd > 0 {

				MBRModificar = MBRAuxiliar.Particion3MBR
				NuevoSize = int(MBRAuxiliar.Particion3MBR.SizePart) + CantidadAdd

				//Crear Disco
				Metodos.LimpiaDisco()
				Metodos.CreaDisco(int(MBRAuxiliar.SizeMbr) + 200)
				Metodos.LlenaDisco(0, 200)
				Metodos.LLenarParticionesAdd(MBRAuxiliar, 3)
				Metodos.LlenaDisco(int(MBRAuxiliar.Particion3MBR.InicioPart), 1)
				Metodos.GeneraEspacios()

				for Contador := 0; Contador <= 200 - 1; Contador++ {

					if Metodos.EspaciosDisponibles[Contador].Disponible {

						if Metodos.EspaciosDisponibles[Contador].P1 == int(MBRAuxiliar.Particion3MBR.InicioPart) + 1 {

							if NuevoSize <= Metodos.EspaciosDisponibles[Contador].Tamano + 1 {

								MBRModificar.SizePart = int64(NuevoSize)
								Bandera = true
								break

							}

						}

					}

				}

				MBRAuxiliar.Particion3MBR = MBRModificar

			} else if CantidadAdd < 0 {

				if NuevoSize > 0 {

					MBRAuxiliar.Particion3MBR.SizePart = int64(NuevoSize)
					Bandera = true

				} else {

					Bandera = false
					MenosEspacio = true

				}

			}

		} else if NumeroParticion == 4 {

			NuevoSize = int(MBRAuxiliar.Particion4MBR.SizePart) + CantidadAdd
			MBRModificar = MBRAuxiliar.Particion4MBR

			if CantidadAdd > 0 {

				MBRModificar = MBRAuxiliar.Particion4MBR
				NuevoSize = int(MBRAuxiliar.Particion4MBR.SizePart) + CantidadAdd

				//Crear Disco
				Metodos.LimpiaDisco()
				Metodos.CreaDisco(int(MBRAuxiliar.SizeMbr) + 200)
				Metodos.LlenaDisco(0, 200)
				Metodos.LLenarParticionesAdd(MBRAuxiliar, 4)
				Metodos.LlenaDisco(int(MBRAuxiliar.Particion4MBR.InicioPart), 1)
				Metodos.GeneraEspacios()

				for Contador := 0; Contador <= 200 - 1; Contador++ {

					if Metodos.EspaciosDisponibles[Contador].Disponible {

						if Metodos.EspaciosDisponibles[Contador].P1 == int(MBRAuxiliar.Particion4MBR.InicioPart) + 1 {

							if NuevoSize <= Metodos.EspaciosDisponibles[Contador].Tamano + 1 {

								MBRModificar.SizePart = int64(NuevoSize)
								Bandera = true
								break

							}

						}

					}

				}

				MBRAuxiliar.Particion4MBR = MBRModificar

			} else if CantidadAdd < 0 {

				if NuevoSize > 0 {

					MBRAuxiliar.Particion4MBR.SizePart = int64(NuevoSize)
					Bandera = true

				} else {

					Bandera = false
					MenosEspacio = true

				}

			}

		}


		if Bandera {

			Metodos.EscribirArchivoBinarioArreglo(MBRAuxiliar)
			color.Success.Println("Particion Modificada Con Exito")
			fmt.Println("")

		} else {

			if MenosEspacio {

				color.HEX("#de4843", false).Println("No Se Puede Quitar El Espacio Indicado Ya Que Es Mayor Al Tamaño De La Particion")
				fmt.Println("")

			} else {

				color.HEX("#de4843", false).Println("No Hay Espacio Disponible Para Extender La Particion")
				fmt.Println("")

			}

		}

	}

	func ComandoFdiskAddParticionLogica(InicioExtendida int64, SizeExtendida int64) {

		//Variables
		var CantidadAdd int
		var UnitAdd int
		var NuevoSize int
		var SizeTotal int
		var InicioPart int64
		var InicioReal int64
		var SizePart int64
		var NombreArray string
		var Bandera bool
		var MenosEspacio bool
		var EBRModificar Variables.EBREstructura
		var ArregloEBR []Variables.EBREstructura

		//Asignaciones
		CantidadAdd, _ = strconv.Atoi(Variables.MapComandos["add"])
		UnitAdd, _ = strconv.Atoi(Variables.MapComandos["unit"])
		CantidadAdd = CantidadAdd * UnitAdd
		EBRModificar = Variables.EBREstructura{}
		Bandera = false
		MenosEspacio = false
		ArregloEBR = make([]Variables.EBREstructura, 0)
		ArregloEBR = ObtenerEBR(InicioExtendida)
		InicioPart = 0
		SizePart = 0
		NuevoSize = 0
		NombreArray = ""

		// Rellenar Particiones Existentes
		for Contador := 0; Contador < len(ArregloEBR); Contador++ {

			NombreArray = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

			if strings.EqualFold(Metodos.Trim(Variables.MapComandos["name"]), NombreArray) {

				EBRModificar = ArregloEBR[Contador]
				NuevoSize = int(ArregloEBR[Contador].SizeEBR) + CantidadAdd


			}
		}


		if CantidadAdd > 0 {

			//Crear Particion Extendida
			Metodos.LimpiaDisco()
			Metodos.CreaDisco(int(SizeExtendida))

			// Rellenar Particiones Existentes
			for Contador := 0; Contador < len(ArregloEBR); Contador++ {

				InicioPart = ArregloEBR[Contador].InicioEBR
				InicioReal = InicioPart - InicioExtendida
				SizePart = ArregloEBR[Contador].SizeEBR
				SizeTotal = 0
				NombreArray = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

				if !strings.EqualFold(Variables.MapComandos["name"], NombreArray) {

					if SizePart != 0 {

						SizeTotal = int(SizePart + int64(unsafe.Sizeof(Variables.EBREstructura{})))

					}
					Metodos.LlenaDisco(int(InicioReal), SizeTotal)

				} else {

					EBRModificar = ArregloEBR[Contador]
					NuevoSize = int(EBRModificar.SizeEBR) + CantidadAdd

				}

			}

			Metodos.LlenaDisco(int(EBRModificar.InicioEBR - InicioExtendida), 1)
			Metodos.GeneraEspacios()
			Metodos.MostrarEspacios()

			for Contador := 0; Contador <= 200 - 1; Contador++ {

				if Metodos.EspaciosDisponibles[Contador].Disponible {

					if Metodos.EspaciosDisponibles[Contador].P1 == int(EBRModificar.InicioEBR) - int(InicioExtendida) + 1 {

						if NuevoSize <= Metodos.EspaciosDisponibles[Contador].Tamano + 1 {

							EBRModificar.SizeEBR = int64(NuevoSize)
							Bandera = true
							break

						}

					}

				}

			}

		} else if CantidadAdd < 0 {

			if NuevoSize > 0 {

				EBRModificar.SizeEBR = int64(NuevoSize)
				Bandera = true

			} else {

				Bandera = false
				MenosEspacio = true

			}

		}

		// Rellenar Particiones Existentes
		for Contador := 0; Contador < len(ArregloEBR); Contador++ {

			NombreArray = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

			if strings.EqualFold(Variables.MapComandos["name"], NombreArray) {

				ArregloEBR[Contador] = EBRModificar

			}
		}

		// Escribir EBR
		for Contador := 0; Contador < len(ArregloEBR); Contador++ {

			//Escribir EBR
			Metodos.EscribirArchivoBinarioEBRAdd(ArregloEBR[Contador], ArregloEBR[Contador].InicioEBR)

		}

		ArregloEBR = make([]Variables.EBREstructura, 0)
		ArregloEBR = ObtenerEBR(InicioExtendida)

		// Crear Disco Virtual
		Metodos.LimpiaDisco()
		Metodos.CreaDisco(int(SizeExtendida))

		// Rellenar Particiones Existentes
		for Contador := 0; Contador < len(ArregloEBR); Contador++ {

			InicioPart = ArregloEBR[Contador].InicioEBR
			InicioReal = InicioPart - InicioExtendida
			SizePart = ArregloEBR[Contador].SizeEBR
			SizeTotal = 0
			if SizePart != 0 {

				SizeTotal = int(SizePart + int64(unsafe.Sizeof(Variables.EBREstructura{})))

			}

			Metodos.LlenaDisco(int(InicioReal), SizeTotal)
		}

		Metodos.GeneraEspacios()
		fmt.Println("")
		fmt.Println("Espacios Vacios:")
		Metodos.MostrarEspacios()
		fmt.Println("")

		if Bandera {

			color.Success.Println("Particion Modificada Con Exito")
			fmt.Println("")

		} else {

			if MenosEspacio {

				color.HEX("#de4843", false).Println("No Se Puede Quitar El Espacio Indicado Ya Que Es Mayor Al Tamaño De La Particion")
				fmt.Println("")

			} else {

				color.HEX("#de4843", false).Println("No Hay Espacio Disponible Para Extender La Particion")
				fmt.Println("")

			}

		}

	}