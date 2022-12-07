package utils

import (
	"bufio"
	"fmt"
	"os"
)

func LeerArchivo(ruta string, cb func(string) bool) {
	var (
		flgContinuo bool = true
		lineaLeida  string
	)

	archivo, err := os.Open(ruta)

	if err != nil {
		return
	}

	defer archivo.Close()

	s := bufio.NewScanner(archivo)

	for s.Scan() && flgContinuo {
		lineaLeida = s.Text()
		flgContinuo = cb(lineaLeida)
	}

	err = s.Err()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
	}
}
