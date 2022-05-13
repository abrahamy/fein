package result

import (
	"errors"
	"testing"
)

func TestOk(t *testing.T) {
	Ok[int, error](1)
}

func TestErr(t *testing.T) {
	Err[int](errors.New("Error!!!"))
}

func TestAnd(t *testing.T) {
	ok := Ok[int, error](1)
	alsoOk := Ok[int, error](2)
	err := Err[int](errors.New("this is an error!"))
	alsoErr := Err[int](errors.New("this is also an error!"))

	if ok.And(alsoOk) != alsoOk {
		t.Errorf("ok.Add(alsoOk) did not return alsoOk")
	}

	if ok.And(err) != err {
		t.Errorf("ok.Add(err) did not return err")
	}

	if err.And(ok) != err {
		t.Errorf("err.Add(ok) did not return err")
	}

	if err.And(alsoErr) != err {
		t.Errorf("err.Add(alsoErr) did not return err")
	}
}

func TestIsErr(t *testing.T) {
	ok := Ok[int, error](1)
	if ok.IsErr() {
		t.Errorf("Ok constructor yielded Err variant")
	}

	err := Err[any](errors.New("this is an error!"))
	if !err.IsErr() {
		t.Errorf("Err constructor yielded Ok variant")
	}
}
