package christmasdb

import "github.com/DevourTech/christmasdb/raft"

func Launch() {
	raft.StartServer()
}
