package main

import (
	"math/rand"
)

var sumIncubation, sumCritical, sumRecovery, sumImmunity int
var countInfected int

// cell becomes infected
func (c *Cell) infected() {

	c.Infected = true
	c.Protected = false
	stats.infected++
	countInfected++
	c.Died = false
	if c.NbTimesInfected > 0 {
		stats.recovered--
	}
	c.NbTimesInfected = c.NbTimesInfected + 1

	// define the infections parameters
	c.Incubation = rand.Intn(params.IncubationMax-params.IncubationMin) + params.IncubationMin
	sumIncubation += c.Incubation
	c.Critical = rand.Intn(params.CriticalMax-params.CriticalMin) + params.CriticalMin
	sumCritical += c.Critical
	c.Recovery = rand.Intn(params.RecoveryMax-params.RecoveryMin) + params.RecoveryMin
	sumRecovery += c.Recovery
	// define immunity as 0 to 2 times the max recovery value
	c.Immunity = rand.Intn(3) * params.RecoveryMax
	sumImmunity += c.Immunity

	meansReport.MeanIncubation = float64(sumIncubation) / float64(countInfected)
	meansReport.MeanCritical = float64(sumCritical) / float64(countInfected)
	meansReport.MeanRecovery = float64(sumRecovery) / float64(countInfected)
	meansReport.MeanImmunity = float64(sumImmunity) / float64(countInfected)
}

// cell recovers and gain a level of immunity
func (c *Cell) recovery() {
	c.Infected = false
	c.Protected = true
	stats.recovered++
	stats.infected--
}

// cell dies :(
func (c *Cell) die() {
	c.Infected = false
	c.Died = true
	c.Protected = true
	stats.dead++
}

// process the infected cell
func (c *Cell) process() {
	if c.Infected {
		// if still in in incubation stage
		if c.Incubation > 0 {
			c.Incubation = c.Incubation - 1
			c.Protected = false
		}
		if c.Critical > 0 && c.Incubation == 0 {
			c.Critical = c.Critical - 1
			c.Protected = false
			if rand.Float64() < params.Deathlevel {
				c.die()
			}
		}
		if c.Recovery > 0 && c.Critical == 0 && c.Incubation == 0 {
			c.recovery()
		}

	} else {
		// if cell is in recovery cannot be infected
		if c.Recovery > 0 && c.Critical == 0 && c.Incubation == 0 && c.NbTimesInfected > 0 {
			c.Recovery = c.Recovery - 1
			c.Protected = true
		}
		// if cell is immunized cannot be infected
		if c.NbTimesInfected > 0 && c.Critical == 0 && c.Incubation == 0 && c.Recovery == 0 && c.Immunity > 0 {
			c.Immunity = c.Immunity - 1
			c.Protected = true
		}
		// if cell is not anymore immunized it can be infected
		if c.NbTimesInfected > 0 && c.Critical == 0 && c.Incubation == 0 && c.Recovery == 0 && c.Immunity == 0 {
			c.Protected = false
		}

	}
}

// create a cell
func createCell(x, y int, empty, protected bool) (c Cell) {
	c = Cell{
		X:               x,
		Y:               y,
		Empty:           empty,
		Infected:        false,
		Protected:       protected,
		NbTimesInfected: 0,
		Incubation:      0,
		Critical:        0,
		Recovery:        0,
		Immunity:        0,
		Confined:        false,
		Died:            false,
	}
	return
}

// create the initial population
func createPopulation(width int) {
	cells = make([]Cell, width*width)
	n := 0

	for i := 1; i <= width; i++ {
		for j := 1; j <= width; j++ {
			p := rand.Float64()
			if p < params.Density {
				cells[n] = createCell(i, j, false, false)
				stats.living++
			} else {
				cells[n] = createCell(i, j, true, true)
			}
			n++
		}
	}
}

// mobility
func moveCells(width int) {
	var emptyList []int // list of empty cells
	var dstIndex int    // cell destination
	var srcIndex int    // cell source
	var tmpCell Cell

	// build the list of empty cells
	for i, cell := range cells {
		if cell.Empty == true {
			emptyList = append(emptyList, i)
		}
	}

	for i := 0; i < len(emptyList); i++ {
		// select a random cell
		srcIndex = rand.Intn((width * width) - 1)
		// choose a non-empty cell to move out
		for cells[srcIndex].Empty == true {
			if srcIndex == (width*width)-1 {
				srcIndex = 1
				continue
			}
			srcIndex++
		}

		// select the dst empty cell
		dstIndex = emptyList[i]

		if rand.Float64() < params.Move && cells[srcIndex].Confined == false {
			tmpCell = cells[srcIndex]
			cells[srcIndex] = cells[dstIndex]
			cells[dstIndex] = tmpCell
			stats.move++
			//update contact list at source
			neighbours := neighboursList(srcIndex, width)
			for i := range neighbours {
				updateContactList(width, i)
			}
			//update contact list at destination
			neighbours = neighboursList(dstIndex, width)
			for i := range neighbours {
				updateContactList(width, i)
			}

		}
	}
}

// choose 1 cell to be patient zero in the middle of the simulation
func infectOneCell(width int) {
	// select a random cell
	i := rand.Intn((width * width) - 1)
	// choose a non-empty cell to infect
	for cells[i].Empty == true {
		if i == (width*width)-1 {
			i = 1
			continue
		}
		i++
	}
	// then infect
	cells[i].infected()
}

// count how many are never infected
func countNeverInfected() int {
	count := 0
	for _, cell := range cells {
		if cell.NbTimesInfected == 0 {
			count++
		}
	}
	return count
}

// count the number of currently infected cells
func countCurrentInfected() int {
	count := 0
	for _, cell := range cells {
		if cell.Infected {
			count++
		}
	}
	return count
}

// count the number of cells that has been infected
func countTotalInfected() int {
	count := 0
	for _, cell := range cells {
		if cell.NbTimesInfected > 0 {
			count++
		}
	}
	return count
}

// count the number of reinfected cells at least one time
func countReinfected() int {
	count := 0
	for _, cell := range cells {
		if cell.NbTimesInfected > 1 {
			count++
		}
	}
	return count
}

// build the full contact list for each cell
func updateFullContactList(width int) {
	for i := range cells {
		neighbours := neighboursList(i, width)
		for n := range neighbours {
			cells[i].ContactsList = append(cells[i].ContactsList, neighbours[n])
		}
	}
}

// update the contact list for one cell
func updateContactList(width, i int) {
	neighbours := neighboursList(i, width)
	for n := range neighbours {
		cells[i].ContactsList = append(cells[i].ContactsList, neighbours[n])
	}
}
