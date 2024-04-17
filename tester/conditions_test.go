package tester

import "testing"

func TestEvaluate(t *testing.T) {
	t.Run("SuccessCondition", func(t *testing.T) {
		s := SuccessCondition{Url: "https://google.com", Name: "Test"}
		res := s.Evaluate()
		if !res {
			t.Errorf("https://google.com said to be not reachable.")
		}

		name := s.GetName()
		if  name != "Test"{
			t.Errorf("GetName() function wrong expexted Test, but got %s", name)
		}
		
		nonReachableURL := "https://asdasldkjaslkdj.dadsasd"

		s = SuccessCondition{Url: nonReachableURL, Name: ""}

		res = s.Evaluate()
		if res {
			t.Errorf("The url %s should not be reachable, but was evaluated true.", nonReachableURL)
		}

	})

	t.Run("NotFoundCondition", func(t *testing.T) {
		s := NotFoundCondition{Url: "https://google.com/404", Name: "Test"}
		res := s.Evaluate()
		if !res {
			t.Errorf("https://google.com/404 said to be not 404.")
		}

		name := s.GetName()
		if name != "Test" {
			t.Errorf("GetName() function wrong expexted Test, but got %s", name)
		}
		
		non404URL := "https://google.com"

		s = NotFoundCondition{Url: non404URL, Name: "Test"}
		res = s.Evaluate()
		if res {
			t.Errorf("The url %s should be reachable, but was evaluated not to be.", non404URL)
		}

	})
}