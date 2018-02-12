package store

import (
	"log"
	"time"

	"github.com/chrislusf/vasto/pb"
	"github.com/chrislusf/vasto/storage/codec"
)

func (ss *storeServer) processMerge(shard *shard, mergeRequest *pb.MergeRequest) *pb.WriteResponse {

	key := mergeRequest.Key
	nowInNano := uint64(time.Now().UnixNano())
	entry := codec.NewMergeEntry(mergeRequest, nowInNano)

	resp := &pb.WriteResponse{
		Ok: true,
	}

	// log.Printf("shard %d put key: %v\n", shard.id, string(mergeRequest.KeyValue.Key))

	err := shard.db.Merge(key, entry.ToBytes())
	if err != nil {
		resp.Ok = false
		resp.Status = err.Error()
	} else {
		shard.logMerge(mergeRequest, nowInNano)
	}

	return resp
}

func (s *shard) logMerge(mergeRequest *pb.MergeRequest, updatedAtNs uint64) {

	// println("logMerge1", mergeRequest.String())

	if s.lm == nil {
		return
	}

	// println("logMerge2", mergeRequest.String())

	err := s.lm.AppendEntry(&pb.LogEntry{
		UpdatedAtNs: updatedAtNs,
		Merge:       mergeRequest,
	})

	if err != nil {
		log.Printf("append put log entry: %v", err)
	}

	// println("logMerge3", mergeRequest.String())

}
