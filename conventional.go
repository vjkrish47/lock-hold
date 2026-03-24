type Participant struct {
	id int
	lock sync.Mutex
}
func prepare(p *Participant, ready chan bool) {
	p.lock.Lock()
	time.Sleep(30 * time.Millisecond)
	ready <- true
}
func commit(p *Participant, done chan bool, start time.Time) {
	time.Sleep(20 * time.Millisecond)
	p.lock.Unlock()
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("Participant", p.id, "lock hold time(ms)", elapsed)
	done <- true
}
func main() {
	participants := []*Participant{
		{1, sync.Mutex{}},
		{2, sync.Mutex{}},
		{3, sync.Mutex{}},
	}
	ready := make(chan bool)
	done := make(chan bool)
	start := time.Now()
	for _, p := range participants {
		go prepare(p, ready)
	}
	for range participants {
		<-ready
	}
	for _, p := range participants {
		go commit(p, done, start)
	}
	for range participants {
		<-done
	}
}
