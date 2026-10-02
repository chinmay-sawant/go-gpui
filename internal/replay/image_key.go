package replay

import (
	"hash/fnv"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// imageBytesKey identifies one encoded payload with FNV-1a. The IsJPEG flag is
// folded in so identical bytes tagged two ways never share an entry.
func imageBytesKey(op *layout.DisplayOp) uint64 {
	if op == nil {
		return 0
	}

	data, _, _ := op.ImageBytes()

	hash := fnv.New64a()
	if op.IsJPEG {
		_, _ = hash.Write([]byte{1})
	}

	_, _ = hash.Write(data)

	return hash.Sum64()
}

// imageGeom maps the decoded image's w by h pixels onto op's point rectangle,
// moved by the page offset. It scales then translates; the replay gate rejects
// non-identity transforms before any draw.
func imageGeom(op *layout.DisplayOp, dx, dy float64, w, h int) ebiten.GeoM {
	var geom ebiten.GeoM

	if w > 0 {
		geom.Scale(op.W*pxPerPt/float64(w), 1)
	}

	if h > 0 {
		geom.Scale(1, op.H*pxPerPt/float64(h))
	}

	geom.Translate(op.X*pxPerPt+dx, op.Y*pxPerPt+dy)

	return geom
}
