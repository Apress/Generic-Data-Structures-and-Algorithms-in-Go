// Discrete event simulation of waiting line
package main 

import (
	"math/rand"
	"math"
	"fmt"
	"time"
	"example.com/nodequeue"
)

const (
	arrivalRate = 0.25 // average customer arrivals per minute
	lowerBoundServiceTime = 0.5 / arrivalRate 
	upperBoundServicetime = 2.0 / arrivalRate
	quitTime = 480 // Minutes in an 8 hour day
)

func InterArrivalInterval(arrivalRate float64) float64 {
	// Models a Poisson process and returns
	rn := rand.Float64() // random float between 0.0 and 1
	return -math.Log(1.0 - rn) / arrivalRate
}

func ServiceTime() float64 {
	// Uniform distribution
	rn := rand.Float64() // rn between 0.0 and 1.0
	return lowerBoundServiceTime + (upperBoundServicetime - lowerBoundServiceTime) * rn
}

type Customer struct {
	arrivalTime float64 
	serviceDuration float64 
}

// ADT for Statistics
type Statistics struct {
	waitTimes []float64 
	queueTime float64 // Accumulated time * queue size
	longestQueue int
	longestWaitTime float64
}

func (s *Statistics) AddWaitTime(wait float64) {
	s.waitTimes = append(s.waitTimes, wait)
	if wait > s.longestWaitTime {
		s.longestWaitTime = wait
	}
}

func (s *Statistics) AddQueueSizeTime(queueSize int, timeAtSize float64) {
	s.queueTime += float64(queueSize) * timeAtSize
}

func (s *Statistics) AddLength(length int) {
	if length > s.longestQueue {
		s.longestQueue = length
	}
}

var lastArrivalTime, departureTime, lastEventTime float64

func main() {
	rand.Seed(time.Now().UnixNano())
	lastEventTime := 0.0 // beginning of day 
	line := nodequeue.Queue[Customer]{}
	statistics := Statistics{}
	// Start simulation 
	for {
		lastArrivalTime = lastArrivalTime + InterArrivalInterval(arrivalRate)
		if lastArrivalTime > quitTime {
			break
		}
		if line.Size() == 0 {
			lastEventTime = lastArrivalTime
			serviceTime := ServiceTime()
			customer := Customer{lastArrivalTime, serviceTime}
			line.Insert(customer)
			statistics.AddLength(line.Size())
			departureTime = lastArrivalTime + serviceTime
		} else {
			if lastArrivalTime < departureTime { // next event is an arrival
				customer := Customer{lastArrivalTime, ServiceTime()}
				statistics.AddQueueSizeTime(line.Size(), lastArrivalTime -  lastEventTime)
				lastEventTime = lastArrivalTime
				line.Insert(customer)
				statistics.AddLength(line.Size())
			} else { // next event is a departure
				statistics.AddQueueSizeTime(line.Size(), departureTime - lastEventTime)
				departingCustomer := line.Remove()
				statistics.AddWaitTime(departureTime - departingCustomer.arrivalTime)
				lastEventTime = departureTime
				if line.Size() > 0 {
					departureTime = lastEventTime + line.First().serviceDuration
				}
			}
		}
	}
	totalWaitTime := 0.0 
	for i := 0; i < len(statistics.waitTimes); i++ {
		totalWaitTime += statistics.waitTimes[i]
	}
	averageWaitTime := totalWaitTime / float64(len(statistics.waitTimes))
	fmt.Printf("\nAverage Time from Arrival to Departure: %0.2f minutes", averageWaitTime)
	fmt.Printf("\nAverage size of waiting line: %0.2f", statistics.queueTime / lastEventTime)
	fmt.Printf("\nLongest queue during the day: %d", statistics.longestQueue)
	fmt.Printf("\nLongest wait time during the day: %0.2f minutes", statistics.longestWaitTime)
}
/* An output
Average Time from Arrival to Departure: 14.70 minutes
Average size of waiting line: 2.75
Longest queue during the day: 14
Longest wait time during the day: 41.03 minutes
*/

