package provider

import (
	"testing"
)

func TestLchomp(t *testing.T) {
	src := map[string]interface{}{
		"example.com":         "1",
		"example.org":         "2",
		"sub.example.org":     "3",
		"alt.example.org":     "4",
		"sub.alt.example.org": "5",
	}

	tests := []struct {
		key    string
		sep    string
		want   string
		wantOk bool
	}{
		{
			key:    "example.com",
			sep:    ".",
			want:   "example.com",
			wantOk: true,
		},
		{
			key:    "sub.example.com",
			sep:    ".",
			want:   "example.com",
			wantOk: true,
		},
		{
			key:    "sub.exempli-gratia.com",
			sep:    ".",
			want:   "",
			wantOk: false,
		},
		{
			key:    "sub.example.org",
			sep:    ".",
			want:   "sub.example.org",
			wantOk: true,
		},
		{
			key:    "sub.example.org",
			sep:    "/",
			want:   "sub.example.org",
			wantOk: true,
		},
		{
			key:    "no-entry.sub.example.org",
			sep:    "/",
			want:   "",
			wantOk: false,
		},
		{
			key:    "a.b.c.d.e.f.g.sub.example.org",
			sep:    ".",
			want:   "sub.example.org",
			wantOk: true,
		},
		{
			key:    "",
			sep:    ".",
			want:   "",
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, ok := lchomp(tt.key, tt.sep, src)
			if got != tt.want {
				t.Errorf("key error: got \"%v\", want \"%v\"", got, tt.want)
			}
			if ok != tt.wantOk {
				t.Errorf("ok error: got %v, want %v", ok, tt.wantOk)
			}
		})
	}
}
