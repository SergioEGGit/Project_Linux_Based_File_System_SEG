
//------------------------------------Paquetes E Imports-------------------------------------------

    package Metodos

    import (
        "../Variables"
        "bufio"
        "fmt"
        "github.com/gookit/color"
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

                color.HEX("#de4843", false).Println("Error Al Generar El Reporte")
                fmt.Println("")

            }

            //Cerrar Archivo
            defer Archivo.Close()


            BanderaOs = true

        } else if Variables.SistemaOperativo == "linux"  {

            //Abrir Archivo Y Generar Try Catch Error
            RutaArchivo = Directorio + NombreReporte + ".txt"
            Archivo, AvisoError = os.Create(RutaArchivo)

            //Catch Error

            if AvisoError != nil {

                color.HEX("#de4843", false).Println("Error Al Generar El Reporte")
                fmt.Println("")

            }

            //Cerrar Archivo
            defer Archivo.Close()

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

                color.HEX("#de4843", false).Println("Error Al Generar El Reporte")
                fmt.Println("")

            }

            _ = Writer.Flush()

        } else {

            color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
            fmt.Println("")

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

                    color.HEX("#de4843", false).Println("Error Al Generar El Reporte")
                    fmt.Println("")

                } else {

                    color.Success.Println("Reporte Generado Con Exito")
                    fmt.Println("")

                    GvizCommand = FileOutput + " &"
                    Command = exec.Command("cmd", "/C", GvizCommand)
                    Command.Stdout = os.Stdout
                    AvisoError = Command.Run()

                    if AvisoError != nil {

                        color.HEX("#de4843", false).Println("Error Al Abrir La Imagen")
                        fmt.Println("")

                    }

                }

            } else {

                color.Success.Println("Reporte Generado Con Exito")
                fmt.Println("")

                GvizCommand = FileOutput + " &"
                Command = exec.Command("cmd", "/C", GvizCommand)
                Command.Stdout = os.Stdout
                AvisoError = Command.Run()

                if AvisoError != nil {

                    color.HEX("#de4843", false).Println("Error Al Abrir La Imagen")
                    fmt.Println("")

                }

            }

        } else if Variables.SistemaOperativo == "linux"  {

            RutaGraphviz = "dot "
            Parametros = "-Tpng -o "
            FileInput = " " + Directorio + NombreReporte + ".txt"
            FileOutput = Directorio + Archivo + " "

            GvizCommand = RutaGraphviz + Parametros + FileOutput + FileInput

            Command = exec.Command("/bin/bash", "-c", GvizCommand)
            Command.Stdout = os.Stdout
            Command.Stderr = os.Stderr
            AvisoError = Command.Run()

            //Catch Error
            if AvisoError != nil {

                color.HEX("#de4843", false).Println("Error Al Generar El Reporte")
                fmt.Println("")

            } else {

                color.Success.Println("Reporte Generado Con Exito")
                fmt.Println("")

                GvizCommand = FileOutput
                Command = exec.Command("xdg-open", GvizCommand)
                Command.Stdout = os.Stdout
                AvisoError = Command.Run()

                if AvisoError != nil {

                    color.HEX("#de4843", false).Println("Error Al Abrir La Imagen")
                    fmt.Println("")

                }

            }

        } else {

            color.HEX("#de4843", false).Println("Sistema Operativo No Soportado")
            fmt.Println("")

        }
    }
