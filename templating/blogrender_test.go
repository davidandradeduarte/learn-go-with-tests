package blogrender

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	post := &Post{
		Title:       "title",
		Body:        "body",
		Description: "desc",
		Tags:        []string{"t1", "t2"},
	}

	t.Run("converts post into html", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := Render(&buf, post)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
		// got := buf.String()
		// want := `<h1>title</h1><p>desc</p>Tags:<ul><li>t1</li><li>t2</li></ul>`

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})
}
