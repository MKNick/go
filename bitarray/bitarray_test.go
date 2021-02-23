package bitarray_test
import(
	"test/bitarray"
	"testing"
)

func TestHas(t *testing.T) {
	var x bitarray.IntSet
	x.Add(5)
	if !x.Has(5) {
		t.Error("bitarray func Has check faild")
	}
	if x.Has(6) {
		t.Error("bitarray func Has check faild")
	}
}