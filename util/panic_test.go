package util

import (
	"testing"
)

func TestHelpPanic_Display(t *testing.T) {
	type fields struct {
		Reason string
		Help   string
		Error  error
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "generate error display",
			fields: fields{
				Reason: "test",
				Help:   "",
				Error:  nil,
			},
			args: args{},
		},
		// TODO: add tests for other sub templates and failures
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HelpPanic{
				Reason: tt.fields.Reason,
				Help:   tt.fields.Help,
				Error:  tt.fields.Error,
			}
			if got, err := p.Display(tt.args.args...); err != nil {
				t.Errorf("HelpPanic.Display() = \n%v", got)
			}
		})
	}
}
