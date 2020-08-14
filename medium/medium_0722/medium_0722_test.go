package medium_0722

import (
	"reflect"
	"testing"
)

func TestRemoveComments(t *testing.T) {
	var rc = []struct {
		source   []string
		expected []string
	}{
		//{[]string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test",
		//	"   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"},
		//	[]string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"}},
		//{[]string{"a/*comment", "line", "more_comment*/b"}, []string{"ab"}},
		//{[]string{"struct Node{", "    /*/ declare members;/**/", "    int size;", "    /**/int val;", "};"},
		//	[]string{"struct Node{", "    ", "    int size;", "    int val;", "};"}},
		//{[]string{"main() {", "/* here is commments", "  // still comments */", "   double s = 33;", "   cout << s;", "}"},
		//	[]string{"main() {", "   double s = 33;", "   cout << s;", "}"}},
		//{[]string{"main() {", "  Node* p;", "  /* declare a Node", "  /*float f = 2.0", "   p->val = f;", "   /**/", "   p->val = 1;", "   //*/ cout << success;*/", "}", " "},
		//	[]string{"main() {", "  Node* p;", "  ", "   p->val = 1;", "   ", "}", " "}},
		//{[]string{"a//*b//*c", "blank", "d//*e/*/f"}, []string{"a", "blank", "d"}},
		//{[]string{"a//*b//*c", "blank", "d*//e*//f"}, []string{"a", "blank", "d*"}},
		//{[]string{"a/*/b//*c", "blank", "d//*e/*/f"}, []string{"af"}},
		//{[]string{"a/*/b//*c", "blank", "d/*/e*//f"}, []string{"ae*"}},
	}

	for _, val := range rc {
		actual := removeComments(val.source)
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! actual :=%s, expected := %s", actual, val.expected)
		}
	}
}
