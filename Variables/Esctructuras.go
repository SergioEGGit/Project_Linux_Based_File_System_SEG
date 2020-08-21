
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