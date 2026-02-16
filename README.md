# matrix-determinant-go

Matrix Determinant in Go

This is a simple Go program to calculate the **determinant of a 3x3 matrix**.

## Features

- Reads a 3x3 matrix from user input.
- Calculates the determinant using the standard formula:

\[
\text{det}(A) = a(ei - fh) - b(di - fg) + c(dh - eg)
\]

where the matrix is:
| a b c |
| d e f |
| g h i |


- Prints the determinant to the console.

## Usage

1. Clone the repository:

```bash
git clone <repo-url>
cd <repo-folder>
go run main.go

Enter the matrix elements 3x3:
1 2 3
4 5 6
7 8 9

Determinant = 0
```
Example

Matrix:
| 2 1 3 |
| 0 -1 4 |
| 5 2 1 |

Output:
Determinant = 29


Requirements

Go 1.18+ installed on your system.
Compatible with Windows, Linux, and macOS.
