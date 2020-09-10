
//---------------------------------------------Paquetes E Imports-------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"bytes"
		"encoding/binary"
		"fmt"
		"github.com/asaskevich/govalidator"
		"github.com/gookit/color"
		"os"
		"path/filepath"
		"strconv"
		"strings"
		"time"
		"unsafe"
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
		var Archivo string
		var FechaActual time.Time
		var NumeroEstructuras int64
		var SizeParticion int64
		var SizeArbolDirectorio int64
		var SizeDetalleDirectorio int64
		var SizeTablaInodos int64
		var SizeBQ int64
		var SizeBI int64
		var SizeSuperBloque int64
		var InicioParticion int64
		var Bandera bool
		var CeroBinario int8
		var CeroByte *int8
		var ArchivoDisco *os.File
		var AvisoError error
		var CadenaBinariaSuperBoot bytes.Buffer
		var CadenaBinariaCero bytes.Buffer
		var NombreArchivo []string
		var SuperBoot Variables.SuperBloqueEstructura
		var ParticionMontada Variables.MountEstructura

		//Asignacion
		Archivo = ""
		NumeroEstructuras = 0
		SizeParticion = 0
		SizeArbolDirectorio = int64(unsafe.Sizeof(Variables.AVDEstructura{}))
		SizeDetalleDirectorio = int64(unsafe.Sizeof(Variables.DDEstructura{}))
		SizeTablaInodos = int64(unsafe.Sizeof(Variables.TablaInodoEstructura{}))
		SizeBQ = int64(unsafe.Sizeof(Variables.BloquesEstructura{}))
		SizeBI = int64(unsafe.Sizeof(Variables.BitacoraEstructura{}))
		SizeSuperBloque = int64(unsafe.Sizeof(Variables.SuperBloqueEstructura{}))
		InicioParticion = 0
		Bandera = false
		NombreArchivo = make([]string, 0)
		ParticionMontada = Variables.MountEstructura{}
		SuperBoot = Variables.SuperBloqueEstructura{}

		//Verificar Particon Montada
		for Contador := 0; Contador < len(Variables.ArregloParticionesMontadas); Contador++ {

			if strings.EqualFold(Variables.MapComandos["id"], Variables.ArregloParticionesMontadas[Contador].IdentificadorMount) {

				ParticionMontada = Variables.ArregloParticionesMontadas[Contador]
				Bandera = true

			}

		}

		if Bandera {

			//Obtener Nombre Disco
			_, Archivo = filepath.Split(Metodos.Trim(ParticionMontada.RutaDiscoMount))
			NombreArchivo = strings.Split(Archivo, ".")

			if ParticionMontada.ParticionMount.SizePart != 0 {

				SizeParticion = ParticionMontada.ParticionMount.SizePart
				InicioParticion = ParticionMontada.ParticionMount.InicioPart


			} else if ParticionMontada.EBRMount.SizeEBR != 0 {

				SizeParticion = ParticionMontada.EBRMount.SizeEBR
				InicioParticion = ParticionMontada.EBRMount.InicioEBR + int64(unsafe.Sizeof(Variables.EBREstructura{}))

			}

			//Formula Para Calcular Numero De Estructuras
			NumeroEstructuras = (SizeParticion - (2 * SizeSuperBloque)) / (27 + SizeArbolDirectorio + SizeDetalleDirectorio + (5 * SizeTablaInodos + (20 * SizeBQ) + SizeBI))

			//Feaha Actual
			FechaActual = time.Now()

			// Apuntadores
			ApuntadorBitmapAVD := InicioParticion + SizeSuperBloque
			ApuntadorAVD := ApuntadorBitmapAVD + NumeroEstructuras
			ApuntadorBitmapDD := ApuntadorAVD + (SizeArbolDirectorio * NumeroEstructuras)
			ApuntadorDD := ApuntadorBitmapDD + NumeroEstructuras
			ApuntadorBitmapTI := ApuntadorDD + (SizeDetalleDirectorio * NumeroEstructuras)
			ApuntadorTI := ApuntadorBitmapTI + (5 * NumeroEstructuras)
			ApuntadorBitmapBQ := ApuntadorTI + (SizeTablaInodos * (5 * NumeroEstructuras))
			ApuntadorBQ := ApuntadorBitmapBQ + (4 * (5 * NumeroEstructuras))
			ApuntadorBI := ApuntadorBQ + (SizeBQ * (4 * (5 * NumeroEstructuras)))

			println(NombreArchivo[0])

			// Rellenar SuperBoot
			// Nombre Disco
			copy(SuperBoot.NombreDiscoSuperBloque[:], NombreArchivo[0])
			// Cantidad De Estructura
			SuperBoot.ArbolCountSuperBloque = NumeroEstructuras
			SuperBoot.DetalleDirectorioCountSuperBloque = NumeroEstructuras
			SuperBoot.InodosCountSuperBloque = 5 * NumeroEstructuras
			SuperBoot.BloquesCountSuperBloque = 4 * (5 * NumeroEstructuras)
			// Estructuras Libres
			SuperBoot.ArbolFreeSuperBloque = NumeroEstructuras
			SuperBoot.DetalleFreeSuperBloque = NumeroEstructuras
			SuperBoot.InodosFreeSuperBloque = 5 * NumeroEstructuras
			SuperBoot.BloquesFreeSuperBloque = 4 * (5 * NumeroEstructuras)
			// Fechas
			copy(SuperBoot.FechaCreacionSuperBloque[:], FechaActual.String())
			copy(SuperBoot.FechaUltimoMontajeSuperBloque[:], FechaActual.String())
			//Numero De Montajes
			SuperBoot.MontajesSuperBloque = 1
			//Apuntadores
			SuperBoot.PBitmapArbolSuperBloque = ApuntadorBitmapAVD
			SuperBoot.PArbolSuperBloque = ApuntadorAVD
			SuperBoot.PBitmapDetalleSuperBloque = ApuntadorBitmapDD
			SuperBoot.PDetalleSuperBloque = ApuntadorDD
			SuperBoot.PBitmapTablaSuperBloque = ApuntadorBitmapTI
			SuperBoot.PTablaSuperBloque = ApuntadorTI
			SuperBoot.PBitmapBloquesSuperBloque = ApuntadorBitmapBQ
			SuperBoot.PBloquesSuperBloque = ApuntadorBQ
			SuperBoot.PLogSuperBloque = ApuntadorBI
			// Tamaño Estructura
			SuperBoot.ArbolSizeSuperBloque = SizeArbolDirectorio
			SuperBoot.DetalleSizeSuperBloque = SizeDetalleDirectorio
			SuperBoot.InodoSizeSuperBloque = SizeTablaInodos
			SuperBoot.BloquesSizeSuperBloque = SizeBQ
			// Free Bit
			SuperBoot.ArbolFreeBitSuperBloque = ApuntadorBitmapAVD
			SuperBoot.DetalleFreeBitSuperBloque = ApuntadorBitmapDD
			SuperBoot.TablaFreeBitSuperBloque = ApuntadorBitmapTI
			SuperBoot.BloquesFreeBitSuperBloque = ApuntadorBitmapBQ
			// Magic Num
			SuperBoot.MagicNumSuperBloque = 201801628

			// Abrir Archivo Binario
			ArchivoDisco, AvisoError = os.OpenFile(ParticionMontada.RutaDiscoMount, os.O_WRONLY, os.ModePerm)

			if AvisoError != nil {

				color.HEX("#de4843", false).Println("Error Al Leer El Disco")
				color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
				fmt.Println("")

			} else {

				// Mover Puntero Inicio Archivo
				_, _ = ArchivoDisco.Seek(0, 0)

				// Mover Puntero Inicio Particion
				_, _ = ArchivoDisco.Seek(InicioParticion, 0)

				// Escribir SuperBoot

				//Asignación
				SuperBootDireccion := &SuperBoot

				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaSuperBoot, binary.BigEndian, SuperBootDireccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaSuperBoot.Bytes())

				// Escribir Bitmap AVD

				//Asignación
				CeroBinario = 0
				CeroByte = &CeroBinario

				//Mover A Posicion Final
				_, _ = ArchivoDisco.Seek(ApuntadorBitmapAVD, 0)

				for Contador := 0; Contador < int(NumeroEstructuras); Contador++ {

					_ = binary.Write(&CadenaBinariaCero, binary.BigEndian, CeroByte)
					Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaCero.Bytes())

				}

				_ = ArchivoDisco.Close()

			}

		} else {

			color.HEX("#de4843", false).Println("No Existe El Id Indicado")
			fmt.Println("")

		}

		
	}

	func ComandoMkfsAdd() {

		print("Comando Mkfs")

	}