Knowledge about basic bit operations with &, |, ~, ^, is required.  

Bit shifting:  
Let k = number of shifting bits, x = number to be shifted.  
left shift << append zeros to the right  
result = 2^k * x (number x is of the right)  
1 << 1 = 2  
1 << 3 = 8  
3 << 2 = 12  
3 << 3 = 2  
right shift << append zeros to the left  
result = x/2^k (number x is of the left)  
4 >> 1 = 2  
16 >> 1 = 8  
16 >> 2 = 4  
16 >> 3 = 2  

1 << uint(x) generates bit with one bit set and x-1 zero bits. Subtracting 1 from it results in full set bit
Example: 1 << uint(x-1) = (1000), hence (1000) - 1 = (111). 8 - 1 = 7  

Concatenation "trick"  
x & (x-1) = find the rightmost 1 in x and flip it and the rest to the right
x = 6 = (110), x - 1 = 5 = (101). x^(x-1) = (110) & (101) = (100)
binary representation of x is the same as binary representation of x & (x-1) except the rightmost 1
i.e. x = 6 = (110), flip only the rightmost 1 and you get (100)  

Function Odd/Even:
Bits ending with 1 are odd.
1 == 00000001, hence x & 00000001 will give us 00000001 or 00000000, that is 1 or 0  

Finding specific number: ("Negation trick")
Formula: ~x = -(x+1)
Example: ~x = -(x+1) = 0 <=> x = -1  

Function CheckIfTheBitIsSet:
Let's use 2^i; i is the bit no and belongs to [0..n];
1 = (1), 2 = (10), 3= (100), 4 =(1000)...  
So if x & 2^i = 2^i, it means the i-th bit is set.  

Function IsPowerOfTwo:  
Numbers which are powers of 2 will all be zeros for the x & (x - 1)  

Other uses:  
Switching between 1 and 2:  
Let x=1 or 2, then x^3 = 2 or 1