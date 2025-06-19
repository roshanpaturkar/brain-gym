# 657. Robot Return to Origin

This LeetCode problem involves determining whether a robot, which starts at the origin (0, 0) on a 2D plane, returns to its starting point after executing a series of moves. The series of moves is given as a string where each character corresponds to a move in one of four possible directions: 'R' (right), 'L' (left), 'U' (up), and 'D' (down). The task is to analyze this string and return true if the robot ends up at the origin after all the moves or false otherwise. The key point is to keep track of the robot's position relative to the origin, regardless of how it's facing, and verify if it returns to (0, 0).

## Result
- Time Complexity: O(n), where n is the length of the moves string.
- Space Complexity: O(1), as we only use a constant amount of space for the position coordinates.
