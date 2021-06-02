// Package command implement the operation controller
package command

// FIXME temporary workaround and it belongs to configurations provided to the functionality
var stages map[string][]string = map[string][]string{
	"mm": {
		"prod",
		"dev",
	},
	"tv": {
		"prod",
		"staging",
		"dev",
	},
	"readr": {
		"prod",
		"dev",
	},
}
