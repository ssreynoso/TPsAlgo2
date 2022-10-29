package customTDAs

type ListaDNIs interface {
	InsertarDNI(int)

	PadronFraudulento(int) bool
}
