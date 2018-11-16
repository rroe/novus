### Comparing summing all numbers (via list reduction) between [1, 10000)

They all got the same results using incredibly similar code, so what do we get?

#### PHP
php test.php  0.08s user 0.01s system 97% cpu 0.092 total

#### Python
python test.py  0.02s user 0.03s system 93% cpu 0.051 total

#### Novus
./novus test.novus  1.84s user 0.27s system 175% cpu 1.201 total

Scaled up (times 100), thats:

PHP: 8
Python: 2
Novus: 184

So Novus is 23 times slower than PHP, and 92 times slower than Python.

Speed sorta needs to be worked on. Just a bit.