package parser

import (
	"strings"
	"testing"

	"github.com/code-gorilla-au/odize"
	"github.com/code-gorilla-au/pyrotic/internal/testfixtures"
)

func Test_withTemplates(t *testing.T) {
	type args struct {
		fileSuffix string
		dirPath    string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "should return inject_after.tmpl",
			args: args{
				fileSuffix: "tmpl",
				dirPath:    "../../example/_templates/fakr",
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "should return error if not exists",
			args: args{
				fileSuffix: "tmpl",
				dirPath:    "../../example/_templates/flat",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := withTemplates(tt.args.dirPath, tt.args.fileSuffix)
			if (err != nil) != tt.wantErr {
				t.Errorf("withTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			odize.AssertEqual(t, tt.want, len(got))
			for _, tmp := range got {
				odize.AssertTrue(t, len(tmp) > 0)
			}
		})
	}
}

func TestTmplEngine_Parse(t *testing.T) {
	group := odize.NewGroup(t, nil)

	expected := TemplateData{
		Name: "hello",
	}

	var te *TemplateEngine

	group.BeforeEach(func() {
		strGoFile := `---
		to: elo.go
		---
		blah
		`

		strNonGoFile := `---
		to: {{ "elo" | caseSnake }}
		---
		blah
		`

		te = &TemplateEngine{
			templates: map[string]string{
				"tmp":  strGoFile,
				"tmp2": strNonGoFile,
			},
			funcs: defaultFuncs,
		}
	})

	err := group.
		Test("should parse go file and format", func(t *testing.T) {
			data, err := te.Parse(expected)
			odize.AssertNoError(t, err)

			result, found := testfixtures.Find(data, func(d TemplateData) bool {
				return d.To == "elo.go"
			})
			odize.AssertTrue(t, found)

			odize.AssertEqual(t, expected.Name, result.Name)
			odize.AssertEqual(t, "elo.go", result.To)
			odize.AssertEqual(t, "blah", strings.TrimSpace(string(result.Output)))
		}).
		Test("should parse non go file", func(t *testing.T) {
			data, err := te.Parse(expected)

			result, found := testfixtures.Find(data, func(d TemplateData) bool {
				return d.To == "elo"
			})
			odize.AssertTrue(t, found)

			odize.AssertNoError(t, err)
			odize.AssertEqual(t, expected.Name, result.Name)
			odize.AssertEqual(t, "elo", result.To)
			odize.AssertEqual(t, "blah", strings.TrimSpace(string(result.Output)))
		}).
		Run()
	odize.AssertNoError(t, err)
}

func TestTmplEngine_Parse_missing_funcs_should_fail_on_meta_parse(t *testing.T) {
	strTmp := `---
	to: {{ "elo" | caseSnake }}
	---
	blah
	`
	expected := TemplateData{
		Name: "hello",
	}
	te := &TemplateEngine{
		templates: map[string]string{"tmp": strTmp},
	}
	_, err := te.Parse(expected)
	odize.AssertError(t, err)
}

func TestTmplEngine_Parse_missing_funcs_should_fail_on_template_parse(t *testing.T) {
	strTmp := `---
	to: elo
	---
	blah {{ "foo" | caseSnake }}
	`
	expected := TemplateData{
		Name: "hello",
	}
	te := &TemplateEngine{
		templates: map[string]string{"tmp": strTmp},
	}
	_, err := te.Parse(expected)
	odize.AssertError(t, err)
}
