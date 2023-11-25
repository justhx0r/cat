package plaintxt_test

import (
	"testing"

	"github.com/justhx0r/cat/plaintxt"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.\n"

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func TestToStr(t *testing.T) {
	txt, err := plaintxt.ToStr("../test/test.txt")
	if err != nil {
		t.Error(err)
	} else if txt == test {
		t.Log("success")
	} else {
		t.Error("does not match test:", txt, test)
	}
	_, err = plaintxt.ToStr("./test/nonexistent")
	if err == nil {
		t.Error("Nonexisting file does not throw error")
	}
}
