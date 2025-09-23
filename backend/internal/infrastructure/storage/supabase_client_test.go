package storage

import "testing"

func TestNormalizeSupabaseBaseURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "already normalized",
			input: "https://example.supabase.co",
			want:  "https://example.supabase.co",
		},
		{
			name:  "trailing slash",
			input: "https://example.supabase.co/",
			want:  "https://example.supabase.co",
		},
		{
			name:  "storage api path included",
			input: "https://example.supabase.co/storage/v1/s3",
			want:  "https://example.supabase.co",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeSupabaseBaseURL(tt.input); got != tt.want {
				t.Fatalf("normalizeSupabaseBaseURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewSupabaseClientNormalizesBaseURL(t *testing.T) {
	rawBase := "https://example.supabase.co/storage/v1/s3"
	client := NewSupabaseClient(rawBase, "bucket", "key", nil)
	sc, ok := client.(*supabaseClient)
	if !ok {
		t.Fatalf("NewSupabaseClient() did not return *supabaseClient")
	}
	wantBase := "https://example.supabase.co"
	if sc.baseURL != wantBase {
		t.Fatalf("baseURL = %q, want %q", sc.baseURL, wantBase)
	}
	wantPrefix := wantBase + "/storage/v1/object/public/bucket/"
	if sc.publicURLPrefix != wantPrefix {
		t.Fatalf("publicURLPrefix = %q, want %q", sc.publicURLPrefix, wantPrefix)
	}
}
