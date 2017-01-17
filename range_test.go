package srg

import "testing"

func TestParseServers(t *testing.T) {
	cases := []struct {
		in   string
		want []string
		err  error
	}{
		{"abc", []string{"abc"}, nil},
		{"abc,xyz", []string{"abc", "xyz"}, nil},
		{"1,10,13~15", []string{"1", "10", "13", "14", "15"}, nil},
		{"abc1~abc3,xyz", []string{"abc1", "abc2", "abc3", "xyz"}, nil},
		{"abc1~abc3,abc4,abc5~abc5", []string{"abc1", "abc2", "abc3", "abc4", "abc5"}, nil},
	}

	for _, c := range cases {
		got, err := ParseRange(c.in)

		if err != c.err {
			t.Errorf("ParseServers(%q) expected error %q but get %q", c.in, c.err, err)
		}

		if diff(c.want, got) {
			t.Errorf("ParseServers(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func diff(want []string, got []string) bool {

	if len(want) != len(got) {
		return true
	}

	for i, _ := range want {
		if want[i] != got[i] {
			return true
		}
	}

	return false
}
