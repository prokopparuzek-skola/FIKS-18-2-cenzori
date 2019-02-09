package main

import "fmt"

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

func oneTask(in *input_t, S []uint64) (uint64, uint64, uint64) {
	var min, max, sum uint64
	for i := range S {
		switch in.t {
		case 0:
			if S[i] < min {
				min = S[i]
			}
			if S[i] > max {
				max = S[i]
			}
			sum += S[i]
		case 1:
			S[i] += in.A
		case 2:
			S[i] = in.A
		}
	}
	return min, max, sum
}

func solve(gen *generator_t, t, N uint64) {
	var in input_t
	var min, sum, max uint64
	var minX, sumX, maxX uint64
	S := make([]uint64, N)

	for i := uint64(0); i < t; i++ {
		gen.genInput(&in, N)
		min, max, sum = oneTask(&in, S[in.b:in.e+1])
		minX ^= min
		maxX ^= max
		sumX ^= sum
		//fmt.Println(in)
		//fmt.Println(S)
	}
	fmt.Printf("%d\n%d\n%d\n", minX, maxX, sumX)
}

func main() {
	var T uint8
	fmt.Scanf("%d", &T)
	for i := uint8(0); i < T; i++ {
		var generator generator_t
		var N, t uint64

		fmt.Scanf("%d %d %d %d %d", &t, &N, &generator.a, &generator.b, &generator.x)
		solve(&generator, t, N)
	}
}
