# Reverse Number

## Problem Statement

Given an integer N, return the reverse of the given number.

**Note:** If a number has trailing zeros, then its reverse will not include them. For example, reverse of 10400 will be 401 instead of 00401.

## Examples

**Example 1:**
```
Input: N = 12345
Output: 54321
Explanation: The reverse of 12345 is 54321.
```

**Example 2:**
```
Input: N = 7789
Output: 9877
Explanation: The reverse of number 7789 is 9877.
```

**Example 3:**
```
Input: N = 10400
Output: 401
Explanation: The reverse of 10400 is 401 (trailing zeros are not included).
```

**Example 4:**
```
Input: N = -123
Output: -321
Explanation: The reverse of -123 is -321 (sign is preserved).
```

## Constraints

- `-2^31 <= N <= 2^31 - 1`
- Handle integer overflow cases appropriately

## Function Signature

```go
func reverseNumber(n int) int
```

## Task

Complete the implementation of the `reverseNumber` function in `ReverseNumber.go`.

## Approach Ideas

1. **Mathematical Approach**: Extract digits one by one and build the reverse
2. **String Conversion**: Convert to string, reverse, and convert back (try to avoid this)
3. **Modulo and Division**: Use `%` and `/` operations to extract and build digits

## Key Points to Remember

- Handle negative numbers (preserve the sign)
- Trailing zeros in the original number should not appear as leading zeros in the result
- Be careful about integer overflow
- Think about the special case when `N = 0`

## Algorithm Hint

1. Extract the last digit using `n % 10`
2. Remove the last digit using `n / 10`
3. Build the reverse number by multiplying previous result by 10 and adding the extracted digit
4. Repeat until `n` becomes 0

## How to Test

### Run your implementation:
```bash
go run ReverseNumber.go
```

### Run comprehensive tests:
```bash
go test -v
```

### Run specific test function:
```bash
go test -run TestReverseNumber
go test -run TestReverseNumberEdgeCases
```

## Follow-up Questions

1. What is the time complexity of your solution?
2. What is the space complexity of your solution?
3. How would you handle integer overflow?
4. Can you solve this without converting to string?

## Expected Behavior

When you run `go run ReverseNumber.go`, you should see your function tested with basic examples.

When you run `go test -v`, you should see comprehensive test results including edge cases like trailing zeros and negative numbers.

Good luck! 🚀
