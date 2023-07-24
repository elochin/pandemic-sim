package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func displayCell(cells []Cell, t, count, width, numDays int, gui int) {

	// icons
	//const hourglass = "\xE2\x8F\xB3"
	//const skull = "\xF0\x9F\x92\x80"
	//const emomask = "\xF0\x9F\x98\xB7"
	//const emosick = "\xF0\x9F\x98\xA1"
	//const emosmile = "\xF0\x9F\x98\x83"
	//const house = "\xF0\x9F\x8F\xA0"

	//w, h := termbox.Size()
	//termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// draw the icons, we use 2*cell.X because each icon takes 2 characters
	if gui == 1 {
		for _, cell := range cells {
			if cell.Empty {
				termbox.SetCell(2*cell.X, cell.Y, ' ', termbox.ColorWhite, termbox.ColorDefault)
				termbox.SetCell(2*cell.X+1, cell.Y, ' ', termbox.ColorWhite, termbox.ColorDefault)
			} else {
				if cell.Died {
					termbox.SetCell(2*cell.X, cell.Y, '💀', termbox.ColorBlack, termbox.ColorDefault)
				} else {
					if cell.Confined {
						termbox.SetCell(2*cell.X, cell.Y, '🏠', termbox.ColorBlue, termbox.ColorDefault)
					} else {
						if cell.Infected {
							if cell.Incubation > 0 {
								termbox.SetCell(2*cell.X, cell.Y, '⏳', termbox.ColorGreen, termbox.ColorDefault)
							} else {
								termbox.SetCell(2*cell.X, cell.Y, '😡', termbox.ColorRed, termbox.ColorDefault)
							}
						} else {

							if cell.Recovery > 0 && cell.Critical == 0 && cell.Incubation == 0 {
								termbox.SetCell(2*cell.X, cell.Y, '😷', termbox.ColorBlue, termbox.ColorDefault)
							} else {
								termbox.SetCell(2*cell.X, cell.Y, '😃', termbox.ColorGreen, termbox.ColorDefault)
							}
						}

					}
				}
			}
		}
	}

	if gui == 2 {
		for _, cell := range cells {
			if cell.Empty {
				if cell.Y%2 != 0 {
					termbox.SetChar(cell.X, (cell.Y+1)/2, ' ')
					termbox.SetBg(cell.X, (cell.Y+1)/2, termbox.ColorBlack)
				} else {
					termbox.SetChar(cell.X, cell.Y/2, '▄')
					termbox.SetFg(cell.X, cell.Y/2, termbox.ColorBlack)
				}
			} else {
				if cell.Died {
					if cell.Y%2 != 0 {
						termbox.SetChar(cell.X, (cell.Y+1)/2, ' ')
						termbox.SetBg(cell.X, (cell.Y+1)/2, termbox.ColorWhite)
					} else {
						termbox.SetChar(cell.X, cell.Y/2, '▄')
						termbox.SetFg(cell.X, cell.Y/2, termbox.ColorWhite)
					}
				} else {
					if cell.Confined {
						if cell.Y%2 != 0 {
							termbox.SetChar(cell.X, (cell.Y+1)/2, ' ')
							termbox.SetBg(cell.X, (cell.Y+1)/2, termbox.ColorBlue)
						} else {
							termbox.SetChar(cell.X, cell.Y/2, '▄')
							termbox.SetFg(cell.X, cell.Y/2, termbox.ColorBlue)
						}
					} else {
						if cell.Infected {
							if cell.Incubation > 0 {
								if cell.Y%2 != 0 {
									termbox.SetChar(cell.X, (cell.Y+1)/2, ' ')
									termbox.SetBg(cell.X, (cell.Y+1)/2, termbox.ColorYellow)
								} else {
									termbox.SetChar(cell.X, cell.Y/2, '▄')
									termbox.SetFg(cell.X, cell.Y/2, termbox.ColorYellow)
								}
							} else {
								if cell.Y%2 != 0 {
									termbox.SetChar(cell.X, (cell.Y+1)/2, ' ')
									termbox.SetBg(cell.X, (cell.Y+1)/2, termbox.ColorRed)
								} else {
									termbox.SetChar(cell.X, cell.Y/2, '▄')
									termbox.SetFg(cell.X, cell.Y/2, termbox.ColorRed)
								}
							}
						} else {

							if cell.Recovery > 0 && cell.Critical == 0 && cell.Incubation == 0 {
								if cell.Y%2 != 0 {
									termbox.SetChar(cell.X, (cell.Y+1)/2, ' ')
									termbox.SetBg(cell.X, (cell.Y+1)/2, termbox.ColorCyan)
								} else {
									termbox.SetChar(cell.X, cell.Y/2, '▄')
									termbox.SetFg(cell.X, cell.Y/2, termbox.ColorCyan)
								}
							} else {
								if cell.Y%2 != 0 {
									termbox.SetChar(cell.X, (cell.Y+1)/2, ' ')
									termbox.SetBg(cell.X, (cell.Y+1)/2, termbox.ColorGreen)
								} else {
									termbox.SetChar(cell.X, cell.Y/2, '▄')
									termbox.SetFg(cell.X, cell.Y/2, termbox.ColorGreen)
								}
							}
						}

					}
				}
			}
		}

	}

	columnPos := 10
	if gui == 1 {
		columnPos = 2*width + 10
	}
	if gui == 2 {
		columnPos = width + 10
	}

	if t == numDays {
		printf_tb(columnPos, 2, termbox.ColorWhite, termbox.ColorDefault, "====== Final Result ====== ")
	} else {
		printf_tb(columnPos, 2, termbox.ColorWhite, termbox.ColorDefault, "Current infected: %d cells              ", count)
	}
	statsDisplay(t, width, numDays, gui)
	termbox.Flush()
}

func statsDisplay(t, width, numDays, gui int) {
	linePos := 3
	columnPos := 10
	if gui == 1 {
		columnPos = 2*width + 10
	}
	if gui == 2 {
		columnPos = width + 10
	}

	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Time            : %d/%d days", t, numDays)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Infected        : %2.2f%%", float64(stats.infected)*100.0/float64(stats.living))
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Re-infected     : %2.2f%%", float64(stats.reInfected)*100.0/float64(stats.living))
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Died            : %2.2f%%", float64(stats.dead)*100.0/float64(stats.living))
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Recovered       : %2.2f%%", float64(stats.recovered)*100.0/float64(stats.living))
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Mean incubation : %2.2f days", meansReport.MeanIncubation)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Mean critical   : %2.2f days", meansReport.MeanCritical)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Mean recovery   : %2.2f days", meansReport.MeanRecovery)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Mean immunity   : %2.2f days", meansReport.MeanImmunity)
	linePos++
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "=== Simulation Parameters ===")
	linePos++
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Density                : %2.0f%% populated", params.Density*100)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Number of move         : %d ", stats.move)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Infection rate         : %2.2f%% ", params.Rate*100)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Max/Min incubation     : %d/%d days", params.IncubationMax, params.IncubationMin)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Max/Min critical phase : %d/%d days", params.CriticalMax, params.CriticalMin)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Max/Min recovery phase : %d/%d days", params.RecoveryMax, params.RecoveryMin)
	linePos++
	printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Death rate             : %2.2f%%", params.Deathlevel*100)
	if params.ConfineStart < numDays {
		linePos++
		printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Confinement introduced : %dth day", params.ConfineStart)
		linePos++
		printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Confinement respect    : %2.2f%%", params.ConfineRespect*100)
		linePos++
		printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Confinement ended      : %dth day", params.ConfineEnd)
	}
	linePos = linePos + 2
	if t == numDays {
		printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "  Press any key to quit.  ")
	} else {
		printf_tb(columnPos, linePos, termbox.ColorWhite, termbox.ColorDefault, "Ctrl-Q to quit simulation.")
	}

	// Captions
	if gui == 1 {
		linePos = linePos + 2
		termbox.SetCell(columnPos, linePos, '😃', termbox.ColorDefault, termbox.ColorDefault)
		printf_tb(columnPos+3, linePos, termbox.ColorWhite, termbox.ColorDefault, "Non contaminated")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '⏳', termbox.ColorDefault, termbox.ColorDefault)
		printf_tb(columnPos+3, linePos, termbox.ColorWhite, termbox.ColorDefault, "Incubating")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '😡', termbox.ColorDefault, termbox.ColorDefault)
		printf_tb(columnPos+3, linePos, termbox.ColorWhite, termbox.ColorDefault, "Infected")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '😷', termbox.ColorDefault, termbox.ColorDefault)
		printf_tb(columnPos+3, linePos, termbox.ColorWhite, termbox.ColorDefault, "Recovered")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '💀', termbox.ColorDefault, termbox.ColorDefault)
		printf_tb(columnPos+3, linePos, termbox.ColorWhite, termbox.ColorDefault, "Died")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '🏠', termbox.ColorDefault, termbox.ColorDefault)
		printf_tb(columnPos+3, linePos, termbox.ColorWhite, termbox.ColorDefault, "Confined")
	}

	if gui == 2 {
		linePos = linePos + 2
		termbox.SetCell(columnPos, linePos, '▄', termbox.ColorGreen, termbox.ColorDefault)
		printf_tb(columnPos+2, linePos, termbox.ColorWhite, termbox.ColorDefault, "Non contaminated")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '▄', termbox.ColorYellow, termbox.ColorDefault)
		printf_tb(columnPos+2, linePos, termbox.ColorWhite, termbox.ColorDefault, "Incubating")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '▄', termbox.ColorRed, termbox.ColorDefault)
		printf_tb(columnPos+2, linePos, termbox.ColorWhite, termbox.ColorDefault, "Infected")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '▄', termbox.ColorCyan, termbox.ColorDefault)
		printf_tb(columnPos+2, linePos, termbox.ColorWhite, termbox.ColorDefault, "Recovered")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '▄', termbox.ColorWhite, termbox.ColorDefault)
		printf_tb(columnPos+2, linePos, termbox.ColorWhite, termbox.ColorDefault, "Died")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '▄', termbox.ColorBlue, termbox.ColorDefault)
		printf_tb(columnPos+2, linePos, termbox.ColorWhite, termbox.ColorDefault, "Confined")
		linePos++
		linePos++
		termbox.SetCell(columnPos, linePos, '▄', termbox.ColorBlack, termbox.ColorDefault)
		printf_tb(columnPos+2, linePos, termbox.ColorWhite, termbox.ColorDefault, "Empty")
	}

}

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}
