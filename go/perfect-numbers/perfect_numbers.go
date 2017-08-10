package perfect

import "errors"

const testVersion = 1

var (
	ErrOnlyPositive = errors.New("Only positive Numbers")
)

type Classification int

const (
	ClassificationPerfect = iota
	ClassificationAbundant
	ClassificationDeficient
	ClassificationFailed
)

func Classify(num uint64) (Classification, error) {
	if num == 0 {
		return ClassificationFailed, ErrOnlyPositive
	}
	if num == 1 {
		return ClassificationDeficient, nil
	}
	var sum uint64 = 1
	for i := uint64(2); i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	if sum == num {
		return ClassificationPerfect, nil
	}
	if sum > num {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}
