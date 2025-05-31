# Product of Array Except Self

Given an integer array nums, return an array answer such that <code>answer[i]</code> is equal to the product of all the elements of nums except <code>nums[i]</code>.

Input: [1, 2, 3, 4].

Output: [24, 12, 8, 6].

## Approach
The problem can be solved using a two-pass approach with prefix and suffix products.
- Calculate prefix products
- Calculate suffix products
- Multiply prefix and suffix products to get the final result

## Return the result
- Time Complexity: O(n)
- Space Complexity: O(1) (excluding the output array)
