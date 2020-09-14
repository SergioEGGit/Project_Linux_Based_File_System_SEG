
//---------------------------------------------Paquetes E Imports-------------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Reportes"
		"../Variables"
		"fmt"
		"github.com/gookit/color"
		"strings"
	)

//-----------------------------------------------------Métodos----------------------------------------------------------

	func VerificarComandoRep() {

		//Variables
		var name bool
		var Path bool
		var Id bool
		var Ruta bool
		var ParametroExtra bool
		var ArregloParametros []string
		var Contadorname int
		var ContadorPath int
		var ContadorId int
		var ContadorRuta int

		//Asignación
		name = false
		Path = false
		Id = false
		Ruta = true
		ParametroExtra = false
		Contadorname = 0
		ContadorPath = 0
		ContadorId = 0
		ContadorRuta = 0
		Variables.MapComandos = make(map[string]string)
		Variables.MapComandos["ruta"] = "none"

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			for Contador := 1; Contador <= len(Variables.ArregloComandos) - 1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				switch ArregloParametros[0] {

				case "name":

				  if Contadorname == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])
							Variables.MapComandos["name"] = Metodos.Trim(ArregloParametros[1])
							name = true

							Contadorname++

						} else {

							name = false

						}

					} else {

						Contadorname++

					}

				case "path":

					if ContadorPath == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.QuitarComillas(ArregloParametros[1])
							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Variables.MapComandos["path"] = ArregloParametros[1]
							Path = true

							ContadorPath++

						} else {

							Path = false

						}

					} else {

						ContadorPath++

					}

				case "id":

					if ContadorId == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])

							Variables.MapComandos["id"] = Metodos.Trim(ArregloParametros[1])
							Id = true

							ContadorId++

						} else {

							Id = false

						}

					} else {

						ContadorId++

					}

				case "ruta":

					if ContadorRuta == 0 {

						if len(ArregloParametros) > 1 {

							ArregloParametros[1] = Metodos.QuitarComillas(ArregloParametros[1])
							ArregloParametros[1] = Metodos.Trim(ArregloParametros[1])
							Variables.MapComandos["ruta"] = Metodos.Trim(ArregloParametros[1])

							ContadorRuta++

						} else {

							Ruta = false

						}

					} else {

						ContadorRuta++

					}

				default:

					ParametroExtra = true

				}
			}
		}


		if Path && Id && name && Ruta && !ParametroExtra && Contadorname == 1 && ContadorPath == 1 && (ContadorRuta == 1 || ContadorRuta == 0) && ContadorId == 1 {

			VerificarNombreReporte()

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "2). -name->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "3). -id->    (Obligatorio)")
				color.HEX("#de4843", false).Println( "4). -ruta->    (Opcional)")
				fmt.Println("")

			}

			if !Path {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -path-> o")
				color.HEX("#de4843", false).Println("Existe Error En Sintaxis")
				fmt.Println("")

			}

			if !name {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -name-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")
			}

			if !Id {

				color.HEX("#de4843", false).Println("No Se Encuentra el Parametro -id-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if Contadorname > 1 || ContadorPath > 1 || ContadorRuta > 1 || ContadorId > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}

	func ObtenerEBRComandoRep(InicioListaExtendida int64, Ruta string) []Variables.EBREstructura {

		//Variables
		var Contador int
		var Bandera bool
		var EBRAuxiliar Variables.EBREstructura
		var ArregloEBR []Variables.EBREstructura

		//Asignación
		Contador = 0

		for {

			//Leer EBR
			EBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioEBR(Ruta, InicioListaExtendida)

			//Lista Corrupta
			if !Bandera {

				return ArregloEBR

			}

			//fmt.Println("Size: ", EBRAuxiliar.SizeEBR, "Inicio: ", EBRAuxiliar.InicioEBR, "Siguiente: ", EBRAuxiliar.SiguienteEBR, "name: ", string(EBRAuxiliar.NameEBR[:]))

			ArregloEBR = append(ArregloEBR, EBRAuxiliar)
			InicioListaExtendida = ArregloEBR[Contador].SiguienteEBR
			Contador++

			if EBRAuxiliar.SiguienteEBR == -1 {

				break

			}

		}

		return ArregloEBR

	}

	func VerificarNombreReporte() {

		//Variables
		var NumeroMount int
		var Bandera bool
		var ExisteExtendida bool
		var InicioExtendida int64
		var MBRAuxiliar Variables.MBREstructura
		var ArregloEBR []Variables.EBREstructura
		var ParticionMontada Variables.MountEstructura

		//Asignacion
		Bandera = false
		NumeroMount = 0
		ExisteExtendida = false
		InicioExtendida = 0
		ArregloEBR = make([]Variables.EBREstructura, 0)
		ParticionMontada = Variables.MountEstructura{}

		// Verificar Si Existe Id
		for Contador := 0; Contador < len(Variables.ArregloParticionesMontadas); Contador++ {

			if strings.EqualFold(Variables.ArregloParticionesMontadas[Contador].IdentificadorMount, Variables.MapComandos["id"]) {

				Bandera = true
				NumeroMount = Contador
				ParticionMontada = Variables.ArregloParticionesMontadas[Contador]

			}

		}

		if Bandera {

			if Variables.MapComandos["name"] == "mbr" {

				Bandera = false

				MBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioArraglo(Metodos.Trim(Variables.ArregloParticionesMontadas[NumeroMount].RutaDiscoMount))

				if Bandera {

					// Verificar Si Hay EBR
					if MBRAuxiliar.Particion1MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion1MBR.InicioPart

					} else if MBRAuxiliar.Particion2MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion2MBR.InicioPart

					} else if MBRAuxiliar.Particion3MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion3MBR.InicioPart

					} else if MBRAuxiliar.Particion4MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion4MBR.InicioPart

					}

					if ExisteExtendida {

						ArregloEBR = ObtenerEBRComandoRep(InicioExtendida, Metodos.Trim(Variables.ArregloParticionesMontadas[NumeroMount].RutaDiscoMount))

					}

					Reportes.ReporteMBR(MBRAuxiliar, Metodos.Trim(Variables.MapComandos["path"]), ArregloEBR)

				} else {

					color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando rep")
					color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
					fmt.Println("")

				}

			} else if Variables.MapComandos["name"] == "disk" {

				Bandera = false

				MBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioArraglo(Metodos.Trim(Variables.ArregloParticionesMontadas[NumeroMount].RutaDiscoMount))

				if Bandera {

					// Verificar Si Hay EBR
					if MBRAuxiliar.Particion1MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion1MBR.InicioPart

					} else if MBRAuxiliar.Particion2MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion2MBR.InicioPart

					} else if MBRAuxiliar.Particion3MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion3MBR.InicioPart

					} else if MBRAuxiliar.Particion4MBR.TipoPart == 'e' {

						ExisteExtendida = true
						InicioExtendida = MBRAuxiliar.Particion4MBR.InicioPart

					}

					if ExisteExtendida {

						ArregloEBR = ObtenerEBRComandoRep(InicioExtendida, Metodos.Trim(Variables.ArregloParticionesMontadas[NumeroMount].RutaDiscoMount))

					}

					Reportes.ReporteDisco(MBRAuxiliar, Metodos.Trim(Variables.MapComandos["path"]), ArregloEBR)

				} else {

					color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando rep")
					color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
					fmt.Println("")

				}

			} else if Variables.MapComandos["name"] == "sb" {

				Reportes.ReporteSuperBoot(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

			} else if Variables.MapComandos["name"] == "bm_arbdir" {

				Reportes.ReporteBitmapAVD(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

			} else if Variables.MapComandos["name"] == "bm_detdir" {

				Reportes.ReporteBitmapDD(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

			} else if Variables.MapComandos["name"] == "bm_inode" {

				Reportes.ReporteBitmapTI(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

			} else if Variables.MapComandos["name"] == "bm_block" {

				Reportes.ReporteBitmapBQ(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

			} else if Variables.MapComandos["name"] == "bitacora" {

				Reportes.ReporteBitacora(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

			} else if Variables.MapComandos["name"] == "directorio" {

				Reportes.ReporteDirectorio(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

			} else if Variables.MapComandos["name"] == "tree_file" {

				if Variables.MapComandos["ruta"] == "none" {

					color.HEX("#de4843", false).Println("Debe Indicar Un Archivo Para Generar El Reporte")
					fmt.Println("")

				} else if Variables.MapComandos["ruta"] == "/users.txt" {

					Reportes.ReporteTreeFile(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

				} else if Variables.MapComandos["ruta"] == "" {

					color.HEX("#de4843", false).Println("Debe De Indicar Un Archivo Valido")
					fmt.Println("")

				} else {

					color.HEX("#de4843", false).Println("No Existe El Archivo Indicado")
					fmt.Println("")

				}

			} else if Variables.MapComandos["name"] == "tree_directorio" {

				if Variables.MapComandos["ruta"] == "none" {

					color.HEX("#de4843", false).Println("Debe Indicar Un Directorio/os Para Generar El Reporte")
					fmt.Println("")

				} else if Variables.MapComandos["ruta"] == "/" {

					Reportes.ReporteTreeDirectorio(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

				} else if Variables.MapComandos["ruta"] == "" {

					color.HEX("#de4843", false).Println("Debe De Indicar Un Directorio/os Valido")
					fmt.Println("")

				} else {

					color.HEX("#de4843", false).Println("No Existe La Ruta Indicada")
					fmt.Println("")

				}

			} else if Variables.MapComandos["name"] == "tree_complete" {

				Reportes.ReporteArbolCompletoAVD(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]))

		    } else if Variables.MapComandos["name"] == "ls" {

				if Variables.MapComandos["ruta"] == "none" {

					color.HEX("#de4843", false).Println("Debe Indicar Un Directorio/os o Archivos Para Generar El Reporte")
					fmt.Println("")

				} else if Variables.MapComandos["ruta"] == "/" {

					Reportes.ReporteLS(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]), 0)

				} else if Variables.MapComandos["ruta"] == "/users.txt" {

					Reportes.ReporteLS(ParticionMontada, Metodos.Trim(Variables.MapComandos["path"]), 1)

				} else if Variables.MapComandos["ruta"] == "" {

					color.HEX("#de4843", false).Println("Debe De Indicar Un Directorio/os O Archivos Valido")
					fmt.Println("")

				} else {

					color.HEX("#de4843", false).Println("No Existe La Ruta Indicada")
					fmt.Println("")

				}

			} else {

				color.HEX("#de4843", false).Println("No Existe El Reporte Indicado")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("No Existe El Id Indicado")
			fmt.Println("")

		}

	}