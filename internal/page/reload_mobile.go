//go:build js || android || ios

package page

// watchHosts is false on wasm and mobile: a file is read once at New.
const watchHosts = false
