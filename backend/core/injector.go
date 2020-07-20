package core

import (
	"go.uber.org/dig"
)

// Injector wires all the application dependencies. This is done at runtime.
//
// It allows an easier setup in main, as well as more private class definitions (to enforce as most consistency as
// possible in the architecture).
//
//Given that it reuses the same standard constructors, it also lives perfectly alongside the tests.
//
var Injector = dig.New(dig.DeferAcyclicVerification())
