package job

import "sync"

type Progress struct {
	step    int
	total   int
	percent float64

	job     *Job
	manager *Manager

	mux sync.Mutex
}

func (m *Manager) newProgress(job *Job) *Progress {
	return &Progress{
		manager: m,
		job:     job,
		total:   -1,
		step:    -1,
	}
}

func (p *Progress) SetTotal(total int) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.total = total
	p.calcPercent()
	p.update()
}

func (p *Progress) AddStep(processed int) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.step += processed
	p.calcPercent()
	p.update()
}

func (p *Progress) SetStep(processed int) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.step = processed
	p.update()
}

func (p *Progress) SetPercent(percent float64) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.percent = percent
	p.update()
}

func (p *Progress) NextStep() {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.step++
	p.calcPercent()
	p.update()
}

func (p *Progress) calcPercent() {
	if p.total <= 0 {
		p.percent = 0
		return
	}

	percent := float64(p.step) / float64(p.total)
	if percent > 1 {
		percent = 1
	}

	p.percent = percent
}

func (p *Progress) update() {
	p.job.Progress = p.percent
	p.manager.notifyJobUpdate(p.job)
}
