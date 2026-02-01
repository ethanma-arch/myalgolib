package factory

import "testing"

func TestNewProduct(t *testing.T) {
	p := NewProduct("A")
	if p == nil {
		t.Fatal("NewProduct(\"A\") should not be nil")
	}
	if got := p.Name(); got != "ProductA" {
		t.Errorf("got %q, want ProductA", got)
	}

	p = NewProduct("B")
	if p == nil {
		t.Fatal("NewProduct(\"B\") should not be nil")
	}
	if got := p.Name(); got != "ProductB" {
		t.Errorf("got %q, want ProductB", got)
	}

	p = NewProduct("unknown")
	if p != nil {
		t.Errorf("NewProduct(\"unknown\") should be nil, got %v", p)
	}
}
