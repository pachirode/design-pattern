package template_method

import "fmt"

type Cooker interface {
	Fire()
	AddCondiment()
	Cooking()
	Finish()
}

type ConcreteCooker struct {
}

func (cc *ConcreteCooker) Fire() {
	fmt.Println("开火")
}

func (cc *ConcreteCooker) AddCondiment() {
	fmt.Println("加佐料")
}

func (cc *ConcreteCooker) Cooking() {
	fmt.Println("开始烹饪")
}

func (cc *ConcreteCooker) Finish() {
	fmt.Println("完成")
}

type Scrambled struct {
	ConcreteCooker
}

func NewScrambled() *Scrambled {
	return &Scrambled{}
}

func (s *Scrambled) AddCondiment() {
	fmt.Println("添加鸡蛋")
}

func DoMake(cooker Cooker) {
	cooker.Fire()
	cooker.AddCondiment()
	cooker.Cooking()
	cooker.Finish()
}
