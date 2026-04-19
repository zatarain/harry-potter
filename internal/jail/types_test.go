package jail

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStateConstants(t *testing.T) {
	cases := map[string]State{
		"running":  StateRunning,
		"stopped":  StateStopped,
		"creating": StateCreating,
		"removing": StateRemoving,
		"paused":   StatePaused,
		"exited":   StateExited,
		"error":    StateError,
	}
	for want, got := range cases {
		if string(got) != want {
			t.Errorf("State constant = %q, want %q", got, want)
		}
	}
}

func TestRestartPolicyConstants(t *testing.T) {
	cases := map[string]RestartPolicy{
		"no":             RestartNo,
		"always":         RestartAlways,
		"on-failure":     RestartOnFailure,
		"unless-stopped": RestartUnlessStopped,
	}
	for want, got := range cases {
		if string(got) != want {
			t.Errorf("RestartPolicy constant = %q, want %q", got, want)
		}
	}
}

func TestJailJSONRoundtrip(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	j := Jail{
		ID:            "abc123",
		Name:          "test-jail",
		Image:         "FreeBSD:14.0",
		Hostname:      "test",
		State:         StateRunning,
		Command:       []string{"/bin/sh"},
		RestartPolicy: RestartAlways,
		Dataset:       "zroot/jails/test-jail",
		ConfigPath:    "/var/db/harry-potter/test-jail.json",
		CreatedAt:     now,
	}

	data, err := json.Marshal(j)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var decoded Jail
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if decoded.ID != j.ID {
		t.Errorf("ID = %q, want %q", decoded.ID, j.ID)
	}
	if decoded.State != j.State {
		t.Errorf("State = %q, want %q", decoded.State, j.State)
	}
	if decoded.RestartPolicy != j.RestartPolicy {
		t.Errorf("RestartPolicy = %q, want %q", decoded.RestartPolicy, j.RestartPolicy)
	}
}

func TestJailOmitEmptyFields(t *testing.T) {
	j := Jail{
		ID:        "x",
		Name:      "minimal",
		Image:     "FreeBSD:14.0",
		Hostname:  "minimal",
		State:     StateStopped,
		Command:   []string{"/bin/sh"},
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(j)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	for _, field := range []string{"jid", "exit_code", "entrypoint", "user", "work_dir",
		"env", "labels", "mounts", "networks", "ports", "started_at", "finished_at"} {
		if _, present := raw[field]; present {
			t.Errorf("field %q should be omitted when zero, but it was present in JSON", field)
		}
	}
}

func TestNetworkJSONRoundtrip(t *testing.T) {
	n := Network{
		ID:     "net1",
		Name:   "bridge0",
		Driver: "bridge",
		Subnet: "10.0.0.0/24",
		IPv6:   false,
	}

	data, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var decoded Network
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if decoded.Name != n.Name {
		t.Errorf("Name = %q, want %q", decoded.Name, n.Name)
	}
	if decoded.Driver != n.Driver {
		t.Errorf("Driver = %q, want %q", decoded.Driver, n.Driver)
	}
}

func TestVolumeJSONRoundtrip(t *testing.T) {
	v := Volume{
		Name:       "data",
		Driver:     "local",
		Mountpoint: "/var/db/harry-potter/volumes/data",
		Scope:      "local",
	}

	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var decoded Volume
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if decoded.Name != v.Name {
		t.Errorf("Name = %q, want %q", decoded.Name, v.Name)
	}
	if decoded.Mountpoint != v.Mountpoint {
		t.Errorf("Mountpoint = %q, want %q", decoded.Mountpoint, v.Mountpoint)
	}
}
