
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

    func GenerarArchivoTxt(NombreReporte, CadenaArchivo string) {

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
            RutaArchivo = "C:\\GraficasMIA\\" + NombreReporte + ".txt"
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

            Writer.Flush()

        } else {

            print("Sistema Operativo No Soportado")

        }
    }

    func GenerarReporte(NombreReporte string) {

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
            FileInput = "C:\\GraficasMIA\\" + NombreReporte + ".txt"
            FileOutput = "C:\\GraficasMIA\\" + NombreReporte + ".png "

            GvizCommand = RutaGraphviz + Parametros + FileOutput + FileInput

            Command = exec.Command("cmd", "/c", GvizCommand)
            Command.Stdout = os.Stdout
			Command.Stderr = os.Stderr
            AvisoError = Command.Run()

            //Catch Error
            if AvisoError != nil {

				RutaGraphviz = "dot "
                GvizCommand = RutaGraphviz + Parametros + FileOutput + FileInput
                Command = exec.Command("cmd", "/C", GvizCommand)
                Command.Stdout = os.Stdout
                AvisoError = Command.Run()
                

                if AvisoError != nil {

                    fmt.Println("Error Al Generar El Reporte")
                    fmt.Scanln()

                } else {

                    fmt.Print("Abro La Imagen")

                }

            } else {

                print("Abro La Imagen")

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
