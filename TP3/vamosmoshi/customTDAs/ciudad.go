package customTDAs

type Ciudad interface {
	getNombre() string
	getCoordenadas() (string, string)
}

type ciudad struct {
	nombre   string
	latitud  string
	longitud string
}

func CrearCiudad(nombre, latitud, longitud string) Ciudad {
	c := new(ciudad)
	c.nombre = nombre
	c.latitud = latitud
	c.longitud = longitud
	return c
}

func (c ciudad) getNombre() string {
	return c.nombre
}

func (c ciudad) getCoordenadas() (string, string) {
	return c.latitud, c.longitud
}
