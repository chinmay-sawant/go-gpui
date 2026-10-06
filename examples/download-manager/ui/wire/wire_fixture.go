package wire

import (
	"log"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/fixture"
)

// startFixture opens the local HTTP fixture service and logs the paths the
// add form can use.
func (b *Backend) startFixture() {
	b.fixture = fixture.NewServer()

	for _, path := range []string{
		fixture.PathOK, fixture.PathSlow, fixture.PathChunked,
		fixture.PathRange, fixture.PathChanging, fixture.PathRedirect,
		fixture.PathInterrupt,
	} {
		log.Printf("download-manager: fixture %s%s", b.fixture.URL, path)
	}
}
