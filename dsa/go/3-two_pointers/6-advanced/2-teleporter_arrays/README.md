# Teleporter Arrays

You are given two sorted arrays of distinct integers, <code>arr1</code>, and <code>arr2</code>. Your goal is to start from the beginning of one array, and arrive at the end of one array (it could be the same array or not).

For each step, you can either move forward a step on an array, or move to a square in the other array where the number is the same as the number in your current square ("teleportation"). Your total score is defined as the sum of all unique numbers that you have been on.

Find the maximum score that you can get given the above rules. Since the result might be very large and cause overflow, return the maximum score modded by <code>10^9 + 7</code>.


## Parameters
 - <code>arr1</code>: A list of ordered, distinct integers.
 - <code>arr2</code>: Another list of ordered, distinct integers.
 
 
 ## Result
 - The maximum score possible, modded by <code>10^9 + 7</code>.


 ## Examples:
 <b>Input</b>: <code>arr1 = [2, 4, 5, 8, 10], arr2 = [4, 6, 8, 9]</code>
<br>
 <b>Output</b>: <code>30</code>

<b>Input</b>: <code>arr1 = [1, 3, 5, 7], arr2 = [3, 5, 6]</code>
<br>
 <b>Output</b>: <code>15</code>

## Constraints:
 - <code>1 <= len(arr1), len(arr2) <= 50000</code>
 - <code>1 <= arr1[i], arr2[i] <= 10^7</code>
 - <code>arr1[i] < arr1[j]</code> for all <code>i < j</code>. Same goes for <code>arr2</code>.


## Time Complexity:
 - <code>O(n + m)</code>, where <code>n</code> is the length of <code>arr1</code> and <code>m</code> is the length of <code>arr2</code>.


## Space Complexity:
 - <code>O(1)</code>, since we are not using any extra space that grows with the input size.