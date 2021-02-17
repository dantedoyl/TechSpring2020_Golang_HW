package calculate

import (
	"testing"
	)

func TestCalculate(t *testing.T) {
	testStr := "(5*(((9+8)*(4*6))+7))/25-3"

	testRes, err := Calculate(testStr)
	expected := 80

	if err != nil{
		t.Error(err.Error())
	}

	if testRes != expected{
		t.Error("Wrong value in result.")
	}
}

func TestCalculate2(t *testing.T) {
	testStr := "5*(((9+8)*(4*6))+7))))"

	_, err := Calculate(testStr)

	if err == nil{
		t.Error("Error was expexted")
	}

}

func TestCalculate3(t *testing.T) {
	testStr := "3+*7"

	_, err := Calculate(testStr)

	if err == nil{
		t.Error("Error was expexted")
	}
}

func TestCalculate4(t *testing.T) {
	testStr := "(3/7"

	_, err := Calculate(testStr)

	if err == nil{
		t.Error("Error was expexted")
	}
}

func TestCalculate5(t *testing.T) {
	testStr := "dadsd"

	_, err := Calculate(testStr)

	if err == nil{
		t.Error("Error was expexted")
	}
}

func TestCalculate6(t *testing.T) {
	testStr := "+++++"

	_, err := Calculate(testStr)

	if err == nil{
		t.Error("Error was expexted")
	}
}

func TestCalculate7(t *testing.T) {
	testStr := "5^5"

	_, err := Calculate(testStr)
	//expected := 0

	if err == nil{
		t.Error("Error was expexted")
	}

	//if testRes != expected{
	//	t.Error("Wrong value in result.")
	//}
}