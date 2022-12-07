package operaciones

import (
	"algogram/customTDAs"
	"algogram/errores"
	"fmt"
	"os"
)

func LikearPost(
	input []string,
	usuario_sesion *customTDAs.Usuario,
	postsPublicados customTDAs.Posts,
) {
	postID := input[0]

	if *usuario_sesion == nil || !postsPublicados.Pertenece(postID) {
		error := new(errores.ErrorLikearPost)
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}

	usuario := (*usuario_sesion).GetID()
	postsPublicados.GuardarLike(usuario, postID)

	fmt.Fprintf(os.Stdout, "Post likeado\n")
}
