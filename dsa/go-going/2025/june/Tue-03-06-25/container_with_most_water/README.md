# Container With Most Water

Given an array representing heights of vertical lines, find the maximum area of water trapped between two lines.

Input: [1,8,6,2,5,4,8,3,7].

Output: 49.

# Approach
1. Use two pointers, one at the start and one at the end of the array.
2. Calculate the area between the two pointers.
3. Move the pointer pointing to the shorter line inward, as this will potentially increase the area.

# Result
 - Time Complexity: O(n)
 - Space Complexity: O(1)

