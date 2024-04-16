Divide: The unsorted array or slice is recursively divided into smaller subarrays until each subarray contains only one element.
Conquer: Individual elements and subarrays are treated as sorted.
Merge: Sorted subarrays are merged to create larger sorted subarrays until the entire array or slice is sorted.

The important step of merge sort lies in the merging step, where sorted subarrays are combined to form larger sorted arrays. This approach ensures stability (equal elements maintain their original order) and guarantees a time complexity of O(n log n) for average, best, and worst-case scenarios. Moreover, it has a space complexity of O(n) due to the temporary storage required for merging subarrays during its divide and conquer process.

