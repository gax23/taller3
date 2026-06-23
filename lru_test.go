package main

import "testing"

func TestLRU(t *testing.T) {
	lru := NuevoLRU(2) // Capacidad máxima de 2

	lru.Put(1, 10)
	lru.Put(2, 20)

	// Accedemos al elemento 1 para que el elemento 2 quede como el menos usado
	lru.Get(1)

	// Insertamos un 3er elemento. Como la capacidad es 2, debe expulsar al 2
	lru.Put(3, 30)

	if _, ok := lru.Get(2); ok {
		t.Errorf("El elemento 2 debió ser expulsado de la caché por política LRU")
	}

	if val, ok := lru.Get(1); !ok || val != 10 {
		t.Errorf("El elemento 1 fue usado recientemente, debería seguir existiendo")
	}
}
