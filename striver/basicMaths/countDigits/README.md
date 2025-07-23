# Count Digits in a Number

## Problem Statement

You are given an integer `n`. You need to return the number of digits in the number.

The number will have no leading zeroes, except when the number is 0 itself.

## Examples

**Example 1:**
```
Input: n = 4
Output: 1
Explanation: There is only 1 digit in 4.
```

**Example 2:**
```
Input: n = 14
Output: 2
Explanation: There are 2 digits in 14.
```

**Example 3:**
```
Input: n = -123
Output: 3
Explanation: There are 3 digits in -123 (sign is not counted).
```

**Example 4:**
```
Input: n = 0
Output: 1
Explanation: There is 1 digit in 0.
```

## Constraints

- `-2^31 <= n <= 2^31 - 1`

## Function Signature

```go
func countDigits(n int) int
```

## Task

Complete the implementation of the `countDigits` function in `CountDigits.go`.

## Approach Ideas

1. **Mathematical Division**: Keep dividing by 10 and count iterations
2. **String Conversion**: Convert to string and get length (try to avoid this)
3. **Logarithmic**: Use mathematical properties with log10

## Key Points to Remember

- Handle the special case when `n = 0`
- Negative numbers: the sign doesn't count as a digit
- Think about time and space complexity

## How to Test

### Run your implementation:
```bash
go run CountDigits.go
```

### Run comprehensive tests:
```bash
go test -v
```

### Run specific test function:
```bash
go test -run TestCountDigits
go test -run TestCountDigitsEdgeCases
```

## Follow-up Questions

1. What is the time complexity of your solution?
2. What is the space complexity of your solution?
3. Can you solve this without converting to string?
4. Can you think of multiple approaches?

## Expected Behavior

When you run `go run CountDigits.go`, you should see your function tested with basic examples.

When you run `go test -v`, you should see comprehensive test results including edge cases.

Good luck! 🚀
