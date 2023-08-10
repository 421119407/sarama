package sarama

import (
	"testing"
	"time"

	"go.uber.org/goleak"
)

var deleteAclsResponse = []byte{
	0, 0, 0, 100,
	0, 0, 0, 1,
	0, 0, // no error
	255, 255, // no error message
	0, 0, 0, 1, // 1 matching acl
	0, 0, // no error
	255, 255, // no error message
	2, // resource type
	0, 5, 't', 'o', 'p', 'i', 'c',
	0, 9, 'p', 'r', 'i', 'n', 'c', 'i', 'p', 'a', 'l',
	0, 4, 'h', 'o', 's', 't',
	4,
	3,
}

func TestDeleteAclsResponse(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	resp := &DeleteAclsResponse{
		ThrottleTime: 100 * time.Millisecond,
		FilterResponses: []*FilterResponse{{
			MatchingAcls: []*MatchingAcl{{
				Resource: Resource{ResourceType: AclResourceTopic, ResourceName: "topic"},
				Acl:      Acl{Principal: "principal", Host: "host", Operation: AclOperationWrite, PermissionType: AclPermissionAllow},
			}},
		}},
	}

	testResponse(t, "", resp, deleteAclsResponse)
}
