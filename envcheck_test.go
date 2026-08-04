package envcheck

import (
	"reflect"
	"testing"
)

func TestMissingReportsUnsetAndEmpty(t *testing.T) {
	env := map[string]string{"SET": "value", "EMPTY": ""}
	got := Missing([]string{"SET", "EMPTY", "ABSENT"}, env)
	want := []string{"EMPTY", "ABSENT"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Missing = %q, want %q", got, want)
	}
}

func TestMissingNoneWhenAllPresent(t *testing.T) {
	if got := Missing([]string{"A"}, map[string]string{"A": "1"}); len(got) != 0 {
		t.Errorf("Missing = %q, want empty", got)
	}
}
