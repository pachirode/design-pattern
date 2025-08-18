package chain_responsibility

import "testing"

func TestChainResponsibility(t *testing.T) {
	startStation := NewStartStationReferee()
	endStation := NewEndStationReferee()
	startStation.SetNext(endStation)
	startStation.Approve(10)
	startStation.Approve(0)
}
