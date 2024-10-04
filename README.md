<h1 align="center">LEM-IN</h1>

## About The Project
This project is meant to make you code a digital version of an ant farm.

## Installation
```
git clone https://learn.zone01dakar.sn/git/alogou/lem-in.git
cd lem-in
```
## Usage
```
go run . fileName
or
go run main.go fileName
```

## Example
```
go run . example00.txt
```


## Project Features:

    Read input from a file: The program reads input data from a file that describes the layout of the ant farm, including the number of ants, the rooms, and the connections between them.

    Find the quickest paths: Lem-in employs an algorithm to determine the quickest paths that ants can take to travel from the start room to the end room while avoiding obstacles and taking into account the connectivity of the colony.

    Simulate ant movement: Once the optimal paths are found, the program simulates the movement of the ants from room to room, step by step, while ensuring that each room is occupied by only one ant at a time.

    Display simulation results: The program displays the content of the input file and showcases each step of ant movement from room to room. It provides a visual representation of how the ants navigate through the colony to reach the end room.

## Rules and Challenges:

    Ant placement: All ants start in the designated start room (##start). The objective is to guide them to the end room (##end) using the shortest paths.

    Pathfinding complexity: The shortest path is not always the simplest, and some colonies may have intricate networks of rooms and tunnels.

    Invalid data handling: The program must handle various scenarios, such as colonies with no valid path between start and end rooms, rooms that link to themselves, and errors in input data formatting. In such cases, the program should provide appropriate error messages

## Authors

* [@Balde17](https://github.com/balde17)
* [@Stéphane](https://github.com/badStephane)
* Alain Gildas Ogou
* Abel Soglada Dakodo

