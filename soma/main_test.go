package main

import "testing"

func TestSoma(t *testing.T) {
	casos := []struct {
		a, b     int
		esperado int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{0, 5, 5},
	}

	for _, caso := range casos {
		resultado := Soma(caso.a, caso.b)
		if resultado != caso.esperado {
			t.Errorf("Esperado %v, mas obteve %v", caso.esperado, resultado)
		}
	}
}
