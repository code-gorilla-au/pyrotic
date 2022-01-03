package writer

import (
	"testing"
)

func TestInject_Validate(t *testing.T) {
	type fields struct {
		Matcher string
		Clause  InjectClause
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should return false if both are missing",
			fields: fields{
				Matcher: "",
				Clause:  "",
			},
			want: false,
		},
		{
			name: "should return false if both are missing",
			fields: fields{
				Matcher: "// flash",
				Clause:  "",
			},
			want: true,
		},
		{
			name: "should return false if both are missing",
			fields: fields{
				Matcher: "",
				Clause:  "// thunder",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inject{
				Matcher: tt.fields.Matcher,
				Clause:  tt.fields.Clause,
			}
			if got := i.Validate(); got != tt.want {
				t.Errorf("Inject.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
