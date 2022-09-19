# Build a Wall

### Build

Just type this in the main directory

```bash
$ go build
```

### Usage

```bash
$ ./build-a-wall <rows> <bricks>
```

Number of rows and bricks must be a positive integer.

### Examples

Here you can see some examples of the expected output:

```bash
$> ./build-a-wall 5 5
     ■■|■■|■■|■■|■■
     ■|■■|■■|■■|■■|■
     ■■|■■|■■|■■|■■
     ■|■■|■■|■■|■■|■
     ■■|■■|■■|■■|■■

$> ./build-a-wall 10 7
     ■|■■|■■|■■|■■|■■|■■|■
     ■■|■■|■■|■■|■■|■■|■■
     ■|■■|■■|■■|■■|■■|■■|■
     ■■|■■|■■|■■|■■|■■|■■
     ■|■■|■■|■■|■■|■■|■■|■
     ■■|■■|■■|■■|■■|■■|■■
     ■|■■|■■|■■|■■|■■|■■|■
     ■■|■■|■■|■■|■■|■■|■■
     ■|■■|■■|■■|■■|■■|■■|■
     ■■|■■|■■|■■|■■|■■|■■

$> ./build-a-wall "eight" [3]
    null
    
$> ./build-a-wall 12 -4
    null
    
$> ./build-a-wall 123 1024
    Naah, too much...here's my resignation.
```

