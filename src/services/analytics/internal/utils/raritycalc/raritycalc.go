package raritycalc

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/pkg/errors"
)

type Weighted interface {
	comparable
	GetWeight() int32
}

type Percentaged interface {
	comparable
	GetPercentage() int32
}

func SelectByPercentage[T Percentaged](objectsSortedByPercentage []T) (T, error) {
	var selected T
	r, err := rand.Int(rand.Reader, big.NewInt(101))
	if err != nil {
		return selected, errors.Wrapf(err, "failed generating rand int for select")
	}

	// Select item rarity
	for _, object := range objectsSortedByPercentage {
		if r.Int64() >= int64(object.GetPercentage()) {
			continue
		}

		return object, nil
	}

	return selected, fmt.Errorf("logical error when selecting by percentage")
}

func SelectByWeight[T Weighted](amount, target int32, objects []T) ([]T, error) {
	// Calculate the total weight of all objects
	totalWeight := int64(0)
	for _, obj := range objects {
		totalWeight += int64(obj.GetWeight())
	}

	selected := make([]T, 0, amount)
	for len(selected) < int(amount) {
		// Generate a random value within the total weight range
		randValueBig, err := rand.Int(rand.Reader, big.NewInt(totalWeight))
		if err != nil {
			return nil, errors.Wrapf(err, "failed generating rand int")
		}

		randValue := randValueBig.Int64()

		// Iterate over the objects and subtract their weights from the random value until it becomes negative
		for _, obj := range objects {
			randValue -= int64(obj.GetWeight())
			if randValue < 0 {
				selected = append(selected, obj)
				break
			}
		}
	}

	return selected, nil
}

func SpreadByPercentage[T Percentaged](amount int32, objectsSortedByPercentage []T, guaranteeBestRarityPer int32) ([]T, error) {
	var selected []T
	bestRarityCount := int32(0)

	for i := int32(0); i < amount; i++ {
		object, err := SelectByPercentage(objectsSortedByPercentage)
		if err != nil {
			return nil, err
		}

		selected = append(selected, object)

		if object == objectsSortedByPercentage[len(objectsSortedByPercentage)-1] {
			bestRarityCount++
		}
	}

	// Guarantee at least X best rarity items
	if guaranteeBestRarityPer != 0 {
		guaranteedBestRarityAmount := amount / guaranteeBestRarityPer
		for i := int32(0); i < guaranteedBestRarityAmount-bestRarityCount; i++ {
			selected[i] = objectsSortedByPercentage[len(objectsSortedByPercentage)-1]
		}
	}

	return selected, nil
}

func SpreadByWeight[T Weighted](amount int32, objects []T, filterFunc func(*T) bool) ([]T, error) {
	var selected []T

	// First, we'll calculate the total chance of all the item types
	var totalChance int32
	for _, object := range objects {
		object := object
		if !filterFunc(&object) {
			continue
		}

		totalChance += object.GetWeight()
	}

	for i := int32(0); i < amount; i++ {
		randomNumBig, err := rand.Int(rand.Reader, big.NewInt(101))
		if err != nil {
			return nil, errors.Wrapf(err, "failed generating random int for spread by weight")
		}

		randomNum := float64(randomNumBig.Int64()) / 100 * float64(totalChance)

		for _, object := range objects {
			object := object
			if !filterFunc(&object) {
				continue
			}

			randomNum -= float64(object.GetWeight())
			if randomNum <= 0 {
				selected = append(selected, object)
				break
			}
		}
	}

	return selected, nil
}

// SelectNumber generates a number between min and max.
// The number will be closer to max the bigger the weight is, with diminishing returns in relation to maxWeight
func SelectNumber[T int | int32 | int64 | uint | uint32 | uint64](min, max, weight, maxWeight T) (T, error) {
	var result T
	// Calculate the size of the range
	rangeSize := max - min

	// Normalize the weight between 0 and 1
	normalizedWeight := float64(weight) / float64(maxWeight)

	// Calculate the weighted range based on the normalized weight
	weightedRange := float64(rangeSize-1) * normalizedWeight

	// Calculate the maximum adjusted range that doesn't exceed the maximum value of uint32
	maxAdjustedRange := math.Min(float64(rangeSize-1)+1.0, math.MaxUint32-float64(min))

	// Adjust the range by adding 1 to ensure inclusivity of the maximum value
	adjustedRange := math.Min(weightedRange+1.0, maxAdjustedRange)

	// Generate a random value between 0 and the adjusted range
	randomValueBig, err := rand.Int(rand.Reader, big.NewInt(101))
	if err != nil {
		return result, err
	}

	randomValue := float64(randomValueBig.Int64()) / 100 * adjustedRange

	// Round the random value to the nearest integer
	roundedValue := math.Round(randomValue)

	// Convert the rounded value to an integer and add it to the minimum value
	result = min + T(roundedValue)

	// Return the generated number
	return result, nil
}
