# ModTables
GoLang program to generate modular addition and multiplication tables, as well as latin squares from these.

# Usage
2 command line arguments, 1st, whether you want a multiplicative, additive, or latin_square table, and 2nd, the number to use for modulo, and also the number of rows and columns wih the same number.

# To-Do
- Add symetry checker.
- Implement interactive mode.

# Example Output
```modtable.exe multiplicative 9
| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| 0 | 2 | 4 | 6 | 8 | 1 | 3 | 5 | 7 |
| 0 | 3 | 6 | 0 | 3 | 6 | 0 | 3 | 6 |
| 0 | 4 | 8 | 3 | 7 | 2 | 6 | 1 | 5 |
| 0 | 5 | 1 | 6 | 2 | 7 | 3 | 8 | 4 |
| 0 | 6 | 3 | 0 | 6 | 3 | 0 | 6 | 3 |
| 0 | 7 | 5 | 3 | 1 | 8 | 6 | 4 | 2 |
| 0 | 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1 |
