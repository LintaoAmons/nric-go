package nric_test

import (
	"fmt"
	"testing"

	. "github.com/LintaoAmons/nric-go/nric"
)

func TestGenerateNric(t *testing.T) {
	nric := GenerateNric()
	fmt.Println(nric)
}
