package runner

import "github.com/spf13/cobra"

type ServerRunOptions struct {
	V        uint32
	Endpoint string
	Path     string
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{}
}

func (s *ServerRunOptions) Flags(cmd *cobra.Command) {
	fs := cmd.Flags()

	fs.Uint32VarP(&s.V, "verbosity", "v", 0, "logger verbosity level")
	fs.StringVarP(&s.Endpoint, "endpoint", "e", "", "etcd configuration center endpoint")
	fs.StringVarP(&s.Path, "path", "p", "", "etcd configuration center service config key")
}
