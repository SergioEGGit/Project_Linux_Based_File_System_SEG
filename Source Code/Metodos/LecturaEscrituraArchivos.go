
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

    package Metodos

    import (
        "../Variables"
        "bufio"
        "fmt"
        "github.com/gookit/color"
        "os"
        "strings"
    )

//------------------------------------------------------Métodos---------------------------------------------------------

    func LeerArchivoEntrada(Ruta string) (bool, []string) {

        //Variables
        var Archivo *os.File
        var Scan *bufio.Scanner
        var AvisoError error
        var Extension []string
        var LineasArchivo []string
        var Cadena string

        //Asignación
        Ruta = Trim(Ruta)
        Extension = make([]string, 0)
        LineasArchivo = make([]string, 0)
        Extension = SplitArchivo(Ruta)

        if Extension[1] == "mia" {

            Archivo, AvisoError = os.Open(Ruta)

            //Catch Error

            if AvisoError != nil {

                color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
                fmt.Println("")
                return false, nil

            }

            Scan = bufio.NewScanner(Archivo)

            //Leer Linea Por Linea

            for Scan.Scan() {

                Cadena += Trim(Scan.Text()) + "\n"

            }

            LineasArchivo = SplitContenidoArchivo(Cadena)

            // Catch Error

            if AvisoError = Scan.Err(); AvisoError != nil {

                color.HEX("#de4843", false).Println("Error Al Abrir El Archivo")
                fmt.Println("")
                return false, nil

            }

            _ = Archivo.Close()

        } else {

            color.HEX("#de4843", false).Println("La Extension Del Archivo No Es La Correcta")
            color.HEX("#de4843", false).Println("Extensión Válida: mia")
            fmt.Println("")
            return false, nil
        }

        return true, LineasArchivo
    }

    func LimpiarArreglo() {

        for Contador := 0; Contador < len(Variables.ArregloArchivo); Contador++ {

            Variables.ArregloArchivo[Contador] = ""

        }

    }

    func RecuperarLDComando(ArregloAuxiliar []string) {

        //Variables
        var Final bool
        var ContadorAuxiliar int

        //Asignación
        Final = false
        ContadorAuxiliar = -1

        //Limpiar Arreglo
        LimpiarArreglo()

        //Comienza Recuperación
        for Con := 0; Con < len(ArregloAuxiliar); Con++ {

            if Trim(ArregloAuxiliar[Con]) != "" {

                if !Final {

                    ContadorAuxiliar++
                    Variables.ArregloArchivo[ContadorAuxiliar] = Trim(ArregloAuxiliar[Con])

                } else {

                    if BuscarPrefijo(Trim(ArregloAuxiliar[Con])) {

                        Variables.ArregloArchivo[ContadorAuxiliar] = Variables.ArregloArchivo[ContadorAuxiliar] + " " + Trim(ArregloAuxiliar[Con])

                    } else {

                        Variables.ArregloArchivo[ContadorAuxiliar] = Variables.ArregloArchivo[ContadorAuxiliar] + Trim(ArregloAuxiliar[Con])

                    }
                }

                if BuscarSeparador(Trim(ArregloAuxiliar[Con])) {

                    Variables.ArregloArchivo[ContadorAuxiliar] = Trim(strings.Replace(Trim(Variables.ArregloArchivo[ContadorAuxiliar]), "\\*", "", 1))
                    Final = true

                } else {

                    Final = false

                }

            }

        }

    }



