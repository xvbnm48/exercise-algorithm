soal 3
ways to sum

An automated packaging system is responsible for packing boxes. A box is certified to hold a certain weight. Given an integer total, calculate the number of possible ways to achieve total as sum of the weights of items weighing interger weights from 1 to k, inclusive.

example
total = 8
k = 2

to reach a weight of 8, there are 5 different ways that items with weightr between 1 and 3 can e combined:
[1, 1, 1, 1, 1, 1, 1, 1]
[1, 1, 1, 1, 1, 1, 2]
[1, 1, 1, 1, 2, 2]
[1, 1, 2, 2, 2]
[2, 2, 2, 2]

functions description:
Complete the function ways in the editor below
ways has the following paramters(s):
int total: the value to sum to
int k: the maximum of the range of integers to considers when summing to total
returns
int: the number of ways to sum to the total: the number might be very large, so return the integer modulo 1000000007 (10⁹+7)