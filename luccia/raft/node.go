package raft

import (
	"sync"
	"time"

	"math/rand"
)

// NodeState represents the state of a Raft node
type NodeState int

const (
	Follower NodeState = iota
	Candidate
	Leader
)

// Config holds configuration for a Raft node
type Config struct {
	ID                string
	ElectionTimeout   time.Duration
	HeartbeatInterval time.Duration
	Peers             []string
}

// Node represents a single node in the Raft cluster
type Node struct {
	mu sync.Mutex

	// Node configuration
	config Config

	// Persistent state
	currentTerm uint64
	votedFor    string
	log         []LogEntry

	// Volatile state
	state       NodeState
	leaderId    string
	commitIndex uint64
	lastApplied uint64

	// Leader volatile state
	nextIndex  map[string]uint64
	matchIndex map[string]uint64

	// Channels for communication
	electionTimer  *time.Timer
	heartbeatTimer *time.Timer
}

// LogEntry represents a single entry in the Raft log
type LogEntry struct {
	Term    uint64
	Index   uint64
	Command interface{}
}

// NewNode creates a new Raft node
func NewNode(config Config) *Node {
	n := &Node{
		config:     config,
		state:      Follower,
		nextIndex:  make(map[string]uint64),
		matchIndex: make(map[string]uint64),
	}

	// Initialize timers
	n.electionTimer = time.NewTimer(randomTimeout(config.ElectionTimeout))
	n.heartbeatTimer = time.NewTimer(config.HeartbeatInterval)

	return n
}

// Start begins the Raft node operation
func (n *Node) Start() {
	go n.run()
}

func (n *Node) run() {
	for {
		switch n.state {
		case Follower:
			n.runFollower()
		case Candidate:
			n.runCandidate()
		case Leader:
			n.runLeader()
		}
	}
}

// randomTimeout returns a random duration between timeout and 2*timeout
func randomTimeout(timeout time.Duration) time.Duration {
	return timeout + time.Duration(rand.Int63n(int64(timeout)))
}

// becomeFollower transitions the node to follower state
func (n *Node) becomeFollower(term uint64) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.state = Follower
	n.currentTerm = term
	n.votedFor = ""
	n.leaderId = ""

	// Reset election timer
	n.resetElectionTimer()
}

func (n *Node) resetElectionTimer() {
	panic("unimplemented")
}

// becomeCandidate transitions the node to candidate state
func (n *Node) becomeCandidate() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.state = Candidate
	n.currentTerm++
	n.votedFor = n.config.ID
	n.leaderId = ""

	// Reset election timer
	n.resetElectionTimer()
}

// becomeLeader transitions the node to leader state
func (n *Node) becomeLeader() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.state = Leader
	n.leaderId = n.config.ID

	// Initialize leader state
	for _, peer := range n.config.Peers {
		n.nextIndex[peer] = uint64(len(n.log))
		n.matchIndex[peer] = 0
	}

	// Stop election timer and start heartbeat timer
	n.electionTimer.Stop()
	n.resetHeartbeatTimer()
}

func (n *Node) resetHeartbeatTimer() {
	panic("unimplemented")
}

func (n *Node) runFollower() {
	for {
		select {
		case <-n.electionTimer.C:
			n.becomeCandidate()
			return
		}
	}
}

func (n *Node) runCandidate() {
	for {
		select {
		case <-n.electionTimer.C:
			n.becomeCandidate()
			return
		}
	}
}

func (n *Node) runLeader() {
	for {
		select {
		case <-n.heartbeatTimer.C:
			n.sendHeartbeats()
		}
	}
}

func (n *Node) sendHeartbeats() {
	panic("unimplemented")
}
