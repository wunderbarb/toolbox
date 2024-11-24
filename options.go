// v0.1.0
// Author: DIEHL E.
// ©  Sep 2023

package toolbox

type options struct {
	withDir     bool
	orderedSize bool
	ext         string
}

// Option allows parameterizing function
type Option func(opts *options)

func collectOptions(opts ...Option) *options {
	oo := &options{}
	for _, option := range opts {
		option(oo)
	}
	return oo
}

func WithExtension(ext string) Option {
	return func(op *options) {
		op.ext = ext
	}
}

// WithSubDir allows specifying that the subdirectories should be also listed.
func WithSubDir() Option {
	return func(op *options) {
		op.withDir = true
	}
}

// WithOrderedSize allows to specify that the directory size should be computed
func WithOrderedSize() Option {
	return func(op *options) {
		op.orderedSize = true
	}
}
