package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("Input number should be at least 1")
    }
    i := 0
	for n != 1 {
        n = nextCollatz(n)
        i++
    }
    return i, nil
}

func nextCollatz(n int) (int) {
    if n % 2 == 0 {
        n /= 2
    } else {
        n *= 3
        n++
    }
    return n 
}