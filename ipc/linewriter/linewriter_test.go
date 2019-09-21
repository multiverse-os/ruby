package linewriter

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	testCases := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"1\n", "2\n"},
			want:  []string{"1", "2"},
		},
		{
			input: []string{"1\n", "\n", "2\n"},
			want:  []string{"1", "", "2"},
		},
		{
			input: []string{"1\n2\n", "3\n"},
			want:  []string{"1", "2", "3"},
		},
		{
			input: []string{"1", "2\n"},
			want:  []string{"12"},
		},
		{
			// Data with no newline yet is omitted.
			input: []string{"1\n", "2\n", "3"},
			want:  []string{"1", "2"},
		},
	}

	for _, c := range testCases {
		var lines [][]byte

		w := NewWriter(func(p []byte) {
			// We must not retain p, so we must make a copy.
			b := make([]byte, len(p))
			copy(b, p)

			lines = append(lines, b)
		})

		for _, in := range c.input {
			n, err := w.Write([]byte(in))
			if err != nil {
				t.Errorf("[error] erite(%q) err got %v want nil (case %+v)", in, err, c)
			}
			if n != len(in) {
				t.Errorf("[error] write(%q) b got %d want %d (case %+v)", in, n, len(in), c)
			}
		}

		if len(lines) != len(c.want) {
			t.Errorf("[error] len(lines) got %d want %d (case %+v)", len(lines), len(c.want), c)
		}

		for i := range lines {
			if !bytes.Equal(lines[i], []byte(c.want[i])) {
				t.Errorf("[error] item %d got %q want %q (case %+v)", i, lines[i], c.want[i], c)
			}
		}
	}
}
