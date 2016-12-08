package awsxray

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// getRandom generates a random hex encoded string
func getRandom(i int) string {
	b := make([]byte, i)
	for {
		// keep trying till we get it
		if _, err := rand.Read(b); err != nil {
			continue
		}
		return fmt.Sprintf("%x", b)
	}
}

// GetTraceId returns trace id from header or blank
func GetTraceId(header string) string {
	for _, h := range strings.Split(header, ";") {
		th := strings.TrimSpace(h)
		if strings.HasPrefix(th, "Root=") {
			return strings.TrimPrefix(th, "Root=")
		}
	}

	if len(header) > 0 {
		// return as is
		return header
	}

	return ""
}

// GetParentId returns parent id from header or blank
func GetParentId(header string) string {
	for _, h := range strings.Split(header, ";") {
		th := strings.TrimSpace(h)
		if strings.HasPrefix(th, "Parent=") {
			return strings.TrimPrefix(th, "Parent=")
		}
	}

	return ""
}

// SetParentId will set the parent id of a trace in the header
func SetParentId(header, parentId string) string {
	parentHeader := fmt.Sprintf("Parent=%s", parentId)

	// no existing header?
	if len(header) == 0 {
		return parentHeader
	}

	headers := strings.Split(header, ";")

	for i, h := range headers {
		th := strings.TrimSpace(h)

		if len(th) == 0 {
			continue
		}

		// get Parent=Id match
		if strings.HasPrefix(th, "Parent=") {
			// set parent header
			headers[i] = parentHeader
			// return entire header
			return strings.Join(headers, "; ")
		}
	}

	// no match; set new parent header
	return strings.Join(append(headers, parentHeader), "; ")
}

// SetTraceId will set the trace id in the header
func SetTraceId(header, traceId string) string {
	traceHeader := fmt.Sprintf("Root=%s", traceId)

	// no existing header?
	if len(header) == 0 {
		return traceHeader
	}

	headers := strings.Split(header, ";")

	for i, h := range headers {
		th := strings.TrimSpace(h)

		if len(th) == 0 {
			continue
		}

		// get Root=Id match
		if strings.HasPrefix(th, "Root=") {
			// set trace header
			headers[i] = traceHeader
			// return entire header
			return strings.Join(headers, "; ")
		}
	}

	// no match; set new trace header as first entry
	return strings.Join(append([]string{traceHeader}, headers...), "; ")
}
