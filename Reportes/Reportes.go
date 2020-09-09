
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package Reportes

	import (
		"../Metodos"
		"../Variables"
		"bytes"
		"fmt"
		"github.com/gookit/color"
		"path/filepath"
		"sort"
		"strconv"
	)

//----------------------------------------------------Métodos-----------------------------------------------------------

	//Reporte MBR

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
						"<td bgcolor=\"#E6E6FA\">" + "Campo Nombre" + "</td> \n" +
						"<td bgcolor=\"#E6E6FA\">" + "Campo Valor" + "</td> \n" +
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
		Libre = true
		LibreEBR = true
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
				//ArregloDiscoEBRSize = append(ArregloDiscoEBRSize, int(ArregloParticiones[Contador].SizePart))

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
				//ArregloDiscoSize = append(ArregloDiscoSize, Metodos.EspaciosDisponibles[Contador].Tamano)

			}

		}

		if SizeExtendida != 0 {

			// Buscar Espacios Vacios
			Metodos.LimpiaDisco()
			Metodos.CreaDisco(SizeExtendida)

			for Contador := 0; Contador < len(ArregloEBR); Contador ++ {

				Metodos.LlenaDisco(int(ArregloEBR[Contador].InicioEBR) - InicioExtendida, int(ArregloEBR[Contador].SizeEBR))
				ArregloDiscoEBR = append(ArregloDiscoEBR, int(ArregloEBR[Contador].InicioEBR) - InicioExtendida)
				//ArregloDiscoEBRSize = append(ArregloDiscoEBRSize, int(ArregloEBR[Contador].SizeEBR))

			}

			Metodos.GeneraEspacios()

			for Contador := 0; Contador <= 200 - 1; Contador++ {

				if Metodos.EspaciosDisponibles[Contador].Disponible {

					ArregloDiscoEBR = append(ArregloDiscoEBR, Metodos.EspaciosDisponibles[Contador].P1)
					//ArregloDiscoEBRSize = append(ArregloDiscoEBRSize, Metodos.EspaciosDisponibles[Contador].Tamano)

				}

			}

		}

		sort.Ints(ArregloDisco)

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

					} else if ArregloParticiones[Contador].TipoPart == 'e' {

						Cadena += "<td bgcolor=\"#FF3F59\" height = \"100\"> \n" +
							"<table color='#FFFFFF' cellspacing='0'> \n" +
							"<tr><td bgcolor=\"#FF3F59\" height = \"50\" colspan=\"" + strconv.Itoa(len(ArregloDiscoEBR)) + "\">" + "Extendida_" + NombreParticion + ": " +  strconv.FormatFloat(float64(ArregloParticiones[Con].SizePart * 100) / float64(MBRAuxiliar.SizeMbr + 200), 'f', 2, 64) + "%</td></tr> \n" +
							"<tr> \n"

						sort.Ints(ArregloDiscoEBR)

						for Count := 0; Count < len(ArregloDiscoEBR); Count++ {

							for CountEBR := 0; CountEBR < len(ArregloEBR); CountEBR++ {

								if ArregloDiscoEBR[Count] == int(ArregloEBR[CountEBR].InicioEBR) - InicioExtendida {

									NombreParticion = string(bytes.Trim(ArregloEBR[CountEBR].NameEBR[:], "\x00"))

									Cadena += "<td bgcolor=\"#FF2828\" height = \"30\"> EBR: " + strconv.FormatFloat(float64(48 * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"
									Cadena += "<td bgcolor=\"#FF9058\" height = \"30\"> Logica_ " + NombreParticion + ": " + strconv.FormatFloat(float64(ArregloEBR[CountEBR].SizeEBR * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

								}

								LibreEBR = false

							}

							if LibreEBR {

								Cadena += "<td bgcolor=\"#EE0947\" height = \"50\"> Libre: " + strconv.FormatFloat(float64(10 * 100) / float64(SizeExtendida), 'f', 2, 64) + "%</td> \n"

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

				Cadena += "<td bgcolor=\"#EE0947\" height = \"50\"> Libre: " + strconv.FormatFloat(float64(10 * 100) / float64(MBRAuxiliar.SizeMbr + 200), 'f', 2, 64) + "%</td> \n"

			}

			Libre = true

		}

		_ = ArregloDiscoSize
		_ = ArregloDiscoEBRSize

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
