package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailSum(t *testing.T) {
	kata := "kasur ini rusak"
	result := Sum(kata)

	require.Equal(t, false, result, "Bukan Polindrom")

	fmt.Println("TestSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	kata := "kasur ini rusak"
	result := Sum(kata)

	assert.Equal(t, true, result, "Polindrom")

	fmt.Println("TestSum Eksekusi Terhenti")
}
