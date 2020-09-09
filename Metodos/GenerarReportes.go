
//------------------------------------Paquetes E Imports-------------------------------------------

    package Metodos

    import (
	    "../Variables"
	    "bufio"
	    "fmt"
	    "log"
	    "os"
	    "os/exec"
    )

//-----------------------------------------Métodos-------------------------------------------------

    func GenerarArchivoTxt(NombreReporte, CadenaArchivo, Directorio string) {

        //Variables
        var BytesArray int
        var RutaArchivo string
        var Archivo *os.File
        var AvisoError error
        var Writer *bufio.Writer
        var BanderaOs bool

        //Rutas Y Archivos
        //Seleccion Sistema Operativo

        if Variables.SistemaOperativo == "windows" {

            //Abrir Archivo Y Generar Try Catch Error
            RutaArchivo = Directorio + NombreReporte + ".txt"
            Archivo, AvisoError = os.Create(RutaArchivo)

            //Catch Error

            if AvisoError != nil {

                fmt.Print("Error Al Generar El Archivo")

            }

            //Cerrar Archivo
            defer Archivo.Close()


            BanderaOs = true

        } else if Variables.SistemaOperativo == "linux"  {

            BanderaOs = true

        } else {

            BanderaOs = false

        }

        if BanderaOs == true {

            //Escribir Archivo

            Writer = bufio.NewWriter(Archivo)
            BytesArray, AvisoError = Writer.WriteString(CadenaArchivo)
            _ = BytesArray

            //Catch Error

            if AvisoError != nil {

                fmt.Print("Error Al Escribir El Archivo")

            }

            _ = Writer.Flush()

        } else {

            print("Sistema Operativo No Soportado")

        }
    }

    func GenerarReporte(NombreReporte, Directorio, Archivo string) {

        //Variables
        var RutaGraphviz string
        var Parametros string
        var FileInput string
        var FileOutput string
        var GvizCommand string
        var Command *exec.Cmd
        var AvisoError error

        //Rutas Y Archivos
        //Seleccion Sistema Operativo

        if Variables.SistemaOperativo == "windows" {

			RutaGraphviz = "dot "
            Parametros = "-Tpng -o "
            FileInput = " " + Directorio + NombreReporte + ".txt"
            FileOutput = Directorio + Archivo + " "

            GvizCommand = RutaGraphviz + Parametros + FileOutput + FileInput

            Command = exec.Command("cmd", "/c", GvizCommand)
            Command.Stdout = os.Stdout
			Command.Stderr = os.Stderr
            AvisoError = Command.Run()

            //Catch Error
            if AvisoError != nil {

				RutaGraphviz = "\"C:\\Program Files (x86)\\Graphviz2.38\\bin\\dot.exe\""
                GvizCommand = RutaGraphviz + Parametros + FileOutput + FileInput
                Command = exec.Command("cmd", "/C", GvizCommand)
                Command.Stdout = os.Stdout
                AvisoError = Command.Run()
                

                if AvisoError != nil {

                    fmt.Println("Error Al Generar El Reporte")
                    fmt.Println("")

                } else {

                    GvizCommand = FileOutput + " &"
                    Command = exec.Command("cmd", "/C", GvizCommand)
                    Command.Stdout = os.Stdout
                    AvisoError = Command.Run()

                }

            } else {

                GvizCommand = FileOutput + " &"
                Command = exec.Command("cmd", "/C", GvizCommand)
                Command.Stdout = os.Stdout
                AvisoError = Command.Run()

            }

        } else if Variables.SistemaOperativo == "linux"  {

            Command = exec.Command("clear")
            Command.Stdout = os.Stdout
            AvisoError = Command.Run()
            log.Printf("Command finished with error: %v", AvisoError)

        } else {

            print("Sistema Operativo No Soportado")

        }
    }
