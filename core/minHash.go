package core

import (
	"hash/fnv"
)

func generateShingles(text string, n int) []string {
	shingles := []string{}

	for i := 0; i <= len(text)-n; i++ {
		shingles = append(shingles, text[i:i+n])
	}
	return shingles
}

func hashValue(shingle string, seed uint32) uint32 {
	h := fnv.New32a()
    h.Write([]byte(shingle))
    a := seed * 12345 + 67890
    b := seed * 54321 + 9876
    return (a * h.Sum32() + b) % 15485863
}

func computeMinHashSignature(shingles []string, numHashFunctions int) []uint32 {
	signature := make([]uint32, numHashFunctions)

	for i := 0; i < numHashFunctions; i++ {
		seed := uint32(i)
		minHash := uint32(^uint32(0))

		for _, shingle := range shingles {
			hash := hashValue(shingle, seed)
			if hash < minHash {
				minHash = hash
			}
		}
		signature[i] = minHash
	}
	return signature
}

func computeJaccardSimilarity(signature1, signature2 []uint32) float64 {
	matches := 0
	
	for i := range signature1 {
		if signature1[i] == signature2[i] {
			matches++
		}
	}
	return float64(matches) / float64(len(signature1))
}

func MinHash(textOne, textTwo string) float64 {
	formattedOne := RemoveStopWords(textOne)
	formattedTwo := RemoveStopWords(textTwo)
	shinglesOne := generateShingles(formattedOne, 3)
	shinglesTwo := generateShingles(formattedTwo, 3)

	signatureOne := computeMinHashSignature(shinglesOne, 100)
	signatureTwo := computeMinHashSignature(shinglesTwo, 100)

	similarity := computeJaccardSimilarity(signatureOne, signatureTwo)

	return similarity * 100
}