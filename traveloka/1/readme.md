
soal 1
1. Slowest Key Press

Engineers have redesigned a keypad used by ambulance drivers in urban areas. In order to determine which key takes the longest time to press, the keypad is tested by a driver. Given the results of that test, determine which key takes the longest to press.

examples
keyTimes = [[0, 2], [1, 5], [0, 9], [2, 15]]

Elements in keyTimes[i][0]  represent encoded characters in the range  where a = 0, b = 1, ..., z = 25. The second element, keyTimes[i][1] represents the time the key is pressed since the start of the test.   The elements will be given in ascending time order. In the example, keys pressed, in order are 0102[encoded] = abac at times 2, 5, 9, 15.  From the start time, it took 2 - 0 = 2 to press the first key, 5 - 2 = 3 to press the second, and so on. The longest time it took to press a key was key 2 or 'c' at 15 - 9 = 6

Complete the function slowestKey  in the editor below.
slowestKey has the following parameter(s):
int keyTimes[n][2]: the first column contains the encoded key pressed, the second contains the time at which it was pressed

returns
char: the key that took the longerst time to press


solution:
To solve this problem in Go, we can follow these steps:

Initialize two variables to keep track of the longest duration and the key associated with that duration. Let's call these variables maxDuration and slowestKey, respectively.

Iterate through the keyTimes array to calculate the duration for each key press. Since the times are in ascending order, the duration for each key press can be calculated by subtracting the time of the current key press from the time of the previous key press. For the first key press, the duration is simply its press time since we consider the start time to be 0.

Update maxDuration and slowestKey whenever we find a key press duration that is longer than the current maxDuration.

After iterating through all the key presses, convert the slowestKey from its encoded form (where 0 represents 'a', 1 represents 'b', and so on) to the actual character. This can be done by adding the ASCII value of 'a' to the slowestKey.

Return the character associated with the slowestKey.