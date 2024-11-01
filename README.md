Intersection of Two Linked Lists


Given two linked lists, find the node at which they intersect.

A: A1 -> A2 -> A3 ->-
                     | -> C1, C2, C3
B: B1 -> B2 -> B3 ->-
Here the linked lists A and B intersect at C1.

If the two lists do not intersect, return null.

Testing
Input Format
The first line contains ‘T’ denoting the number of independent test cases.

For each test case, the input has five lines:

Two space-separated integers L1 and L2 denoting the size of the lists A and B before they intersect.
L1 space-separated integers denoting the node values for the first L1 nodes of A.
L2 space-separated integers denoting the node values for the first L2 nodes of B.
An integer L3 denoting the length of the list after A and B intersect.
L3 space-separated integers denoting the node values for the common nodes of A and B.
Output Format
For each test case, a line with the value of the node where A and B intersect. If they do not intersect, output -1.

Sample Input
3
3 3
4 5 6
7 8 9
3
1 2 3
0 4

4 5 6 7
2
8 9
2 3
1 2
3 4 5
0

Expected Output
1
8
-1