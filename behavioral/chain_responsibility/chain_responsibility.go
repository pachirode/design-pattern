package chain_responsibility

import "fmt"

type IStation interface {
	SetNext(station IStation) IStation
	Approve(speed int)
}

type Station struct {
	limit int
	next  IStation
}

func NewStation(limit int) *Station {
	return &Station{
		limit: limit,
	}
}

func (s *Station) SetNext(station IStation) IStation {
	s.next = station
	return station
}

func (s *Station) Approve(speed int) {

}

type StartStationReferee struct {
	*Station
}

func NewStartStationReferee() *StartStationReferee {
	return &StartStationReferee{
		Station: NewStation(0),
	}
}

func (s *StartStationReferee) Approve(speed int) {
	if speed > s.limit {
		fmt.Println("抢跑取消成绩")
	} else if s.next != nil {
		s.next.Approve(speed)
	} else {
		fmt.Println("选手缺席")
	}
}

type EndStationReferee struct {
	*Station
}

func NewEndStationReferee() *EndStationReferee {
	return &EndStationReferee{
		Station: NewStation(0),
	}
}

func (s *EndStationReferee) Approve(speed int) {
	if speed <= s.limit {
		fmt.Println("未按时到达终点")
	} else {
		fmt.Println("成绩正常")
	}
}
