
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package Reportes

import (
	"../Metodos"
	"../Variables"
	"bytes"
	"fmt"
	"path/filepath"
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

		//Asignacion
		Cadena = ""
		Directorio = ""
		Archivo = ""
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
					"</tr> \n" +

					//Particion 1

					"<tr> \n" +
						"<td bgcolor = \" #FFA07A\" colspan=\" 2\">" + "Particion 1" + "</td> \n" +
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
					"</tr> \n" +

					//Particion 2

					"<tr> \n" +
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
					"</tr> \n" +

					//Particion 3

					"<tr> \n" +
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
					"</tr> \n" +

					//Particion 4

					"<tr> \n" +
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

		println(Directorio, "-------", Archivo)
		fmt.Scanln()

		Metodos.GenerarArchivoTxt("Reporte_MBR", Cadena)
		Metodos.GenerarReporte("Reporte_MBR", Ruta)

	}