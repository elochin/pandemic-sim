package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// set initial parameters
	params.IncubationMin = 4
	params.IncubationMax = 7
	params.CriticalMin = 5
	params.CriticalMax = 10
	params.RecoveryMin = 10
	params.RecoveryMax = 30

	// simulation configurations
	var numDays int      // number of simulation days
	var width int        // the number of cells on one side of the square
	var filename string  // name of data file
	var gui int          // gui 0: disable 1: smileys 2: colors
	var timerSleep int64 // granularity of the step of the GUI simulation in ms

	// generate a random seed
	rand.Seed(time.Now().UTC().UnixNano())

	// capturing user input
	flag.Float64Var(&params.Deathlevel, "d", 0.02, "probability of deathlevel")
	flag.IntVar(&gui, "gui", 0, "enable GUI animation 0: disable 1: smileys 2: colors")
	flag.IntVar(&numDays, "t", 365, "number of simulation days")
	flag.IntVar(&width, "w", 50, "the number of cells on one side of the image")
	flag.StringVar(&filename, "f", "data", "file name of data and PNG file")
	flag.Float64Var(&params.Rate, "i", 0.15, "probability of infection")
	flag.Float64Var(&params.Move, "m", 0.15, "probability of moving")
	flag.Float64Var(&params.Density, "p", 0.7, "percentage of simulation grid that is populated")
	flag.IntVar(&params.ConfineStart, "cs", 77, "day when confine is introduced")
	flag.IntVar(&params.ConfineEnd, "ce", 132, "day when confine is over")
	flag.Float64Var(&params.ConfineRespect, "r", 0.5, "respect of the confine")
	flag.Int64Var(&timerSleep, "g", 10, "granularity of the step of the GUI simulation in ms")
	flag.Parse()

	// using termbox to control the simulation
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	//dead = 0

	// poll for keyboard events in another goroutine
	events := make(chan termbox.Event, 1000)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	// create the initial population
	createPopulation(width)
	displayCell(cells, 0, 0, width, numDays, gui)
	infectOneCell(width)
	displayCell(cells, 0, 0, width, numDays, gui)

	// main simulation loop
	var t int
	for t = 0; t < numDays; t++ {
		// capture the ctrl-q key to end the simulation
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyCtrlQ {
					termbox.Close()
					os.Exit(0)
				}
			}
			if ev.Type == termbox.EventResize {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				displayCell(cells, t, 0, width, numDays, gui)
			}

		default:
		}
		//debugMap(t)

		updateFullContactList(width)

		for n := range cells {
			// if cell is empty, died, go to the next cell
			if cells[n].Died == true || cells[n].Empty == true {
				continue
			}

			// process infected cells
			cells[n].process()

			// confine has been started
			if !cells[n].Confined && (t > params.ConfineStart) && (t < params.ConfineEnd) &&
				(rand.Float64() < params.ConfineRespect) {
				cells[n].Confined = true
				cells[n].Protected = true
			}

			// check nobody is confined after the end of confinment
			if cells[n].Confined && (t > params.ConfineEnd) {
				cells[n].Confined = false
				cells[n].Protected = false
			}

			// unless confined, infect neighbours
			if !cells[n].Protected && cells[n].Infected {
				// find all the cell's neighbours
				neighbours := neighboursList(n, width)
				//debugNeighbours(t, n, neighbours)
				// for every neighbours
				for _, neighbour := range neighbours {
					//if the cell is empty, died, already infected or cannot be infected, go to the next neighbour
					if cells[neighbour].Died == true || cells[neighbour].Empty == true ||
						cells[neighbour].Protected == true || cells[neighbour].Infected == true {
						continue
					}
					// if probability less than infection rate then cell gets infected
					if rand.Float64() < params.Rate {
						cells[neighbour].infected()
					}
				}
			}
			// mobility simulation
		}

		moveCells(width)

		// stats
		stats.neverInfected = countNeverInfected()
		stats.reInfected = countReinfected()
		count := countCurrentInfected()
		total := countTotalInfected()
		displayCell(cells, t, count, width, numDays, gui)
		time.Sleep(time.Duration(timerSleep) * time.Millisecond)

		// collect simulation data
		stats.infecteds = append(stats.infecteds, count)
		stats.infections = append(stats.infections, total)
		stats.deaths = append(stats.deaths, stats.dead)
	}
	//Final results
	displayCell(cells, t, 0, width, numDays, gui)
	saveStats(filename, stats.infections, stats.deaths, stats.infecteds, numDays)

	//Wait for keypressed
	<-events
	//termbox.Close()
	//fmt.Println()
}

func debugMap(t int) {
	f, err := os.Create(fmt.Sprintf("debug%d.txt", t))
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, cell := range cells {
		_, err = fmt.Fprintf(w, "%v\n", cell)
		check(err)
	}
	w.Flush()
}

func debugNeighbours(t, n int, neighbours []int) {
	f, err := os.Create(fmt.Sprintf("debug%dN%d.txt", t, n))
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, cell := range neighbours {
		_, err = fmt.Fprintf(w, "%v\n", cell)
		check(err)
	}
	w.Flush()
}

// save simulation data
func saveStats(filename string, infections, deaths, infecteds []int, numDays int) {
	f, err := os.Create(fmt.Sprintf("%s.txt", filename))
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	for i := 0; i < len(infections); i++ {
		_, err = fmt.Fprintf(w, "%d %d %d\n", infecteds[i], deaths[i], infections[i])
		check(err)
	}
	w.Flush()

	// plot

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = ""
	p.X.Label.Text = "Days"
	p.Y.Label.Text = "#"

	infectedPts := make(plotter.XYs, numDays)
	deathPts := make(plotter.XYs, numDays)

	for i := 0; i < numDays; i++ {
		infectedPts[i].X = float64(i)
		infectedPts[i].Y = float64(infecteds[i])
		deathPts[i].X = float64(i)
		deathPts[i].Y = float64(deaths[i])
	}

	err = plotutil.AddLinePoints(p,
		"Number of infected", infectedPts,
		"Number of death", deathPts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, fmt.Sprintf("%s.png", filename)); err != nil {
		panic(err)
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
