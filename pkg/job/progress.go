package job

import "sync"

type Progress struct {
	processed int
	total     int

	job     *Job
	manager *Manager

	mux sync.Mutex
}

func (m *Manager) newProgress(job *Job) *Progress {
	return &Progress{
		manager:   m,
		job:       job,
		total:     -1,
		processed: -1,
	}
}

func (p *Progress) Percent() float64 {
	if p.total <= 0 {
		return -1
	}

	percent := float64(p.processed) / float64(p.total)
	if percent > 1 {
		percent = 1
	}

	return percent
}

func (p *Progress) SetTotal(total int) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.total = total
}

func (p *Progress) notifyListeners() {

}
