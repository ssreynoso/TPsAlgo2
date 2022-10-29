package customTDAs

import (
	"fmt"
	"os"
	"rerepolez/errores"
	"rerepolez/utils"
)

type listaPadrones struct {
	padrones []int
}

func CrearListaPadrones(ruta string) (ListaPadrones, error) {
	lista, errorLeerArchivo := utils.ProcesarArchivoPadrones(ruta)
	if errorLeerArchivo != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errorLeerArchivo.Error())
		return nil, errorLeerArchivo
	}
	nuevaListaPadrones := new(listaPadrones)
	nuevaListaPadrones.padrones = lista

	return nuevaListaPadrones, nil
}

func (lista *listaPadrones) VerificarDNI(dni int) bool {
	errARetornar := new(errores.DNIFueraPadron)
	posDNI, err := utils.BusquedaBinaria(lista.padrones, dni, 0, len(lista.padrones), errARetornar)
	if err != nil || posDNI < 0 {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return false
	}
	return true
}
