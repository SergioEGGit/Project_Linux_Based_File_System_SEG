# Project_Linux_Based_File_System_SEG

## Language: GO

File System Based On Linux SEG: This system seeks to simulate a file system based on linux, in addition to having a virtual hard disk manager, 
the system is a console application that works through commands

Command List (Sintaxis: Command_Name -ParmName1->ParamValue1 -Parm2->ParamaValue2 -ParamN->ParamaValueN)

| Command | Param (PathExample /home/Example/File.mia)                                                             | Action                                   |  
|---------|--------------------------------------------------------------------------------------------------------|------------------------------------------|
| exec    | -path->FilePath                                                                                        | Load a file with commands                | 
| pause   | -                                                                                                      | Pause between commands                   |   
| Mkdisk  | -size->number -path->DestinationPath -name->name only accept numbers,letters and _ -unit->Kb or MB     | Create a virtual disk                    |
| Rmdisk  | -path->LocationPath                                                                                    | Delete a virtual disk                    |
| Fdisk   | -size->number -unit->Kb or MB or B -path->LocationPath -Type->Primary or Extended or Logic             | Create a partition in a virtual disk     |
| Mount   | -path->DestinationPath -name->name only accept numbers,letters and _                                   | Mount a partition                        |
| Unmount | -IDn->ID (Example: -ID1->vda1 -ID2->vdx1 -IDn->vdaN                                                    | Unmount partitions                       |
| MKFS    | -ID->IDMount -Type->Fast or Full                                                                       | Make LWH format to partition             |

To place comments you can place # followed by any text.
