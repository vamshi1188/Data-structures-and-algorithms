# Armstrong Number

## Problem Statement

Given an integer N, return true if it is an Armstrong number otherwise return false.

An Armstrong number is a number that is equal to the sum of its own digits each raised to the power of the number of digits.

## Examples

**Example 1:**
```
Input: N = 153
Output: True
Explanation: 1³ + 5³ + 3³ = 1 + 125 + 27 = 153
```

**Example 2:**
```
Input: N = 371
Output: True
Explanation: 3³ + 7³ + 1³ = 27 + 343 + 1 = 371
```

**Example 3:**
```
Input: N = 1634
Output: True
Explanation: 1⁴ + 6⁴ + 3⁴ + 4⁴ = 1 + 1296 + 81 + 256 = 1634
```

**Example 4:**
```
Input: N = 123
Output: False
Explanation: 1³ + 2³ + 3³ = 1 + 8 + 27 = 36 ≠ 123
```

## Constraints

- `0 <= N <= 10^9`
- Consider only non-negative integers

## Function Signature

```go
func isArmstrong(n int) bool
```

## Task

Complete the implementation of the `isArmstrong` function in `ArmstrongNumber.go`.

## Algorithm Steps

1. **Count the digits** in the number
2. **Extract each digit** one by one
3. **Calculate the power** of each digit raised to the number of digits
4. **Sum all the powers**
5. **Compare** the sum with the original number

## Key Points to Remember

- Single digits (0-9) are Armstrong numbers (e.g., 5¹ = 5)
- The power is equal to the number of digits in the number
- For 3-digit numbers: each digit is raised to power 3
- For 4-digit numbers: each digit is raised to power 4
- Use appropriate data types to handle large calculations

## Common Armstrong Numbers

- **1-digit**: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
- **3-digit**: 153, 371, 407
- **4-digit**: 1634, 8208, 9474
- **5-digit**: 54748, 92727, 93084

## Algorithm Hint

```go
// Step 1: Count digits
digits := countDigits(n)

// Step 2: Extract digits and calculate sum
temp := n
sum := 0
while temp > 0 {
    digit := temp % 10
    sum += power(digit, digits)  // digit^digits
    temp = temp / 10
}

// Step 3: Check if sum equals original number
return sum == n
```

## How to Test

### Run your implementation:
```bash
go run ArmstrongNumber.go
```

### Run comprehensive tests:
```bash
go test -v
```

### Run specific test function:
```bash
go test -run TestIsArmstrong
go test -run TestIsArmstrongEdgeCases
```

## Follow-up Questions

1. What is the time complexity of your solution?
2. What is the space complexity of your solution?
3. How would you handle very large numbers?
4. Can you optimize the digit counting or power calculation?

## Expected Behavior

When you run `go run ArmstrongNumber.go`, you should see your function tested with basic examples.

When you run `go test -v`, you should see comprehensive test results including various Armstrong numbers and edge cases.

Good luck! 🚀
