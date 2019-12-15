package medium_0130

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {

	var solveTest = []struct {
		in       [][]byte
		expected [][]byte
	}{
		//{[][]byte{[]byte("XXXX"), []byte("XOOX"), []byte("XXOX"), []byte("XOXX")}, [][]byte{[]byte("XXXX"), []byte("XXXX"), []byte("XXXX"), []byte("XOXX")}},
		{[][]byte{[]byte("XOXOXO"), []byte("OXOXOX"), []byte("XOXOXO"), []byte("OXOXOX")}, [][]byte{[]byte("XXXX"), []byte("XXXX"), []byte("XXXX"), []byte("XOXX")}},
	}

	for _, st := range solveTest {
		solve(st.in)

		if !reflect.DeepEqual(st.in, st.expected) {
			t.Errorf("Test Error!")
		}
	}

}
