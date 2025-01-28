package core

const (
	k           = 5
	base uint64 = 256
	mod  uint64 = 1e9 + 7
)

func computerHashes(text string, n int) map[uint64]bool {
	hashes := make(map[uint64]bool)
	for i := 0; i < len(text)-n; i++ {
		chunk := text[i : i+n]
		hash := rabinKarpHash(chunk, base, mod)
		hashes[hash] = true
	}
	return hashes
}

func rabinKarpHash(s string, base, mod uint64) uint64 {
	var hash uint64
	for _, c := range s {
		hash = (hash*base + uint64(c)) % mod
	}
	return hash
}

func similarity(set1, set2 map[uint64]bool) float64 {
	matches := 0
	for h := range set1 {
		if set2[h] {
			matches++
		}
	}
	union := len(set1) + len(set2) - matches
	if union == 0 {
		return 0.0
	}
	return float64(matches) / float64(union)
}

func RabinKarp(textOne, textTwo string) float64 {
	hashSetOne := computerHashes(textOne, k)
	hashSetTwo := computerHashes(textTwo, k)

	similarity := similarity(hashSetOne, hashSetTwo)

	return similarity * 100
}
