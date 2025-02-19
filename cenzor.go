package main

import (
	"bufio"
	"fmt"
	"math"
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
	min uint64
	max uint64
	sum uint64
	inc uint64
	set int64
}

type answer_t struct {
	min uint64
	max uint64
	sum uint64
}

type query_t struct {
	i uint64
	j uint64
}

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

func increase(S []leave, vrchol uint64, where query_t, what query_t, inc uint64) {
	eval(S, vrchol)
	if where.i == what.i && where.j == what.j {
		S[vrchol].inc = inc
		return
	}
	stred := (where.i + where.j) / 2
	if what.j <= stred {
		increase(S, vrchol*2, query_t{where.i, stred}, what, inc)
	} else if what.i >= stred {
		increase(S, vrchol*2+1, query_t{stred, where.j}, what, inc)
	} else {
		increase(S, (vrchol)*2, query_t{where.i, stred}, query_t{what.i, stred}, inc)
		increase(S, (vrchol)*2+1, query_t{stred, where.j}, query_t{stred, what.j}, inc)
	}

	eval(S, vrchol*2)
	eval(S, vrchol*2+1)
	if S[vrchol*2].min < S[vrchol*2+1].min {
		S[vrchol].min = S[vrchol*2].min
	} else {
		S[vrchol].min = S[vrchol*2+1].min
	}
	if S[vrchol*2].max < S[vrchol*2+1].max {
		S[vrchol].max = S[vrchol*2+1].max
	} else {
		S[vrchol].max = S[vrchol*2].max
	}
	S[vrchol].sum = S[vrchol*2].sum + S[vrchol*2+1].sum
}

func setValue(S []leave, vrchol uint64, where, what query_t, set int64) {
	eval(S, vrchol)
	if where.i == what.i && where.j == what.j {
		S[vrchol].set = set
		return
	}
	stred := (where.i + where.j) / 2
	if what.j <= stred {
		setValue(S, vrchol*2, query_t{where.i, stred}, what, set)
	} else if what.i >= stred {
		setValue(S, vrchol*2+1, query_t{stred, where.j}, what, set)
	} else {
		setValue(S, (vrchol)*2, query_t{where.i, stred}, query_t{what.i, stred}, set)
		setValue(S, (vrchol)*2+1, query_t{stred, where.j}, query_t{stred, what.j}, set)
	}

	eval(S, vrchol*2)
	eval(S, vrchol*2+1)
	if S[vrchol*2].min < S[vrchol*2+1].min {
		S[vrchol].min = S[vrchol*2].min
	} else {
		S[vrchol].min = S[vrchol*2+1].min
	}
	if S[vrchol*2].max < S[vrchol*2+1].max {
		S[vrchol].max = S[vrchol*2+1].max
	} else {
		S[vrchol].max = S[vrchol*2].max
	}
	S[vrchol].sum = S[vrchol*2].sum + S[vrchol*2+1].sum
}

func eval(S []leave, vrchol uint64) {
	if S[vrchol].set != -1 {
		set := S[vrchol].set
		S[vrchol].set = -1
		S[vrchol].max = uint64(set)
		S[vrchol].min = uint64(set)
		hladina := math.Floor(math.Log2(float64(vrchol)))
		pod := math.Pow(2, math.Log2(float64(len(S)/2))-hladina)
		S[vrchol].sum = uint64(pod) * uint64(set)
		if vrchol < uint64(len(S)/2) {
			if S[vrchol*2].inc != 0 {
				S[vrchol*2].inc = 0
			}
			if S[vrchol*2+1].inc != 0 {
				S[vrchol*2+1].inc = 0
			}
			S[vrchol*2].set = set
			S[vrchol*2+1].set = set
		}
	}

	if S[vrchol].inc != 0 {
		inc := S[vrchol].inc
		S[vrchol].inc = 0
		S[vrchol].max += inc
		S[vrchol].min += inc
		hladina := math.Floor(math.Log2(float64(vrchol)))
		pod := math.Pow(2, math.Log2(float64(len(S)/2))-hladina)
		S[vrchol].sum = uint64(pod)*uint64(inc) + S[vrchol].sum
		if vrchol < uint64(len(S)/2) {
			if S[vrchol*2].set != -1 {
				eval(S, vrchol*2)
			}
			if S[vrchol*2+1].set != -1 {
				eval(S, vrchol*2+1)
			}
			S[vrchol*2].inc += inc
			S[vrchol*2+1].inc += inc
		}

	}
}

func search(S []leave, vrchol uint64, where query_t, what query_t) answer_t {
	var ans answer_t

	eval(S, vrchol)
	if (where.i == what.i) && (where.j == what.j) {
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
		ans2 := search(S, (vrchol)*2+1, query_t{stred, where.j}, query_t{stred, what.j})
		if ans1.max > ans2.max {
			ans.max = ans1.max
		} else {
			ans.max = ans2.max
		}
		if ans1.min < ans2.min {
			ans.min = ans1.min
		} else {
			ans.min = ans2.min
		}
		ans.sum = ans1.sum + ans2.sum
		return ans
	}
}

func oneTask(in *input_t, S []leave) answer_t {
	var ans answer_t

	switch in.t {
	case 0:
		ans = search(S, 1, query_t{1, uint64(len(S)/2) + 1}, query_t{in.b + 1, in.e + 2})
	case 1:
		increase(S, 1, query_t{1, uint64(len(S)/2) + 1}, query_t{in.b + 1, in.e + 2}, in.A)
	case 2:
		setValue(S, 1, query_t{1, uint64(len(S)/2 + 1)}, query_t{in.b + 1, in.e + 2}, int64(in.A))
	}
	return ans
}

func solve(gen generator_t, t, N uint64, w *bufio.Writer) {
	var in input_t
	var minX, sumX, maxX uint64
	var ans answer_t
	var size uint64

	size = uint64(math.Pow(2, math.Ceil(math.Log2(float64(N)))) * 2)
	S := make([]leave, size)
	for i := range S {
		S[i].set = -1
	}

	for i := uint64(0); i < t; i++ {
		gen.genInput(&in, N)
		ans = oneTask(&in, S)
		if in.t == 0 {
			minX ^= ans.min
			maxX ^= ans.max
			sumX ^= ans.sum
		}
		//fmt.Println(in)
		//fmt.Println(S)
		//fmt.Printf("%d %d %d\n", ans.min, ans.max, ans.sum)
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
		solve(generator, t, N, w)
	}
}
