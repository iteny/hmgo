package hmgo

import (
	"regexp"
)

// Tree has three elements: FixRouter/wildcard/leaves
// fixRouter sotres Fixed Router
// wildcard stores params
// leaves store the endpoint information
type Tree struct {
	//prefix set for static router
	prefix string
	//search fix route first
	fixrouters []*Tree
	//if set, failure to match fixrouters search then search wildcard
	wildcard *Tree
	//if set, failure to match wildcard search
	leaves []*leafInfo
}
type leafInfo struct {
	// names of wildcards that lead to this leaf. eg, ["id" "name"] for the wildcard ":id" and ":name"
	wildcards []string

	// if the leaf is regexp
	regexps *regexp.Regexp

	runObject interface{}
}
