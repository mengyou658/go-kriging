package helper

import "sync"

type DistanceRun struct {
	distanceChan chan [][2]float64
	doneChan     chan struct{}
	n            sync.WaitGroup
}

//
//func (self DistanceRun) distanceRun(fn func()) [][2]float64 {
//
//}
//
//func distanceRun(fn func()) [][2]float64 {
//	distanceChan := make(chan [][2]float64)
//	var doneChan = make(chan struct{})
//	var nGroup sync.WaitGroup
//
//
//}
