package main

func plusOne(digits []int) []int {
	reverse(digits)
}

func reverse(l []int) {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}

func main() {

}
