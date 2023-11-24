package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_E2E_Simple_Inputs(t *testing.T) {
	solver, _ := CreateSolver("meetmebythetree", "sconessconessco")
	assert.Equal(t, solver.Encode(), "egsgqwtahuiljgs")
}
