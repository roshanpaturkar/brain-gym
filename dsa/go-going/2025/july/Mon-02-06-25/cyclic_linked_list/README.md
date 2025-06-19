# Linked List Cycle

Given a linked list with potentially a loop, determine whether the linked list from the first node contains a cycle in it. For bonus points, do this with constant space.

Parameters
- nodes: The first node of a linked list with potentially a loop.
  
Result
- Whether there is a loop contained in the linked list.

Constraints
- 1 <= len(nodes) <= 10^5

## Approach
- Using Floyd's Cycle-Finding Algorithm (Tortoise and Hare)
- Tortoise moves one step at a time, hare moves two steps at a time
- If there is a cycle, they will eventually meet
- If there is no cycle, hare will reach the end of the list
## Return the result
- Time Complexity: O(n)
- Space Complexity: O(1)