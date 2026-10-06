package store

import (
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/reader"
)

func openReader(src entry.Source, last Position, pol entry.Policy) (reader.Reader, error) {
	switch src.Kind {
	case entry.KindFile:
		pos, gen := src.Position, src.Generation
		if last.OK && last.Partial && last.Generation == src.Generation {
			pos, gen = last.Offset, last.Generation
		}

		return reader.NewFile(reader.FileOptions{
			Path: src.Path, Position: pos, Generation: gen,
			Identity: src.Identity, HeadHash: src.HeadHash,
			HeadLen: src.HeadLen, Policy: pol,
		})
	case entry.KindDummy:
		return reader.NewDummy(reader.StreamOptions{
			Seed: src.Seed, Key: src.Path, Start: nextRecord(src.Position),
			Rate: src.Rate, Policy: pol,
		}), nil
	case entry.KindBurst:
		return reader.NewBurst(reader.StreamOptions{
			Seed: src.Seed, Key: src.Path, Start: nextRecord(src.Position),
			Count: src.Total, Rate: src.Rate, Policy: pol,
		}), nil
	}

	return nil, fmt.Errorf("store: source %d has unknown kind %q", src.ID, src.Kind)
}

func nextRecord(position int64) int64 {
	if position < 1 {
		return 1
	}

	return position
}
