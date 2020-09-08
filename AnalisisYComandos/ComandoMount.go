
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package AnalisisYComandos

	import (
		"../Metodos"
		"../Variables"
		"bytes"
		"fmt"
		"github.com/gookit/color"
		"path/filepath"
		"strconv"
		"strings"
	)

//--------------------------------------------------------Métodos-------------------------------------------------------

	func VerificarComandoMount() {

		//Variables
		var Path bool
		var Name bool
		var ParametroExtra bool
		var ListaDeParticiones bool
		var ArregloParametros []string
		var ContadorPath int
		var ContadorName int

		//Asignación
		Path = false
		Name = false
		ParametroExtra = false
		ContadorPath = 0
		ContadorName = 0
		Variables.MapComandos = make(map[string]string)

		//Verificación De Parametros
		if len(Variables.ArregloComandos) > 1 {

			ListaDeParticiones = false

			for Contador := 1; Contador <= len(Variables.ArregloComandos)-1; Contador++ {

				//Obtener Parametro
				Variables.ArregloComandos[Contador] = Metodos.Trim(Variables.ArregloComandos[Contador])
				ArregloParametros = Metodos.SplitParametro(Variables.ArregloComandos[Contador])

				ArregloParametros[0] = strings.ToLower(ArregloParametros[0])
				ArregloParametros[0] = Metodos.Trim(ArregloParametros[0])

				switch ArregloParametros[0] {

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
		} else {

			ListaDeParticiones = true
			Path = true
			Name = true
			ParametroExtra = false
			ContadorPath = 1
			ContadorName = 1

		}

		if Path && Name && !ParametroExtra && ContadorPath == 1 && ContadorName == 1 {

			if ListaDeParticiones {

				ComandoMountListaParticiones()

			} else {

				ComandoMount()

			}

		} else {

			if ParametroExtra {

				color.HEX("#de4843", false).Println("Parametro Especificado No Valido")
				color.HEX("#de4843", false).Println("Parametros Validos: ")
				color.HEX("#de4843", false).Println("1). -path->    (Obligatorio)")
				color.HEX("#de4843", false).Println("2). -name->    (Obligatorio)")
				fmt.Println("")

			}

			if !Path {

				color.HEX("#de4843", false).Println("No Se Encuentra El Parametro -path-> o")
				color.HEX("#de4843", false).Println("No Existe El Archivo")
				fmt.Println("")

			}

			if !Name {

				color.HEX("#de4843", false).Println("No Se Encuentra el Parametro -name-> o")
				color.HEX("#de4843", false).Println("Existe Error En La Sintaxis")
				fmt.Println("")

			}

			if ContadorPath > 1 || ContadorName > 1 {

				color.HEX("#de4843", false).Println("Existen Demasiados Parametros")
				fmt.Println("")

			}

		}

	}

	func VerificarNombreParticion() (bool, Variables.EBREstructura, Variables.ParticionEstructura) {

		//Variables
		var Nombre string
		var NombreArray1 string
		var NombreArray2 string
		var NombreArray3 string
		var NombreArray4 string
		var NombreExtendida string
		var Bandera bool
		var ExisteNombre bool
		var InicioExtendida int64
		var MBRAuxiliar Variables.MBREstructura
		var EBRAuxiliar Variables.EBREstructura
		var PartAuxiliar Variables.ParticionEstructura
		var ArregloEBR []Variables.EBREstructura

		//Asignacion
		MBRAuxiliar, Bandera = Metodos.LeerArchivoBinarioArraglo(Variables.MapComandos["path"])
		Nombre = Metodos.Trim(strings.ToLower(Variables.MapComandos["name"]))
		ExisteNombre = false
		InicioExtendida = 0
		NombreArray1 = string(bytes.Trim(MBRAuxiliar.Particion1MBR.NamePart[:], "\x00"))
		NombreArray2 = string(bytes.Trim(MBRAuxiliar.Particion2MBR.NamePart[:], "\x00"))
		NombreArray3 = string(bytes.Trim(MBRAuxiliar.Particion3MBR.NamePart[:], "\x00"))
		NombreArray4 = string(bytes.Trim(MBRAuxiliar.Particion4MBR.NamePart[:], "\x00"))

		if Bandera {

			if MBRAuxiliar.Particion1MBR.SizePart != 0 {

				if strings.EqualFold(Nombre, NombreArray1) {

					ExisteNombre = true

					if MBRAuxiliar.Particion1MBR.TipoPart == 'e' {

						InicioExtendida = MBRAuxiliar.Particion1MBR.InicioPart

					}

					PartAuxiliar = MBRAuxiliar.Particion1MBR

				}

			}

			if MBRAuxiliar.Particion2MBR.SizePart != 0 {

				if strings.EqualFold(Nombre, NombreArray2)  {

					ExisteNombre = true

				}

				if MBRAuxiliar.Particion2MBR.TipoPart == 'e' {

					InicioExtendida = MBRAuxiliar.Particion2MBR.InicioPart

				}

				PartAuxiliar = MBRAuxiliar.Particion2MBR

			}

			if MBRAuxiliar.Particion3MBR.SizePart != 0 {

				if strings.EqualFold(Nombre, NombreArray3)   {

					ExisteNombre = true

				}

				if MBRAuxiliar.Particion3MBR.TipoPart == 'e' {

					InicioExtendida = MBRAuxiliar.Particion3MBR.InicioPart

				}

				PartAuxiliar = MBRAuxiliar.Particion3MBR

			}

			if MBRAuxiliar.Particion4MBR.SizePart != 0 {

				if strings.EqualFold(Nombre, NombreArray4)   {

					ExisteNombre = true

				}

				if MBRAuxiliar.Particion4MBR.TipoPart == 'e' {

					InicioExtendida = MBRAuxiliar.Particion4MBR.InicioPart

				}

				PartAuxiliar = MBRAuxiliar.Particion4MBR

			}

			if InicioExtendida != 0 {

				ArregloEBR = ObtenerEBR(InicioExtendida)

				for Contador := 0; Contador < len(ArregloEBR); Contador++ {

					NombreExtendida = string(bytes.Trim(ArregloEBR[Contador].NameEBR[:], "\x00"))

					if strings.EqualFold(Variables.MapComandos["name"], NombreExtendida) {

						ExisteNombre = true
						EBRAuxiliar = ArregloEBR[Contador]

					}

				}

			}

		} else {

			color.HEX("#de4843", false).Println("Error Al Ejecutar El Comando fdisk")
			color.HEX("#de4843", false).Println("El Disco Se Encuentra Corrupto")
			fmt.Println("")

		}

		return ExisteNombre, EBRAuxiliar, PartAuxiliar

	}

	func GenerarIdentificadorDisco() string {

		//Variables
		var Archivo string
		var NombreDisco string
		var NuevoId string
		var Existe bool
		var ArrayDisco []string
		var ArrayLetra [26]string
		var IdAuxiliar Variables.IDEstructura

		//Asignacion
		Archivo = ""
		NombreDisco = ""
		NuevoId = ""
		Existe = false
		ArrayDisco = make([]string, 0)
		IdAuxiliar = Variables.IDEstructura{}
		ArrayLetra = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n" +
		                        "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

		//Obtener Nombre Archivo
		_, Archivo = filepath.Split(Metodos.Trim(Variables.MapComandos["path"]))

		//Obtener Nombre Del Disco
		ArrayDisco = strings.Split(Archivo, ".")
		NombreDisco = ArrayDisco[0]

		//Buscar Clave
		IdAuxiliar, Existe = Variables.MapIdentificador[NombreDisco]

		//Verificar Si Ya Existe El Disco Montado
		if Existe {

			IdAuxiliar.ParticionID = IdAuxiliar.ParticionID + 1
			NuevoId = "vd" + IdAuxiliar.LetraID + strconv.Itoa(IdAuxiliar.ParticionID)

			Variables.MapIdentificador[NombreDisco] = IdAuxiliar

		} else {

			//Verificar Letras Utilizadas
			for Disco := range Variables.MapIdentificador {

				for Contador := 0; Contador < len(ArrayLetra); Contador++ {

					if Variables.MapIdentificador[Disco].LetraID == ArrayLetra[Contador] {

						ArrayLetra[Contador] = ""

					}

				}

			}

			for Contador := 0; Contador < len(ArrayLetra); Contador++ {

				if ArrayLetra[Contador] != "" {

					IdAuxiliar.LetraID = ArrayLetra[Contador]
					IdAuxiliar.ParticionID = 1
					NuevoId = "vd" + IdAuxiliar.LetraID + strconv.Itoa(IdAuxiliar.ParticionID)
					Variables.MapIdentificador[NombreDisco] = IdAuxiliar
					break

				}

			}

		}

		return NuevoId

	}

	func LlenarLista(Largo, Size int) string {

		//Variables
		var Espacios string
		var Diferencia int

		//Asignacion
		Espacios = ""
		Diferencia = Size - Largo

		for Contador := 0; Contador < Diferencia; Contador++ {

			Espacios += " "

		}

		return Espacios

	}

	func ComandoMount(){

		//Variables
		var CodigoPart string
		var ExisteNombre bool
		var EBRAuxiliar Variables.EBREstructura
		var PartAuxiliar Variables.ParticionEstructura
		var PartMontada Variables.MountEstructura

		//Asignacion
		CodigoPart = ""
		ExisteNombre = false
		EBRAuxiliar = Variables.EBREstructura{}
		PartAuxiliar = Variables.ParticionEstructura{}
		ExisteNombre, EBRAuxiliar, PartAuxiliar = VerificarNombreParticion()

		if ExisteNombre {

			CodigoPart = GenerarIdentificadorDisco()

			if CodigoPart != "" {

				if EBRAuxiliar.SizeEBR != 0 {

					PartMontada.IdentificadorMount = CodigoPart
					PartMontada.EBRMount = EBRAuxiliar
					PartMontada.ParticionMount = PartAuxiliar
					PartMontada.NombreMount = Metodos.Trim(Variables.MapComandos["name"])
					PartMontada.RutaDiscoMount = Metodos.Trim(Variables.MapComandos["path"])
					PartMontada.TipoMount = "Logica"

				} else {

					PartMontada.IdentificadorMount = CodigoPart
					PartMontada.EBRMount = EBRAuxiliar
					PartMontada.ParticionMount = PartAuxiliar
					PartMontada.NombreMount = Metodos.Trim(Variables.MapComandos["name"])
					PartMontada.RutaDiscoMount = Metodos.Trim(Variables.MapComandos["path"])
					PartMontada.TipoMount = "Primaria"

				}

				Variables.ArregloParticionesMontadas = append(Variables.ArregloParticionesMontadas, PartMontada)
				color.Success.Println("Particion Montada Con Exito!")
				fmt.Println("")

			} else {

				color.HEX("#de4843", false).Println("Error No Existe Espacio Para Montar La Particion")
				fmt.Println("")

			}

		} else {

			color.HEX("#de4843", false).Println("Error No Existe La Particion Indicada")
			fmt.Println("")

		}

	}

	func ComandoMountListaParticiones() {

		fmt.Println("")
		fmt.Println("")
		color.HEX("#21C68A", false).Println("                                    " +
			"Lista De Particiones Montadas")
		fmt.Println("")
		color.HEX("#2194C6", false).Println(" ------ID------", "--------------------" +
			"Ruta--------------------", "------------Nombre------------")

		for Contador := 0; Contador < len(Variables.ArregloParticionesMontadas); Contador++ {

			color.HEX("#3CE90D", false).Println(" ",
				Variables.ArregloParticionesMontadas[Contador].IdentificadorMount,
				LlenarLista(len(Variables.ArregloParticionesMontadas[Contador].IdentificadorMount), 13),
				Variables.ArregloParticionesMontadas[Contador].RutaDiscoMount,
				LlenarLista(len(Variables.ArregloParticionesMontadas[Contador].RutaDiscoMount), 43),
				Variables.ArregloParticionesMontadas[Contador].NombreMount)

		}

		fmt.Println("")
		fmt.Println("")

	}
