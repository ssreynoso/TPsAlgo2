package customTDAs

type ListaPadrones interface {

	// VerificarDNI realiza una búsqueda en la lista de padrones para ver si el dni se
	// encuentra en ella.
	VerificarDNI(int) bool
}
