package main

import "fmt"


func prime(n int) bool {
	i := 2
	for i * i <= n {
		if n % i == 0 {
			return false
		}
		i += 1
	}
	return true
}

func main() {
	b, c := 81, 81
	b = (100 * b) + 100000
	c = b + 17000
	sum := 0
	for n := b; n <= c; n += 17 {
		if !prime(n) {
			sum++
		}
	}
	print(sum)
}

func xmain() {
	a, b, c, d, e, f, g, h := 1, int64(0), int64(0), int64(0), int64(0), int64(0), int64(0), int64(0)
	b = 81
	c = b
	if a != 0 {
		b = (100 * b) + 100000
		c = b + 17000
	}
	for {
		f = 1
		d = 2
		skip := false
		for {
			e = 2
			for {
				//fmt.Println(a, b, c, d, e, f, g, h)
				g = (d * e) - b
				if g == 0 {
					f = 0
					skip = true
					break
				}
				e++
				g = e - b
				if g == 0 {
					break
				}
			}
			if skip {
				break
			}
			d++
			g = d - b
			if g == 0 {
				break
			}
		}
		if f == 0 {
			h++
		}
		g = b
		g -= c
		if g != 0 {
			b += 17
		} else {
			break
		}
	}
	fmt.Println(a, b, c, d, e, f, g, h)
}

func mainsdf() {
	a, b, c, d, e, f, g, h := 0, 0, 0, 0, 0, 0, 0, 0
	// 1. set b 81
	b = 81
	// 2. set c b
	c = b
	// 3. jnz a 2
	if a != 0 {
		b = (100 * b) + 100000
		c = b + 17000
	}
	for {
		// 9. set f 1
		f = 1
		// 10. set d 2
		d = 2
	l11:
	// 11. set e 2
		e = 2
		if g%10000 == 0 {
			//fmt.Println("g", g)
		}
	l12:
		fmt.Println(b, c, d, e, f, g)
		g = (d * e) - b
		if g == 0 {
			f = 0
		}
		e++
		g = e - b
		if g != 0 {
			goto l12
		}
		d++
		g = d - b
		if g != 0 {
			goto l11
		}
		if f == 0 {
			h++
			g = b
		}
		g -= c
		if g != 0 {
			b -= -17
		} else {
			break
		}
	}
	fmt.Println(a, b, c, d, e, f, g, h)
}

func main2() {
	a, b, c, d, e, f, g, h := 0, 0, 0, 0, 0, 0, 0, 0
	// 1. set b 81
	b = 81
	// 2. set c b
	c = b
	// 3. jnz a 2
	if a != 0 {
		goto l5
	}
	// 4. jnz 1 5
	goto l9
l5:
// 5. mul b 100
	b *= 100
	// 6. sub b -100000
	b -= -100000
	// 7. set c b
	c = b
	// 8. sub c -17000
	c -= -17000
l9:
// 9. set f 1
	f = 1
	// 10. set d 2
	d = 2
l11:
// 11. set e 2
	e = 2
l12:
// 12. set g d
	g = d
	// 13. mul g e
	g *= e
	// 14. sub g b
	g -= b
	// 15. jnz g 2
	if g != 0 {
		goto l17
	}
	// 16. set f 0
	f = 0
l17:
// 17. sub e -1
	e -= -1
	// 18. set g e
	g = e
	// 19. sub g b
	g -= b
	// 20. jnz g -8
	if g != 0 {
		goto l12
	}
	// 21. sub d -1
	d -= -1
	// 22. set g d
	g = d
	// 23. sub g b
	g -= b
	// 24. jnz g -13
	if g != 0 {
		goto l11
	}
	// 25. jnz f 2
	if f != 0 {
		goto x6
	}
	// 26. sub h -1
	h -= -1
	// 27. set g b
	g = b
x6:
// 28. sub g c
	g -= c
	// 29. jnz g 2
	if g != 0 {
		goto l31
	}
	// 30. jnz 1 3
	goto l33
l31:
// 31. sub b -17
	b -= -17
	// 32. jnz 1 -23
	goto l9
l33:
	fmt.Println(a, b, c, d, e, f, g, h)
}
