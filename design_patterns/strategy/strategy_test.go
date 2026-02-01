package strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	ctx := NewContext(AddStrategy{})
	if got := ctx.Execute(3, 5); got != 8 {
		t.Errorf("AddStrategy: got %d, want 8", got)
	}

	ctx.SetStrategy(SubStrategy{})
	if got := ctx.Execute(10, 4); got != 6 {
		t.Errorf("SubStrategy: got %d, want 6", got)
	}

	ctx.SetStrategy(MulStrategy{})
	if got := ctx.Execute(3, 4); got != 12 {
		t.Errorf("MulStrategy: got %d, want 12", got)
	}
}
