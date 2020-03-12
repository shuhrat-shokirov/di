package main

import "testing"

func TestFunc_NoParameter(t *testing.T)  {
	err := callMe(sample)
	if err != nil {
		t.Fatalf("error must be nil: %v", err)
	}
}

func TestFunc_NoParameterError(t *testing.T)  {
	err := callMe(5)
	if err == nil {
		t.Fatalf("error must not be nil: %v", err)
	}
}

func TestFunc_WithParameter(t *testing.T)  {
	err := callMeAdvanced(sample2, "args")
	if err != nil {
		t.Fatalf("error must be nil: %v", err)
	}
}

func TestFunc_WithParameterError(t *testing.T)  {
	err := callMeAdvanced(false)
	if err == nil {
		t.Fatalf("error must not be nil: %v", err)
	}
}

func TestFunc_WithOutputParameter(t *testing.T)  {
	param, err := callMeAdvancedWithResult(sample3, "args")
	if err != nil {
		t.Fatalf("error must be nil: %v", err)
	}

	if param != "args" {
		t.Fatalf("param just be \"args\": %s", param)
	}
}

func TestFunc_WithOutputParameterError(t *testing.T)  {
	_, err := callMeAdvancedWithResult("sample3", "args")
	if err == nil {
		t.Fatalf("error must not be nil: %v", err)
	}
}