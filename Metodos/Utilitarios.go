
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

    package Metodos

	import (
		"../Variables"
		"bytes"
		"encoding/binary"
		"fmt"
		"github.com/gookit/color"
		"os"
		"regexp"
		"runtime"
		"strings"
		"unsafe"
	)

//----------------------------------------------------Métodos-----------------------------------------------------------

    //Obtener Nombre Del Sistema Operativo : windows/darwin(MacOs)/linux

    func ObtenerNombreOs() {

	    Variables.SistemaOperativo = runtime.GOOS

    }
    
    func SplitComando(Comando string) []string {

        var ArregloAuxiliar []string
        ArregloAuxiliar = strings.Split(Comando, " -")

        return ArregloAuxiliar

    }

    func SplitParametro(Parametro string) []string {

		var ArregloAuxiliar []string
		ArregloAuxiliar = strings.Split(Parametro, "->")

		return ArregloAuxiliar

	}

	func SplitArchivo(Archivo string) []string {

		var ArregloAuxiliar []string
		ArregloAuxiliar = strings.Split(Archivo, ".")

		return ArregloAuxiliar

	}

	func SplitContenidoArchivo(ContenidoArchivo string) []string {

		var ArregloAuxiliar []string
		ArregloAuxiliar = strings.Split(ContenidoArchivo, "\n")

		return ArregloAuxiliar

	}

	func BuscarPrefijo(Cadena string) bool {

		//Variables
		var Bandera bool

		Bandera = strings.HasPrefix(Cadena, "-")

		return Bandera

	}

	func Trim(Cadena string) string {

		//Variables
        var EspaciosEnLados *regexp.Regexp
        var EspaciosEnMedio *regexp.Regexp
        var TrimString string

        //Asignaciones
    	EspaciosEnLados = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`) //Conseguir Todos Los Espacios INiciales Y Finales
		EspaciosEnMedio = regexp.MustCompile(`[\s\p{Zs}]{2,}`) //Conseguir 2 O Mas Espacios Entre La Cadena

		TrimString = EspaciosEnLados.ReplaceAllString(Cadena, "")
		TrimString = EspaciosEnMedio.ReplaceAllString(TrimString, " ")

    	return TrimString
	}

	func BuscarSeparador(Cadena string) bool {

		//Variables
		var Bandera bool

		Bandera = strings.Contains(Cadena, "\\*")

		return Bandera

	}

	func QuitarComillas(Ruta string) string {

		//Variables
		var CadenaComillas string

		//Asignación
		CadenaComillas = Ruta

		CadenaComillas = Trim(CadenaComillas)
		CadenaComillas = strings.TrimPrefix(CadenaComillas, "\"")
		CadenaComillas = strings.TrimSuffix(CadenaComillas, "\"")

		return CadenaComillas

	}

	func ExisteRuta(Ruta string) bool {

		//Variables
		var Archivo os.FileInfo
		var AvisoError error

		//Asignación
		Ruta = QuitarComillas(Ruta)
		Variables.NoExisteArchivo = false

        //Verificar Existencia
		if Archivo, AvisoError = os.Stat(Ruta); AvisoError != nil {

			if os.IsNotExist(AvisoError) {

				Variables.NoExisteArchivo = true
				return false

			} else {

				Variables.NoExisteArchivo = true
				return false

			}

		}

		//Verificar Si Es Archivo
		if Archivo.IsDir() {

			Variables.NoExisteArchivo = true
			return false

		}

		Variables.NoExisteArchivo = false
		return true
	}

	func VerificarYCrearRutas(Ruta string) bool {

		//Variables
		var Archivo os.FileInfo
		var AvisoError error
		_ = Archivo

		//Asignación
		Ruta = Trim(Ruta)
		Ruta = QuitarComillas(Ruta)

		if Archivo, AvisoError = os.Stat(Ruta); os.IsNotExist(AvisoError) {

			AvisoError = os.MkdirAll(Ruta, os.ModePerm)

			if AvisoError != nil {

				return false

			} else {

			    Variables.CreeDirectorio = true

			}
		}

		return true

	}

    func ContineCaractereProhibidos(Cadena string) bool {

    	//Variables
    	var Bandera bool
        var ArregloCaracteres [34]string


    	//Asignación
    	Bandera = false
    	ArregloCaracteres[0] = "!"
		ArregloCaracteres[1] = "#"
		ArregloCaracteres[2] = "$"
		ArregloCaracteres[3] = "%"
		ArregloCaracteres[4] = "&"
		ArregloCaracteres[5] = "/"
		ArregloCaracteres[6] = "("
		ArregloCaracteres[7] = ")"
		ArregloCaracteres[8] = "="
		ArregloCaracteres[9] = "?"
		ArregloCaracteres[10] = "´"
		ArregloCaracteres[11] = "\\"
		ArregloCaracteres[12] = "¡"
		ArregloCaracteres[13] = "¿"
		ArregloCaracteres[14] = "'"
		ArregloCaracteres[15] = "|"
		ArregloCaracteres[16] = "°"
		ArregloCaracteres[17] = "¬"
		ArregloCaracteres[18] = "+"
		ArregloCaracteres[19] = "~"
		ArregloCaracteres[20] = "*"
		ArregloCaracteres[21] = "["
		ArregloCaracteres[22] = "]"
		ArregloCaracteres[23] = "{"
		ArregloCaracteres[24] = "}"
		ArregloCaracteres[25] = "^"
		ArregloCaracteres[26] = "`"
		ArregloCaracteres[27] = ","
		ArregloCaracteres[28] = "-"
		ArregloCaracteres[29] = ":"
		ArregloCaracteres[30] = ";"
		ArregloCaracteres[31] = "<"
		ArregloCaracteres[32] = ">"
		ArregloCaracteres[33] = "@"

		for Con := 0; Con < len(ArregloCaracteres); Con++ {

			Bandera = strings.Contains(Cadena, ArregloCaracteres[Con])

			if Bandera {

				return false

			}

		}

    	return true

	}

	func VerificarExtension(Cadena string) bool {

		//Variables
		var ArregloAuxiliar []string
		ArregloAuxiliar = strings.Split(Cadena, ".")

		if len(ArregloAuxiliar) > 1 {

			if ArregloAuxiliar[1] == "dsk" {

				return true

			} else {

				return false

			}

		} else {

			return false

		}

	}

    func ExisteArchivo(Cadena string) bool {

		//Variables
		var Archivo os.FileInfo
		var AvisoError error
		_ = Archivo

		//Asignación
		Cadena = QuitarComillas(Cadena)

		//Verificar Existencia
		if Archivo, AvisoError = os.Stat(Cadena); AvisoError != nil {

			if os.IsNotExist(AvisoError) {

				return false

			}

		}

		return true
	}

	func EscribirArchivoBinario(Archivo *os.File, ArregloBytes []byte) {

        //Variables
		var IntAuxiliar int
		var AvisoError error

		//Asignación
		_ = IntAuxiliar

		//Escribir en Archivo
		IntAuxiliar, AvisoError = Archivo.Write(ArregloBytes)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Escribir El Archivo")
		    fmt.Println("")

		}

	}

	func LeerArchivoBinario(Archivo *os.File,  Cantidad int) []byte {

		//Variables
		var Auxiliar int
		var AvisoError error
		var ArregloBytes []byte

		//Asignación
		ArregloBytes = make([]byte, Cantidad)
		_ = Auxiliar

		//Leer Bytes
		Auxiliar, AvisoError = Archivo.Read(ArregloBytes)

		//Catch Error

		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Leer El Archivo")
			fmt.Println("")

		}

		return ArregloBytes

	}

    func LeerArchivoBinarioArraglo(Ruta string) (Variables.MBREstructura, bool) {

		//Variables
		var Archivo *os.File
		var AvisoError error
		var MBRAuxiliar Variables.MBREstructura
		var PosicionMBR int
		var ArregloBytes []byte
		var Decodificador *bytes.Buffer

		//Abrir El Archivo
		Archivo, AvisoError = os.Open(Ruta)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
			fmt.Println("")
			return MBRAuxiliar, false

		} else {

			//Estructura Auxiliar
			MBRAuxiliar = Variables.MBREstructura{}
			MBRDireccion := &MBRAuxiliar

			//Obtener Posicion Del MBR
			PosicionMBR = int(unsafe.Sizeof(MBRAuxiliar))

			//Obtener Arreglo De Bytes
			ArregloBytes = LeerArchivoBinario(Archivo, PosicionMBR)

			//Decodificar Binario
			Decodificador = bytes.NewBuffer(ArregloBytes)

			AvisoError = binary.Read(Decodificador, binary.BigEndian, MBRDireccion)

			if AvisoError != nil {

				color.HEX("#de4843", false).Println("Error Al Leer El Archivo")
				fmt.Println("")
				return MBRAuxiliar, false

			} else {

				Archivo.Close()
				return MBRAuxiliar, true

			}

		}

	}

	func EscribirArchivoBinarioArreglo(MBRAuxiliar Variables.MBREstructura) {

		//Variables
		var Archivo *os.File
		var CadenaBinariaMBR bytes.Buffer
		var AvisoError error
		var MBRModificado Variables.MBREstructura

		//Abrir El Archivo
		Archivo, AvisoError = os.OpenFile(Variables.MapComandos["path"], os.O_WRONLY, os.ModePerm)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
			fmt.Println("")

		} else {

			//Mover Puntero
			_, _ = Archivo.Seek(0, 0)

			//Asignación
			MBRModificado = MBRAuxiliar
			MBRDireccion := &MBRModificado

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaMBR, binary.BigEndian, MBRDireccion)
			EscribirArchivoBinario(Archivo, CadenaBinariaMBR.Bytes())

			Archivo.Close()

		}

	}

	func EscribirArchivoBinarioArregloDelete(MBRAuxiliar Variables.MBREstructura, PosicionInicial int64, PosicionFinal int64) {

		//Variables
		var CeroBinario int
		var CeroByte *int
		var Archivo *os.File
		var CadenaBinariaInicio bytes.Buffer
		var CadenaBinariaMBR bytes.Buffer
		var AvisoError error
		var MBRModificado Variables.MBREstructura

		//Abrir El Archivo
		Archivo, AvisoError = os.OpenFile(Variables.MapComandos["path"], os.O_WRONLY, os.ModePerm)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
			fmt.Println("")

		} else {

			//Mover Puntero
			_, _ = Archivo.Seek(0, 0)

			//Asignación
			MBRModificado = MBRAuxiliar
			MBRDireccion := &MBRModificado

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaMBR, binary.BigEndian, MBRDireccion)
			EscribirArchivoBinario(Archivo, CadenaBinariaMBR.Bytes())

			//Mover A Posicion Inicial Particion
			_, _ = Archivo.Seek(PosicionInicial, 0)

			//Asignación
			CeroBinario = 0
			CeroByte = &CeroBinario

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaInicio, binary.BigEndian, CeroByte)
			EscribirArchivoBinario(Archivo, CadenaBinariaInicio.Bytes())

			//Mover A Posicion Final Particion
			_, _ = Archivo.Seek(PosicionFinal, 0)

			//Asignación
			CeroBinario = 0
			CeroByte = &CeroBinario

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaInicio, binary.BigEndian, CeroByte)
			EscribirArchivoBinario(Archivo, CadenaBinariaInicio.Bytes())


			Archivo.Close()

		}

	}

	func EscribirArchivoBinarioEBR(EBRAuxiliar Variables.EBREstructura, PosicionInicial int64) {

		//Variables
		var Archivo *os.File
		var CadenaBinariaEBR bytes.Buffer
		var AvisoError error
		var EBRModificado Variables.EBREstructura

		//Abrir El Archivo
		Archivo, AvisoError = os.OpenFile(Variables.MapComandos["path"], os.O_WRONLY, os.ModePerm)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
			fmt.Println("")

		} else {

			//Mover Puntero
			_, _ = Archivo.Seek(PosicionInicial, 0)

			//Asignación
			EBRModificado = EBRAuxiliar
			EBRDireccion := &EBRModificado

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaEBR, binary.BigEndian, EBRDireccion)
			EscribirArchivoBinario(Archivo, CadenaBinariaEBR.Bytes())

			Archivo.Close()

		}

	}

	func EscribirArchivoBinarioEBRDelete(EBREliminar Variables.EBREstructura, PosicionInicial int64) {

		//Variables
		var CeroBinario int
		var PosicionFinal int64
		var CeroByte *int
		var Archivo *os.File
		var CadenaBinariaEBR bytes.Buffer
		var CadenaBinariaEBRFinal bytes.Buffer
		var AvisoError error		

		//Abrir El Archivo
		Archivo, AvisoError = os.OpenFile(Variables.MapComandos["path"], os.O_WRONLY, os.ModePerm)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
			fmt.Println("")

		} else {

			//Mover Puntero
			_, _ = Archivo.Seek(PosicionInicial, 0)

			//Asignación
			CeroBinario = 0
			CeroByte = &CeroBinario

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaEBR, binary.BigEndian, CeroByte)
			EscribirArchivoBinario(Archivo, CadenaBinariaEBR.Bytes())

			//Posicion Final
			PosicionFinal = PosicionInicial + int64(unsafe.Sizeof(Variables.EBREstructura{})) + EBREliminar.SizeEBR

			//Mover A Posicion Final Particion
			_, _ = Archivo.Seek(PosicionFinal, 0)

			//Asignación
			CeroBinario = 0
			CeroByte = &CeroBinario

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaEBRFinal, binary.BigEndian, CeroByte)
			EscribirArchivoBinario(Archivo, CadenaBinariaEBRFinal.Bytes())


			Archivo.Close()

		}

	}

	func EscribirArchivoBinarioEBRAdd(EBRModificar Variables.EBREstructura, PosicionInicial int64) {

		//Variables
		var Archivo *os.File
		var CadenaBinariaEBR bytes.Buffer
		var AvisoError error
		var EBRModificado Variables.EBREstructura

		//Abrir El Archivo
		Archivo, AvisoError = os.OpenFile(Variables.MapComandos["path"], os.O_WRONLY, os.ModePerm)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
			fmt.Println("")

		} else {

			//Mover Puntero
			_, _ = Archivo.Seek(PosicionInicial, 0)

			//Asignación
			EBRModificado = EBRModificar
			EBRDireccion := &EBRModificado

			//Escribir Archivo
			_ = binary.Write(&CadenaBinariaEBR, binary.BigEndian, EBRDireccion)
			EscribirArchivoBinario(Archivo, CadenaBinariaEBR.Bytes())

			Archivo.Close()

		}

	}

	func LeerArchivoBinarioEBR(Ruta string, PosicionExtendida int64) (Variables.EBREstructura, bool) {

		//Variables
		var Archivo *os.File
		var AvisoError error
		var EBRAuxiliar Variables.EBREstructura
		var PosicionEBR int
		var ArregloBytes []byte
		var Decodificador *bytes.Buffer

		//Abrir El Archivo
		Archivo, AvisoError = os.Open(Ruta)

		//Catch Error
		if AvisoError != nil {

			color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
			fmt.Println("")
			return EBRAuxiliar, false

		} else {

			//Estructura Auxiliar
			EBRAuxiliar = Variables.EBREstructura{}
			EBRDireccion := &EBRAuxiliar

			//Obtener Posicion Del MBR
			PosicionEBR = int(unsafe.Sizeof(EBRAuxiliar))

			//Mover Puntero
			_, _ = Archivo.Seek(PosicionExtendida, 0)

			//Obtener Arreglo De Bytes
			ArregloBytes = LeerArchivoBinario(Archivo, PosicionEBR)

			//Decodificar Binario
			Decodificador = bytes.NewBuffer(ArregloBytes)

			AvisoError = binary.Read(Decodificador, binary.BigEndian, EBRDireccion)

			if AvisoError != nil {

				color.HEX("#de4843", false).Println("Error Al Leer El Archivo")
				fmt.Println("")
				return EBRAuxiliar, false

			} else {

				Archivo.Close()
				return EBRAuxiliar, true

			}

		}

	}