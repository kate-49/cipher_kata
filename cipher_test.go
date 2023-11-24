package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Keyword(t *testing.T) {
	
}

func Test_Encode(t *testing.T) {
	solver, _ := CreateSolver()
	result := solver.Encode("meetmebythetree", "scones")
	assert.Equal(t, "egsgqwtahuiljgs", result)
}

func Test_Encode_2(t *testing.T) {
	solver, _ := CreateSolver()
	result := solver.Encode("meetmeontuesdayeveningatseven", "vigilance")
	assert.Equal(t, "hmkbxebpxpmyllyrxiiqtoltfgzzv", result)
}

//func Test_Decode(t *testing.T) {
//	solver, _ := CreateSolver("hmkbxebpxpmyllyrxiiqtoltfgzzv", "vigilance")
//	assert.Equal(t, "vigilancevigilancevigilancevi", strings.Join(solver.Keyword, ""))
//	assert.Equal(t, "hmkbxebpxpmyllyrxiiqtoltfgzzv", solver.Encode())
//}
