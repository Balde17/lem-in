package main

type Room struct {
	name    string
	coord_x int
	coord_y int
	Visited bool
}

type LinksInRooms struct {
	room1 string
	room2 string
	
}

type InformationsInFile struct {
	number_of_ants int
	rooms          []Room
	links          []LinksInRooms
	start          Room
	end            Room
}

type Ant struct{
	ID int
	Position Room
}