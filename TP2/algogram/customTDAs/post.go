package customTDAs

import (
	ABB "algogram/TDAs/ABB"
	"fmt"
	"os"
	"strings"
	"time"
)

type Post interface {
	GuardarLike(string)
	MostrarLikes()
	CantidadLikes() int
	GetUsuario() Usuario
	GetFecha() time.Time
	Informacion()
}

type post struct {
	id                int
	usuario           Usuario
	contenido         string
	fecha_publicacion time.Time
	likes             ABB.DiccionarioOrdenado[string, string] // usuario como clave y valor.
}

func CrearPost(id int, usuario Usuario, contenido string) Post {
	post := new(post)
	post.id = id
	post.usuario = usuario
	post.contenido = contenido
	post.likes = ABB.CrearABB[string, string](strings.Compare)
	post.fecha_publicacion = time.Now()
	return post
}

func (p post) GetUsuario() Usuario {
	return p.usuario
}

func (p post) GetFecha() time.Time {
	return p.fecha_publicacion
}

func (p *post) GuardarLike(usuario string) {
	if p.likes.Pertenece(usuario) {
		return
	}
	p.likes.Guardar(usuario, usuario)
}

func (p post) MostrarLikes() {
	fmt.Fprintf(os.Stdout, "El post tiene %d likes:\n", p.likes.Cantidad())
	mostrarLikes := func(usuario, _ string) bool {
		fmt.Fprintf(os.Stdout, "\t%s\n", usuario)
		return true
	}

	p.likes.Iterar(mostrarLikes)
}

func (p post) CantidadLikes() int {
	return p.likes.Cantidad()
}

func (p post) Informacion() {
	fmt.Fprintf(os.Stdout, "Post ID %d\n", p.id)
	fmt.Fprintf(os.Stdout, "%s dijo: %s\n", p.usuario.GetID(), p.contenido)
	fmt.Fprintf(os.Stdout, "Likes: %d\n", p.likes.Cantidad())
}
