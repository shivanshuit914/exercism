package secret

const testVersion = 1

func Handshake(number int) []string {
	if number < 0 {
		return nil
	}
	var out []string
	if number&1 != 0 {
		out = append(out, "wink")
	}
	if number&2 != 0 {
		out = append(out, "double blink")
	}
	if number&4 != 0 {
		out = append(out, "close your eyes")
	}
	if number&8 != 0 {
		out = append(out, "jump")
	}

	if number&16 != 0 {
		reverse(out)
	}

	return out
}

func reverse(out []string) {
	for left, right := 0, len(out)-1; left < right; left, right = left+1, right-1 {
		out[left], out[right] = out[right], out[left]
	}
}
