package domain

import ()

type Nullable[T any] struct {
	Value *T
	Set   bool
}
