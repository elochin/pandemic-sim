package main

// simulation statistics
type Statistics struct {
	living, dead, recovered, infected, neverInfected, reInfected int
	infections, deaths, infecteds                                []int
	move                                                         int // number of cells move
}

// pandemic simulation parameters
type PandemicParameters struct {
	IncubationMin  int // {min, max} days healthy carrier - not contagious
	IncubationMax  int
	CriticalMin    int // {min, max} days critical phase (either die during this period or recover - contagious
	CriticalMax    int
	RecoveryMin    int // {min, max} days recovery period - no contagious/resilient
	RecoveryMax    int
	Rate           float64 // probability of infection
	Move           float64 // probability of moving
	Density        float64 // percentage of simulation grid that is populated
	Deathlevel     float64 // probability to die
	ConfineStart   int     // the day when confine is introduced
	ConfineEnd     int     // the day when confine is stopped
	ConfineRespect float64 // respect of the confine
}

// Cell is a representation of a cell within the grid
type Cell struct {
	X               int
	Y               int
	Empty           bool  // define whether the cell is empty or contains a live or dead cell
	Infected        bool  // infected or not?
	Protected       bool  // can be infected or not?
	NbTimesInfected int   // nb of times the person has been infected
	Incubation      int   // incubation phase in days
	Critical        int   // critical phase in days
	Recovery        int   // recovery phase in days
	Immunity        int   // number of immunity days computed as a function of the end of the recovery period
	Confined        bool  // confined
	Died            bool  // died
	ContactsList    []int // list of contacts
}

type Means struct {
	MeanIncubation float64
	MeanCritical   float64
	MeanRecovery   float64
	MeanImmunity   float64
}

var meansReport Means
var params PandemicParameters
var stats Statistics

// the simulation grid
var cells []Cell
