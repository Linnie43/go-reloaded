# go-reloaded

A simple Go-based text processing tool for text completion, editing, and auto-correction. This project demonstrates file handling, string manipulation, and text formatting in Go.

---

## Features

This tool can:

- Convert numbers from **hexadecimal** `(hex)` or **binary** `(bin)` to decimal.
- Change text **case**:
  - `(up)` → uppercase  
  - `(low)` → lowercase  
  - `(cap)` → capitalized  
  - Apply to multiple words: `(up, N)`, `(low, N)`, `(cap, N)`  
- Correct spacing and placement of punctuation: `. , ! ? : ;`  
- Handle multi-character punctuation like `...` and `!?`  
- Properly format text inside **single quotes `' '`**  
- Replace the article `a` with `an` if the next word starts with a vowel or `h`.

---

## Usage

You will have to provide some input files for this.

```bash
$ go run . input.txt output.txt
```

### Example 1: Case formatting
Input:  it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom...
Output: It was the best of times, it was the worst of TIMES, it was the age of wisdom...
### Example 2: Number conversion
Input: Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
Output: Simply add 66 and 2 and you will see the result is 68.
### Example 3: Punctuation handling
Input: Punctuation tests are ... kinda boring ,what do you think ?
Output: Punctuation tests are... kinda boring, what do you think?
