| Aspect | Nil Slice | Empty Slice |
| --- | --- | --- |
| Definition | A slice that has not been initialized or assigned any value. | A slice that has been initialized but contains no elements. |
| Declaration | `var s []int` | `s := []int{}` or `s := make([]int, 0)` |
| Length (`len(s)`) | 0 | 0 |
| Capacity (`cap(s)`) | 0 | 0 |
| Underlying Array | None | Exists but has zero length. |
| Comparison to `nil` | `true` | `false` |
| Memory Allocation | No memory is allocated. | Memory is allocated for the slice header but not for elements. |
| Appending Elements | Creates a new underlying array. | Uses the existing slice header and allocates memory for elements. |
| Usage | Typically used to indicate the absence of a slice. | Used to represent an initialized but empty collection. |
| Example | `var s []int` | `s := []int{}` or `s := make([]int, 0)` |
