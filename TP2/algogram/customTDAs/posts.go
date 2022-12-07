package customTDAs

import (
	Hash "algogram/TDAs/Hash"
	"fmt"
)

type Posts interface {
	Agregar(Usuario, string) Post
	Pertenece(string) bool
	GuardarLike(string, string)
	GetCantidadLikes(string) int
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
