package skiplist

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/option"
	"golang.org/x/exp/constraints"
)

// SkiplistMaxLevel maximum levels allocated for each Skiplist
// next pointer arrays are of constant size
const SkiplistMaxLevel = 30

// MinProb Minimum probability of bernoulli trial success
// for random function
const MinProb = 0.01

type Node[A any] struct {
	value       A
	next        [SkiplistMaxLevel]*Node[A]
	marked      bool
	fullyLinked bool
	mux         sync.Mutex
	topLevel    int
}

// * Skiplist : The Skiplist structure, must be initialised before use.
type Skiplist[A any] struct {
	ord interfaces.Order[A]
	log interfaces.Logger

	nLevels    int
	head       *Node[A]
	nElements  int
	prob       float64
	maxLevels  int
	lock       sync.RWMutex
	fastRandom bool
}

func (list *Skiplist[A]) Height() int {
	/* current max level */
	defer list.lock.RUnlock()
	list.lock.RLock()
	return list.nLevels
}

func (list *Skiplist[A]) Len() int {
	/* current dataAmount of inserted elements */
	defer list.lock.RUnlock()
	list.lock.RLock()
	return list.nElements
}

func (list *Skiplist[A]) setMaxLevels(levels int) {
	defer list.lock.Unlock()
	list.lock.Lock()
	list.maxLevels = min(max(1, levels), SkiplistMaxLevel)
}

func (list *Skiplist[A]) setProb(prob float64) {
	defer list.lock.Unlock()
	list.lock.Lock()
	list.prob = min(max(MinProb, prob), 1.0)
}

func (list *Skiplist[A]) setFastRandom(isSet bool) {
	defer list.lock.Unlock()
	list.lock.Lock()
	list.fastRandom = isSet
}

func New[A any](prob float64, maxLevels int, fastRandom bool) *Skiplist[A] {
	list := new(Skiplist[A])
	if prob < 0 {
		prob = 0.5
		fmt.Println("Init: Probability given less than zero, set to 0.5 instead")
	}
	if maxLevels > SkiplistMaxLevel {
		fmt.Println("Init: Max level given more than supported dataAmount of",
			SkiplistMaxLevel, " setting to ", SkiplistMaxLevel, "instead")
		maxLevels = SkiplistMaxLevel
	}
	list.nLevels = 1
	list.prob = prob
	list.maxLevels = maxLevels
	list.fastRandom = fastRandom

	newHead := new(Node[A])
	newHead.fullyLinked = true
	newHead.marked = false

	list.nElements = 0
	list.head = newHead
	return list
}

func (list *Skiplist[A]) ToSortedArray() []A {
	/* make a sorted array out of the Skiplist
	   returns the lowest level               */
	arr := make([]A, list.nElements, list.nElements)
	counter := 0
	for currentNode := list.head.next[0]; currentNode != nil; currentNode = currentNode.next[0] {
		arr[counter] = currentNode.value
		counter++
	}
	return arr
}

func (list *Skiplist[A]) findNextLowest(val A) (node *Node[A]) {
	// could be modified by inserts
	list.lock.RLock()
	level := list.nLevels - 1
	list.lock.RUnlock()
	// much faster than starting at max

	pred := list.head
	var curr *Node[A]
	// traverse vertically
	for ; level >= 0; level-- {
		// horizontally
		curr = pred.next[level]
		for curr != nil && list.ord.Compare(curr.value, val) < 0 {
			pred = curr
			curr = pred.next[level]
		}
		// next of where it should be
		if curr != nil && list.ord.Compare(curr.value, val) == 0 {
			break
		}
	}
	return curr
}

func (list *Skiplist[A]) Find(val A, prev, next []*Node[A]) (foundLevel int) {
	// could be modified by inserts
	list.lock.RLock()
	level := list.nLevels - 1
	list.lock.RUnlock()
	// much faster than starting at max

	pred := list.head
	foundLevel = -1
	var curr *Node[A]

	// traverse vertically
	for ; level >= 0; level-- {
		// horizontally
		curr = pred.next[level]
		for curr != nil && list.ord.Compare(curr.value, val) < 0 {
			pred = curr
			curr = pred.next[level]
		}
		// next of where it should be
		if curr != nil && list.ord.Compare(curr.value, val) == 0 && foundLevel == -1 {
			foundLevel = level
		}
		// previous of where the item should be
		prev[level] = pred
		next[level] = curr
	}
	return foundLevel
}

func (list *Skiplist[A]) Contains(val A) bool {

	list.lock.RLock()
	level := list.nLevels - 1
	list.lock.RUnlock()

	pred := list.head
	var curr *Node[A]
	// vertically
	for ; level >= 0; level-- {
		// horizontally
		curr = pred.next[level]
		for curr != nil && list.ord.Compare(curr.value, val) < 0 {
			pred = curr
			curr = pred.next[level]
		}
		//found something or have to go down

		// is the next element what I seek
		if curr != nil && list.ord.Compare(curr.value, val) == 0 {
			node := curr
			return node.fullyLinked && !node.marked
		}
	}
	// not found
	return false
}

func (list *Skiplist[A]) Get(val A) option.Option[A] {
	list.lock.RLock()
	level := list.nLevels - 1
	list.lock.RUnlock()

	pred := list.head
	var curr *Node[A]
	// vertically
	for ; level >= 0; level-- {
		// horizontally
		curr = pred.next[level]
		for curr != nil && list.ord.Compare(curr.value, val) < 0 {
			pred = curr
			curr = pred.next[level]
		}
		//found something or have to go down

		// is the next element what I seek
		if curr != nil && list.ord.Compare(curr.value, val) == 0 && curr.fullyLinked && !curr.marked {
			return option.From(&curr.value)
		}
	}
	// not found
	return option.From[A](nil)
}

func (list *Skiplist[A]) Insert(v A) bool {
	// insert element

	// highest level of insertion
	// the list.fast property should not be modified after init
	topLevel := coinTosses(list.prob, list.maxLevels, list.fastRandom)

	// check if list must become taller
	list.lock.Lock()
	if topLevel > list.nLevels {
		list.nLevels = topLevel
	}
	list.lock.Unlock()

	// buffers to store prev and next pointers
	var prev, next []*Node[A]
	prev = make([]*Node[A], SkiplistMaxLevel)
	next = make([]*Node[A], SkiplistMaxLevel)

	for {
		// find insertion point and previous and next nodes
		foundLevel := list.Find(v, prev, next)

		// already in Skiplist
		if foundLevel != -1 {
			// should be the node with value v
			nodeFound := next[foundLevel]
			// if node is not set for removal
			if !nodeFound.marked {
				// wait until stable
				for !nodeFound.fullyLinked {
				}
				//don't insert
				return false
			}
			// try again
			continue
		}
		// highest level locked
		highestLocked := -1
		var pred, succ *Node[A]
		var prevPred *Node[A]

		valid := true

		// validate that new node can be added
		// by checking previous and next nodes
		for level := 0; valid && level < topLevel; level++ {
			pred = prev[level]
			succ = next[level]

			// avoid locking same node twice
			// if two or more levels
			// connected to same node
			if pred != prevPred {
				pred.mux.Lock()

				highestLocked = level
				prevPred = pred
			}
			// can the insertion proceed
			// node is locked so we can check next
			valid = !pred.marked && (succ == nil || !succ.marked) && pred.next[level] == succ
		}

		// cannot add
		if !valid {
			// unlock to try again
			prevPred = nil
			for i := highestLocked; i >= 0; i-- {
				if prevPred != prev[i] {
					prev[i].mux.Unlock()
				}
				prevPred = prev[i]
			}
			// restart attempt
			continue
		}

		// try to add new node
		newNode := new(Node[A])
		newNode.value = v
		newNode.topLevel = topLevel - 1
		newNode.marked = false

		for level := 0; level < topLevel; level++ {
			newNode.next[level] = next[level]
			prev[level].next[level] = newNode
		}
		// new node is ok
		newNode.fullyLinked = true

		//unlock
		prevPred = nil
		for i := highestLocked; i >= 0; i-- {
			if prevPred != prev[i] {
				prev[i].mux.Unlock()
			}
			prevPred = prev[i]
		}
		list.lock.Lock()
		list.nElements = list.nElements + 1
		list.lock.Unlock()
		return true
	}
}

func (list *Skiplist[A]) Remove(val A) bool {
	/* remove node */
	var nodeToDelete *Node[A]
	isMarked := false
	topLevel := -1

	var prev, next [SkiplistMaxLevel]*Node[A]

	for {
		// try to find node
		foundLevel := list.Find(val, prev[:], next[:])

		// if not found or already marked for deletion
		// return false
		if isMarked || (foundLevel != -1 && canDelete(next[foundLevel], foundLevel)) {
			// not already marked
			if !isMarked {
				// get node
				nodeToDelete = next[foundLevel]
				topLevel = nodeToDelete.topLevel
				// lock it
				nodeToDelete.mux.Lock()
				// did some other routine
				// mark it first?
				if nodeToDelete.marked {
					// yes, unlock and abort
					nodeToDelete.mux.Unlock()
					return false
				}
				// no mark it for deletion
				nodeToDelete.marked = true
				isMarked = true
			}

			// now locked
			highestLocked := -1
			var pred, succ *Node[A]
			var prevPred *Node[A]

			// validate levels up to topLevel
			valid := true
			for level := 0; valid && level <= topLevel; level++ {
				pred = prev[level]
				succ = next[level]

				if pred != prevPred {
					pred.mux.Lock()
					highestLocked = level
					prevPred = pred
				}
				valid = !pred.marked && pred.next[level] == succ
			}

			// can't delete try again
			if !valid {
				// unlock to try again
				prevPred = nil
				for i := highestLocked; i >= 0; i-- {
					if prevPred != prev[i] {
						prev[i].mux.Unlock()
					}
					prevPred = prev[i]
				}
				continue
			}
			// actually delete node
			for level := topLevel; level >= 0; level-- {
				prev[level].next[level] = nodeToDelete.next[level]
			}

			nodeToDelete.mux.Unlock()

			// cleanup and unlock
			prevPred = nil
			for i := highestLocked; i >= 0; i-- {
				if prevPred != prev[i] {
					prev[i].mux.Unlock()
				}
				prevPred = prev[i]
			}
			// update element count
			list.lock.Lock()
			list.nElements--
			list.lock.Unlock()
			return true
		}
		return false
	}
}

func min[A constraints.Ordered](a, b A) A {
	if a < b {
		return a
	}
	return b
}

func max[A constraints.Ordered](a, b A) A {
	if a > b {
		return a
	}
	return b
}

const mask = ((1 << SkiplistMaxLevel) - 1)

func coinTosses(prob float64, maxLevels int, fast bool) (counter int) {
	counter = 1
	// very fast with probability 0.5
	// only one call to rand
	// find first zero in random float bit representation
	// geometric distribution
	if fast {
		resMask := rand.Uint64() & mask
		// find first zero in float representation
		for ; resMask&1 == 0; resMask >>= 1 {
			counter++
		}
		return counter
	}
	// supports probability
	// slower
	res := rand.Float64()
	for res < prob {
		res = rand.Float64()
		counter++
	}
	return counter
}

func canDelete[A any](candidate *Node[A], foundLevel int) bool {
	return candidate.fullyLinked && candidate.topLevel == foundLevel && !candidate.marked
}
