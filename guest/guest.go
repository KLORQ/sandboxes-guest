// Package guest exposes guest-side primitives.
package guest

// Info describes the guest runtime.
type Info struct {
	OS       string
	Arch     string
	Hostname string
}
