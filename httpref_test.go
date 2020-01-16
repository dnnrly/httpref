package httpref

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReferences_ByName(t *testing.T) {
	Statuses = append(Statuses, Reference{Name: "501-extended"})
	type args struct {
		code string
	}
	tests := []struct {
		name string
		want int
	}{
		{name: "1", want: 5},
		{name: "40", want: 10},
		{name: "501", want: 1},
		{name: "501*", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Statuses.ByName(tt.name); len(got) != tt.want {
				t.Errorf("References.ByName() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestReference_Summarize(t *testing.T) {
	r := Reference{
		Name:        "name",
		Summary:     "summary",
		Description: "description",
	}

	s := r.Summarize()
	assert.Equal(t, "                          name summary", s)

	r = Reference{
		Name:        "title name",
		IsTitle:     true,
		Summary:     "title summary",
		Description: "title description",
	}

	s = r.Summarize()
	assert.Equal(t, "                    title name title summary", s)
}

func TestReference_Describe(t *testing.T) {
	description := Headers.ByName("Headers")[0].Describe()

	expected := `HTTP headers let the client and the server pass additional information with an HTTP request or
response. An HTTP header consists of its case-insensitive name followed by a colon (:), then by its
value. Whitespace before the value is ignored.

Custom proprietary headers have historically been used with an X- prefix, but this convention was
:b3
1:bd
standard in RFC 6648; others are listed in an IANA registry, whose original content was defined in
RFC 4229. IANA also maintains a registry of proposed new HTTP headers.

Headers can be grouped according to their contexts:

    General headers apply to both requests and responses, but with no relation to the data
    transmitted in the body.
    Request headers contain more information about the resource to be fetched, or about the client
    requesting the resource.
    Response headers hold additional information about the response, like its location or about the
    server providing it.
    Entity headers contain information about the body of the resource, like its content length or
    MIME type.

Headers can also be grouped according to how proxies handle them:

    Connection
    Keep-Alive
    Proxy-Authenticate
    Proxy-Authorization
    TE
    Trailer
    Transfer-Encoding
    Upgrade.

End-to-end headers
    These headers must be transmitted to the final recipient of the message: the server for a
    request, or the client for a response. Intermediate proxies must retransmit these headers
    unmodified and caches must store them.
Hop-by-hop headers
    These headers are meaningful only for a single transport-level connection, and must not be
    retransmitted by proxies or cached. Note that only hop-by-hop headers may be set using the
	Connection general header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers`

	assert.Equal(t, expected, description)
}

func TestReferences_Titles(t *testing.T) {
	n := Statuses.Titles()
	assert.Equal(t, 5, len(n))
}
