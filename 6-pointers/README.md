Maps in Go:

A map is implemented as a pointer to a struct in the Go runtime
When you pass a map to a function, you're actually copying a pointer
This means modifications to the map inside the function will affect the original map
The author cautions against using maps in public APIs because:

They lack explicit structure
It's hard to know what keys exist without tracing through the code
They are not self-documenting



The author recommends:

Use structs instead of maps when possible
Maps are best used when the keys are not known at compile time

Slices in Go:
A slice is implemented as a struct with three fields:

An integer for length
An integer for capacity
A pointer to a block of memory

Slice Behavior When Passed to Functions:

When a slice is copied, a new copy is made of:

The length
The capacity
The pointer to the underlying memory



Modifying Slice Contents:

Changing values in the slice affects the shared underlying memory
This means changes are visible in both the original and the copied slice

Appending to Slices:
Scenario 1 (Enough Capacity):

If you append to a slice copy and there's enough capacity
New values are stored in the shared memory block
However, the length of the original slice remains unchanged
The original slice cannot "see" the new values beyond its original length

Scenario 2 (Not Enough Capacity):

If you append to a slice copy and there's not enough capacity
A new, larger memory block is allocated
Values are copied to the new block
The pointer, length, and capacity in the copy are updated
Changes to these fields are NOT reflected in the original slice

Key Takeaways:

Slice contents can be modified in a function
Slice cannot be resized in a function
Always document whether a function modifies a slice's contents
Slices are great for passing around linear data structures in Go

The author suggests treating slices as not modifiable by default and explicitly documenting any modifications in your function's documentation.

A garbage collector doesn’t immediately reclaim memory as soon as it is no longer
referenced. Doing so would seriously impact performance. Instead, it lets the
garbage pile up for a bit. The heap almost always contains both live data and
memory that’s no longer needed. The Go runtime provides users a couple of
settings to control the heap’s size. The first is the GOGC environment variable. The
garbage collector looks at the heap size at the end of a garbage-collection cycle and
uses the formula CURRENT_HEAP_SIZE + CURRENT_HEAP_SIZE*GOGC/100 to calculate
the heap size that needs to be reached to trigger the next garbage-collection cycle.
Note
The GOGC heap size calculation is a little more complicated than just
described. It takes into account not just the heap size, but the sizes of all
the stacks of all the goroutines and the memory set aside to hold package-
level variables. Most of the time, the heap size is far bigger than the size of
these other memory areas, but in some situations they do have an effect.
By default, GOGC is set to 100, which means that the heap size that triggers the next
collection is roughly double the heap size at the end of the current collection.
Setting GOGC to a smaller value will decrease the target heap size, and setting it to a
larger value will increase it. As a rough estimate, doubling the value of GOGC will
halve the amount of CPU time spent on GC.