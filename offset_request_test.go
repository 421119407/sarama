package sarama

import (
	"testing"

	"go.uber.org/goleak"
)

var (
	offsetRequestNoBlocksV1 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00,
	}

	offsetRequestNoBlocksV2 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00,
		0x00,
	}

	offsetRequestOneBlock = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x03, 'f', 'o', 'o',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
	}

	offsetRequestOneBlockV1 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x03, 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}

	offsetRequestOneBlockReadCommittedV2 = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x01, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x03, 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}

	offsetRequestReplicaID = []byte{
		0x00, 0x00, 0x00, 0x2a,
		0x00, 0x00, 0x00, 0x00,
	}

	offsetRequestV4 = []byte{
		0xff, 0xff, 0xff, 0xff, // replicaID
		0x01, // IsolationLevel
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x04,
		0x64, 0x6e, 0x77, 0x65, // topic name
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x09, // partitionID
		0xff, 0xff, 0xff, 0xff, // leader epoch
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // timestamp
	}
)

func TestOffsetRequest(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	request := new(OffsetRequest)
	testRequest(t, "no blocks", request, offsetRequestNoBlocksV1)

	request.AddBlock("foo", 4, 1, 2)
	testRequest(t, "one block", request, offsetRequestOneBlock)
}

func TestOffsetRequestV1(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	request := new(OffsetRequest)
	request.Version = 1
	testRequest(t, "no blocks", request, offsetRequestNoBlocksV1)

	request.AddBlock("bar", 4, 1, 2) // Last argument is ignored for V1
	testRequest(t, "one block", request, offsetRequestOneBlockV1)
}

func TestOffsetRequestV2(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	request := new(OffsetRequest)
	request.Version = 2
	testRequest(t, "no blocks", request, offsetRequestNoBlocksV2)

	request.IsolationLevel = ReadCommitted
	request.AddBlock("bar", 4, 1, 2) // Last argument is ignored for V1
	testRequest(t, "one block", request, offsetRequestOneBlockReadCommittedV2)
}

func TestOffsetRequestReplicaID(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	request := new(OffsetRequest)
	replicaID := int32(42)
	request.SetReplicaID(replicaID)

	if found := request.ReplicaID(); found != replicaID {
		t.Errorf("replicaID: expected %v, found %v", replicaID, found)
	}

	testRequest(t, "with replica ID", request, offsetRequestReplicaID)
}

func TestOffsetRequestV4(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	request := new(OffsetRequest)
	request.Version = 4
	request.IsolationLevel = ReadCommitted
	request.AddBlock("dnwe", 9, -1, -1)
	testRequest(t, "V4", request, offsetRequestV4)
}
