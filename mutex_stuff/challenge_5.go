/*
Challenge 5: The Ride-Hailing Matchmaker (Expert / Fine-Grained Locking)
Now for the final boss. We are moving from simple resource coordination to high-throughput system optimization.

The Scenario
You are building the matching engine for a ride-hailing app like Bolt or Uber in Lagos during peak Friday traffic. You have thousands of active drivers updating their locations, and thousands of riders requesting rides at the exact same time.

If you wrap your entire drivers map in one massive struct with a single sync.RWMutex, every time a rider locks the map to find a driver, all other riders and drivers in the entire city get blocked from making updates. This is called Coarse-Grained Locking, and it creates massive performance bottlenecks.

To fix this, we use Fine-Grained Locking (or Sharding/Partitioning), where instead of locking the whole universe, we break the data down so threads only lock the specific slice they actually need.

Your Task
Build a Matchmaker system that handles a map of drivers.

Structure it so that finding or updating a driver in one area (e.g., Ikeja) does not block or lock out a driver update happening in another area (e.g., Lekki).

Implement thread-safe methods to UpdateDriverLocation and RequestRide.
*/