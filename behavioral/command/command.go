package command

type Command interface {
	Execute() string
}

type Runner struct {
}

func (r *Runner) Run() string {
	return "run"
}

func (r *Runner) Stop() string {
	return "stop"
}

type RunCommand struct {
	runner *Runner
}

func (rc *RunCommand) Execute() string {
	return rc.runner.Run()
}

type StopCommand struct {
	runner *Runner
}

func (sc *StopCommand) Execute() string {
	return sc.runner.Stop()
}

type Invoker struct {
	command Command
}

func (i *Invoker) Execute() string {
	return i.command.Execute()
}
