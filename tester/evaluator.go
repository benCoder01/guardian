package tester

import "log"

type Evaluator struct {
	Conditions []Condition
}

func (e Evaluator) Run() {
	log.Printf("Beginne Prüfung der Conditions. Insgesamt %d", len(e.Conditions))
	for _, condition := range e.Conditions {
		log.Printf("Start of Testing for Condition %s", condition.GetName())
		res := condition.Evaluate()
		if !res {
			log.Printf("Condition %s failed to execute.", condition.GetName())
		}else {
			log.Printf("Condition %s was tested successfully", condition.GetName())
		}

		log.Printf("End of Testing for Condition %s", condition.GetName())
	}

}