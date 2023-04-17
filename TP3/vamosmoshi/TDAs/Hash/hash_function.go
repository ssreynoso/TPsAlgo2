package diccionarioHASH

import (
	// https://github.com/spaolacci/murmur3
	"fmt"
	"math/rand"
	"time"
	"vamosmoshi/TDAs/Hash/murmur3"
)

func InicializarHashFunction[K comparable](largo int) func(K) uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var (
		seed   uint32 = 0
		random int32
	)

	for {
		random = r.Int31()
		if random > 0 {
			seed = uint32(random)
			break
		}
	}

	h := murmur3.New64WithSeed(seed)

	return func(key K) uint64 {
		h.Reset()
		keyBytes := convertirABytes(key)
		h.Write(keyBytes)
		return h.Sum64() % uint64(largo)
	}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
