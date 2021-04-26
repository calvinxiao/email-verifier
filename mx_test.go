package emailverifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMxOK(t *testing.T) {
	domain := "github.com"

	mx, err := verifier.CheckMX(domain)
	assert.NoError(t, err)
	assert.True(t, mx.HasMXRecord)
}

func TestCheckNoMxOK(t *testing.T) {
	domain := "githubexists.com"

	_, err := verifier.CheckMX(domain)
	assert.Error(t, err, ErrNoSuchHost)
}

func BenchmarkCheckMxOK(b *testing.B) {
	domain := "github.com"

	for i := 0; i < b.N; i++ {
		mx, err := verifier.CheckMX(domain)
		assert.NoError(b, err)
		assert.True(b, mx.HasMXRecord)
	}
}
