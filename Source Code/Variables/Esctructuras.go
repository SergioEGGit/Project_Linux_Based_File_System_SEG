
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
	
	//Arbol Virtual De Directorio 
	type AVDEstructura struct {

		FechaCreacionAVD [19]byte
		NombreDirectorioAVD [16]byte
		PArraySubDirectoriosAVD [6]int64
		PDetalleDirectorioAVD int64
		PArbolVirtualDirectorio int64
		PropietarioAVD [16]byte
		GrupoAVD [16]byte
		PermisosAVD int64

	}
	
	//Informacion Detalle De Directorio
	type DDInformacionEstructura struct {

		NombreArchivoDDInformacion [16]byte
		PInodoArchivoDDInformacion int64
		FechaCreacionArchivoDDInformacion [19]byte
		FechaModificacionArchivoDDInformacion [19]byte

	}

	//Detalle De Directorio
	type DDEstructura struct {

		ArrayArchivosDD [5]DDInformacionEstructura
		PDetalleDirectorioDD int64
	}

	//Tabla Inodos
	type TablaInodoEstructura struct {

		NumeroInodoTI int64
		SizeArchivoTI int64
		NumeroBloquesTI int64
		ArrayBloquesTI [4]int64
		PTabalInodosTI int64
		PropietarioTI [16]byte
		GrupoTI [16]byte
		PermisosTI int64

	}

	//Bloques
	type BloquesEstructura struct {

		InformacionBQ [25]byte

	}

	//Bitacora
	type BitacoraEstructura struct {

		TipoOperacionBT [16]byte
		TipoArchivoDirectorioBT int64
		NombreArchivoDirectorioBT [16]byte
		ContenidoBT [100]byte
		FechaTransaccionBT [19]byte

	}

