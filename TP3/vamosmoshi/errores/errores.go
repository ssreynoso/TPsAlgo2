package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

// Errores login

type ErrorNoSeEncontroRecorrido struct{}

func (e ErrorNoSeEncontroRecorrido) Error() string {
	return "No se encontro recorrido"
}
