
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"bytes"
		"encoding/binary"
		"fmt"
		"github.com/gookit/color"
		"math/rand"
		"os"
		"strconv"
		"strings"
		"time"
		"unsafe"
	)

//--------------------------------------------------------Métodos-------------------------------------------------------

	func VerificarComandoMkdisk() {

		//Variables
		var Size bool
		var Path bool
		var Name bool
		var Unit bool
		var ParametroExtra bool
		var ArregloParametros []string
		var ContadorSize int
        var ContadorPath int
		var ContadorName int
		var ContadorUnit int
		var AvisoError error

		//Asignación
		Size = false
		Path = false
		Name = false
		Unit = true
		ParametroExtra = false
		ContadorSize = 0
		ContadorPath = 0
		ContadorName = 0
		ContadorUnit = 0
		Variables.MapComandos = make(map[string]string)
		Variables.MapComandos["unit"] = "1048576"
		Variables.CreeDirectorio = false

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

							Path = Metodos.VerificarYCrearRutas(ArregloParametros[1])

							if Path {

                                Variables.MapComandos["path"] = ArregloParametros[1]

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

							Name = Metodos.ContineCaractereProhibidos(ArregloParametros[1])

							if Name {

								Name = Metodos.VerificarExtension(ArregloParametros[1])

								if Name {

									Variables.MapComandos["name"] = ArregloParametros[1]

								} else {

									color.HEX("#de4843", false).Println("En El Parametro Name La Extension Del Archivo No Es la Indicada Debe De Ser .dsk")
									fmt.Println("")
								}

							} else {

								color.HEX("#de4843", false).Println("En El Parametor Name El Nombre Del Disco Contiene Carcteres Prohibidos")
								fmt.Println("")

							}

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

							} else {

								color.HEX("#de4843", false).Println("En El Parametro Unit Debe De Ingresar La Letra m (Megabytes) O La Letra k (Kylobytes)")
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


		if Path && Size && Name && Unit && !ParametroExtra && ContadorPath == 1 && ContadorSize == 1 && ContadorName == 1 && (ContadorUnit == 1 || ContadorUnit == 0) {

			ObtenerRutaArchivo()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "2). -size->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "3). -name->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "4). -unit->    (Opcional)")
				fmt.Println("")

			}

			if !Path {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -path-> o")
				color.HEX("#de4843", false).Println("Error Al Crear El Directorio")
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

			if ContadorSize > 1 || ContadorPath > 1 || ContadorName > 1 || ContadorUnit > 1{

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

			if Variables.CreeDirectorio {

				AvisoError = os.Remove(Metodos.Trim(Variables.MapComandos["path"]))
				_ = AvisoError

			}

		}

	}

	func ObtenerRutaArchivo() {

		//Variables
		var ExisteBarra bool
		var RutaArchivo string

		//Verificar Sistema Operativo

		if Variables.SistemaOperativo == "windows" {

			ExisteBarra = strings.HasSuffix(Metodos.Trim(Variables.MapComandos["path"]), "\\")

			if ExisteBarra {

				RutaArchivo = Metodos.Trim(Variables.MapComandos["path"]) + Metodos.Trim(Variables.MapComandos["name"])

			} else {

				RutaArchivo = Metodos.Trim(Variables.MapComandos["path"]) + "\\" + Metodos.Trim(Variables.MapComandos["name"])

			}

			VerificarSiExisteDisco(RutaArchivo)

		} else if Variables.SistemaOperativo == "linux" {

			ExisteBarra = strings.HasSuffix(Metodos.Trim(Variables.MapComandos["path"]), "/")

			if ExisteBarra {

				RutaArchivo = Metodos.Trim(Variables.MapComandos["path"]) + Metodos.Trim(Variables.MapComandos["name"])

			} else {

				RutaArchivo = Metodos.Trim(Variables.MapComandos["path"]) + "/" + Metodos.Trim(Variables.MapComandos["name"])

			}

			VerificarSiExisteDisco(RutaArchivo)

		} else {

			color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
			fmt.Println("")

		}

	}

	func VerificarSiExisteDisco(RutaArchivo string) {

		//Variables
		var Bandera bool

		//Asignación
		Bandera = false

		//Verificar Si No Existe El Disco
		Bandera = Metodos.ExisteArchivo(Metodos.Trim(RutaArchivo))

		if !Bandera {

            ComandoMkdisk(RutaArchivo)

		} else {

			color.HEX("#de4843", false).Println("Ya Existe Un Disco Con El Mismo Nombre En La Ruta Indicada")
			fmt.Println("")

		}

	}

	func ComandoMkdisk(RutaArchivo string) {

	    //Variables
		var Archivo *os.File
		var AvisoError error
		var CeroBinario int8
		var CeroByte *int8
		var CadenaBinaria bytes.Buffer
		var CadenaBinariaFinal bytes.Buffer
		var CadenaBinariaMBR bytes.Buffer
		var Posicion int64
		var Unit int
		var Size int
		var Fecha time.Time
		var MBRAuxiliar = Variables.MBREstructura{}

		//Creando Archivo
		Archivo, AvisoError = os.Create(RutaArchivo)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error al Generar Al Archivo")
			fmt.Println("")

		} else {

			//Asignación
			CeroBinario = 0
			CeroByte = &CeroBinario

			//Escribir Archivo
			_ = binary.Write(&CadenaBinaria, binary.BigEndian, CeroByte)
			Metodos.EscribirArchivoBinario(Archivo, CadenaBinaria.Bytes())

			//Posicion Final
			Unit, _ = strconv.Atoi(Metodos.Trim(Variables.MapComandos["unit"]))
			Size, _ = strconv.Atoi(Metodos.Trim(Variables.MapComandos["size"]))
			Posicion = int64((Unit * Size) - 1)

			//Mover A Posicion Final
			_, _ = Archivo.Seek(Posicion, 0)

			//Escribir Al Final Del Archivo
			_ = binary.Write(&CadenaBinariaFinal, binary.BigEndian, CeroByte)
			Metodos.EscribirArchivoBinario(Archivo, CadenaBinariaFinal.Bytes())

			//Escribir MBR

			//Posicionar Al Inicio Para Escribir El MBR
			_, _ = Archivo.Seek(0, 0)

			//Rellenar Esctructura
			MBRAuxiliar = Variables.MBREstructura{}
			//Tamaño Total Del Disco TamañoTotal - Tamaño MBR
			MBRAuxiliar.SizeMbr = Posicion + 1 - int64(unsafe.Sizeof(Variables.MBREstructura{}))

			//Fecha Actual
			Fecha = time.Now()
			copy(MBRAuxiliar.FCreacionMBR[:], Fecha.String())
			//Identificador Unico Disco
			MBRAuxiliar.IDMBR = int64(rand.Intn(100000000))
			//Particiones
			MBRAuxiliar.Particion1MBR = Variables.ParticionEstructura{}
			MBRAuxiliar.Particion2MBR = Variables.ParticionEstructura{}
			MBRAuxiliar.Particion3MBR = Variables.ParticionEstructura{}
			MBRAuxiliar.Particion4MBR = Variables.ParticionEstructura{}

			//Asignación
			MBRDireccion := &MBRAuxiliar

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaMBR, binary.BigEndian, MBRDireccion)
			Metodos.EscribirArchivoBinario(Archivo, CadenaBinariaMBR.Bytes())

			color.Success.Println("Disco Creado Con Exito!")
			fmt.Println("")

			_ = Archivo.Close()

		}

	}