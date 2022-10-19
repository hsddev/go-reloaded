package reloaded

import (
	"reflect"
	"reloaded"
	"testing"
)

type basicConvertTest struct {
	got, expected []string
}

var basicConvertTests = []basicConvertTest{
	{reloaded.BasicConvert([]string{"Simply","add", "42","(hex)"}),[]string{"Simply","add", "66"}},
}
func TestAlterText(t *testing.T) {
	for _, test := range basicConvertTests {
		if output := reloaded.BasicConvert(test.got); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}