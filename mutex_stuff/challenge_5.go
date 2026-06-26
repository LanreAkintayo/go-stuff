/*
Challenge 5: The Ride-Hailing Matchmaker (Expert / Fine-Grained Locking)
We are moving to a high-throughput system optimization.

The Scenario
You are building the matching engine for a ride-hailing app like Bolt or Uber in Lagos during peak Friday traffic. You have thousands of active drivers updating their locations, and thousands of riders requesting rides at the exact same time.

If you wrap your entire drivers map in one massive struct with a single sync.RWMutex, every time a rider locks the map to find a driver, all other riders and drivers in the entire city get blocked from making updates. This is called Coarse-Grained Locking, and it creates massive performance bottlenecks.

To fix this, we use Fine-Grained Locking (or Sharding/Partitioning), where instead of locking the whole universe, we break the data down so threads only lock the specific slice they actually need.

Your Task
Build a Matchmaker system that handles a map of drivers.

Structure it so that finding or updating a driver in one area (e.g., Ikeja) does not block or lock out a driver update happening in another area (e.g., Lekki).

Implement thread-safe methods to UpdateDriverLocation and RequestRide.
*/

package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"sync"
)

type Driver struct {
	driverId    string
	location    string
	isAvailable bool
}

type Shard struct {
	areaDrivers map[string]map[string]*Driver // location -> driverId --> Driver
	keyLock     sync.RWMutex
}

type MatchMaker struct {
	shards     []*Shard
	noOfShards int
}

func getShardIndex(location string, noOfShards int) int {
	// Initialize our FNV-1a hashing engine
	h := fnv.New32a()

	// We convert our string location to bytes array and then feed the hashing engine
	h.Write([]byte(location))

	// We execute the hash function that returns a 32-bit unsigned integer
	locationUint := h.Sum32()

	return int(locationUint % uint32(noOfShards))
}

func (m *MatchMaker) UpdateDriverLocation(location string, newLocation string, driverId string) {
	// Get driver shard
	oldShardIndex := getShardIndex(location, m.noOfShards)
	newShardIndex := getShardIndex(newLocation, m.noOfShards)

	oldShard := m.shards[oldShardIndex]
	newShard := m.shards[newShardIndex]

	// The two locations can be in the same shard.
	if oldShardIndex == newShardIndex {
		oldShard.keyLock.Lock()
		defer oldShard.keyLock.Unlock()

		driver, exist := oldShard.areaDrivers[location][driverId]

		if exist {
			delete(oldShard.areaDrivers[location], driverId)
			driver.location = newLocation

			innerMap := oldShard.areaDrivers[newLocation]
			if innerMap == nil {
				innerMap = make(map[string]*Driver)
				oldShard.areaDrivers[newLocation] = innerMap
			}
			innerMap[driverId] = driver

		}
		return
	}
	oldShard.keyLock.Lock()

	driver, exist := oldShard.areaDrivers[location][driverId]

	if exist {
		driver.location = newLocation
		delete(oldShard.areaDrivers[location], driverId)
	}
	oldShard.keyLock.Unlock()

	if exist {
		newShard.keyLock.Lock()

		if newShard.areaDrivers[newLocation] == nil {
			newShard.areaDrivers[newLocation] = make(map[string]*Driver)
		}
		newShard.areaDrivers[newLocation][driverId] = driver

		newShard.keyLock.Unlock()
	}

}

func (m *MatchMaker) RequestRide(fromLocation string) (*Driver, error) {

	shardIndex := getShardIndex(fromLocation, m.noOfShards)
	currentLocationShard := m.shards[shardIndex]

	currentLocationShard.keyLock.Lock()
	defer currentLocationShard.keyLock.Unlock()

	drivers, exist := currentLocationShard.areaDrivers[fromLocation]

	if exist {
		for _, driver := range drivers {
			if driver.isAvailable {
				driver.isAvailable = false
				return driver, nil
			}
		}

		return nil, errors.New("Driver is not available in location")

	} else {
		return nil, errors.New("Location does not exist")
	}
}

func Challenge5() {
	// Initialize a match maker with 5 distinct shards
	noOfShards := 5
	matchMaker := MatchMaker{
		noOfShards: noOfShards,
		shards:     make([]*Shard, 5),
	}

	// Explicitly allocate memory for every single shard and its inner map
	for i := range matchMaker.shards {
		matchMaker.shards[i] = &Shard{areaDrivers: make(map[string]map[string]*Driver)}
	}

	// Let's seed an initial driver directly into the system state
	newDriver := Driver{
		driverId:    "001",
		location:    "Ikeja",
		isAvailable: true,
	}
	ikejaIndex := getShardIndex("Ikeja", noOfShards)
	ikejaShard := matchMaker.shards[ikejaIndex]
	ikejaShard.areaDrivers["Ikeja"] = make(map[string]*Driver)
	ikejaShard.areaDrivers["Ikeja"]["001"] = &newDriver

	fmt.Println("Driver 001 is available in Ikeja")

	var goTrackers sync.WaitGroup
	goTrackers.Add(2) // Wait for 2 goroutines

	// First one will request a ride in Ikeja
	go func() {
		defer goTrackers.Done()
		fmt.Println("Requesting a ride...")
		driverPointer, err := matchMaker.RequestRide("Ikeja")
		if driverPointer != nil {
			fmt.Printf("Driver %s requested successfully\n", driverPointer.driverId)
		} else {
			fmt.Println("Error encountered: ", err)
		}
	}()

	// Second one will update driver location from "Ikeja" to "Lagos"
	go func() {
		defer goTrackers.Done()
		fmt.Println("Changing location from Ikeja to Lekki")
		matchMaker.UpdateDriverLocation("Ikeja", "Lekki", "001")
		fmt.Println("Location updated successfully!")
	}()

	goTrackers.Wait()

	fmt.Println("Done!")

}
