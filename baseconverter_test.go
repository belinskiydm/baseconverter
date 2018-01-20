package baseconverter

import (
	"fmt"
	"testing"
)

func TestConversion(t *testing.T) {
	//Correct result testing
	num, err := Conversion("0", 2, 10)
	if num != "0" && err != nil {
		t.Error("Expected 0,", "got ", num)
	}

	num, err = Conversion("56f", 16, 2)
	if num != "10101101111" && err != nil {
		t.Error("Expected 10101101111,", "got ", num)
	}

	num, err = Conversion("-58586655672", 10, 62)
	if num != "-11WTDte" && err != nil {
		t.Error("Expected -11WTDte,", "got ", num)
	}

	num, err = Conversion("11WTDte", 62, 2)
	if num != "110110100100000010010110101110111000" && err != nil {
		t.Error("Expected 110110100100000010010110101110111000,", "got ", num)
	}

	//Errors testing
	_, err = Conversion("3334%", 10, 16)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = Conversion("", 10, 16)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = Conversion("", 1, 16)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = Conversion("", 10, 0)
	if err == nil {
		t.Error("Expecting an error")
	}

}

func TestConvertToDec(t *testing.T) {
	//Correct result testing
	num, err := ConvertToDec("453fa", 16)
	if num != "283642" && err != nil {
		t.Error("Expected 283642,", "got ", num)
	}
	num, err = ConvertToDec("-4t3fa", 55)
	if num != "-41437285" && err != nil {
		t.Error("Expected -41437285,", "got ", num)
	}

	//Errors testing
	_, err = ConvertToDec("453$fa", 16)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = ConvertToDec("", 16)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = ConvertToDec("", -1)
	if err == nil {
		t.Error("Expecting an error")
	}
}

func TestConvertFromDec(t *testing.T) {
	//Correct result testing
	num, err := ConvertFromDec("-6455537849492", 62)
	if num != "-1PEvwVPC" && err != nil {
		t.Error("Expected -1PEvwVPC,", "got ", num)
	}
	num, err = ConvertFromDec("57483", 2)
	if num != "1110000010001011" && err != nil {
		t.Error("Expected 1110000010001011,", "got ", num)
	}

	//Errors testing
	_, err = ConvertFromDec("453$fa", 16)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = ConvertFromDec("", 16)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = ConvertFromDec("", 1)
	if err == nil {
		t.Error("Expecting an error")
	}
}

func ExampleConversion() {
	fmt.Println(Conversion("11WTDte", 62, 2))
	// Output:
	// 110110100100000010010110101110111000
}

func ExampleConvertToDec() {
	fmt.Println(ConvertToDec("453fa", 16))
	// Output:
	// 283642
}

func ExampleConvertFromDec() {
	fmt.Println(ConvertFromDec("6455537849492", 62))
	// Output:
	// 1PEvwVPC
}

func BenchmarkConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conversion("11WTDtejxjvosjsshs3455", 62, 2)
	}
}
