package ninjascript
import (
math_rand "math/rand"
)
func init() {if _, ok := Api["math/rand"]; !ok {
   Api["math/rand"] = map[string]interface{}{}
}
Api["math/rand"]["ExpFloat64"] = math_rand.ExpFloat64
Api["math/rand"]["Float32"] = math_rand.Float32
Api["math/rand"]["Float64"] = math_rand.Float64
Api["math/rand"]["Int"] = math_rand.Int
Api["math/rand"]["Int31"] = math_rand.Int31
Api["math/rand"]["Int31n"] = math_rand.Int31n
Api["math/rand"]["Int63"] = math_rand.Int63
Api["math/rand"]["Int63n"] = math_rand.Int63n
Api["math/rand"]["Intn"] = math_rand.Intn
Api["math/rand"]["NormFloat64"] = math_rand.NormFloat64
Api["math/rand"]["Perm"] = math_rand.Perm
Api["math/rand"]["Read"] = math_rand.Read
Api["math/rand"]["Seed"] = math_rand.Seed
Api["math/rand"]["Shuffle"] = math_rand.Shuffle
Api["math/rand"]["Uint32"] = math_rand.Uint32
Api["math/rand"]["Uint64"] = math_rand.Uint64
Api["math/rand"]["Rand"] = math_rand.Rand{}
Api["math/rand"]["New"] = math_rand.New
Api["math/rand"]["NewSource"] = math_rand.NewSource
Api["math/rand"]["Zipf"] = math_rand.Zipf{}
Api["math/rand"]["NewZipf"] = math_rand.NewZipf
Api["math/rand"]["ExpFloat64"] = math_rand.ExpFloat64
Api["math/rand"]["Float32"] = math_rand.Float32
Api["math/rand"]["Float64"] = math_rand.Float64
Api["math/rand"]["Int"] = math_rand.Int
Api["math/rand"]["Int31"] = math_rand.Int31
Api["math/rand"]["Int31n"] = math_rand.Int31n
Api["math/rand"]["Int63"] = math_rand.Int63
Api["math/rand"]["Int63n"] = math_rand.Int63n
Api["math/rand"]["Intn"] = math_rand.Intn
Api["math/rand"]["NormFloat64"] = math_rand.NormFloat64
Api["math/rand"]["Perm"] = math_rand.Perm
Api["math/rand"]["Read"] = math_rand.Read
Api["math/rand"]["Seed"] = math_rand.Seed
Api["math/rand"]["Shuffle"] = math_rand.Shuffle
Api["math/rand"]["Uint32"] = math_rand.Uint32
Api["math/rand"]["Uint64"] = math_rand.Uint64
Api["math/rand"]["New"] = math_rand.New
Api["math/rand"]["NewSource"] = math_rand.NewSource
Api["math/rand"]["NewZipf"] = math_rand.NewZipf

}