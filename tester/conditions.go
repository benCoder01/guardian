package tester

import (
	"log"
	"net/http"
)

type Condition interface {
	Evaluate() bool
	GetName() string
}

// SuccessCondition 
type SuccessCondition struct {
	Url string
	Name string
}

func (c SuccessCondition) Evaluate() bool {
	log.Println("Starte Evaluation")
	resp, err := http.Get(c.Url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (c SuccessCondition) GetName() string {
	return c.Name
}