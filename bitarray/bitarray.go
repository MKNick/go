package bitarray
import "fmt"
type  IntSet struct{
	words []uint64
}

func (s *IntSet)Has(val int) bool{
	word, bit := val/64, uint(val%64)
	if word > len(s.words){
		return false
	}

	return s.words[word] & (1<<bit) != 0
}

func (s * IntSet)Add(val int){
	word, bit := val/64, uint(val%64)
	for  word >= len(s.words){
		slice := make([]uint64, word-len(s.words)+1)
		s.words=append(s.words,slice...)
	}
	s.words[word] |= 1<<bit
	fmt.Println("Add index:",word,"value:",1<<bit,"total:",s.words[word])
}

func (s *IntSet)Remove(val int){
	word, bit := val/64, uint(val%64)
	if word > len(s.words) {
		return
	}

	if s.words[word] & (1<<bit) !=0 {
		s.words[word] -= 1<<bit
	}
}

func (s * IntSet)Count() int{
	total := 0
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if word&(1<<i) != 0 {
				total++
			}
		}
	}
	return total
}

func (s *IntSet)Clear(){
	s.words=s.words[0:0]
}

func (s *IntSet)UnionWith(t *IntSet){
	for i,word :=range t.words{
		if i < len(s.words) {
			s.words[i] |=word
		} else{
			s.words = append(s.words, word)
		}
	}
}

