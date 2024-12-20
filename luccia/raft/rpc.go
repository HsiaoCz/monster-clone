package raft

// RequestVoteArgs represents the arguments for a RequestVote RPC
type RequestVoteArgs struct {
	Term         uint64
	CandidateId  string
	LastLogIndex uint64
	LastLogTerm  uint64
}

// RequestVoteReply represents the response from a RequestVote RPC
type RequestVoteReply struct {
	Term        uint64
	VoteGranted bool
}

// AppendEntriesArgs represents the arguments for an AppendEntries RPC
type AppendEntriesArgs struct {
	Term         uint64
	LeaderId     string
	PrevLogIndex uint64
	PrevLogTerm  uint64
	Entries      []LogEntry
	LeaderCommit uint64
}

// AppendEntriesReply represents the response from an AppendEntries RPC
type AppendEntriesReply struct {
	Term    uint64
	Success bool
}

// RequestVote handles the RequestVote RPC
func (n *Node) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	reply.Term = n.currentTerm
	reply.VoteGranted = false

	// If request term is less than current term, reject
	if args.Term < n.currentTerm {
		return nil
	}

	// If request term is greater than current term, become follower
	if args.Term > n.currentTerm {
		n.becomeFollower(args.Term)
	}

	// If we haven't voted for anyone yet or already voted for this candidate
	if n.votedFor == "" || n.votedFor == args.CandidateId {
		// Check if candidate's log is at least as up-to-date as ours
		lastLogIndex := uint64(len(n.log) - 1)
		var lastLogTerm uint64
		if lastLogIndex >= 0 {
			lastLogTerm = n.log[lastLogIndex].Term
		}

		if args.LastLogTerm > lastLogTerm ||
			(args.LastLogTerm == lastLogTerm && args.LastLogIndex >= lastLogIndex) {
			reply.VoteGranted = true
			n.votedFor = args.CandidateId
			n.resetElectionTimer()
		}
	}

	return nil
}

// AppendEntries handles the AppendEntries RPC
func (n *Node) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	reply.Term = n.currentTerm
	reply.Success = false

	// If request term is less than current term, reject
	if args.Term < n.currentTerm {
		return nil
	}

	// If request term is greater than current term, become follower
	if args.Term > n.currentTerm {
		n.becomeFollower(args.Term)
	}

	// Reset election timer since we received a valid AppendEntries
	n.resetElectionTimer()

	// Reply false if log doesn't contain an entry at prevLogIndex whose term matches prevLogTerm
	if args.PrevLogIndex >= uint64(len(n.log)) {
		return nil
	}

	if args.PrevLogIndex >= 0 && n.log[args.PrevLogIndex].Term != args.PrevLogTerm {
		return nil
	}

	// If existing entries conflict with new entries, delete all existing entries starting with first conflicting entry
	newEntries := make([]LogEntry, len(args.Entries))
	copy(newEntries, args.Entries)

	logInsertIndex := args.PrevLogIndex + 1
	newEntriesIndex := 0

	for logInsertIndex < uint64(len(n.log)) && newEntriesIndex < len(newEntries) {
		if n.log[logInsertIndex].Term != newEntries[newEntriesIndex].Term {
			n.log = n.log[:logInsertIndex]
			break
		}
		logInsertIndex++
		newEntriesIndex++
	}

	// Append any new entries not already in the log
	if newEntriesIndex < len(newEntries) {
		n.log = append(n.log, newEntries[newEntriesIndex:]...)
	}

	// Update commit index if leader commit is greater than current commit index
	if args.LeaderCommit > n.commitIndex {
		n.commitIndex = min(args.LeaderCommit, uint64(len(n.log)-1))
	}

	reply.Success = true
	return nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
