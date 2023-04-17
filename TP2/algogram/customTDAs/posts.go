package customTDAs

import (
	Hash "algogram/TDAs/Hash"
	"fmt"
)

type Posts interface {

	// Recibe un usuario autor del nuevo post, el contenido del post,
	// crea un nuevo post, lo agrega al diccionario de posts y lo devuelve.
	Agregar(Usuario, string) Post

	// Recibe el id de un post y devuelve true si el post existe; en caso contrario, false.
	Pertenece(string) bool

	// Recibe el id de un usuario, el id de un post y guarda el like de ese usuario en ese post.
	GuardarLike(string, string)

	// Recibe el id de un post y devuelve la cantidad de likes que este tiene.
	GetCantidadLikes(string) int

	// Recibe el id de un post y muestra sus likes.
	MostrarLikes(string)
}

type diccionarioPosts struct {
	ultID int
	dic   Hash.Diccionario[string, Post]
}

func CrearDiccionarioPosts() Posts {
	d := new(diccionarioPosts)
	d.dic = Hash.CrearHash[string, Post]()
	return d
}

func (d *diccionarioPosts) Agregar(usuario Usuario, contenido string) Post {
	post := CrearPost(d.ultID, usuario, contenido)
	id := fmt.Sprintf("%d", d.ultID)
	d.dic.Guardar(id, post)
	d.ultID++
	return post
}

func (d diccionarioPosts) Pertenece(id string) bool {
	return d.dic.Pertenece(id)
}

func (d diccionarioPosts) GetCantidadLikes(id string) int {
	return d.dic.Obtener(id).CantidadLikes()
}

func (d diccionarioPosts) GuardarLike(usuario, id string) {
	d.dic.Obtener(id).GuardarLike(usuario)
}

func (d diccionarioPosts) MostrarLikes(id string) {
	d.dic.Obtener(id).MostrarLikes()
}
