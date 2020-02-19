package integrations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTraefikIntegration(t *testing.T) {

	t.SkipNow()

	t.Run("List Services from Traefik", func(t *testing.T) {
		services, err := TraefikIntegrator.List()
		require.Nil(t, err)
		assert.NotEqual(t, 0, len(services))
	})

	t.Run("Confirm Services from Traefik", func(t *testing.T) {
		services, err := TraefikIntegrator.List()
		require.Nil(t, err)
		for _, s := range services {
			t.Log(s)
		}
	})

}
