package state

import "fmt"

type State interface {
	Run(runner *Runner)
	Walk(runner *Runner)
	Stop(runner *Runner)
}

type Runner struct {
	state State
}

type Run struct {
}

type Walk struct {
}

type Stop struct {
}

func NewRunner() *Runner {
	return &Runner{state: NewStop()}
}

func (r *Runner) SetState(state State) {
	r.state = state
}

func (r *Runner) Run() {
	r.state.Run(r)
}
func (r *Runner) Walk() {
	r.state.Walk(r)
}
func (r *Runner) Stop() {
	r.state.Stop(r)
}

func NewRun() *Run {
	return &Run{}
}

func (r *Run) Stop(*Runner) {
	fmt.Println("错误，跑步不可以切换为停止")
}

func (r *Run) Run(*Runner) {
	fmt.Println("错误，已经在跑步")
}

func (r *Run) Walk(runner *Runner) {
	fmt.Println("切换为走路")
	runner.SetState(NewWalk())
}

func NewWalk() *Walk {
	return &Walk{}
}

func (w *Walk) Run(runner *Runner) {
	fmt.Println("切换为跑步")
	runner.SetState(NewRun())
}

func (w *Walk) Walk(*Runner) {
	fmt.Println("错误，正在走路")
}
func (w *Walk) Stop(*Runner) {
	fmt.Println("切换为停止")
}

func NewStop() *Stop {
	return &Stop{}
}

func (s *Stop) Run(*Runner) {
	fmt.Println("错误，停止不可以切换为跑步")
}
func (s *Stop) Walk(runner *Runner) {
	fmt.Println("切换为走路")
	runner.SetState(NewWalk())
}

func (s *Stop) Stop(*Runner) {
	fmt.Println("错误，正在停止")
}
