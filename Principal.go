
//------------------------------------------------Paquetes E Imports----------------------------------------------------

    package main

    import (
        "./AnalisisYComandos"
        "./Metodos"
        "./Variables"
        "bufio"
        "fmt"
        "github.com/gookit/color"
        "os"
    )

//------------------------------------------------------Métodos---------------------------------------------------------

    func main() {

        //Variables
        var Lectura *bufio.Reader
        var Cadena string
        var AvisoError error

        //Asignaciones
        Lectura = bufio.NewReader(os.Stdin)
        Variables.Salir = false

        //Obtiene El Nombre Del Sistema Operativo
        Metodos.ObtenerNombreOs()
        
        //Presentación
        color.HEX("#1fbfad", false).Println("                 <Bienvenido Al Sistema SEG>")
        color.HEX("#2fa4ed", false).Println("                        <201801628>")
        color.HEX("#2fa4ed", false).Println("                     <Sergio Echigoyen>")
        fmt.Println("")
        fmt.Println("")

        //Flujo Principal Aplicación

        for {

            color.HEX("#5260de", false).Print("<Seg: Ingrese Un Comando>")

            Cadena, AvisoError = Lectura.ReadString('\n')
            _ = AvisoError

            AnalisisYComandos.AnalisisComando(Cadena)

            if Variables.Salir {

                break

            }
        }

        AnalisisYComandos.ComandoCls()
        color.HEX("#9152de", false).Println("\n", "Gracias Por Utilizar Nuestro Sistema Vuelva Pronto...!")
        _, _ = fmt.Scanln()

    }



















