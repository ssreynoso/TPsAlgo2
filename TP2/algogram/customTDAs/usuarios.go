package customTDAs

import (
	Hash "algogram/TDAs/Hash"
	"algogram/utils"
)

type Usuarios interface {

	// Pertenece devuelve true si el usuario existe; en caso contrario, false.
	Pertenece(string) bool

	// NuevoPost recibe el id del usuario autor del post (string) y
	// guarda este nuevo post en el feed de cada uno de los usuarios en el sistema.
	NuevoPost(string, Post)

	// Devuelve el TDA Usuario de un usuario dado de alta.
	GetUsuario(string) Usuario
}

type diccionarioUsuarios struct {
	usuarios Hash.Diccionario[string, Usuario]
}

func CrearDiccionarioUsuarios(rutaListaUsuarios string) Usuarios {
	d := new(diccionarioUsuarios)
	d.usuarios = Hash.CrearHash[string, Usuario]()
	i := 0

	agregarUsuarios := func(u string) bool {
		nuevoUsuario := CrearUsuario(u, i)
		d.usuarios.Guardar(u, nuevoUsuario)
		i++
		return true
	}

	utils.LeerArchivo(rutaListaUsuarios, agregarUsuarios)

	return d
}

func (d diccionarioUsuarios) Pertenece(id string) bool {
	return d.usuarios.Pertenece(id)
}

func (d diccionarioUsuarios) NuevoPost(autor string, post Post) {
	guardarPostEnCadaUsuario := func(idUsuario string, usuario Usuario) bool {
		if idUsuario != autor {
			usuario.AgregarPostAlFeed(&post)
		}
		return true
	}

	d.usuarios.Iterar(guardarPostEnCadaUsuario)
}

func (d diccionarioUsuarios) GetUsuario(id string) Usuario {
	return d.usuarios.Obtener(id)
}
