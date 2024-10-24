# go-longread-canvas

- trimmomatic for the long reads, no heap memory allocation required, runs on goroutines and later adding the waitgroup.
- uses the same as a single line one pattern only file, see the pattern.txt as how to define the pattern. 
- parsed 100GB fastq file in few seconds. 
- give a pattern files and ports the goroutines for the long reads. 

```
╭─gauavsablok@gauravsablok ~/Desktop/go/go-longread-canvas ‹main●›
╰─$ go run main.go -h
look for the matching patterns

Usage:
  longread [flags]

Flags:
  -h, --help              help for longread
  -L, --longread string   long read file to be checked (default "path to the long read file")
  -P, --pattern string    pattern file (default "path to the file containing the patterns")
exit status 1
```
-detailed usage
```
╭─gauavsablok@gauravsablok ~/Desktop/go/go-longread-canvas ‹main●›
╰─$ go run main.go -L sample.fastq -P pattern.txt
The canavssed reads have been written
exit status 1
╭─gauavsablok@gauravsablok ~/Desktop/go/go-longread-canvas ‹main●›
╰─$ cat sample.fastq                                                                                                                                                1 ↵
@ERR10930361.1 magdelm64071_201030_115446/27
ATACTTTAAATTTTAGTTACTATTATTATTATTTAAAAAAAAAAAAAAAATTGAAAGTATATCCAAACTAGCACTCAATTAATGCAAACAATATTGTAACTTGGAAAAAGTTCAAAGGCATGAAAATCAAGATTTTTCATTTAATCAACCGGGTCAAGCAAAGTATGA
@ERR10930361.1 magdelm64071_201030_115446/27
ATACTTTAAATTTTAGTTACTATTATTATTATTTAAAAAAAAAAAAAAAATTGAAAGTATATCCAAACTAGCACTCAATTAATGCAAACAATATTGTAACTTGGAAAAAGTTCAAAGGCATGAAAATCAAGATTTTTCATTTAATCAACCGGGTCAAGCAAAGTATGA
╭─gauavsablok@gauravsablok ~/Desktop/go/go-longread-canvas ‹main●›
╰─$ cat canvased.fastq
@ERR10930361.1
ATACTTTAAATTTTAGTTACTATTATTATTATTTTTGAAAGTATATCCAAACTAGCACTCAATTAATGCAAACAATATTGTAACTTGGAAAAAGTTCAAAGGCATGAAAATCAAGATTTTTCATTTAATCAACCGGGTCAAGCAAAGTATGA
@ERR10930361.1
ATACTTTAAATTTTAGTTACTATTATTATTATTTTTGAAAGTATATCCAAACTAGCACTCAATTAATGCAAACAATATTGTAACTTGGAAAAAGTTCAAAGGCATGAAAATCAAGATTTTTCATTTAATCAACCGGGTCAAGCAAAGTATGA
```

Gaurav Sablok
