// Date: 10/10/21

// # Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// # You may assume that each input would have exactly one solution, and you may not use the same element twice.
// # You can return the answer in any order.

// # Example 1:
// # Input: nums = [2,7,11,15], target = 9
// # Output: [0,1]
// # Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

const twoSum = (nums, target) => {
    const numsMap = new Map();
    
    for (let i = 0; i < nums.length; i++) {
        if (numsMap.get(target - nums[i]) !== undefined) {
            return [numsMap.get(target - nums[i]), i];
        }
        numsMap.set(nums[i], i);
    }
    return null;
};

console.log(twoSum([2, 7, 11, 15], 9));
console.log(twoSum([3, 2, 4], 6));
console.log(twoSum([3,3], 6));