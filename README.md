[![Build Status](https://travis-ci.org/edbond/angrygo.svg?branch=master)](https://travis-ci.org/edbond/angrygo)

[![codecov](https://codecov.io/gh/edbond/angrygo/branch/master/graph/badge.svg)](https://codecov.io/gh/edbond/angrygo)

[![Go Report Card](https://goreportcard.com/badge/github.com/edbond/angrygo)](https://goreportcard.com/report/github.com/edbond/angrygo)

# Description

The code is example of brute force solver for angry birds puzzle.
Puzzle consists of a board and four figures.
Board contains images of pigs, figures should be positioned on board
to hide pigs. Challenge list pigs that should be left uncovered.

# Board

![Board](https://github.com/edbond/angrygo/blob/master/angrybirds/board.jpg)

# Figures

![Figures](https://github.com/edbond/angrygo/blob/master/angrybirds/figures.jpg)

# Challenge

![Challenge](https://github.com/edbond/angrygo/blob/master/angrybirds/challenge.jpg)

# Pigs

A
![Angry](https://github.com/edbond/angrygo/blob/master/angrybirds/A%20-%20Angry.png)

B
![Beaten](https://github.com/edbond/angrygo/blob/master/angrybirds/B%20-%20Beaten.png)

H
![Hat](https://github.com/edbond/angrygo/blob/master/angrybirds/H%20-%20Hat.png)

P
![Pig](https://github.com/edbond/angrygo/blob/master/angrybirds/P%20-%20Pig.png)

R
![Red](https://github.com/edbond/angrygo/blob/master/angrybirds/R%20-%20Red.png)

S
![Sleepy](https://github.com/edbond/angrygo/blob/master/angrybirds/S%20-%20Sleepy.png)


# Example output

Running `/usr/bin/time -lp ./angrybirds`

```
....
Left on board map[B:1 S:2 A:1 R:1]
Left on board map[H:1 A:2 S:2]
Left on board map[H:1 B:1 P:1 R:1 S:1]
Left on board map[H:1 P:2]
Found solution!
VVP
VVXXX
VZXYX
HZYYP
ZZZYY
Total variants checked: 2333550
Total valid variants: 200
real         4.46
user         4.73
sys          0.14
   7770112  maximum resident set size```
