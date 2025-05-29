# Subarray Sum Total

Find the total number of subarrays that sums up to target.

Example:

Consider the array [1, 2, 3] with a target of 3.

- Start with {0: 1} in the hashmap (sum of empty subarray).
- At index 0: sum = 1, add {1: 1} to hashmap.
- At index 1: sum = 3, it equals target, count = 1. Add {3: 1} to hashmap.
- At index 2: sum = 6, (6 - 3) = 3 exists in hashmap, count += 1. Update {6: 1} in hashmap.
- Final count: 2 (subarrays [1, 2] and [3])