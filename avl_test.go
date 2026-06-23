package main

import "testing"

func TestAVL(t *testing.T) {
	var raiz *Nodo

	// Insertamos datos ordenados secuencialmente.
	// En un árbol normal esto lo deformaría, en un AVL debe activar las rotaciones.
	valores := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	for _, v := range valores {
		raiz = Insertar(raiz, v, Pelicula{ID: int(v), Rating: v})
	}

	// Comprobamos la función exigida por la rúbrica
	if !estaBalanceado(raiz) {
		t.Errorf("El árbol AVL se desbalanceó al recibir datos ordenados")
	}

	// Probamos la búsqueda por rango
	res := []Pelicula{}
	Buscar(raiz, 2.0, 4.0, &res)
	if len(res) != 3 {
		t.Errorf("La búsqueda por rango falló, se esperaban 3 películas, se obtuvieron %d", len(res))
	}
}
