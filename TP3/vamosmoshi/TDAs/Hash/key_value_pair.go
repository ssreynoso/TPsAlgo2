package diccionarioHASH

type KeyValuePair[K comparable, V any] struct {
	key   K
	value V
}

func CrearKeyValuePair[K comparable, V any](clave K, dato V) KeyValuePair[K, V] {
	pair := new(KeyValuePair[K, V])
	pair.key = clave
	pair.value = dato
	return *pair
}
