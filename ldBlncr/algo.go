package ldBlncr

func (l *loadBalancer) algoRoundRobin() Server {
	var s Server

	for {
		if l.servers[l.roundRobinCount].IsAlive() {
			s = l.servers[l.roundRobinCount]
		}
		l.roundRobinCount = (1 + l.roundRobinCount) % len(l.servers)

		if s != nil {
			break
		}
	}

	return s
}
