package generator_test

import (
	"testing"

	"github.com/reedom/convergen/pkg/generator"
	"github.com/reedom/convergen/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestGenerator_ArgRetReceiver(t *testing.T) {
	t.Parallel()

	const pre = `package simple

import (
	"github.com/reedom/convergen/pkg/tests/fixtures/data/domain"
	"github.com/reedom/convergen/pkg/tests/fixtures/data/model"
)
`

	const header = "// Code generated by github.com/reedom/convergen\n// DO NOT EDIT.\n\n"

	cases := []struct {
		name     string
		fn       *model.Function
		expected string
	}{
		{
			name: "src:ptr/dst:ptr,arg/rhs:simple",
			fn: &model.Function{
				Comments:     []string{"// comment 1", "// comment 2"},
				Name:         "ToModel",
				Receiver:     "",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: false,
				DstVarStyle:  model.DstVarArg,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID"}},
				},
			},
			expected: header + pre + `
// comment 1
// comment 2
func ToModel(dst *model.Pet, src *domain.Pet) {
	dst.ID = src.ID
}
`,
		},
		{
			name: "src:ptr/dst:ptr,return/rhs:simple",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: false,
				DstVarStyle:  model.DstVarReturn,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID"}},
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst *model.Pet) {
	dst = &model.Pet{}
	dst.ID = src.ID

	return
}
`,
		},
		{
			name: "src:ptr/dst:copy,return/rhs:simple",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: false},
				ReturnsError: false,
				DstVarStyle:  model.DstVarReturn,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID"}},
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst model.Pet) {
	dst.ID = src.ID

	return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:copy,return/rhs:simple",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "src",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: false},
				ReturnsError: false,
				DstVarStyle:  model.DstVarReturn,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID"}},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel() (dst model.Pet) {
	dst.ID = src.ID

	return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:ptr,arg/rhs:simple",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "src",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: false,
				DstVarStyle:  model.DstVarArg,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID"}},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel(dst *model.Pet) {
	dst.ID = src.ID
}
`,
		},
		{
			name: "src:ptr,receiver/dst:ptr,arg/error/rhs:simple",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "src",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: true,
				DstVarStyle:  model.DstVarArg,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID()", Error: true}},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel(dst *model.Pet) (err error) {
	dst.ID, err = src.ID()
	if err != nil {
		return
	}

	return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:ptr,return/error/rhs:simple",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "src",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: true,
				DstVarStyle:  model.DstVarReturn,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID()", Error: true}},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel() (dst *model.Pet, err error) {
	dst = &model.Pet{}
	dst.ID, err = src.ID()
	if err != nil {
		return nil, err
	}

	return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:val,return/error/rhs:simple",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "src",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: false},
				ReturnsError: true,
				DstVarStyle:  model.DstVarReturn,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID()", Error: true}},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel() (dst model.Pet, err error) {
	dst.ID, err = src.ID()
	if err != nil {
		return
	}

	return
}
`,
		},
		{
			name: "src:ptr/dst:ptr,arg/error/rhs:skip",
			fn: &model.Function{
				Name:         "ToModel",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: true,
				DstVarStyle:  model.DstVarArg,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SkipField{}},
				},
			},
			expected: header + pre + `
func ToModel(dst *model.Pet, src *domain.Pet) (err error) {
	// skip: dst.ID

	return
}
`,
		},
		{
			name: "src:ptr/dst:ptr,return/error/rhs:nomatch",
			fn: &model.Function{
				Name:         "ToModel",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: true,
				DstVarStyle:  model.DstVarReturn,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.NoMatchField{}},
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst *model.Pet, err error) {
	dst = &model.Pet{}
	// no match: dst.ID

	return
}
`,
		},
		{
			name: "preprocess/postprocess",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: true},
				ReturnsError: false,
				DstVarStyle:  model.DstVarArg,
				Assignments: []*model.Assignment{
					{LHS: "dst.ID", RHS: model.SimpleField{Path: "src.ID"}},
				},
				PreProcess: &model.Manipulator{
					Name:         "PreProcess",
					Src:          model.Var{Name: "rhs", Pointer: true},
					Dst:          model.Var{Name: "lhs", Pointer: true},
					ReturnsError: false,
				},
				PostProcess: &model.Manipulator{
					Pkg:          "domain",
					Name:         "PostProcess",
					Src:          model.Var{Name: "rhs", Pointer: true},
					Dst:          model.Var{Name: "lhs", Pointer: true},
					ReturnsError: false,
				},
			},
			expected: header + pre + `
func ToModel(dst *model.Pet, src *domain.Pet) {
	PreProcess(dst, src)
	dst.ID = src.ID
	domain.PostProcess(dst, src)
}
`,
		},
		{
			name: "postprocess/dst,val/src,val/rhs,ptr/error",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: false},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: false},
				ReturnsError: true,
				DstVarStyle:  model.DstVarReturn,
				PostProcess: &model.Manipulator{
					Name:         "PostProcess",
					Src:          model.Var{Name: "rhs", Pointer: true},
					Dst:          model.Var{Name: "lhs", Pointer: true},
					ReturnsError: true,
				},
			},
			expected: header + pre + `
func ToModel(src domain.Pet) (dst model.Pet, err error) {
	err = PostProcess(&dst, &src)
	if err != nil {
		return
	}

	return
}
`,
		},
		{
			name: "postprocess/dst,val/src,ptr/rhs,val/error",
			fn: &model.Function{
				Name:         "ToModel",
				Receiver:     "",
				Src:          model.Var{Name: "src", PkgName: "domain", Type: "Pet", Pointer: true},
				Dst:          model.Var{Name: "dst", PkgName: "model", Type: "Pet", Pointer: false},
				ReturnsError: true,
				DstVarStyle:  model.DstVarReturn,
				PostProcess: &model.Manipulator{
					Name:         "PostProcess",
					Src:          model.Var{Name: "rhs", Pointer: false},
					Dst:          model.Var{Name: "lhs", Pointer: true},
					ReturnsError: true,
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst model.Pet, err error) {
	err = PostProcess(&dst, *src)
	if err != nil {
		return
	}

	return
}
`,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			code := model.Code{
				Pre:       pre,
				Functions: []*model.Function{tt.fn},
			}
			g := generator.NewGenerator(code)
			actual, err := g.Generate("temp.gen.go", false, true)
			if assert.Nil(t, err) {
				assert.Equal(t, tt.expected, string(actual))
			}
		})
	}
}
