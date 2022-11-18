package operaciones

import (
	"fmt"
	"os"
	"rerepolez/cola"
	"rerepolez/customTDAs"
	errores "rerepolez/errores"
	"rerepolez/votos"
	"strconv"
)

func Ingresar(
	data []string,
	listaPadrones customTDAs.ListaPadrones,
	colaVotantes cola.Cola[votos.Votante],
) {
	if len(data) < 1 {
		err := new(errores.ErrorParametros)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	dniString := data[0]
	dni, err := strconv.Atoi(dniString)
	if err != nil || dni <= 0 {
		err := new(errores.DNIError)
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	if !listaPadrones.VerificarDNI(dni) {
		return
	}

	nuevoVotante := votos.CrearVotante(dni)
	colaVotantes.Encolar(nuevoVotante)
	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}
