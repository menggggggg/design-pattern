package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("ConcreteProduct1", func(t *testing.T) {
		assert.Equal(t, "concreteProduct1", MakeProduct("1").show())
	})

	t.Run("ConcreteProduc21", func(t *testing.T) {
		assert.Equal(t, "concreteProduct2", MakeProduct("2").show())
	})
}
