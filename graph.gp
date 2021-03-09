set xlabel 'Days' 
set ylabel '#'
set xrange [0:365]
set key top left
set style line 1 lt 1 lw 3
set style line 11 lt 1 lw 1
set style line 2 lt 2 lw 3
set style line 3 lt 3 lw 3

#set arrow from 77, graph 0 to 77, graph 1 nohead
set style rect fc lt -1 fs solid 0.15 noborder
set obj rect from 77, graph 0 to 132, graph 1

plot \
'data.txt' using :1 t 'Infected' with lines ls 1,\
'data.txt' using :2 t 'Death' with lines ls 2, \
'data.txt' using :3 t 'Total number of infections' with lines ls 3

set term post color "Helvetica" 16
set encoding iso_8859_1
#set size 1.0, 0.7
set output "data.eps"
pause -1
replot
exit
