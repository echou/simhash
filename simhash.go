package simhash

const (
  HashSize = 64
)

func CalculateSimHash(input string) int64 {
  tokeniser := NewOverlappingStringTokeniser(4, 3)
  hashedTokens := getHashTokens(tokeniser.Tokenise(input))
  vector := make([]int, HashSize)
  for i, _ := range vector {
    vector[i] = 0
  }

  for _, v := range hashedTokens {
    for i, _ := range vector {
      if isBitSet(uint64(v), uint64(i)) {
        vector[i] += 1
      } else {
        vector[i] -= 1
      }
    }
  }

  fingerprint := uint64(0)
  for i, v := range vector {
    if v > 0 {
      fingerprint += 1 << uint64(i)
    }
  }

  return fingerprint
}

func GetLikenessValue(needle, haystack string) float64 {
  needleSimHash := CalculateSimHash(needle)
  hayStackSimHash := CalculateSimHash(haystack)
  return float64(HashSize-GetHammingDistance(needleSimHash, hayStackSimHash)) / float64(HashSize)
}
