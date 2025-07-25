# Greatest Common Divisor (GCD) Problem

## Problem Statement
Given two integers N1 and N2, find their greatest common divisor.

The Greatest Common Divisor of any two integers is the largest number that divides both integers.

## Examples

### Example 1:
- **Input:** N1 = 9, N2 = 12
- **Output:** 3
- **Explanation:** 
  - Factors of 9: 1, 3, 9
  - Factors of 12: 1, 2, 3, 4, 6, 12
  - Common Factors: 1, 3 
  - Greatest common factor: 3

### Example 2:
- **Input:** N1 = 20, N2 = 15
- **Output:** 5
- **Explanation:**
  - Factors of 20: 1, 2, 4, 5, 10, 20
  - Factors of 15: 1, 3, 5, 15
  - Common Factors: 1, 5
  - Greatest common factor: 5

## Approach

### 1. Euclidean Algorithm (Iterative)
The most efficient approach uses the Euclidean algorithm:
- **Time Complexity:** O(log(min(a,b)))
- **Space Complexity:** O(1)

**Algorithm:**
```
while b ≠ 0:
    temp = b
    b = a % b
    a = temp
return a
```

### 2. Euclidean Algorithm (Recursive)
Same algorithm implemented recursively:
- **Time Complexity:** O(log(min(a,b)))
- **Space Complexity:** O(log(min(a,b))) due to recursion stack

**Algorithm:**
```
if b = 0:
    return a
else:
    return gcd(b, a % b)
```

### 3. Brute Force (Alternative)
Check all numbers from min(a,b) down to 1:
- **Time Complexity:** O(min(a,b))
- **Space Complexity:** O(1)

## Edge Cases Handled
- One or both numbers are 0
- Negative numbers
- Same numbers
- Coprime numbers (GCD = 1)

## Files
- `GCD.go` - Main implementation with both iterative and recursive solutions
- `GCD_test.go` - Comprehensive test cases including benchmarks
- `go.mod` - Go module file

## How to Run

### Run the program:
```bash
go run GCD.go
```

### Run tests:
```bash
go test
```

### Run tests with verbose output:
```bash
go test -v
```

### Run benchmarks:
```bash
go test -bench=.
```

## Key Learning Points
1. **Euclidean Algorithm** is the most efficient method for finding GCD
2. **Recursive vs Iterative** - Both have same time complexity, but iterative uses less space
3. **Edge case handling** is crucial for robust implementation
4. **Mathematical properties** - GCD(a,0) = a, GCD(a,b) = GCD(b, a%b)
