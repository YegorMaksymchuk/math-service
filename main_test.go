package mathservice

import "testing"

func TestGenerateQuestions(t *testing.T)  {
	expectedResult := 2
	actualResult:= len(generateQuestions(2))
	if expectedResult != actualResult{
		t.Error("Exception actual result is 2 but got a ",actualResult)
	}

}