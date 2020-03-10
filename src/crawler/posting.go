package crawler

import "fmt"

// Posting is an Instagram post type.
type Posting struct {
	URL  string `json:"url,omitempty"`
	SRC  string `json:"src,omitempty"`
	Tags string `json:"tags,omitempty"`
	Like int    `json:"like,string,omitempty"`
}

// String implements fmt.Stringer interface.
func (p Posting) String() string {
	return fmt.Sprintf("url: %s\tsrc: %s\ttags: %s\tlike: %d", p.URL, p.SRC, p.Tags, p.Like)
}
