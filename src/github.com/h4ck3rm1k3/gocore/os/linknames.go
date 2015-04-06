package os
import "github.com/h4ck3rm1k3/gocore/run_time"

func sigpipe() {
	run_time.Os_sigpipe()
}

func run_time_args() []string {
	return run_time.Os_run_time_args()
}

