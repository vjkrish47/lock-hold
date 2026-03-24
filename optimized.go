type Node struct {
	id int
}

func (n Node) prepare(tx int, wg *sync.WaitGroup, votes chan<- bool) {
	time.Sleep(time.Duration(10+n.id*5) * time.Millisecond)
	votes <- true
	wg.Done()
}

func commitTransaction(nodes []Node, tx int) {
	var wg sync.WaitGroup
	votes := make(chan bool, len(nodes))
	wg.Add(len(nodes))
	start := time.Now()

	for _, n := range nodes {
		go n.prepare(tx, &wg, votes)
	}

	go func() {
		wg.Wait()
		close(votes)
	}()

	ack := 0
	for v := range votes {
		if v {
			ack++
		}
	}

	if ack == len(nodes) {
		fmt.Println("Committed tx", tx, "nodes", len(nodes), "time", time.Since(start))
	}
}

func main() {
	for _, size := range []int{3, 5, 7, 9, 11} {
		nodes := []Node{}
		for i := 0; i < size; i++ {
			nodes = append(nodes, Node{id: i})
		}
		commitTransaction(nodes, 1)
	}
}
