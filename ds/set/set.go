package set

/* 
 *  Quick unordered Set data structure using a map.
 *  The Set is optimized towards lookups and insertions
 *  which should run in O(1) time. Order is not guarenteed.
 *  Might need to change the implementation of it
 *  if I find out using a bool type is not that efficient.
 */

type Set struct {
    set map[interface{}]bool
}

// Initializes a new Set struct and returns a pointer to it.
func NewSet() *Set {
    s := new(Set)
    s.set = make(map[interface{}]bool)
    return s
}

// Checks if an element is contained in the map. Returns true or false.
func (s Set) Contains(item interface{}) bool {
    return s.set[item] == true;
}

// Adds an element into the set. Since the set is unordered this will
// take constant time but will not be in order necessarily.
func (s *Set) Add(element interface{}) {
    s.set[element] = true
}

// Removes the given element if it exists. Might change this to return the element.
func (s *Set) Remove(element interface{}) {
    delete(s.set, element)
}

// Returns the size of the Set
func (s Set) Size() int {
    return len(s.set)
}

// Checks whether the Set is empty or not
func (s Set) IsEmpty() bool {
    return len(s.set) == 0
}
