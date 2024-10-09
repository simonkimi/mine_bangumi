package job

type ChangeNotifier struct {
	NewJob     <-chan Job
	RemovedJob <-chan Job
	UpdatedJob <-chan Job

	newJob     chan Job
	removedJob chan Job
	updatedJob chan Job
}

func newChangeNotifier() *ChangeNotifier {
	c := &ChangeNotifier{
		newJob:     make(chan Job, 100),
		removedJob: make(chan Job, 100),
		updatedJob: make(chan Job, 100),
	}

	c.NewJob = c.newJob
	c.RemovedJob = c.removedJob
	c.UpdatedJob = c.updatedJob

	return c
}

func (c *ChangeNotifier) close() {
	close(c.newJob)
	close(c.removedJob)
	close(c.updatedJob)
}

func (m *Manager) notifyJobUpdate(j *Job) {
	if j.Status == StatusCancelled || j.Status == StatusFinished {
		return
	}
	for _, listener := range m.listeners {
		select {
		case listener.updatedJob <- *j:
		default:
		}
	}
}

func (m *Manager) notifyJobAdd(j *Job) {
	for _, listener := range m.listeners {
		select {
		case listener.newJob <- *j:
		default:
		}
	}
}

func (m *Manager) notifyJobRemove(j *Job) {
	for _, listener := range m.listeners {
		select {
		case listener.removedJob <- *j:
		default:
		}
	}
}
