package engine

import (
	"errors"
	"testing"
)

func TestStorageEngine_GetMissing(t *testing.T) {
	storage := NewStorageEngine()

	value, err := storage.Get("missing-key")

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("error not notfound: %v", err)
	}
	if value != nil {
		t.Errorf("value not nil %v", value)
	}
}

func TestStorageEngine_Set(t *testing.T) {
	storage := NewStorageEngine()

	err := storage.Set("a", "b")

	if err != nil {
		t.Fatalf("error not nil: %v", err)
	}
}

func TestStorageEngine_SetAndGet(t *testing.T) {
	storage := NewStorageEngine()

	err := storage.Set("a", "b")
	if err != nil {
		t.Fatalf("error set not nil: %v", err)
	}
	value, err := storage.Get("a")
	if err != nil {
		t.Fatalf("error get not nil: %v", err)
	}
	if value != "b" {
		t.Fatalf("value not equal b: %v", value)
	}
}

func TestStorageEngine_Delete(t *testing.T) {
	storage := NewStorageEngine()

	err := storage.Set("a", "b")
	if err != nil {
		t.Fatalf("error set not nil: %v", err)
	}
	err = storage.Delete("a")
	if err != nil {
		t.Fatalf("error delete not nil: %v", err)
	}
	v, err := storage.Get("a")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("error get not notfound: %v", err)
	}
	if v != nil {
		t.Fatalf("value not nil: %v", v)
	}
}
