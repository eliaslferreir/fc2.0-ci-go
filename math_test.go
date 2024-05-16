//math_test.go
package main

import "testing"

func TestSoma(t *testing.T) {
  total := Soma(15, 15)

  if total != 30 {
    t.Errorf("Resultado da Soma Inválido: Resultado %d. Esperado %d", total, 30)

  }
}