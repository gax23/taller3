package main

import "testing"

func TestColaFIFO(t *testing.T) {
	c := &Cola{}

	if c.Len() != 0 {
		t.Errorf("La cola debería inicializarse vacía")
	}

	c.Enqueue(100)
	c.Enqueue(200)

	if l := c.Len(); l != 2 {
		t.Errorf("El tamaño de la cola debería ser 2, pero es %d", l)
	}

	f, ok := c.Front()
	if !ok || f != 100 {
		t.Errorf("El elemento en el frente debería ser 100")
	}

	v, ok := c.Dequeue()
	if !ok || v != 100 {
		t.Errorf("Dequeue debería haber retornado 100, retornó %d", v)
	}

	if c.Len() != 1 {
		t.Errorf("El tamaño de la cola tras el Dequeue debería ser 1")
	}
}
