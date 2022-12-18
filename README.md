# Huffman
##### By Vedaant Kuchhal, Anmol Sandhu
Implementation of the Huffman encoding algorithm in Go.

## What is Huffman encoding?

Huffman encoding is a lossless data compression algorithm. The idea is to assign variable-length encoding to input characters where lengths of the assigned codes are based on the frequencies of corresponding characters. The most frequent character gets the smallest encoding and the least frequent character has the largest encoding, therefore saving space.

## Use cases:
A common use case for Huffman encoding is in image and video file formats, where it is often used to compress the data before it is stored or transmitted. For example, the JPEG image format and the H.264 video format both use Huffman encoding as part of their compression schemes.

Huffman encoding is also often used in data transmission systems, such as wireless communication networks and satellite systems, where bandwidth is limited and it is important to transmit data efficiently.

## How it works:

Taking the following string as an example:
```
aaaaaaaaaakkkkkkkkdjdjdjdjdjdjdjdjaaaaaadjdjf
```
Without Huffman encoding, this string, in ASCII format, would be represented using 8 bits per character, for a total of 8 * 45 = 360 bits
### Step 1:
Calculate the frequency of each character in the string
| 16 | 8 | 10 | 10 | 1 |
|---|---|---|---|---|
| a | k | d | j | f |

We do this via a simple loop through the string, aggregating data using a map:
```go
freq_map := make(map[byte]int)
for i:=0; i<len(data); i++ {
    freq_map[byte(data[i])] += 1
}
```
Note how the characters are stored as type `byte`; this is essentially an unsigned 8-bit integer that Go can use to store characters in ASCII format - it's being typecast from the default 32-bit `rune` format which stores characters in UTF-8 encoding.

### Step 2:
The encoding process utilizes two different abstract data types - a min-heap and a binary tree. To reduce redundancy and simplify implementation, we use the same base node struct `Node` for both the min-heap and the tree. The structure is the following:
```go
type Node struct {
    symbol byte
    freq int
    left  *Node
    right *Node
}
```
The `Node` holds the character as a byte (see previous note) and its frequency as calculated in [Step 1](#Step-1). The `left` and `right` variables are used in the binary tree.

### Step 3:
The first step is to sort the characters in increasing order of frequency.
<!-- | 1 | 8 | 10 | 10 | 16 |
|---|---|---|---|---|
| f | k | d | j | a | -->
We do this by using a min-heap. We use a min-heap since it has the best worst-case runtime  for inserting and deleting the minimum value - in both operations this is `O(log(n))`. 

The characters and corresponding frequencies are inserted as nodes to the min-heap:
```go
h := NewMinHeap()
for symbol, freq := range agg_data {
    h.Insert(&Node{symbol, freq, nil, nil})
}
```
(The two `nil` values for the `left` and `right` variables will be relevant in the next step.)

### Step 4:
Once we have everything inserted into a min-heap, we build the binary tree we will use for encoding the characters.

The first step is to remove the two nodes with the smallest frequencies from the heap.
```go
// Left and right children in subtree are minimum frequencies
left := h.Pop()
right := h.Pop()
```

Create a new node with a placeholder symbol of `$` with a frequency that is the sum of the two smallest nodes. These nodes become the left and right children of the new aggregate node.
```go
// New parent node is sum of frequencies of children
tmp := &Node{'$', left.freq + right.freq, left, right}
```
Insert this node back into the min-heap:
```go
// Insert parent node into min-heap
h.Insert(tmp)
```
Overall, the new array now looks like the following:
| 9 | 10 | 10 | 16 |
|---|---|---|---|
| $ | d | j | a |

`$` is the root node of the following tree:
```
9/
├── 8
└── 1
```
Note that, for clarity, we are choosing a sorted array table to visualize the min-heap.

Summarized code for this step:
```go
// Left and right children in subtree are minimum frequencies
left := h.Pop()
right := h.Pop()
// New parent node is sum of frequencies of children
tmp := &Node{'$', left.freq + right.freq, left, right}
// Insert parent node into min-heap
h.Insert(tmp)
```
### Step 5:
Repeat this process until there is only one node left in the min-heap, by running the above code snippet in a simple `while` loop:
```go
for h.size != 1 {
    // ... insert above code here
}
```
Since this operation repeats `n-1` times and each step takes `log(n)`, the overall time complexity is `O(nlog(n))`.

The final min-heap array looks like this
|45 |
|---|
| $ |

Where this node is the root node of the following binary tree:

```
45/
├── 26/
│   ├── 16
│   └── 10
└── 19/
    ├── 10
    └── 9/
        ├── 8
        └── 1
```

### Step 6:
To encode the symbols, assign `0` to the left edge and `1` to the right edge of each non-leaf node. The symbols are encoded by walking down from the root to the relevant leaf. For example, using the following tree, the code for `k` is `001`. Huffman uses variable length encoding, so the code for `j` is `10`. The codes can still be decoded since the complete binary tree implementation ensures that all codes have a unique prefix. In this example, there can be no symbol corresponding to `00` since it is not a leaf, so the decoder 'knows' to continue to the next bit until it reaches a valid encoding.

#### Tree:
```
45/
├1── 26/
│   ├1── 16(a)
│   └0── 10(j)
└0── 19/
    ├1── 10(d)
    └0── 9/
        ├1── 8(k)
        └0── 1(f)
```
The table below built using the tree stores the mapping of the symbols to the code. This information (i.e. the first and third column) needs to be included with the compressed data, and counts towards its total size.

#### Table:
| Character       | Frequency | Code   | Size     |
|-----------------|-----------|--------|----------|
| f               | 1         | 000    | 1*3 = 3  |
| k               | 8         | 001    | 8*3 = 24 |
| d               | 10        | 01     | 10*2 = 20|
| j               | 10        | 10     | 10*2 = 20|
| a               | 16        | 11     | 16*2 = 32|
| 5 * 8 = 40 bits |           | 12 bits| 99 bits  |

## Analysis
### Size Calculation:
Huffman encoding is the most effective when the data has less uniqueness and high repetition. In our example, there is a pretty significant memory saving:
```
Total size of compressed data = (message size) + (lookup table size)
Total size of compressed data = 99 + 52 = 151 bits
Percentage size reduction (360 bits -> 151 bits) ~ 58%
```
### Time complexity:
The time complexity of Huffman coding depends on the size of the input data and the number of unique characters in the data. The plot below shows the time complexity for building the tree which, as we mentioned earlier, is `O(nlog(n))`.

```
Timings
Increasing size of set of unique characters
14563
|                                                                              a
|                                                                            a  
|                                                                          a    
|                                                                       a a     
|                                                                   aa a        
|                                                                 a             
|                                                                               
|                                                               a               
|                                                            a a                
|                                                         a a                   
|                                                                               
|                                                      a a                      
|                                                                               
|                                                                               
|                                                     a                         
|                                                                               
|                                                   a                           
|                                                aa                             
|                                           a aa                                
|                                          a                                    
|                                        a                                      
|                                       a                                       
|                                  aa a                                         
|                                a                                              
|                             a                                                 
|                            a  a                                               
|                                                                               
|                          a                                                    
|                     a aa                                                      
|                                                                               
|                    a                                                          
|                 aa                                                            
|               a                                                               
|            a a                                                                
|                                                                               
|         aa                                                                    
|      aa                                                                       
|   aa                                                                          
| a                                                                             
+-------------------------------------------------------------------------------
```

## Citations 

Huffman coding. Programiz. From https://www.programiz.com/dsa/huffman-coding 

Plotting source code from Riccardo Pucella <riccardo.pucella@olin.edu>