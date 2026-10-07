package entry

// ForStream returns the policy a non-replayable stream uses: the same batch
// limits with counted loss instead of a pause.
func (p Policy) ForStream() Policy {
	p.Overflow = OverflowCountLoss

	return p
}
