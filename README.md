# Pandemic simulation using Go 

This is a pandemic simulation tool written in Go, highly configurable (see `struct.go` file) and inspired by Sau Sheong simulator (see [Simulating epidemics using Go and Python](https://towardsdatascience.com/simulating-epidemics-using-go-and-python-101557991b20)). Pandemic-sim also used a color grid (game of life like) but also implement a mobility model based on [random way point](https://en.wikipedia.org/wiki/Random_waypoint_model) and proposed two funny GUI to assess the evolution of the dissemination over time. At the end, a data curve PNG file is produced and raw data are saved in a TXT file (`data.txt`). Using `gnuplot graph.gp` also allows to generate an EPS file with the results. The grey area represented corresponding to the default confinement period.