package mathservice

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/math", func(writer http.ResponseWriter, request *http.Request) {
		amount, _ := strconv.Atoi(request.FormValue("amount"))
		json.NewEncoder(writer).Encode(generateQuestions(amount))
	})
	http.ListenAndServe(":6000", nil)
}

func generateQuestions(amount int) Questions {
	questions := make(Questions, amount);
	for i := 0; i < amount; i++ {
		b := rand.Intn(10);
		a := rand.Intn(10);
		singleQuestion := Question{fmt.Sprintf("%d+%d=?", a, b), fmt.Sprintf("%d", a+b)}
		questions = append(questions[:i], singleQuestion);
	}

	return questions;
}

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Questions []Question