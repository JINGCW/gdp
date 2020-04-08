package test

import (
	"fmt"
	"hash_collections/src/asm_x86/g"
	"testing"
)

func TestGPTR(t *testing.T) {
	gptr := g.GPTR()
	fmt.Printf("%#v", gptr)
}

func TestGID(t *testing.T) {
	gid := g.GID()
	fmt.Printf("%#v", gid)
	// Output: 1
}

func Test_G_innerouter(t *testing.T) {
	if gouter, ginner := g.G_innerouter(); ginner == gouter {
		t.Fatalf("ginner==guter :ginner->%d, gouter->%d", ginner, gouter)
	}
}
