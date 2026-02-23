package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[2:5] //Returns a slice of the original array starting from index 2 to index 4 (5 is exclusive)
	fmt.Println("Slice b:", b)
	c := a[:5] //Returns a slice of the original array starting from index 0 to index 4 (5 is exclusive)
	fmt.Println("Slice c:", c)
	d := a[2:] //Returns a slice of the original array starting from index 2 to the end of the array
	fmt.Println("Slice d:", d)
	e := a[:] //Returns a slice of the original array starting from index 0 to the end of the array
	fmt.Println("Slice e:", e)
	a[3] = 10 //Modifying the original array will affect all slices that reference it
	fmt.Println("Modified array a:", a)
	fmt.Println("Slice b after modification:", b)
	fmt.Println("Slice c after modification:", c)
	fmt.Println("Slice d after modification:", d)
	fmt.Println("Slice e after modification:", e)
	// Thus, slices are reference types and share the same underlying array. Modifying the original array will affect all slices that reference it.
	// Appending an element to the original array creates a new array and does not affect the original array or the slices that reference it. However, if we append to the original array a, it will create a new array and the slices will still reference the old array, which may lead to unexpected results.
	a = append(a, 6)
	fmt.Println("Array a after appending 6:", a)
	fmt.Println("Slice b after appending 6 to array a:", b)
	fmt.Println("Slice c after appending 6 to array a:", c)
	fmt.Println("Slice d after appending 6 to array a:", d)
	fmt.Println("Slice e after appending 6 to array a:", e)

	f := a[:6:7] //Returns a slice of the original array starting from index 0 to index 5 (6 is exclusive) with a capacity of 7
	fmt.Println("Slice f:", f)
	fmt.Printf("Length of slice f: %d, Capacity of slice f: %d\n", len(f), cap(f))

	f = append(f, 100, 120, 130) //Appending elements to slice f will not affect the original array a because it has a capacity of 7 and we are appending 3 elements which exceeds the capacity, thus creating a new array for slice f.
	fmt.Println("Slice f after appending 100, 120, 130:", f)
	fmt.Printf("Length of slice f after appending: %d, Capacity of slice f after appending: %d\n", len(f), cap(f))
	fmt.Println("Array a after appending to slice f:", a)

	g := make([]int, 6)
	copy(g, a)
	fmt.Println("Array g (copy of a):", g) // The copy function creates a new slice g and copies the elements of array a into it. Modifying array g will not affect array a or any of the slices that reference it because g is a separate copy of the original array.
	// Modifying array g will not affect array a or any of the slices that reference it because g is a separate copy of the original array.
	g[0] = 100
	fmt.Println("Array g after modification:", g)
	fmt.Println("Array a after modifying g:", a)

	h := [6]int{1, 2, 3, 4, 5, 6}
	i := make([]int, 6)
	copy(i, h[:])
	fmt.Println("Array i (copy of h):", i)
	// Thus copy function also helps in bridging the gap between arrays and slices. We can copy elements from an array to a slice, and the resulting slice will have its own underlying array that is separate from the original array. Modifying the slice will not affect the original array, and modifying the original array will not affect the slice.

	// Slicing and copy function work on array too, similarly to slices.
}
