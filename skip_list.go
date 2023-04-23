// resource : https://mecha-mind.medium.com/redis-sorted-sets-and-skip-lists-4f849d188a33
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Sample struct {
	values  *[]int
	weights *[]float64
	cumSum  *[]float64
}

func (samp *Sample) init(values *[]int, weights *[]float64) {
	sumWeights := (func(arr *[]float64) float64 {
		tempSum := 0.0
		for _, v := range *arr {
			tempSum += v
		}
		return tempSum
	})(weights)
	tempCumSum := make([]float64, len(*values))

	samp.values = values
	samp.weights = (func(weights *[]float64, weightSum float64) *[]float64 {
		tempWeights := make([]float64, len(*weights))
		if weightSum != 0 {
			for index, elem := range *weights {
				tempWeights[index] = elem / weightSum
			}
		}
		return &tempWeights
	})(weights, sumWeights)
	samp.cumSum = &tempCumSum
	samp.computerCumSum()
}
func (samp *Sample) computerCumSum() {
	s := 0.0
	for index, w := range *samp.weights {
		s += w
		(*samp.cumSum)[index] += s
	}
}
func (samp *Sample) getOne() int {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	randProbability := rng.Float64()

	// perform binary search lol
	left, right, lastTrue := 0, len(*samp.values)-1, 0

	for left <= right {
		mid := left + ((right - left) / 2)

		if (*samp.cumSum)[mid] >= randProbability {
			lastTrue = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return (*samp.values)[lastTrue]
}

type SLNode struct {
	key          *string // int or nil
	val          *int    // string or nil
	nextPointers map[int]*SLNode
	prevPointers map[int]*SLNode
	level        *int // int or nil
}

func (sln *SLNode) init(key *string, val *int, level *int) {
	sln.key = key
	sln.val = val
	sln.level = level
	sln.nextPointers = make(map[int]*SLNode)
	sln.prevPointers = make(map[int]*SLNode)
}

type SkipList struct {
	levels        *[]int
	probabilities *[]float64
	sample        *Sample
	head          *SLNode
	tail          *SLNode
	nodeMap       map[string]*SLNode
}

func (sl *SkipList) init(maxSize int) {

	tempLevels := make([]int, 0)
	for i := 1; i < int(math.Log2(float64(maxSize)))+2; i++ {
		tempLevels = append(tempLevels, i)
	}
	sl.levels = &tempLevels

	tempProbabilities := make([]float64, len(*sl.levels))
	for index := range tempProbabilities {
		if index == 0 {
			tempProbabilities[index] = 1.0
		} else {
			tempProbabilities[index] = 0.5 * tempProbabilities[index-1]
		}
	}
	sl.probabilities = &tempProbabilities

	sl.sample = &Sample{}
	sl.sample.init(sl.levels, sl.probabilities)

	node := SLNode{}

	defaultKey := "__head__"
	defaultLen := len(*sl.levels)
	node.init(&defaultKey, nil, &defaultLen)

	for _, elem := range *sl.levels {
		node.nextPointers[elem] = nil
		node.prevPointers[elem] = nil
	}

	sl.head, sl.tail = &node, &node

	sl.nodeMap = make(map[string]*SLNode)
}
func (sl *SkipList) SearchKey(key string) (int, bool) {
	if elem, ok := sl.nodeMap[key]; ok {
		return *(elem.val), true
	}
	return 0, false
}
func (sl *SkipList) SearchVal(val int) bool {
	currentLevel := (*sl.levels)[len(*sl.levels)-1]
	node := sl.head

	for i := currentLevel; i > 0; i-- {
		for node.nextPointers[i] != nil && *(node.nextPointers[i].val) < val {
			node = node.nextPointers[i]
		}
	}
	return node.nextPointers[1] != nil && *(node.nextPointers[1].val) == val
}
func (sl *SkipList) Insert(key string, val int) {
	if _, ok := sl.nodeMap[key]; !ok {
		currentLevel := (*sl.levels)[len(*sl.levels)-1]
		node := sl.head

		randomLevel := sl.sample.getOne()
		updates := make([]*SLNode, randomLevel)
		for index := range updates {
			updates[index] = nil
		}

		for i := currentLevel; i > 0; i-- {
			for node.nextPointers[i] != nil && *(node.nextPointers[i].val) < val {
				node = node.nextPointers[i]
			}
			if i <= randomLevel {
				updates[i-1] = node
			}
		}
		newNode := SLNode{}
		newNode.init(&key, &val, &randomLevel)

		for i := 1; i <= randomLevel; i++ {
			node = updates[i-1]

			newNode.nextPointers[i] = node.nextPointers[i]

			if node.nextPointers[i] != nil {
				node.nextPointers[i].prevPointers[i] = &newNode
			} else {
				if i == 1 {
					sl.tail = &newNode
				}
			}
			node.nextPointers[i] = &newNode
			newNode.prevPointers[i] = node
		}
		sl.nodeMap[key] = &newNode
	} else {
		sl.Delete(key)
		sl.Insert(key, val)
	}
}
func (sl *SkipList) Delete(key string) bool {
	if elem, ok := sl.nodeMap[key]; ok {
		node := elem

		for i := 1; i <= *(node.level); i++ {
			if node.prevPointers[i] != nil {
				node.prevPointers[i].nextPointers[i] = node.nextPointers[i]
			}
			if node.nextPointers[i] != nil {
				node.nextPointers[i].prevPointers[i] = node.prevPointers[i]
			} else {
				if i == 1 {
					sl.tail = node.prevPointers[i]
				}
			}
		}
		delete(sl.nodeMap, key)
		return true
	}
	return false
}

func (sl *SkipList) GetMax() (string, int, bool) {
	if sl.tail != nil {
		if retKey, retVal := sl.tail.key, sl.tail.val; retKey != nil && retVal != nil {
			return *retKey, *retVal, true
		}
	}
	return "", 1, false
}
func (sl *SkipList) PopMax() (string, int, bool) {
	maxElemKey, maxElemVal, maxElemStatus := sl.GetMax()
	if maxElemStatus != false {
		sl.Delete(maxElemKey)
	}
	return maxElemKey, maxElemVal, maxElemStatus
}

func (sl *SkipList) GetMin() (string, int, bool) {
	if sl.head != nil && sl.head.nextPointers[1] != nil {
		if retKey, retVal := sl.head.nextPointers[1].key, sl.head.nextPointers[1].val; retKey != nil && retVal != nil {
			return *retKey, *retVal, true
		}
	}
	return "", 1, false
}
func (sl *SkipList) PopMin() (string, int, bool) {
	maxElemKey, maxElemVal, maxElemStatus := sl.GetMin()
	if maxElemStatus != false {
		sl.Delete(maxElemKey)
	}
	return maxElemKey, maxElemVal, maxElemStatus
}

func (sl *SkipList) GetPredecessor(val int) (string, int, bool) {
	currentLevel := (*sl.levels)[len(*sl.levels)-1]
	node := sl.head

	for i := currentLevel; i > 0; i-- {
		for node.nextPointers[i] != nil && *(node.nextPointers[i].val) < val {
			node = node.nextPointers[i]
		}
	}
	if node != nil {
		if retKey, retVal := node.key, node.val; retKey != nil && retVal != nil {
			return *retKey, *retVal, true
		}
	}
	return "", 1, false
}
func (sl *SkipList) GetSuccessor(val int) (string, int, bool) {
	currentLevel := (*sl.levels)[len(*sl.levels)-1]
	node := sl.head

	for i := currentLevel; i > 0; i-- {
		for node.nextPointers[i] != nil && *(node.nextPointers[i].val) < val {
			node = node.nextPointers[i]
		}
	}

	if node.nextPointers[1] != nil {
		if *(node.nextPointers[1].val) > val {
			if retKey, retVal := node.nextPointers[1].key, node.nextPointers[1].val; retKey != nil && retVal != nil {
				return *retKey, *retVal, true
			}
		} else {
			if node.nextPointers[1].nextPointers[1] != nil {
				if retKey, retVal := node.nextPointers[1].nextPointers[1].key, node.nextPointers[1].nextPointers[1].val; retKey != nil && retVal != nil {
					return *retKey, *retVal, true
				}
			}
		}
	}
	return "", 1, false
}
func (sl *SkipList) PrintList() {
	for i := len(*sl.levels) - 1; i >= 0; i-- {
		p := sl.head
		out := make([]string, 0)

		for p != nil {
			k, v := "", 0

			if pkey, pval := p.key, p.val; pkey != nil || pval != nil {
				if pkey != nil {
					k = *pkey
				}
				if pval != nil {
					v = *pval
				}
			}
			out = append(out, fmt.Sprintf("%v <-> %v", k, v))
			p = p.nextPointers[(*sl.levels)[i]]
		}

		//for _, elem := range out {
		fmt.Println(out)
		//fmt.Println()
		//}
	}
}

func runner() {
	sl := SkipList{}
	sl.init(1000000)

	sl.Insert("hello", 3)
	sl.Insert("world", 11)
	sl.Insert("this", 18)
	sl.Insert("is", 9)
	sl.Insert("me", 14)
	sl.Insert("reaching", 19)
	sl.Insert("out", 17)
	sl.Insert("to", 6)
	sl.Insert("you", 24)
	sl.Insert("to", 27)
	sl.Insert("talk", 10)
	sl.Insert("about", 2)
	sl.Insert("the", 4)
	sl.Insert("extended", 26)
	sl.Insert("warranty", 29)
	sl.Insert("of", 25)
	sl.Insert("your", 23)
	sl.Insert("car", 21)
	sl.Insert("insaurance", 22)
	sl.Insert("please", 20)
	sl.Insert("response", 5)
	sl.Insert("ASAP", 12)
	sl.Insert("because", 13)
	sl.Insert("this", 1)
	sl.Insert("is", 28)
	sl.Insert("a", 16)
	sl.Insert("limited", 8)
	sl.Insert("time", 15)
	sl.Insert("offer", 7)

	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())

	fmt.Println(sl.SearchVal(20))

	/*	fmt.Println(sl.SearchVal("world"))
		fmt.Println(sl.SearchVal("this"))
		fmt.Println(sl.SearchVal("is"))
		fmt.Println(sl.SearchVal("me"))
		fmt.Println(sl.SearchVal("reaching"))
		fmt.Println(sl.SearchVal("out"))
		fmt.Println(sl.SearchVal("to"))
		fmt.Println(sl.SearchVal("you"))
		fmt.Println(sl.SearchVal("to"))
		fmt.Println(sl.SearchVal("talk"))
		fmt.Println(sl.SearchVal("about"))
		fmt.Println(sl.SearchVal("the"))
		fmt.Println(sl.SearchVal("extended"))
		fmt.Println(sl.SearchVal("warranty"))
		fmt.Println(sl.SearchVal("of"))
		fmt.Println(sl.SearchVal("your"))
		fmt.Println(sl.SearchVal("car"))
		fmt.Println(sl.SearchVal("insaurance"))
		fmt.Println(sl.SearchVal("please"))
		fmt.Println(sl.SearchVal("response"))
		fmt.Println(sl.SearchVal("ASAP"))
		fmt.Println(sl.SearchVal("because"))
		fmt.Println(sl.SearchVal("this"))
		fmt.Println(sl.SearchVal("is"))
		fmt.Println(sl.SearchVal("a"))
		fmt.Println(sl.SearchVal("limited"))
		fmt.Println(sl.SearchVal("time"))
		fmt.Println(sl.SearchVal("offer"))
		fmt.Println(sl.SearchVal("SUS"))
	*/

	//sl.PrintList()

	/*fmt.Println(sl.PopMin())
	fmt.Println(sl.PopMin())
	fmt.Println(sl.PopMin())
	fmt.Println(sl.PopMin())

	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())
	fmt.Println(sl.PopMax())*/

	fmt.Println(sl.GetPredecessor(2))
	fmt.Println(sl.GetPredecessor(13))
	fmt.Println(sl.GetPredecessor(20))

}
