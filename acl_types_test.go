package sarama

import (
	"testing"

	"go.uber.org/goleak"
)

func TestAclOperationTextMarshal(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	for i := AclOperationUnknown; i <= AclOperationIdempotentWrite; i++ {
		text, err := i.MarshalText()
		if err != nil {
			t.Errorf("couldn't marshal %d to text: %s", i, err)
		}
		var got AclOperation
		err = got.UnmarshalText(text)
		if err != nil {
			t.Errorf("couldn't unmarshal %s to acl operation: %s", text, err)
		}
		if got != i {
			t.Errorf("got %d, want %d", got, i)
		}
	}
}

func TestAclPermissionTypeTextMarshal(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	for i := AclPermissionUnknown; i <= AclPermissionAllow; i++ {
		text, err := i.MarshalText()
		if err != nil {
			t.Errorf("couldn't marshal %d to text: %s", i, err)
		}
		var got AclPermissionType
		err = got.UnmarshalText(text)
		if err != nil {
			t.Errorf("couldn't unmarshal %s to acl permission: %s", text, err)
		}
		if got != i {
			t.Errorf("got %d, want %d", got, i)
		}
	}
}

func TestAclResourceTypeTextMarshal(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	for i := AclResourceUnknown; i <= AclResourceTransactionalID; i++ {
		text, err := i.MarshalText()
		if err != nil {
			t.Errorf("couldn't marshal %d to text: %s", i, err)
		}
		var got AclResourceType
		err = got.UnmarshalText(text)
		if err != nil {
			t.Errorf("couldn't unmarshal %s to acl resource: %s", text, err)
		}
		if got != i {
			t.Errorf("got %d, want %d", got, i)
		}
	}
}

func TestAclResourcePatternTypeTextMarshal(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	})
	for i := AclPatternUnknown; i <= AclPatternPrefixed; i++ {
		text, err := i.MarshalText()
		if err != nil {
			t.Errorf("couldn't marshal %d to text: %s", i, err)
		}
		var got AclResourcePatternType
		err = got.UnmarshalText(text)
		if err != nil {
			t.Errorf("couldn't unmarshal %s to acl resource pattern: %s", text, err)
		}
		if got != i {
			t.Errorf("got %d, want %d", got, i)
		}
	}
}
