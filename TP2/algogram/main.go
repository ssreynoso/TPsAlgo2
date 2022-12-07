package main

import (
	"algogram/customTDAs"
	"algogram/errores"
	"algogram/operaciones"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	params := os.Args[1:] // No me importa el nombre del programa

	if len(params) < 1 {
		errorParamentos := new(errores.ErrorParametros)
		fmt.Fprintf(os.Stdout, "%s\n", errorParamentos.Error())
		return
	}

	// Incializo usuarios
	rutaListaUsuarios := params[0] // Ruta del archivo donde estarán los usuarios
	usuariosRegistrados := customTDAs.CrearDiccionarioUsuarios(rutaListaUsuarios)

	// Inicializo posts
	postsPublicados := customTDAs.CrearDiccionarioPosts()

	// Usuario de la sesión
	var usuario_sesion customTDAs.Usuario = nil

	// Leo input del usuario
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		menu(s.Text(), usuariosRegistrados, &usuario_sesion, postsPublicados)
	}

	err := s.Err()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
	}
}

func menu(
	input string,
	usuariosRegistrados customTDAs.Usuarios,
	usuario_sesion *customTDAs.Usuario,
	postsPublicados customTDAs.Posts,
) {
	params := strings.Split(input, " ") // Comando a evaluar
	cmd := params[0]
	data := params[1:]

	switch cmd {
	case "login":
		operaciones.Login(data, usuariosRegistrados, usuario_sesion)
	case "logout":
		operaciones.Logout(usuario_sesion)
	case "publicar":
		operaciones.Publicar(data, usuariosRegistrados, usuario_sesion, postsPublicados)
	case "ver_siguiente_feed":
		operaciones.VerSiguienteFeed(usuario_sesion)
	case "likear_post":
		operaciones.LikearPost(data, usuario_sesion, postsPublicados)
	case "mostrar_likes":
		operaciones.MostrarLikes(data, postsPublicados)
	default:
		fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("ERROR: [%s] el parámetro ingresado no es válido.", cmd))
	}
}
