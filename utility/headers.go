package bns

import (
	"net/http"
	"net/textproto"
)

// AppendHeader copies all header values from src to dst with canonicalized keys.
// No-op if either src or dst is nil.
func AppendHeader(src http.Header, dst http.Header) {
	if src == nil || dst == nil {
		return
	}

	for key, vals := range src {
		if key == "" || len(vals) == 0 {
			continue
		}
		ck := textproto.CanonicalMIMEHeaderKey(key)

		dst[ck] = append(dst[ck], vals...)
	}
}

// CopyHeaderFrom creates a deep copy of the provided header.
// Returns empty header if src is nil.
func CopyHeaderFrom(src http.Header) http.Header {
	out := make(http.Header)
	if src == nil {
		return out
	}

	for key, vals := range src {
		if key == "" || len(vals) == 0 {
			continue
		}
		ck := textproto.CanonicalMIMEHeaderKey(key)

		out[ck] = append([]string(nil), vals...)
	}

	return out
}
