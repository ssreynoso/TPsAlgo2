package utils

func BusquedaBinaria(arreglo []int, target int, ini int, fin int, tipoError error) (int, error) {
	if fin < ini || len(arreglo) == 0 {
		err := tipoError
		return -1, err
	}
	mid := int(ini + (fin-ini)/2)
	if arreglo[mid] > target {
		return BusquedaBinaria(arreglo, target, ini, mid-1, tipoError)
	} else if arreglo[mid] < target {
		return BusquedaBinaria(arreglo, target, mid+1, fin, tipoError)
	} else {
		return mid, nil
	}
}
