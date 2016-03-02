package gago

import "math/rand"

// The Initializer is here to create the first generation of individuals in a
// deme. It applies to an individual level and instantiates it's genome gene by
// gene.
type Initializer interface {
	apply(individual *Individual, generator *rand.Rand)
}

// FloatUniform generates random floating point values between given boundaries.
type FloatUniform struct {
	lower, upper float64
}

// Apply the FloatUniform initializer.
func (fu FloatUniform) apply(indi *Individual, generator *rand.Rand) {
	for i := range indi.Genome {
		// Decide if positive or negative
		var gene float64
		if generator.Float64() < 0.5 {
			gene = generator.Float64() * fu.lower
		} else {
			gene = generator.Float64() * fu.upper
		}
		indi.Genome[i] = gene
	}
}

// StringUniform generates random string slices based on a given corpus.
type StringUniform struct {
	Corpus []string
}

// Apply the StringUniform initializer.
func (su StringUniform) apply(indi *Individual, generator *rand.Rand) {
	for i := range indi.Genome {
		indi.Genome[i] = su.Corpus[generator.Intn(len(su.Corpus))]
	}
}

// StringUnique generates random string slices based on a given corpus, each
// element from the corpus is only represented once in each slice. The method
// starts by shuffling, it then assigns the elements of the corpus in increasing
// index order to an individual. Usually the length of the individual's genome
// should match the length of the corpus.
type StringUnique struct {
	Corpus []string
}

// Apply the StringUnique initializer.
func (su StringUnique) apply(indi *Individual, generator *rand.Rand) {
	var strings = shuffleStrings(su.Corpus, generator)
	for i := range indi.Genome {
		indi.Genome[i] = strings[i]
	}
}