package store

// dummyNames cycle through the seeded history.
var dummyNames = []string{
	"annual-report-2026.pdf", "ubuntu-live-server.iso", "photos.zip",
	"metrics-dataset.csv", "podcast-episode-142.mp3", "design-assets.tar.gz",
	"invoice-4471.pdf", "release-notes.txt", "backup-2026-09.tar",
	"training-video.mp4",
}

// dummyErrors are the failure texts the seeded history uses.
var dummyErrors = []string{
	"connection reset by peer", "HTTP 404", "checksum mismatch",
	"no data before the stall deadline", "HTTP 503",
}
