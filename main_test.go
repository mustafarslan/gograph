package main

import "testing"

func TestScenario1(t *testing.T) {
	t.Log("Sending ID 1 ...")
	result := totalExecutionTime(1, jobs)
	if result != 120 {
		t.Errorf("Expected total execution times is 120, but it was %d instead", result)
	} else {
		t.Log("Successful")
	}
}

func TestScenario2(t *testing.T) {
	t.Log("Sending ID 2 ...")
	result := totalExecutionTime(2, jobs)
	if result != 30 {
		t.Errorf("Expected total execution times is 30, but it was %d instead", result)
	} else {
		t.Log("Successful")
	}
}

func TestScenario3(t *testing.T) {
	t.Log("Sending ID 3 ...")
	result := totalExecutionTime(3, jobs)
	if result != 20 {
		t.Errorf("Expected total execution times is 20, but it was %d instead", result)
	} else {
		t.Log("Successful")
	}
}

func TestScenario4(t *testing.T) {
	t.Log("Sending ID 5 ...")
	result := totalExecutionTime(5, jobs)
	if result != 200 {
		t.Errorf("Expected total execution times is 200, but it was %d instead", result)
	} else {
		t.Log("Successful")
	}
}

func TestScenario5(t *testing.T) {
	t.Log("Sending ID 8 ...")
	result := totalExecutionTime(8, jobs)
	if result != 90 {
		t.Errorf("Expected total execution times is 90, but it was %d instead", result)
	} else {
		t.Log("Successful")
	}
}

func TestScenario6(t *testing.T) {
	t.Log("Sending ID 7 ...")
	result := totalExecutionTime(7, jobs)
	if result != 210 {
		t.Errorf("Expected total execution times is 210, but it was %d instead", result)
	} else {
		t.Log("Successful")
	}
}

func TestScenario7(t *testing.T) {
	t.Log("Sending ID 9 ...")
	result := totalExecutionTime(9, jobs)
	if result != 160 {
		t.Errorf("Expected total execution times is 160, but it was %d instead", result)
	} else {
		t.Log("Successful")
	}
}
