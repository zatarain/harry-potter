package jail

import "time"

type State string

const (
	StateRunning  State = "running"
	StateStopped  State = "stopped"
	StateCreating State = "creating"
	StateRemoving State = "removing"
	StatePaused   State = "paused"
	StateExited   State = "exited"
	StateError    State = "error"
)

type RestartPolicy string

const (
	RestartNo            RestartPolicy = "no"
	RestartAlways        RestartPolicy = "always"
	RestartOnFailure     RestartPolicy = "on-failure"
	RestartUnlessStopped RestartPolicy = "unless-stopped"
)

// Jail is the runtime representation of a FreeBSD jail.
type Jail struct {
	ID            string            `json:"id"`
	JID           int               `json:"jid,omitempty"`
	Name          string            `json:"name"`
	Image         string            `json:"image"`
	Hostname      string            `json:"hostname"`
	State         State             `json:"state"`
	ExitCode      *int              `json:"exit_code,omitempty"`
	Command       []string          `json:"command"`
	Entrypoint    []string          `json:"entrypoint,omitempty"`
	User          string            `json:"user,omitempty"`
	WorkDir       string            `json:"work_dir,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Mounts        []Mount           `json:"mounts,omitempty"`
	Networks      []NetworkEndpoint `json:"networks,omitempty"`
	Ports         []PortBinding     `json:"ports,omitempty"`
	Resources     ResourceConfig    `json:"resources"`
	RestartPolicy RestartPolicy     `json:"restart_policy"`
	Dataset       string            `json:"dataset"`
	ConfigPath    string            `json:"config_path"`
	CreatedAt     time.Time         `json:"created_at"`
	StartedAt     *time.Time        `json:"started_at,omitempty"`
	FinishedAt    *time.Time        `json:"finished_at,omitempty"`
}

// Mount represents a filesystem mount inside a jail (nullfs or ZFS).
type Mount struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Type        string `json:"type"`
	ReadOnly    bool   `json:"read_only"`
}

// NetworkEndpoint describes a jail's attachment to a virtual network.
type NetworkEndpoint struct {
	NetworkID   string   `json:"network_id"`
	NetworkName string   `json:"network_name"`
	IPAddress   string   `json:"ip_address"`
	IPv6Address string   `json:"ipv6_address,omitempty"`
	MacAddress  string   `json:"mac_address,omitempty"`
	Aliases     []string `json:"aliases,omitempty"`
}

// PortBinding maps a jail-internal port to a host port via pf(4).
type PortBinding struct {
	HostIP   string `json:"host_ip,omitempty"`
	HostPort string `json:"host_port"`
	JailPort string `json:"jail_port"`
	Protocol string `json:"protocol"`
}

// ResourceConfig mirrors RCTL-based resource limits.
type ResourceConfig struct {
	CPUs       string `json:"cpus,omitempty"`
	CPUSetCPUs string `json:"cpuset_cpus,omitempty"`
	Memory     string `json:"memory,omitempty"`
	MemorySwap string `json:"memory_swap,omitempty"`
	PidsLimit  int64  `json:"pids_limit,omitempty"`
}

// Image is a FreeBSD base system template or a custom-built jail image.
type Image struct {
	ID         string            `json:"id"`
	Repository string            `json:"repository"`
	Tag        string            `json:"tag"`
	Digest     string            `json:"digest,omitempty"`
	Size       int64             `json:"size"`
	Labels     map[string]string `json:"labels,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	Layers     []string          `json:"layers,omitempty"`
}

// Network is a virtual network backed by epair(4) and bridge(4).
type Network struct {
	ID        string                     `json:"id"`
	Name      string                     `json:"name"`
	Driver    string                     `json:"driver"`
	Subnet    string                     `json:"subnet,omitempty"`
	Gateway   string                     `json:"gateway,omitempty"`
	IPv6      bool                       `json:"ipv6"`
	Internal  bool                       `json:"internal"`
	Labels    map[string]string          `json:"labels,omitempty"`
	Jails     map[string]NetworkEndpoint `json:"jails,omitempty"`
	CreatedAt time.Time                  `json:"created_at"`
}

// Volume is a named storage resource (nullfs directory or ZFS dataset).
type Volume struct {
	Name       string            `json:"name"`
	Driver     string            `json:"driver"`
	Mountpoint string            `json:"mountpoint"`
	Options    map[string]string `json:"options,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
	Scope      string            `json:"scope"`
	CreatedAt  time.Time         `json:"created_at"`
}
