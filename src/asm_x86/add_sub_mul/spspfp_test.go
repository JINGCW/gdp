package add_sub_mul

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAdd_x86(t *testing.T) {
	out := Add_x86(3, 4)
	assert.Equal(t, out, 7, fmt.Sprintf("TestAdd_x86(3,4): want 7, got %d", out))
}

func TestOutput(t *testing.T) {
	o1, o2, o3 := Output(3)
	out := []int{o1, o2, o3}
	assert.Equal(t, out, []int{3, 3, 3})
}

func TestOutput2(t *testing.T) {
	out := Output2(3, 4)
	assert.Equal(t, out, 7)
}

func TestArraySum(t *testing.T) {
	out := ArraySum([]int64{3, 4, 5})
	assert.Equal(t, out, 12)
}

func TestOffsetSB(t *testing.T) {
	fmt.Printf("%v", OffsetSB())
}
