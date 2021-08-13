
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
		var ArchivoDisco *os.File
		var AvisoError error
		var CadenaBinariaSuperBoot bytes.Buffer
		var CadenaBinariaAVD bytes.Buffer
		var CadenaBinariaAVDB bytes.Buffer
		var CadenaBinariaDD bytes.Buffer
		var CadenaBinariaDDB bytes.Buffer
		var CadenaBinariaTI bytes.Buffer
		var CadenaBinariaTIB bytes.Buffer
		var CadenaBinariaBQ1 bytes.Buffer
		var CadenaBinariaBQ2 bytes.Buffer
		var CadenaBinariaBQB bytes.Buffer
		var CadenaBinariaBI1 bytes.Buffer
		var CadenaBinariaBI2 bytes.Buffer
		var NombreArchivo []string
		var SuperBoot Variables.SuperBloqueEstructura
		var ParticionMontada Variables.MountEstructura
		var AVDAuxiliar Variables.AVDEstructura
		var DDAuxiliar Variables.DDEstructura
		var TIAuxiliar Variables.TablaInodoEstructura
		var BQ1Auxiliar Variables.BloquesEstructura
		var BQ2Auxiliar Variables.BloquesEstructura
		var BI1Auxiliar Variables.BitacoraEstructura
		var BI2Auxiliar Variables.BitacoraEstructura
		var ArregloBytes []byte

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
		AVDAuxiliar = Variables.AVDEstructura{}
		DDAuxiliar = Variables.DDEstructura{}
		TIAuxiliar = Variables.TablaInodoEstructura{}
		BQ1Auxiliar = Variables.BloquesEstructura{}
		BQ2Auxiliar = Variables.BloquesEstructura{}
		BI1Auxiliar = Variables.BitacoraEstructura{}
		BI2Auxiliar = Variables.BitacoraEstructura{}
		ArregloBytes = make([]byte, 0)

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

				//Mover A Posicion Final
				_, _ = ArchivoDisco.Seek(ApuntadorBitmapAVD, 0)

				for Contador := 0; Contador < int(NumeroEstructuras); Contador++ {

					if Contador == 0 {

						ArregloBytes = append(ArregloBytes, 1)

					} else {

						ArregloBytes = append(ArregloBytes, 0)

					}

				}

				_ = binary.Write(&CadenaBinariaAVDB, binary.BigEndian, ArregloBytes)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaAVDB.Bytes())

				// Escribir AVD Root

				//Feaha Actual
				FechaActual = time.Now()

				// Rellenar AVD
				copy(AVDAuxiliar.FechaCreacionAVD[:], FechaActual.String())
				copy(AVDAuxiliar.NombreDirectorioAVD[:], "/")
				AVDAuxiliar.PDetalleDirectorioAVD = SuperBoot.PDetalleSuperBloque
				copy(AVDAuxiliar.PropietarioAVD[:], "root")
				copy(AVDAuxiliar.GrupoAVD[:], "root")
				AVDAuxiliar.PermisosAVD = 777

				// Mover Puntero
				_, _ = ArchivoDisco.Seek(SuperBoot.PArbolSuperBloque + (0 * SizeArbolDirectorio), 0)

				//Asignación
				AVDDireccion := &AVDAuxiliar
				

				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaAVD, binary.BigEndian, AVDDireccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaAVD.Bytes())

				// Escribir Bitmap DD

				//Mover A Posicion Final
				_, _ = ArchivoDisco.Seek(ApuntadorBitmapDD, 0)

				ArregloBytes = make([]byte, 0)

				for Contador := 0; Contador < int(NumeroEstructuras); Contador++ {

					if Contador == 0 {

						ArregloBytes = append(ArregloBytes, 1)

					} else {

						ArregloBytes = append(ArregloBytes, 0)

					}

				}

				_ = binary.Write(&CadenaBinariaDDB, binary.BigEndian, ArregloBytes)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaDDB.Bytes())

				// Escribir DD

				//Feaha Actual
				FechaActual = time.Now()

				// Rellenar DD
				copy(DDAuxiliar.ArrayArchivosDD[0].NombreArchivoDDInformacion[:], "users.txt")
				copy(DDAuxiliar.ArrayArchivosDD[0].FechaCreacionArchivoDDInformacion[:], FechaActual.String())
				copy(DDAuxiliar.ArrayArchivosDD[0].FechaModificacionArchivoDDInformacion[:], FechaActual.String())
				DDAuxiliar.ArrayArchivosDD[0].PInodoArchivoDDInformacion = SuperBoot.PTablaSuperBloque
				DDAuxiliar.PDetalleDirectorioDD = 0


				// Mover Puntero
				_, _ = ArchivoDisco.Seek(SuperBoot.PDetalleSuperBloque + (0 * SizeDetalleDirectorio), 0)

				//Asignación
				DDDireccion := &DDAuxiliar


				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaDD, binary.BigEndian, DDDireccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaDD.Bytes())

				// Escribir Bitmap Inodos

				//Mover A Posicion Final
				_, _ = ArchivoDisco.Seek(ApuntadorBitmapTI, 0)

				ArregloBytes = make([]byte, 0)

				for Contador := 0; Contador < int(5 * NumeroEstructuras); Contador++ {

					if Contador == 0 {

						ArregloBytes = append(ArregloBytes, 1)

					} else {

						ArregloBytes = append(ArregloBytes, 0)

					}

				}

				_ = binary.Write(&CadenaBinariaTIB, binary.BigEndian, ArregloBytes)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaTIB.Bytes())

				// Escribir TI
				// Rellenar TI
				TIAuxiliar.NumeroInodoTI = 1
				TIAuxiliar.SizeArchivoTI = 50
				TIAuxiliar.NumeroBloquesTI = 2
				TIAuxiliar.ArrayBloquesTI[0] = ApuntadorBQ + (0 * SizeBQ)
				TIAuxiliar.ArrayBloquesTI[1] = ApuntadorBQ + (1 * SizeBQ)
				TIAuxiliar.PTabalInodosTI = 0
				copy(TIAuxiliar.PropietarioTI[:], "root")
				copy(TIAuxiliar.GrupoTI[:], "root")
				TIAuxiliar.PermisosTI = 770

				// Mover Puntero
				_, _ = ArchivoDisco.Seek(SuperBoot.PTablaSuperBloque + (0 * SizeTablaInodos), 0)

				//Asignación
				TIDireccion := &TIAuxiliar

				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaTI, binary.BigEndian, TIDireccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaTI.Bytes())

				// Escribir Bitmap Bloques

				//Mover A Posicion Final
				_, _ = ArchivoDisco.Seek(ApuntadorBitmapBQ, 0)

				ArregloBytes = make([]byte, 0)

				for Contador := 0; Contador < int(4 * (5 * NumeroEstructuras)); Contador++ {

					if Contador == 0 || Contador == 1 {

						ArregloBytes = append(ArregloBytes, 1)

					} else {

						ArregloBytes = append(ArregloBytes, 0)

					}

				}

				_ = binary.Write(&CadenaBinariaBQB, binary.BigEndian, ArregloBytes)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaBQB.Bytes())

				// Escribir BQ 1
				// Rellenar BQ 1
				copy(BQ1Auxiliar.InformacionBQ[:], "1,G,root\\n1,U,root,root,2")


				// Mover Puntero
				_, _ = ArchivoDisco.Seek(SuperBoot.PBloquesSuperBloque + (0 * SizeBQ), 0)

				//Asignación
				BQ1Direccion := &BQ1Auxiliar

				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaBQ1, binary.BigEndian, BQ1Direccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaBQ1.Bytes())


				// Escribir BQ 2
				// Rellenar BQ 2
				copy(BQ2Auxiliar.InformacionBQ[:], "01801628\\n")


				// Mover Puntero
				_, _ = ArchivoDisco.Seek(SuperBoot.PBloquesSuperBloque + (1 * SizeBQ), 0)

				//Asignación
				BQ2Direccion := &BQ2Auxiliar

				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaBQ2, binary.BigEndian, BQ2Direccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaBQ2.Bytes())

				// Escribir BI 1

				//Feaha Actual
				FechaActual = time.Now()

				// Rellenar BI 1
				copy(BI1Auxiliar.TipoOperacionBT[:], "(Mkdir)")
				BI1Auxiliar.TipoArchivoDirectorioBT = 1
				copy(BI1Auxiliar.NombreArchivoDirectorioBT[:], "/")
				copy(BI1Auxiliar.ContenidoBT[:], "-")
				copy(BI1Auxiliar.FechaTransaccionBT[:], FechaActual.String())


				// Mover Puntero
				_, _ = ArchivoDisco.Seek(SuperBoot.PLogSuperBloque + (0 * SizeBI), 0)

				//Asignación
				BI1Direccion := &BI1Auxiliar

				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaBI1, binary.BigEndian, BI1Direccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaBI1.Bytes())


				// Escribir BI 2

				//Feaha Actual
				FechaActual = time.Now()

				// Rellenar BI 2
				copy(BI2Auxiliar.TipoOperacionBT[:], "(Mkfile)")
				BI2Auxiliar.TipoArchivoDirectorioBT = 0
				copy(BI2Auxiliar.NombreArchivoDirectorioBT[:], "users.txt")
				copy(BI2Auxiliar.ContenidoBT[:], "-")
				copy(BI2Auxiliar.FechaTransaccionBT[:], FechaActual.String())

				// Mover Puntero
				_, _ = ArchivoDisco.Seek(SuperBoot.PLogSuperBloque + (1 * SizeBI), 0)

				//Asignación
				BI2Direccion := &BI2Auxiliar

				//Escribir Archivo
				_ = binary.Write(&CadenaBinariaBI2, binary.BigEndian, BI2Direccion)
				Metodos.EscribirArchivoBinario(ArchivoDisco, CadenaBinariaBI2.Bytes())

				_ = ArchivoDisco.Close()
				
				color.Success.Println("Particion Formateada Con Exito")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("No Existe El Id Indicado")
			fmt.Println("")

		}

		
	}

	func ComandoMkfsAdd() {

		print("Comando Mkfs")

	}