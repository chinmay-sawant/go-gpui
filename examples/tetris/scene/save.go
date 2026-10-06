package scene

import "time"

// saveRetry is how long a failed or refused score save waits before it is
// tried again.
const saveRetry = 2 * time.Second

// saveCompleted takes a finished run from the model and queues it. The
// resume slot is dropped: the run cannot be resumed once it is over.
func (s *Scene) saveCompleted() {
	if s.store == nil {
		return
	}

	res, rep, ok := s.model.TakeCompleted()
	if !ok {
		return
	}

	r := res
	s.pendingSave = &r
	s.pendingReplay = rep
	s.nextSaveTry = time.Time{}
	s.setStatus("SCORE SAVING")
	s.clearResume()
}

// scheduleSave queues a pending score save when the worker has room. A
// refused request is retried after saveRetry.
func (s *Scene) scheduleSave() {
	if s.pendingSave == nil || s.now().Before(s.nextSaveTry) {
		return
	}

	id := s.store.SaveGame(*s.pendingSave, s.pendingReplay)
	if id == 0 {
		s.nextSaveTry = s.now().Add(saveRetry)

		return
	}

	s.saveReq = id
}

// applySaved handles the answer to SaveGame. A failure keeps the result
// queued for another try, so play can continue.
func (s *Scene) applySaved(res Result) {
	if res.ID != s.saveReq {
		return
	}

	if res.Err != nil {
		s.setStatus("SCORE NOT SAVED - RETRYING")
		s.nextSaveTry = s.now().Add(saveRetry)

		return
	}

	s.pendingSave = nil
	s.pendingReplay = nil
	s.setStatus("SCORE SAVED")
}

// savePoint stores a resumable snapshot, or clears the slot when the
// current run cannot resume.
func (s *Scene) savePoint() {
	if s.store == nil {
		return
	}

	snap, ok := s.model.SavePoint()
	if !ok {
		s.clearResume()

		return
	}

	s.snapReq = s.store.SaveSnapshot(snap)
}

// clearResume drops the resume slot.
func (s *Scene) clearResume() {
	if s.store == nil {
		return
	}

	s.snapReq = s.store.ClearSnapshot()
}
