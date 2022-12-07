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

type ErrorYaHayLoggeado struct{}

func (e ErrorYaHayLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

// Errores logout Y publicar

type ErrorNoHayLoggeado struct{}

func (e ErrorNoHayLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

// Errores ver_siguiente_feed

type ErrorVerSigFeed struct{}

func (e ErrorVerSigFeed) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

// Errores likear_post

type ErrorLikearPost struct{}

func (e ErrorLikearPost) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

// Errores mostrar_likes

type ErrorMostrarLikes struct{}

func (e ErrorMostrarLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}
