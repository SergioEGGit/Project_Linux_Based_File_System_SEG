
//-----------------------------------------------Paquetes E Imports-----------------------------------------------------

	package Variables

//----------------------------------------------------Estructuras-------------------------------------------------------

    //Master Boot Record
	type MBREstructura struct {

		SizeMbr int64
        FCreacionMBR [19]byte
		IDMBR int64
		Particion1MBR ParticionEstructura
		Particion2MBR ParticionEstructura
		Particion3MBR ParticionEstructura
		Particion4MBR ParticionEstructura

    }

    //Información De Cada Partición
    type ParticionEstructura struct {

    	StatusPart byte
    	TipoPart byte
    	FitPart byte
    	InicioPart int64
    	SizePart int64
    	NamePart [16]byte

	}

	//Extended Boot Record
	type EBREstructura struct {

		StatusEBR byte
		FitEBR byte
		InicioEBR int64
		SizeEBR int64
		SiguienteEBR int64
		NameEBR [16]byte

	}

	//Discos Montados
	type MountEstructura struct {

		IdentificadorMount string
		EBRMount EBREstructura
		ParticionMount  ParticionEstructura
		NombreMount string
		RutaDiscoMount string
		TipoMount string

	}

	//Identificador Discos
	type IDEstructura struct {

		LetraID string
		ParticionID int

	}

	//Super Bloque
	type SuperBloqueEstructura struct {

		NombreDiscoSuperBloque [16]byte
		ArbolCountSuperBloque int64
		DetalleDirectorioCountSuperBloque int64
		InodosCountSuperBloque int64
		BloquesCountSuperBloque int64
		ArbolFreeSuperBloque int64
		DetalleFreeSuperBloque int64
		InodosFreeSuperBloque int64
		BloquesFreeSuperBloque int64
		FechaCreacionSuperBloque [19]byte
		FechaUltimoMontajeSuperBloque [19]byte
		MontajesSuperBloque int64
		PBitmapArbolSuperBloque int64
		PArbolSuperBloque int64
		PBitmapDetalleSuperBloque int64
		PDetalleSuperBloque int64
		PBitmapTablaSuperBloque int64
		PTablaSuperBloque int64
		PBitmapBloquesSuperBloque int64
		PBloquesSuperBloque int64
		PLogSuperBloque int64
		ArbolSizeSuperBloque int64
		DetalleSizeSuperBloque int64
		InodoSizeSuperBloque int64
		BloquesSizeSuperBloque int64
		ArbolFreeBitSuperBloque int64
		DetalleFreeBitSuperBloque int64
		TablaFreeBitSuperBloque int64
		BloquesFreeBitSuperBloque int64
		MagicNumSuperBloque int64

	}
