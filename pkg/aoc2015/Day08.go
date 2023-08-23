package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/8">Day 8</a>
 */

type Day08 struct{}

func (d Day08) encodeSantasList(input []string) (int, int) {
	oldEncoding, newEncoding := 0, 0
	for _, str := range input {
		dir := str[1:(len(str) - 1)]
		oldEncoding += 2 + d.computeInMemory(dir)
		newEncoding += 4 + d.computeEncrypted(dir)
	}

	return oldEncoding, newEncoding
}

func (d Day08) computeInMemory(input string) int {
	count, flag := 0, 0
	for _, ch := range []rune(input) {
		switch ch {
		case '\\':
			if flag == 1 {
				count += 1
				flag = 0
			} else {
				flag = 1
			}
		case '"':
			if flag == 1 {
				count += 1
			}
			flag = 0
		case 'x':
			if flag == 1 {
				flag = 2
			} else {
				flag = 0
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f':
			switch flag {
			case 3:
				count += 3
				flag = 0
			case 2:
				flag = 3
			default:
				flag = 0
			}
		}
	}

	return count
}

func (d Day08) computeEncrypted(input string) int {
	count, flag := 0, 0
	for _, ch := range []rune(input) {
		switch ch {
		case '\\':
			if flag == 1 {
				count += 2
				flag = 0
			} else {
				flag = 1
			}
		case '"':
			if flag == 1 {
				count += 2
			}
			flag = 0
		case 'x':
			if flag == 1 {
				flag = 2
			} else {
				flag = 0
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f':
			switch flag {
			case 3:
				count += 1
				flag = 0
			case 2:
				flag = 3
			default:
				flag = 0
			}
		}
	}

	return count
}
