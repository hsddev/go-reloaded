package reloaded

import (
	"reflect"
	"reloaded"
	"testing"
)

type advancedConvertTest struct {
	got, expected []string
}

var advancedConvertTests = []advancedConvertTest{
	{reloaded.AdvancedConvert([]string{"it","was","the","age","of","foolishness", "(cap, ","6)"}),[]string{"It","Was","The","Age","Of","Foolishness"}},
}
func AdConvert(t *testing.T) {
	for _, test := range advancedConvertTests {
		if output := reloaded.AdvancedConvert(test.got); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}