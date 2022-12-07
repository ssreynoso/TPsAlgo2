package operaciones

import (
	"algogram/customTDAs"
	"algogram/errores"
	"fmt"
	"os"
)

func MostrarLikes(
	input []string,
	postsPublicados customTDAs.Posts,
) {
	postID := input[0]

	if !postsPublicados.Pertenece(postID) || postsPublicados.GetCantidadLikes(postID) == 0 {
		error := new(errores.ErrorMostrarLikes)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	postsPublicados.MostrarLikes(postID)
}
