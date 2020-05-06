package slice

func TestIntPositive(t *testing.T) {
	listInt := []int{1,2,3}
	b, _ := json.Marshal(listInt)

	found, err := Find(b, 1)
	if err != nil {
		t.Errorf("Unexpected Error %d", err.Error)
	}
	if found == false {
		t.Errorf("%d expected in %d." , 1, listInt)
	}
}

func TestIntNegative(t *testing.T) {
	listInt := []int{1,2,3}
	b, _ := json.Marshal(listInt)

	found, err := Find(b, 4)
	if err != nil {
		t.Errorf("Unexpected Error %d", err.Error)
	}
	if found {
		t.Errorf("4 not expected in %d.", listInt)
	}

	found, err = Find(b, "a")
	if err != nil {
		t.Errorf("Unexpected Error %d", err.Error)
	}
	if found {
		t.Errorf("a not expected in %d.", listInt)
	}
}

func TestStrPositive(t *testing.T) {
	listString := []string{"a","b","c"}
	b, _ := json.Marshal(listString)


	found, err := Find(b, "a")
	if err != nil {
		t.Errorf("Unexpected Error %d", err.Error)
	}
	if found == false {
		t.Errorf("a expected in %d." , listString)
	}
}

func TestStrNegative(t *testing.T) {
	listString := []string{"a","b","c"}
	b, _ := json.Marshal(listString)


	found, err := Find(b, "d")
	if err != nil {
		t.Errorf("Unexpected Error %d", err.Error)
	}
	if found {
		t.Errorf("d not expected in %d." , listString)
	}

	found, err = Find(b, 1)
	if err != nil {
		t.Errorf("Unexpected Error %d", err.Error)
	}
	if found {
		t.Errorf("1 not expected in %d." , listString)
	}
}
