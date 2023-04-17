package operaciones

import (
	"algogram/customTDAs"
	"algogram/errores"
	"fmt"
	"os"
	"strings"
)

func Publicar(
	input []string,
	usuariosRegistrados customTDAs.Usuarios,
	usuario_sesion *customTDAs.Usuario,
	postsPublicados customTDAs.Posts,
) {
	if *usuario_sesion == nil {
		error := new(errores.ErrorNoHayLoggeado)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	usuario_actual := (*usuario_sesion).GetID()
	contenido := strings.Join(input, " ")
	nuevoPost := postsPublicados.Agregar(usuariosRegistrados.GetUsuario(usuario_actual), contenido)

	usuariosRegistrados.NuevoPost(usuario_actual, nuevoPost)

	fmt.Fprintf(os.Stdout, "Post publicado\n")
}
