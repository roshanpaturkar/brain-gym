# Minimum Window Substring

Given two strings, <code>original</code> and <code>check</code>, return the minimum substring of <code>original</code> such that each character in <code>check</code>, including duplicates, are included in this substring. By "minimum", I mean the shortest substring. If two substrings that satisfy the condition have the same length, the one that comes lexicographically first is smaller.


## Parameters
- <code>original</code>: The original string.
- <code>check</code>: The string to check if a window contains it.


## Result
The smallest substring of <code>original</code> that satisfies the condition.


## Examples
<b>Input</b>: <code>original = "cdbaebaecd", check = "abc"</code>
<br>
<b>Output</b>: <code>baec</code>
<br>
<b>Explanation</b>: baec is the shortest substring of <code>original</code> that contains all of a, b, and c.

<b>Input</b>: <code>original = "aabbcc", check = "abc"</code>
<br>
<b>Output</b>: <code>abc</code>
<br>
<b>Explanation</b>: abc is the shortest substring of <code>original</code> that contains all of a, b, and c.

## Constraints
- <code>1 <= len(check), len(original) <= 10^5</code>
- <code>original</code> and <code>check</code> both contain only uppercase and lowercase characters in English. The characters are case sensitive.

## Time Complexity
The time complexity of this algorithm is O(n), where n is the length of the original string. This is because we are using a sliding window approach, and each character is processed at most twice (once when expanding the window and once when contracting it).

## Space Complexity
The space complexity is O(m), where m is the number of unique characters in the check string. This is because we are using a hash map to store the frequency of characters in the check string and another hash map to store the frequency of characters in the current window of the original string.