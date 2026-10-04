//go:build !js && !android && !ios

package page

// watchHosts is true on the desktop and server hosts, which reread a file
// while the process runs.
const watchHosts = true
