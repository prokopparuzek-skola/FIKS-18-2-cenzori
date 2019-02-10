package main

import (
	"bufio"
	"fmt"
	"os"
)

type input_t struct {
	t uint64
	b uint64
	e uint64
	A uint64
}

type generator_t struct {
	a uint64
	b uint64
	x uint64
}

type leave struct {
	value uint64
	min   uint64
	max   uint64
	sum   uint64
}

<<<<<<< Updated upstream
=======
type answer_t struct {
	min uint64
	max uint64
	sum uint64
}

type query_t struct {
	i uint64
	j uint64
}

>>>>>>> Stashed changes
func (gen *generator_t) nextInt() uint64 {
	gen.x = (gen.x*gen.a + gen.b) % uint64(1000000007)
	return gen.x
}

func (gen *generator_t) genInput(in *input_t, N uint64) {
	in.t = gen.nextInt() % 3
	in.b = gen.nextInt() % N
	in.e = gen.nextInt() % N
	if in.b > in.e {
		swap := in.e
		in.e = in.b
		in.b = swap
	}
	in.A = gen.nextInt() % N
}

func search(S []leave, vrchol uint64, where query_t, what query_t) answer_t {
	var ans answer_t

	if where.i == what.i && where.j == what.j {
		ans.max = S[vrchol].max
		ans.min = S[vrchol].min
		ans.sum = S[vrchol].sum
		return ans
	}
	stred := (where.i + where.j) / 2
	if what.j <= stred {
		return search(S, (vrchol)*2, query_t{where.i, stred}, what)
	} else if what.i >= stred {
		return search(S, (vrchol)*2+1, query_t{stred, where.j}, what)
	} else {
		ans1 := search(S, (vrchol)*2, query_t{where.i, stred}, query_t{what.i, stred})
		ans2 := search(S, (vrchol)*2+1, query_t{stred + 1, where.j}, query_t{stred + 1, what.j})
		if ans1.max > ans2.max {
			ans.max = ans1.max
		}
		if ans1.min < ans2.min {
			ans.min = ans1.min
		}
		ans.sum = ans1.sum + ans2.sum
		return ans
	}
}

func oneTask(in *input_t, S []leave) (uint64, uint64, uint64) {
	var min, max, sum uint64
	min = S[0].value
	max = S[0].value
	for i := range S {
		switch in.t {
		case 0:
		case 1:
			S[i].value += in.A
		case 2:
			S[i].value = in.A
		}
	}
	return min, max, sum
}

func solve(gen *generator_t, t, N uint64, w *bufio.Writer) {
	var in input_t
	var min, sum, max uint64
	var minX, sumX, maxX uint64
<<<<<<< Updated upstream
	S := make([]uint64, N*2)
=======
	S := make([]leave, N*2+1)
>>>>>>> Stashed changes

	for i := uint64(0); i < t; i++ {
		gen.genInput(&in, N)
		min, max, sum = oneTask(&in, S)
		if in.t == 0 {
			minX ^= min
			maxX ^= max
			sumX ^= sum
		}
		//fmt.Println(in)
		//fmt.Println(S)
		//fmt.Printf("%d %d %d\n", min, max, sum)
	}
	fmt.Fprintf(w, "%d\n%d\n%d\n", minX, maxX, sumX)
}

func main() {
	var T uint8
	fmt.Scanf("%d", &T)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := uint8(0); i < T; i++ {
		var generator generator_t
		var N, t uint64

		fmt.Scanf("%d %d %d %d %d", &t, &N, &generator.a, &generator.b, &generator.x)
		solve(&generator, t, N, w)
	}
}
