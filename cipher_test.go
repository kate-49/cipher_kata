package cipher

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_E2E_Simple_Inputs(t *testing.T) {
	solver, _ := CreateSolver("meetmebythetree", "scones")
	assert.Equal(t, "sconessconessco", strings.Join(solver.Keyword, ""))
	assert.Equal(t, "egsgqwtahuiljgs", solver.Encode())
}

func Test_E2E_Simple_Inputs_2(t *testing.T) {
	solver, _ := CreateSolver("meetmeontuesdayeveningatseven", "vigilance")
	assert.Equal(t, "vigilancevigilancevigilancevi", strings.Join(solver.Keyword, ""))
	assert.Equal(t, "hmkbxebpxpmyllyrxiiqtoltfgzzv", solver.Encode())
}
