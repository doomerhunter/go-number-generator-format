# go-number-generator-format

Burp Intruder style number generator for fuzzing purposes. Generates sequential or random numbers based on user-provided parameters.
Numbers outputted will follow the min and max number of digits provided (similarly to Burp Intruder).

## Example usage 

Generates all numbers from 1 to 10, with 3 digits, no floats

```
./main -from=15 -to=20 -minint=3 -maxint=3 -float=false -mode=sequential 
015
016
017
018
019
020
```

Generate 3 random numbers between 10 and 20, with 3 integer digits and 3 fraction digits : 

```
./main -from=10 -to=20 -minint=3 -maxint=3 -minfrac 3 -maxfrac 3 -float=true -mode=random -n 3

018.190
020.349
018.052
```

## Compilation

```
go build main.go
```

## Usage

```
./main -h
Usage of ./main:
  -float
    	Include floating point numbers. Default = false
  -from int
    	Start of range. Default = 0
  -maxfrac int
    	Maximum fraction digits. Default = 0
  -maxint int
    	Maximum integer digits. Default = 0
  -minfrac int
    	Minimum fraction digits. Default = 0
  -minint int
    	Minimum integer digits. Default = 0
  -mode string
    	Generation mode (sequential or random). Default = sequential (default "sequential")
  -n int
    	Number of random numbers to generate. Default = 0
  -to int
    	End of range. Default = 0

```
