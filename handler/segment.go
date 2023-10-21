package handler

func GetSegments(s string) []string {
	segments := make([]string, 0)
	prevPtr := 0
	i := 0

	for ; i < len(s); i++ {
		if s[i] == '/' {
			if i != 0 {
				segments = append(segments, s[prevPtr:i])
			}
			prevPtr = i + 1
		}
	}

	if i != prevPtr {
		segments = append(segments, s[prevPtr:i])
	}

	return segments
}
