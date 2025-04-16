package render_test

import (
	"bytes"
	"testing"

	render "github.com/Darshan-1012/TDD_Basics/templating"
)

func TestRender(t *testing.T) {
	var (
		aPost = render.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := render.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1><p>This is description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("\ngot '%s' \n want '%s'", got, want)
		}
	})
}
