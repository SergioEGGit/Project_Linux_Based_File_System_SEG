
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package Reportes

	import (
		"../Metodos"
		"../Variables"
		"bytes"
		"fmt"
		"github.com/gookit/color"
		"os"
		"os/exec"
		"path/filepath"
		"sort"
		"strconv"
		"strings"
		"unsafe"
	)

//----------------------------------------------------Métodos-----------------------------------------------------------

	// Reporte MBR

	func ReporteMBR(MBRAuxiliar Variables.MBREstructura, Ruta string, ArregloEBR []Variables.EBREstructura) {

		//Variables
		var Cadena string
		var SizeMBR string
		var IDMBR string
		var InicioPart1 string
		var SizePart1 string
		var InicioPart2 string
		var SizePart2 string
		var InicioPart3 string
		var SizePart3 string
		var InicioPart4 string
		var SizePart4 string
		var NombrePart1 string
		var NombrePart2 string
		var NombrePart3 string
		var NombrePart4 string
		var Directorio string
		var Archivo string
		var Path bool

		//Asignacion
		Cadena = ""
		Directorio = ""
		Archivo = ""
		Path = false
		SizeMBR = strconv.Itoa(int(MBRAuxiliar.SizeMbr))
		IDMBR = strconv.Itoa(int(MBRAuxiliar.IDMBR))
		InicioPart1 = strconv.Itoa(int(MBRAuxiliar.Particion1MBR.InicioPart))
		SizePart1 = strconv.Itoa(int(MBRAuxiliar.Particion1MBR.SizePart))
		InicioPart2 = strconv.Itoa(int(MBRAuxiliar.Particion2MBR.InicioPart))
		SizePart2 = strconv.Itoa(int(MBRAuxiliar.Particion2MBR.SizePart))
		InicioPart3 = strconv.Itoa(int(MBRAuxiliar.Particion3MBR.InicioPart))
		SizePart3 = strconv.Itoa(int(MBRAuxiliar.Particion3MBR.SizePart))
		InicioPart4 = strconv.Itoa(int(MBRAuxiliar.Particion4MBR.InicioPart))
		SizePart4 = strconv.Itoa(int(MBRAuxiliar.Particion4MBR.SizePart))
		NombrePart1 = string(bytes.Trim(MBRAuxiliar.Particion1MBR.NamePart[:], "\x00"))
		NombrePart2 = string(bytes.Trim(MBRAuxiliar.Particion2MBR.NamePart[:], "\x00"))
		NombrePart3 = string(bytes.Trim(MBRAuxiliar.Particion3MBR.NamePart[:], "\x00"))
		NombrePart4 = string(bytes.Trim(MBRAuxiliar.Particion4MBR.NamePart[:], "\x00"))

		// Comenzar Reporte
		Cadena = "digraph Reporte_MBR { \n" +
			"node [shape = plaintext] \n" +
			"some_node [ \n" +
			"label =< \n" +
				"<table border=\"0\" cellborder=\"1\" cellspacing=\"0\"> \n" +
					"<tr> \n" +
						"<td bgcolor = \" #FFA07A\" colspan=\" 2\">" + "Reporte_MBR" + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
						"<td bgcolor=\"#E6E6FA\">" + "Nombre" + "</td> \n" +
						"<td bgcolor=\"#E6E6FA\">" + "Valor" + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
						"<td bgcolor = \"#1A87E1\" colspan=\" 2\">" + "MBR" + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
						"<td bgcolor=\"#98FB98\">" + "Size_MBR" + "</td> \n" +
						"<td bgcolor=\"#98FB98\">" + SizeMBR + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
						"<td bgcolor=\"#98FB98\">" + "Fecha_Creacion_MBR" + "</td> \n" +
						"<td bgcolor=\"#98FB98\\\">" + string(MBRAuxiliar.FCreacionMBR[:]) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
						"<td bgcolor=\"#98FB98\">" + "Identificador_MBR" + "</td> \n" +
						"<td bgcolor=\"#98FB98\">" + IDMBR + "</td> \n" +
					"</tr> \n"

					//Particion 1

					if MBRAuxiliar.Particion1MBR.SizePart > 0 {

						Cadena += "<tr> \n" +
							"<td bgcolor = \"#FFA07A\" colspan=\" 2\">" + "Particion 1" + "</td> \n" +
							"</tr> \n" +

							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Estado_Particion_1" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion1MBR.StatusPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Tipo_Particion_1" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion1MBR.TipoPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Fit_Particion_1" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion1MBR.FitPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Inicio_Particion_1" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + InicioPart1 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Size_Particion_1" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + SizePart1 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Nombre_Particion_1" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + NombrePart1 + "</td> \n" +
							"</tr> \n"

					}

					//Particion 2

					if MBRAuxiliar.Particion2MBR.SizePart > 0 {

						Cadena += "<tr> \n" +
							"<td bgcolor = \" #FFA07A\" colspan=\" 2\">" + "Particion 2" + "</td> \n" +
							"</tr> \n" +

							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Estado_Particion_2" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion2MBR.StatusPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Tipo_Particion_2" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion2MBR.TipoPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Fit_Particion_2" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion2MBR.FitPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Inicio_Particion_2" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + InicioPart2 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Size_Particion_2" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + SizePart2 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Nombre_Particion_2" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + NombrePart2 + "</td> \n" +
							"</tr> \n"

					}

					//Particion 3

					if MBRAuxiliar.Particion3MBR.SizePart > 0 {

						Cadena += "<tr> \n" +
							"<td bgcolor = \" #FFA07A\" colspan=\" 2\">" + "Particion 3" + "</td> \n" +
							"</tr> \n" +

							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Estado_Particion_3" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion3MBR.StatusPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Tipo_Particion_3" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion3MBR.TipoPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Fit_Particion_3" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion3MBR.FitPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Inicio_Particion_3" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + InicioPart3 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Size_Particion_3" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + SizePart3 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Nombre_Particion_3" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + NombrePart3 + "</td> \n" +
							"</tr> \n"

					}

					//Particion 4

					if MBRAuxiliar.Particion4MBR.SizePart > 0 {

						Cadena += "<tr> \n" +
							"<td bgcolor = \" #FFA07A\" colspan=\" 2\">" + "Particion 4" + "</td> \n" +
							"</tr> \n" +

							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Estado_Particion_4" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion4MBR.StatusPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Tipo_Particion_4" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion4MBR.TipoPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Fit_Particion_4" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + string(MBRAuxiliar.Particion4MBR.FitPart) + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Inicio_Particion_4" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + InicioPart4 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Size_Particion_4" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + SizePart4 + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" + "Nombre_Particion_4" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" + NombrePart4 + "</td> \n" +
							"</tr> \n"

					}

		Cadena += "<tr> \n" +
			"<td bgcolor = \"#1A87E1\" colspan=\" 2\">" + "EBRs" + "</td> \n" +
			"</tr> \n"

		for Contador := 0; Contador < len(ArregloEBR); Contador++ {

			Cadena += "<tr> \n" +
				"<td bgcolor = \" #FFA07A\" colspan=\" 2\">" + "Particion " + strconv.Itoa(Contador + 1) + "</td> \n" +
				"</tr> \n" +

				"<tr> \n" +
				"<td bgcolor=\"#ADD8E6\">" + "Estado_Particion_" + strconv.Itoa(Contador + 1) + "</td> \n" +
				"<td bgcolor=\"#ADD8E6\">" + string(ArregloEBR[Contador].StatusEBR) + "</td> \n" +
				"</tr> \n" +
				"<tr> \n" +
				"<td bgcolor=\"#ADD8E6\">" + "Fit_Particion_" + strconv.Itoa(Contador + 1) + "</td> \n" +
				"<td bgcolor=\"#ADD8E6\">" + string(ArregloEBR[Contador].FitEBR) + "</td> \n" +
				"</tr> \n" +
				"<tr> \n" +
				"<td bgcolor=\"#ADD8E6\">" + "Inicio_Particion_" + strconv.Itoa(Contador + 1) + "</td> \n" +
				"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(ArregloEBR[Contador].InicioEBR)) + "</td> \n" +
				"</tr> \n" +
				"<tr> \n" +
				"<td bgcolor=\"#ADD8E6\">" + "Size_Particion_" + strconv.Itoa(Contador + 1) + "</td> \n" +
				"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(ArregloEBR[Contador].SizeEBR)) + "</td> \n" +
				"</tr> \n" +
				"<tr> \n" +
				"<td bgcolor=\"#ADD8E6\">" + "Siguiente_Particion_" + strconv.Itoa(Contador + 1) + "</td> \n" +
				"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(ArregloEBR[Contador].SiguienteEBR)) + "</td> \n" +
				"</tr> \n" +
				"<tr> \n" +
				"<td bgcolor=\"#ADD8E6\">" + "Nombre_Particion_" + strconv.Itoa(Contador + 1) + "</td> \n" +
				"<td bgcolor=\"#ADD8E6\">" + string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00")) + "</td> \n" +
				"</tr> \n"

		}

		Cadena +=	"</table>> \n" +
			"]; \n" +
			"}"

		// Obtener Directorio
		Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

		Path = Metodos.VerificarYCrearRutas(Directorio)

		if Path {

			Metodos.GenerarArchivoTxt("Reporte_MBR", Cadena, Directorio)
			Metodos.GenerarReporte("Reporte_MBR", Directorio, Archivo)

		} else {

			color.HEX("#de4843", false).Println("Error No Se Genero El Reporte Con Exito")
			fmt.Println("")

		}

	}

	// Reporte Disk

	func ReporteDisco(MBRAuxiliar Variables.MBREstructura, Ruta string, ArregloEBR []Variables.EBREstructura) {

		//Variables
		var Cadena string
		var Directorio string
		var Archivo string
		var NombreParticion string
		var InicioExtendida int
		var SizeExtendida int
		var Path bool
		var Libre bool
		var LibreEBR bool
		var ArregloParticiones []Variables.ParticionEstructura
		var ArregloDisco []int
		var ArregloDiscoSize []int
		var ArregloDiscoEBR []int
		var ArregloDiscoEBRSize []int

		//Asignacion
		Cadena = ""
		Directorio = ""
		Archivo = ""
		Libre = true
		LibreEBR = true
		Path = false
		SizeExtendida = 0
		InicioExtendida = 0
		ArregloParticiones = make([]Variables.ParticionEstructura, 0)
		ArregloDisco = make([]int, 0)
		ArregloDiscoSize = make([]int, 0)
		ArregloDiscoEBR = make([]int, 0)
		ArregloDiscoEBRSize = make([]int, 0)
		ArregloParticiones = append(ArregloParticiones, MBRAuxiliar.Particion1MBR)
		ArregloParticiones = append(ArregloParticiones, MBRAuxiliar.Particion2MBR)
		ArregloParticiones = append(ArregloParticiones, MBRAuxiliar.Particion3MBR)
		ArregloParticiones = append(ArregloParticiones, MBRAuxiliar.Particion4MBR)

		// Buscar Espacios Vacios
		Metodos.LimpiaDisco()
		Metodos.CreaDisco(int(MBRAuxiliar.SizeMbr) + 200)
		Metodos.LlenaDisco(0, 200)

		for Contador := 0; Contador < len(ArregloParticiones); Contador ++ {

			if ArregloParticiones[Contador].SizePart != 0 {

				Metodos.LlenaDisco(int(ArregloParticiones[Contador].InicioPart), int(ArregloParticiones[Contador].SizePart))
				ArregloDisco = append(ArregloDisco, int(ArregloParticiones[Contador].InicioPart))

				if ArregloParticiones[Contador].TipoPart == 'e' {

					InicioExtendida = int(ArregloParticiones[Contador].InicioPart)
					SizeExtendida = int(ArregloParticiones[Contador].SizePart)

				}
			}

		}

		Metodos.GeneraEspacios()

		for Contador := 0; Contador <= 200 - 1; Contador++ {

			if Metodos.EspaciosDisponibles[Contador].Disponible {

				ArregloDisco = append(ArregloDisco, Metodos.EspaciosDisponibles[Contador].P1)

			}

		}

		sort.Ints(ArregloDisco)

		for Contador := 0; Contador < len(ArregloDisco); Contador++ {

			for Con := 0; Con < len(ArregloParticiones); Con ++ {

				if ArregloDisco[Contador] == int(ArregloParticiones[Con].InicioPart) {

					ArregloDiscoSize = append(ArregloDiscoSize, int(ArregloParticiones[Con].SizePart))

				}

			}

			for Con := 0; Con <= 200 - 1; Con++ {

				if Metodos.EspaciosDisponibles[Con].Disponible {

					if ArregloDisco[Contador] == Metodos.EspaciosDisponibles[Con].P1 {

						ArregloDiscoSize = append(ArregloDiscoSize, Metodos.EspaciosDisponibles[Con].Tamano)

					}

				}

			}

		}

		if SizeExtendida != 0 {

			// Buscar Espacios Vacios
			Metodos.LimpiaDisco()
			Metodos.CreaDisco(SizeExtendida)

			for Contador := 0; Contador < len(ArregloEBR); Contador ++ {

				NombreParticion = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

				if ArregloEBR[Contador].SizeEBR != 0 {

					Metodos.LlenaDisco(int(ArregloEBR[Contador].InicioEBR) - InicioExtendida, int(ArregloEBR[Contador].SizeEBR) + 48)
					ArregloDiscoEBR = append(ArregloDiscoEBR, int(ArregloEBR[Contador].InicioEBR) - InicioExtendida)

				}

		    }

			Metodos.GeneraEspacios()

			for Contador := 0; Contador <= 200 - 1; Contador++ {

				if Metodos.EspaciosDisponibles[Contador].Disponible {

					ArregloDiscoEBR = append(ArregloDiscoEBR, Metodos.EspaciosDisponibles[Contador].P1)

				}

			}

			sort.Ints(ArregloDiscoEBR)

			for Contador := 0; Contador < len(ArregloDiscoEBR); Contador++ {

				for Con := 0; Con < len(ArregloEBR); Con ++ {

					if ArregloDiscoEBR[Contador] == int(ArregloEBR[Con].InicioEBR) - InicioExtendida {

							ArregloDiscoEBRSize = append(ArregloDiscoEBRSize, int(ArregloEBR[Con].SizeEBR))

					}

				}

				for Con := 0; Con <= 200 - 1; Con++ {

					if Metodos.EspaciosDisponibles[Con].Disponible {

						if ArregloDiscoEBR[Contador] == Metodos.EspaciosDisponibles[Con].P1 {

							ArregloDiscoEBRSize = append(ArregloDiscoEBRSize, Metodos.EspaciosDisponibles[Con].Tamano)

						}

					}

				}

			}

		}

		Cadena += "digraph Reporte_Disco { \n" +
			"shape = plaintext \n" +
			"label=< \n" +
			"<table border = '1' cellborder = '1'> \n" +
			"<tr><td bgcolor=\"#FF9B48\" height = \"50\"> MBR: " + strconv.FormatFloat(float64(200 * 100) / float64(MBRAuxiliar.SizeMbr + 200), 'f', 2, 64) + "%</td> \n"

		for Contador := 0; Contador < len(ArregloDisco); Contador++ {

			for Con := 0; Con < len(ArregloParticiones); Con++ {

				if ArregloDisco[Contador] == int(ArregloParticiones[Con].InicioPart) {

					NombreParticion = string(bytes.Trim(ArregloParticiones[Con].NamePart[:], "\x00"))

					if ArregloParticiones[Con].TipoPart == 'p' {

						Cadena += "<td bgcolor=\"#B6EE09\" height = \"50\"> Primaria_" + NombreParticion + ": " + strconv.FormatFloat(float64(ArregloParticiones[Con].SizePart * 100) / float64(MBRAuxiliar.SizeMbr + 200), 'f', 2, 64) + "%</td> \n"

					} else if ArregloParticiones[Con].TipoPart == 'e' {

						Cadena += "<td bgcolor=\"#FF3F59\" height = \"100\"> \n" +
							"<table color='#FFFFFF' cellspacing='0'> \n" +
							"<tr><td bgcolor=\"#FF3F59\" height = \"50\" colspan=\"" + strconv.Itoa(len(ArregloDiscoEBR) * 2) + "\">" + "Extendida_" + NombreParticion + ": " +  strconv.FormatFloat(float64(ArregloParticiones[Con].SizePart * 100) / float64(MBRAuxiliar.SizeMbr + 200), 'f', 2, 64) + "%</td></tr> \n" +
							"<tr> \n"

						for Count := 0; Count < len(ArregloDiscoEBR); Count++ {

							for CountEBR := 0; CountEBR < len(ArregloEBR); CountEBR++ {

								if ArregloDiscoEBR[Count] == int(ArregloEBR[CountEBR].InicioEBR) - InicioExtendida {

									NombreParticion = string(bytes.Trim(ArregloEBR[CountEBR].NameEBR[:], "\x00"))

									Cadena += "<td bgcolor=\"#FF2828\" height = \"30\"> EBR: " + strconv.FormatFloat(float64(48 * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

									if strings.EqualFold(NombreParticion, "none") {

										if ArregloDiscoEBRSize[0] == 0 {

											Cadena += "<td bgcolor=\"#EE0947\" height = \"50\"> Libre: " + strconv.FormatFloat(float64((ArregloDiscoEBRSize[Count + 1] - 48) * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

										} else {

											Cadena += "<td bgcolor=\"#EE0947\" height = \"50\"> Libre: " + strconv.FormatFloat(float64((ArregloDiscoEBRSize[Count] - 48) * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

										}

									} else {

										Cadena += "<td bgcolor=\"#FF9058\" height = \"30\"> Logica_ " + NombreParticion + ": " + strconv.FormatFloat(float64(ArregloEBR[CountEBR].SizeEBR * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

									}

									LibreEBR = false

								}

							}

							if LibreEBR {

								if ArregloDiscoEBRSize[0] == 0 {

									if ArregloDiscoEBRSize[Count] != 0 {

										Cadena += "<td bgcolor=\"#EE0947\" height = \"50\"> Libre: " + strconv.FormatFloat(float64(ArregloDiscoEBRSize[Count + 1] * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

									}

								} else {

									if ArregloDiscoEBRSize[Count] != 0 {

										Cadena += "<td bgcolor=\"#EE0947\" height = \"50\"> Libre: " + strconv.FormatFloat(float64(ArregloDiscoEBRSize[Count] * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

									}

								}

							}

							LibreEBR = true

						}

						Cadena += "</tr> \n" +
							"</table> \n" +
							"</td> \n"

					}

					Libre = false
				}

			}


			if Libre {

				Cadena += "<td bgcolor=\"#FF3657\" height = \"50\"> Libre: " + strconv.FormatFloat(float64(ArregloDiscoSize[Contador] * 100) / float64(MBRAuxiliar.SizeMbr + 200), 'f', 2, 64) + "%</td> \n"

			}

			Libre = true

		}

		Cadena += "</tr> \n </table> \n" +
			">; \n" +
			"} \n"


		// Obtener Directorio
		Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

		Path = Metodos.VerificarYCrearRutas(Directorio)

		if Path {

			Metodos.GenerarArchivoTxt("Reporte_Disco", Cadena, Directorio)
			Metodos.GenerarReporte("Reporte_Disco", Directorio, Archivo)

		} else {

			color.HEX("#de4843", false).Println("Error No Se Genero El Reporte Con Exito")
			fmt.Println("")

		}

	}

	// Reporte SuperBoot

	func ReporteSuperBoot(ParticionMontada Variables.MountEstructura, Ruta string) {

		//Variables
		var Cadena string
		var Directorio string
		var Archivo string
		var NombreDisco string
		var FechaCreacion string
		var FechaModificacion string
		var Path bool
		var Bandera bool
		var SBAuxiliar Variables.SuperBloqueEstructura
		var Particion Variables.MountEstructura

		//Asignacion
		Cadena = ""
		Directorio = ""
		Archivo = ""
		Path = false
		Bandera = false
		SBAuxiliar = Variables.SuperBloqueEstructura{}
		Particion = ParticionMontada

		// Verificar SuperBloque
		if Particion.ParticionMount.SizePart != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), Particion.ParticionMount.InicioPart)

		} else if Particion.EBRMount.SizeEBR != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), int64(int(Particion.EBRMount.InicioEBR) + int(unsafe.Sizeof(Variables.EBREstructura{}))))

		}

		if Bandera {

			if SBAuxiliar.MagicNumSuperBloque != 0 {

				NombreDisco = string(bytes.Trim(SBAuxiliar.NombreDiscoSuperBloque[:], "\x00"))
				FechaCreacion = string(bytes.Trim(SBAuxiliar.FechaCreacionSuperBloque[:], "\x00"))
				FechaModificacion = string(bytes.Trim(SBAuxiliar.FechaUltimoMontajeSuperBloque[:], "\x00"))

				// Comenzar Reporte
				Cadena = "digraph Reporte_SB { \n" +
					"node [shape = plaintext] \n" +
					"some_node [ \n" +
					"label =< \n" +
					"<table border=\"0\" cellborder=\"1\" cellspacing=\"0\"> \n" +
					"<tr> \n" +
					"<td bgcolor = \" #FFA07A\" colspan=\" 2\">" + "Reporte_SuperBoot" + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#E6E6FA\">" + "Nombre" + "</td> \n" +
					"<td bgcolor=\"#E6E6FA\">" + "Valor" + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Nombre_Disco" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + NombreDisco + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Arbol_Directorio_Count" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.ArbolCountSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Detalle_Directorio_Count" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.DetalleDirectorioCountSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Inodos_Count" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.InodosCountSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Bloques_Count" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.BloquesCountSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Arbol_Directorio_Free" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.ArbolFreeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Detalle_Directorio_Free" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.DetalleFreeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Inodos_Free" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.InodosFreeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Bloques_Free" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.BloquesFreeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Fecha_Creacion" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + FechaCreacion + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Fecha_Ultimo_Montaje" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + FechaModificacion + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Montaje_Count" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.MontajesSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_BitMap_Arbol_Directorio" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PBitmapArbolSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_Arbol_Directorio" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PArbolSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_BitMap_Detalle_Directorio" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PBitmapDetalleSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_Detalle_Directorio" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PDetalleSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_BitMap_Tabla_Inodos" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PBitmapTablaSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_Tabla_Inodos" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PTablaSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_BitMap_Bloques" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PBitmapBloquesSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_Bloques" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PBloquesSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Apuntador_Bitacora" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.PLogSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Arbol_Directorio_Size" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.ArbolSizeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Detalle_Directorio_Size" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.DetalleSizeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Inodos_Size" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.InodoSizeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Bloques_Size" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.BloquesSizeSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Arbol_Directorio_Free_Bit" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.ArbolFreeBitSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Detalle_Directorio_Free_Bit" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.DetalleFreeBitSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Inodos_Free_Bit" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.TablaFreeBitSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Bloques_Free_Bit" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.BloquesFreeBitSuperBloque)) + "</td> \n" +
					"</tr> \n" +
					"<tr> \n" +
					"<td bgcolor=\"#ADD8E6\">" + "Magic_Num" + "</td> \n" +
					"<td bgcolor=\"#ADD8E6\">" + strconv.Itoa(int(SBAuxiliar.MagicNumSuperBloque)) + "</td> \n" +
					"</tr> \n"

				Cadena +=	"</table>> \n" +
					"]; \n" +
					"}"

				// Obtener Directorio
				Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

				Path = Metodos.VerificarYCrearRutas(Directorio)

				if Path {

					Metodos.GenerarArchivoTxt("Reporte_SB", Cadena, Directorio)
					Metodos.GenerarReporte("Reporte_SB", Directorio, Archivo)

				} else {

					color.HEX("#de4843", false).Println("Error No Se Genero El Reporte Con Exito")
					fmt.Println("")

				}

			} else {

				color.HEX("#de4843", false).Println("La Particion Indicada Aun No Posee El Formato LWH")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Leer El SuperBloque")
			fmt.Println("")

		}

	}

	// Reporte Bitmap Arbol Vitual Directorio

	func ReporteBitmapAVD(ParticionMontada Variables.MountEstructura, Ruta string) {

		//Variables
		var Directorio string
		var Archivo string
		var Comando string
		var ContadorAuxiliar int
		var Path bool
		var Bandera bool
		var ArregloBytes []byte
		var ArchivoFisico *os.File
		var Command *exec.Cmd
		var AvisoError error
		var SBAuxiliar Variables.SuperBloqueEstructura
		var Particion Variables.MountEstructura

		//Asignacion
		ContadorAuxiliar = 1
		Directorio = ""
		Archivo = ""
		Comando = ""
		Path = false
		Bandera = false
		SBAuxiliar = Variables.SuperBloqueEstructura{}
		Particion = ParticionMontada

		// Verificar SuperBloque
		if Particion.ParticionMount.SizePart != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), Particion.ParticionMount.InicioPart)

		} else if Particion.EBRMount.SizeEBR != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), int64(int(Particion.EBRMount.InicioEBR) + int(unsafe.Sizeof(Variables.EBREstructura{}))))

		}

		if Bandera {

			if SBAuxiliar.MagicNumSuperBloque != 0 {

				ArregloBytes, Bandera = Metodos.LeerArchivoBinarioBitmapAVD(Particion.RutaDiscoMount, SBAuxiliar.PBitmapArbolSuperBloque, int(SBAuxiliar.ArbolCountSuperBloque))

				// Obtener Directorio
				Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

				Path = Metodos.VerificarYCrearRutas(Directorio)

				if Path {

					ArchivoFisico, AvisoError = os.Create(Ruta)

					if AvisoError != nil {

						color.HEX("#de4843", false).Println("Error Al Generar Reporte")
						fmt.Println("")

					} else {

						for Contador := 0; Contador < len(ArregloBytes); Contador++ {

							if ContadorAuxiliar == 20 {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + "\n")
									ContadorAuxiliar = 0

							} else {

								if Contador == len(ArregloBytes) - 1 {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])))

								} else {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + " | ")

								}

							}

							ContadorAuxiliar++

						}

						ArchivoFisico.Close()

						if Variables.SistemaOperativo == "windows" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo + " &"
							Command = exec.Command("cmd", "/C", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else if Variables.SistemaOperativo == "linux" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo
							Command = exec.Command("xdg-open", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else {

							color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
							fmt.Println("")

						}

					}

				}


			} else {

				color.HEX("#de4843", false).Println("La Particion Indicada Aun No Posee El Formato LWH")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Leer El SuperBloque")
			fmt.Println("")

		}

	}

	// Reporte Bitmap Detalle Directorio

	func ReporteBitmapDD(ParticionMontada Variables.MountEstructura, Ruta string) {

		//Variables
		var Directorio string
		var Archivo string
		var Comando string
		var ContadorAuxiliar int
		var Path bool
		var Bandera bool
		var ArregloBytes []byte
		var ArchivoFisico *os.File
		var Command *exec.Cmd
		var AvisoError error
		var SBAuxiliar Variables.SuperBloqueEstructura
		var Particion Variables.MountEstructura

		//Asignacion
		ContadorAuxiliar = 1
		Directorio = ""
		Archivo = ""
		Comando = ""
		Path = false
		Bandera = false
		SBAuxiliar = Variables.SuperBloqueEstructura{}
		Particion = ParticionMontada

		// Verificar SuperBloque
		if Particion.ParticionMount.SizePart != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), Particion.ParticionMount.InicioPart)

		} else if Particion.EBRMount.SizeEBR != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), int64(int(Particion.EBRMount.InicioEBR) + int(unsafe.Sizeof(Variables.EBREstructura{}))))

		}

		if Bandera {

			if SBAuxiliar.MagicNumSuperBloque != 0 {

				ArregloBytes, Bandera = Metodos.LeerArchivoBinarioBitmapDD(Particion.RutaDiscoMount, SBAuxiliar.PBitmapDetalleSuperBloque, int(SBAuxiliar.DetalleDirectorioCountSuperBloque))

				// Obtener Directorio
				Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

				Path = Metodos.VerificarYCrearRutas(Directorio)

				if Path {

					ArchivoFisico, AvisoError = os.Create(Ruta)

					if AvisoError != nil {

						color.HEX("#de4843", false).Println("Error Al Generar Reporte")
						fmt.Println("")

					} else {

						for Contador := 0; Contador < len(ArregloBytes); Contador++ {

							if ContadorAuxiliar == 20 {

								_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + "\n")
								ContadorAuxiliar = 0

							} else {

								if Contador == len(ArregloBytes) - 1 {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])))

								} else {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + " | ")

								}

							}

							ContadorAuxiliar++

						}

						ArchivoFisico.Close()

						if Variables.SistemaOperativo == "windows" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo + " &"
							Command = exec.Command("cmd", "/C", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else if Variables.SistemaOperativo == "linux" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo
							Command = exec.Command("xdg-open", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else {

							color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
							fmt.Println("")

						}

					}

				}


			} else {

				color.HEX("#de4843", false).Println("La Particion Indicada Aun No Posee El Formato LWH")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Leer El SuperBloque")
			fmt.Println("")

		}

	}

	// Reporte Bitmap Inodos

	func ReporteBitmapTI(ParticionMontada Variables.MountEstructura, Ruta string) {

		//Variables
		var Directorio string
		var Archivo string
		var Comando string
		var ContadorAuxiliar int
		var Path bool
		var Bandera bool
		var ArregloBytes []byte
		var ArchivoFisico *os.File
		var Command *exec.Cmd
		var AvisoError error
		var SBAuxiliar Variables.SuperBloqueEstructura
		var Particion Variables.MountEstructura

		//Asignacion
		ContadorAuxiliar = 1
		Directorio = ""
		Archivo = ""
		Comando = ""
		Path = false
		Bandera = false
		SBAuxiliar = Variables.SuperBloqueEstructura{}
		Particion = ParticionMontada

		// Verificar SuperBloque
		if Particion.ParticionMount.SizePart != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), Particion.ParticionMount.InicioPart)

		} else if Particion.EBRMount.SizeEBR != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), int64(int(Particion.EBRMount.InicioEBR) + int(unsafe.Sizeof(Variables.EBREstructura{}))))

		}

		if Bandera {

			if SBAuxiliar.MagicNumSuperBloque != 0 {

				ArregloBytes, Bandera = Metodos.LeerArchivoBinarioBitmapTI(Particion.RutaDiscoMount, SBAuxiliar.PBitmapTablaSuperBloque, int(SBAuxiliar.InodosCountSuperBloque))

				// Obtener Directorio
				Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

				Path = Metodos.VerificarYCrearRutas(Directorio)

				if Path {

					ArchivoFisico, AvisoError = os.Create(Ruta)

					if AvisoError != nil {

						color.HEX("#de4843", false).Println("Error Al Generar Reporte")
						fmt.Println("")

					} else {

						for Contador := 0; Contador < len(ArregloBytes); Contador++ {

							if ContadorAuxiliar == 20 {

								_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + "\n")
								ContadorAuxiliar = 0

							} else {

								if Contador == len(ArregloBytes) - 1 {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])))

								} else {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + " | ")

								}

							}

							ContadorAuxiliar++

						}

						ArchivoFisico.Close()

						if Variables.SistemaOperativo == "windows" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo + " &"
							Command = exec.Command("cmd", "/C", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else if Variables.SistemaOperativo == "linux" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo
							Command = exec.Command("xdg-open", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else {

							color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
							fmt.Println("")

						}

					}

				}


			} else {

				color.HEX("#de4843", false).Println("La Particion Indicada Aun No Posee El Formato LWH")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Leer El SuperBloque")
			fmt.Println("")

		}

	}

	// Reporte Bitmap Bloques

	func ReporteBitmapBQ(ParticionMontada Variables.MountEstructura, Ruta string) {

		//Variables
		var Directorio string
		var Archivo string
		var Comando string
		var ContadorAuxiliar int
		var Path bool
		var Bandera bool
		var ArregloBytes []byte
		var ArchivoFisico *os.File
		var Command *exec.Cmd
		var AvisoError error
		var SBAuxiliar Variables.SuperBloqueEstructura
		var Particion Variables.MountEstructura

		//Asignacion
		ContadorAuxiliar = 1
		Directorio = ""
		Archivo = ""
		Comando = ""
		Path = false
		Bandera = false
		SBAuxiliar = Variables.SuperBloqueEstructura{}
		Particion = ParticionMontada

		// Verificar SuperBloque
		if Particion.ParticionMount.SizePart != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), Particion.ParticionMount.InicioPart)

		} else if Particion.EBRMount.SizeEBR != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), int64(int(Particion.EBRMount.InicioEBR) + int(unsafe.Sizeof(Variables.EBREstructura{}))))

		}

		if Bandera {

			if SBAuxiliar.MagicNumSuperBloque != 0 {

				ArregloBytes, Bandera = Metodos.LeerArchivoBinarioBitmapBQ(Particion.RutaDiscoMount, SBAuxiliar.PBitmapTablaSuperBloque, int(SBAuxiliar.BloquesCountSuperBloque))

				// Obtener Directorio
				Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

				Path = Metodos.VerificarYCrearRutas(Directorio)

				if Path {

					ArchivoFisico, AvisoError = os.Create(Ruta)

					if AvisoError != nil {

						color.HEX("#de4843", false).Println("Error Al Generar Reporte")
						fmt.Println("")

					} else {

						for Contador := 0; Contador < len(ArregloBytes); Contador++ {

							if ContadorAuxiliar == 20 {

								_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + "\n")
								ContadorAuxiliar = 0

							} else {

								if Contador == len(ArregloBytes) - 1 {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])))

								} else {

									_, _ = ArchivoFisico.WriteString(strconv.Itoa(int(ArregloBytes[Contador])) + " | ")

								}

							}

							ContadorAuxiliar++

						}

						ArchivoFisico.Close()

						if Variables.SistemaOperativo == "windows" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo + " &"
							Command = exec.Command("cmd", "/C", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else if Variables.SistemaOperativo == "linux" {

							color.Success.Println("Reporte Generado Con Exito")
							fmt.Println("")

							Comando = Archivo
							Command = exec.Command("xdg-open", Comando)
							Command.Stdout = os.Stdout
							AvisoError = Command.Run()

							if AvisoError != nil {

								color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
								fmt.Println("")

							}

						} else {

							color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
							fmt.Println("")

						}

					}

				}


			} else {

				color.HEX("#de4843", false).Println("La Particion Indicada Aun No Posee El Formato LWH")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Leer El SuperBloque")
			fmt.Println("")

		}

	}
	
	// Reporte Arbol Completo

	func ObtenerAVDS(InicioAVDS int64, Ruta string) []Variables.AVDEstructura {

		//Variables
		var Contador int
		var Bandera bool
		var AVDAuxiliar Variables.AVDEstructura
		var ArregloAVDS []Variables.AVDEstructura

		//Asignación
		Contador = 0

		for {

			//Leer AVD
			AVDAuxiliar, Bandera = Metodos.LeerArchivoBinarioAVD(Ruta, InicioAVDS)

			//Lista Corrupta
			if !Bandera {

				return ArregloAVDS

			}

			//fmt.Println("Size: ", EBRAuxiliar.SizeEBR, "Inicio: ", EBRAuxiliar.InicioEBR, "Siguiente: ", EBRAuxiliar.SiguienteEBR, "Nombre: ", string(EBRAuxiliar.NameEBR[:]))

			ArregloAVDS = append(ArregloAVDS, AVDAuxiliar)
			InicioAVDS = ArregloAVDS[Contador].PDetalleDirectorioAVD
			Contador++

			if AVDAuxiliar.PArbolVirtualDirectorio == 0 {

				break

			}

		}

		return ArregloAVDS

	}

	func ObtenerDDS(InicioDDS int64, Ruta string) []Variables.DDEstructura {

		//Variables
		var Contador int
		var Bandera bool
		var DDAuxiliar Variables.DDEstructura
		var ArregloDDS []Variables.DDEstructura

		//Asignación
		Contador = 0

		for {

			//Leer AVD
			DDAuxiliar, Bandera = Metodos.LeerArchivoBinarioDD(Ruta, InicioDDS)

			//Lista Corrupta
			if !Bandera {

				return ArregloDDS

			}

			//fmt.Println("Size: ", EBRAuxiliar.SizeEBR, "Inicio: ", EBRAuxiliar.InicioEBR, "Siguiente: ", EBRAuxiliar.SiguienteEBR, "Nombre: ", string(EBRAuxiliar.NameEBR[:]))

			ArregloDDS = append(ArregloDDS, DDAuxiliar)
			InicioDDS = ArregloDDS[Contador].PDetalleDirectorioDD
			Contador++

			if DDAuxiliar.PDetalleDirectorioDD == 0 {

				break

			}

		}

		return ArregloDDS

	}

	func ObtenerInodos(InicioInodos int64, Ruta string) []Variables.TablaInodoEstructura {

		//Variables
		var Contador int
		var Bandera bool
		var TIAuxiliar Variables.TablaInodoEstructura
		var ArregloInodos []Variables.TablaInodoEstructura

		//Asignación
		Contador = 0

		for {

			//Leer AVD
			TIAuxiliar, Bandera = Metodos.LeerArchivoBinarioTI(Ruta, InicioInodos)

			//Lista Corrupta
			if !Bandera {

				return ArregloInodos

			}

			//fmt.Println("Size: ", EBRAuxiliar.SizeEBR, "Inicio: ", EBRAuxiliar.InicioEBR, "Siguiente: ", EBRAuxiliar.SiguienteEBR, "Nombre: ", string(EBRAuxiliar.NameEBR[:]))

			ArregloInodos = append(ArregloInodos, TIAuxiliar)
			InicioInodos = ArregloInodos[Contador].PTabalInodosTI
			Contador++

			if TIAuxiliar.PTabalInodosTI == 0 {

				break

			}

		}

		return ArregloInodos

	}
	
	func ReporteArbolCompletoAVD(ParticionMontada Variables.MountEstructura, Ruta string) {

		//Variables
		var Directorio string
		var Archivo string
		var Cadena string
		var FechaCreacion string
		var NombreDirectorio string
		var Propietario string
		var Path bool
		var Bandera bool
		var SBAuxiliar Variables.SuperBloqueEstructura
		var Particion Variables.MountEstructura
		var ArregloDirectorios []Variables.AVDEstructura
		var ArregloArchivos []Variables.DDEstructura
		var ArregloTablaInodos []Variables.TablaInodoEstructura

		//Asignacion
		Cadena = ""
		FechaCreacion = ""
		NombreDirectorio = ""
		Propietario = ""
		Directorio = ""
		Archivo = ""
		Path = false
		Bandera = false
		SBAuxiliar = Variables.SuperBloqueEstructura{}
		Particion = ParticionMontada
		ArregloDirectorios = make([]Variables.AVDEstructura, 0)
		ArregloArchivos = make([]Variables.DDEstructura, 0)
		ArregloTablaInodos = make([]Variables.TablaInodoEstructura, 0)

		// Verificar SuperBloque
		if Particion.ParticionMount.SizePart != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), Particion.ParticionMount.InicioPart)

		} else if Particion.EBRMount.SizeEBR != 0 {

			// Verificar SuperBloque
			SBAuxiliar, Bandera = Metodos.LeerArchivoBinarioSB(Metodos.Trim(Particion.RutaDiscoMount), int64(int(Particion.EBRMount.InicioEBR) + int(unsafe.Sizeof(Variables.EBREstructura{}))))

		}

		ArregloDirectorios = ObtenerAVDS(SBAuxiliar.PArbolSuperBloque, Particion.RutaDiscoMount)
		ArregloArchivos = ObtenerDDS(SBAuxiliar.PDetalleSuperBloque, Particion.RutaDiscoMount)
		ArregloTablaInodos = ObtenerInodos(SBAuxiliar.PTablaSuperBloque, Particion.RutaDiscoMount)

		if Bandera {

			if SBAuxiliar.MagicNumSuperBloque != 0 {

				if len(ArregloDirectorios) > 0 {

					for Contador := 0; Contador < len(ArregloDirectorios); Contador++ {

						FechaCreacion = string(bytes.Trim(ArregloDirectorios[Contador].FechaCreacionAVD[:], "\x00"))
						NombreDirectorio = string(bytes.Trim(ArregloDirectorios[Contador].NombreDirectorioAVD[:], "\x00"))
						Propietario = string(bytes.Trim(ArregloDirectorios[Contador].PropietarioAVD[:], "\x00"))

						// Comenzar Reporte
						Cadena = "digraph Reporte_TreeComplete { \n" +
							"node [shape = plaintext] \n" +
							"AVD0 [ \n" +
							"label =< \n" +
							"<table border=\"0\" cellborder=\"1\" cellspacing=\"0\"> \n" +
							"<tr> \n" +
							"<td  colspan = \"2\" bgcolor=\"#05C095\">" +  NombreDirectorio + "</td> \n" +
							"</tr> \n" +
							"<tr> \n" +
							"<td bgcolor=\"#ADD8E6\">" +  "Fecha_Creacion" + "</td> \n" +
							"<td bgcolor=\"#ADD8E6\">" +  FechaCreacion + "</td> \n" +
							"</tr> \n"

							for Con := 0; Con < len(ArregloDirectorios[Contador].PArraySubDirectoriosAVD); Con++ {

								Cadena += "<tr> \n" +
									"<td bgcolor=\"#ADD8E6\">" +  "Apuntador_SubDirectorios_" + strconv.Itoa(Con) + "</td> \n" +
									"<td bgcolor=\"#ADD8E6\">" +  strconv.Itoa(int(ArregloDirectorios[Contador].PArraySubDirectoriosAVD[Con])) + "</td> \n" +
									"</tr> \n"

							}

							Cadena += "<tr> \n" +
								"<td bgcolor=\"#ADD8E6\">" +  "Apuntador_DD" + "</td> \n" +
								"<td bgcolor=\"#ADD8E6\" port=\"1\">" +  strconv.Itoa(int(ArregloDirectorios[Contador].PDetalleDirectorioAVD)) + "</td> \n" +
								"</tr>" +
								"<tr> \n" +
								"<td bgcolor=\"#ADD8E6\">" +  "Apuntador_AVD" + "</td> \n" +
								"<td bgcolor=\"#ADD8E6\">" +  strconv.Itoa(int(ArregloDirectorios[Contador].PArbolVirtualDirectorio)) + "</td> \n" +
								"</tr> \n" +
								"<tr> \n" +
								"<td bgcolor=\"#ADD8E6\">" +  "Propietario" + "</td> \n" +
								"<td bgcolor=\"#ADD8E6\">" +  Propietario + "</td> \n" +
								"</tr> \n" +
								"</table>> \n" +
								"]; \n"

							if len(ArregloArchivos) > 0 {

								for ContadorDD := 0; ContadorDD < len(ArregloArchivos); ContadorDD++ {

									Cadena += "DD0 [ \n" +
										"label =< \n" +
										"<table border=\"0\" cellborder=\"1\" cellspacing=\"0\"> \n" +
										"<tr> \n" +
										"<td  colspan = \"2\" bgcolor=\"#05C095\">" +  "Detalle Directorio" + "</td> \n" +
										"</tr> \n"

									FechaCreacion := string(bytes.Trim(ArregloArchivos[ContadorDD].ArrayArchivosDD[0].FechaCreacionArchivoDDInformacion[:], "\x00"))
									FechaModificacion := string(bytes.Trim(ArregloArchivos[ContadorDD].ArrayArchivosDD[0].FechaModificacionArchivoDDInformacion[:], "\x00"))
									NombreArchivo := string(bytes.Trim(ArregloArchivos[ContadorDD].ArrayArchivosDD[0].NombreArchivoDDInformacion[:], "\x00"))

										Cadena += "<tr> \n" +
											"<td bgcolor=\"#FF627C\">" +  "Nombre_Archivo_1" + "</td> \n" +
											"<td bgcolor=\"#FF627C\" port=\"1\">" +  NombreArchivo + "</td> \n" +
											"</tr> \n"+
											"<tr> \n" +
											"<td bgcolor=\"#FF627C\">" +  "Fecha_Creacion_Archivo_1" + "</td> \n" +
											"<td bgcolor=\"#FF627C\">" +  FechaCreacion + "</td> \n" +
											"</tr> \n" +
											"<tr> \n" +
											"<td bgcolor=\"#FF627C\">" +  "Fecha_Modificacion_Archivo_1" + "</td> \n" +
											"<td bgcolor=\"#FF627C\">" +  FechaModificacion + "</td> \n" +
											"</tr> \n" +
											"<tr> \n" +
											"<td bgcolor=\"#FF627C\">" +  "Apuntador_Inodos" + "</td> \n" +
											"<td bgcolor=\"#FF627C\" port=\"1\">" +  strconv.Itoa(int(ArregloArchivos[ContadorDD].ArrayArchivosDD[0].PInodoArchivoDDInformacion)) + "</td> \n" +
											"</tr> \n"

											for I := 1; I < len(ArregloArchivos[0].ArrayArchivosDD); I++ {

												Cadena += "<tr> \n" +
													"<td bgcolor=\"#FF627C\">" +  "Archivo" + strconv.Itoa(I + 1) + "</td> \n" +
													"<td bgcolor=\"#FF627C\">" +  "-" + "</td> \n" +
													"</tr> \n"

											}

								Cadena += "<tr> \n" +
									"<td bgcolor=\"#FF627C\">" +  "Apuntador_DD" + "</td> \n" +
									"<td bgcolor=\"#FF627C\">" + strconv.Itoa(int(ArregloArchivos[0].PDetalleDirectorioDD)) + "</td> \n" +
									"</tr> \n" +
									"</table>> \n" +
									"]; \n"

								}

							}


							if len(ArregloTablaInodos) > 0 {


								for ContadorTI := 0; ContadorTI < len(ArregloTablaInodos); ContadorTI++ {

									Cadena += "TI0 [ \n" +
										"label =< \n" +
										"<table border=\"0\" cellborder=\"1\" cellspacing=\"0\"> \n" +
										"<tr> \n" +
										"<td  colspan = \"2\" bgcolor=\"#05C095\">" +  "Tabla Inodo" + "</td> \n" +
										"</tr> \n" +
										"<tr> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  "Numero_Inodo" + "</td> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  strconv.Itoa(int(ArregloTablaInodos[ContadorTI].NumeroInodoTI)) + "</td> \n" +
										"</tr> \n" +
										"<tr> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  "Size_Inodo" + "</td> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  strconv.Itoa(int(ArregloTablaInodos[ContadorTI].SizeArchivoTI)) + "</td> \n" +
										"</tr> \n" +
										"<tr> \n" +
										"<td bgcolor=\"#6A9BFA\">" + "Numero_Bloques_Inodo" + "</td> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  strconv.Itoa(int(ArregloTablaInodos[ContadorTI].NumeroBloquesTI)) + "</td> \n" +
										"</tr> \n"

									for ConTI := 0; ConTI < len(ArregloTablaInodos[ContadorTI].ArrayBloquesTI); ConTI++ {

										Cadena += "<tr> \n" +
											"<td bgcolor=\"#6A9BFA\">" +  "Bloque_" + strconv.Itoa(ConTI + 1) + "</td> \n" +
											"<td bgcolor=\"#6A9BFA\">" +  strconv.Itoa(int(ArregloTablaInodos[ContadorTI].ArrayBloquesTI[ConTI])) + "</td> \n" +
											"</tr> \n"

									}

									Propietario = string(bytes.Trim(ArregloTablaInodos[ContadorTI].PropietarioTI[:], "\x00"))


									Cadena += "<tr> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  "Apuntador_Tabla_Inodo" + "</td> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  strconv.Itoa(int(ArregloTablaInodos[ContadorTI].PTabalInodosTI)) + "</td> \n" +
										"</tr> \n" +
										"<tr> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  "Propietario" + "</td> \n" +
										"<td bgcolor=\"#6A9BFA\">" +  Propietario + "</td> \n" +
										"</tr> \n" +
										"</table>> \n" +
										"]; \n"

								}

							}



					}
					Cadena += "\n {rank = same; AVD0; DD0; TI0}\n" +
						"\n AVD0:1 -> DD0 \n" +
						"\n DD0:1 -> TI0 \n" +
						"}"

					// Obtener Directorio
					Directorio, Archivo = filepath.Split(Metodos.Trim(Ruta))

					Path = Metodos.VerificarYCrearRutas(Directorio)

					if Path {

						Metodos.GenerarArchivoTxt("Reporte_Tree_Complete", Cadena, Directorio)
						Metodos.GenerarReporte("Reporte_Tree_Complete", Directorio, Archivo)

					} else {

						color.HEX("#de4843", false).Println("Error No Se Genero El Reporte Con Exito")
						fmt.Println("")

					}

				} else {

					color.HEX("#de4843", false).Println("Error No Se Encuentran Directorios")
					fmt.Println("")

				}

			} else {

				color.HEX("#de4843", false).Println("La Particion Indicada Aun No Posee El Formato LWH")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Leer El SuperBloque")
			fmt.Println("")

		}


	}
