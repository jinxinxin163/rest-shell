// one_test.go
package test

import (
	"testing"
	"fmt"
	"os"
)

const checkMark = "\u2713" //
const ballotX = "\u2717"   //

//

func TestMain(m *testing.M) {
	fmt.Printf("initial func here")
	os.Exit(m.Run())
}

func TestA(t *testing.T) {
	t.Log(checkMark, "TestA OK,Logf")
	t.Log("TestA end")
}

func TestB(t *testing.T) {
    t.Error(ballotX, "TestB Fail, But continue,Errorf")
    t.Log("TestB end")
}
func TestC(t *testing.T) {
    t.Fatal(ballotX, "TestC Fail, Terminated,Fatalf")
    t.Log("TestC Can not run here")
}

