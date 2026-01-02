package day12

type Shape [][]int
type Package struct {
	Shape Shape
	Area  int
}

type Packages []Package

/*
0:
#..
##.
.##

1:
###
.#.
###

2:
###
###
..#

3:
###
.##
##.

4:
###
##.
#..

5:
#.#
#.#
###
*/

func GetPackages() Packages {
	/*
		#..
		##.
		.##
	*/
	package0 := Package{
		Shape: Shape{
			[]int{1, 0, 0},
			[]int{1, 1, 0},
			[]int{0, 1, 1},
		},
		Area: 5,
	}

	/*
		###
		.#.
		###
	*/
	package1 := Package{
		Shape: Shape{
			[]int{1, 1, 1},
			[]int{0, 1, 0},
			[]int{1, 1, 1},
		},
		Area: 7,
	}

	/*
		###
		###
		..#
	*/
	package2 := Package{
		Shape: Shape{
			[]int{1, 1, 1},
			[]int{1, 1, 1},
			[]int{0, 0, 1},
		},
		Area: 7,
	}

	/*
		###
		.##
		##.
	*/
	package3 := Package{
		Shape: Shape{
			[]int{1, 1, 1},
			[]int{0, 1, 1},
			[]int{1, 1, 0},
		},
		Area: 7,
	}

	/*
		###
		##.
		#..
	*/
	package4 := Package{
		Shape: Shape{
			[]int{1, 1, 1},
			[]int{1, 1, 0},
			[]int{1, 0, 0},
		},
		Area: 6,
	}

	/*
		#.#
		#.#
		###
	*/
	package5 := Package{
		Shape: Shape{
			[]int{1, 0, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		Area: 7,
	}

	return Packages{
		package0,
		package1,
		package2,
		package3,
		package4,
		package5,
	}
}
