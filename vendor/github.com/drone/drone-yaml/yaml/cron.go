package yaml

import "errors"

type (
	// Cron is a resource that defines a cron job, used
	// to execute pipelines at scheduled intervals.
	Cron struct {
		Kind string `json:"kind,omitempty"`
		Type string `json:"type,omitempty"`
		Name string `json:"name,omitempty"`

		Spec CronSpec `json:"spec,omitempty"`
	}

	// CronSpec defines the cron job.
	CronSpec struct {
		Schedule string         `json:"schedule,omitempty"`
		Branch   string         `json:"branch,omitempty"`
		Deploy   CronDeployment `json:"deployment,omitempty" yaml:"deployment"`
	}

	// CronDeployment defines a cron job deployment.
	CronDeployment struct {
		Target string `json:"target,omitempty"`
	}
)

// Validate returns an error if the cron is invalid.
func (c Cron) Validate() error {
	switch {
	case c.Spec.Branch == "":
		return errors.New("yaml: invalid cron branch")
	default:
		return nil
	}
}