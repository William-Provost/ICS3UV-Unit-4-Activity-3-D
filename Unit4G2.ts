/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-30
 * @fileoverview This program uses indecies in an array.
 */

// constants and variables
// Initializes an array of 4 floating-point numbers
const decimalValues: number[] = new Array(4);
let index: number = 0;

// input (Initialization of array elements)
decimalValues[0] = 10.5;
decimalValues[1] = 20.5;
decimalValues[2] = 30.5;
decimalValues[3] = 40.5;

// output
// set index variable to 3 so you can see
// how to use a variable for an index value in an array
index = 3;

// Prints decimalValues[3]
console.log(`Cell 3: ${decimalValues[index].toFixed(2)}`);
// Prints decimalValues[3 - 1] = decimalValues[2]
console.log(`Cell 2: ${decimalValues[index - 1].toFixed(2)}`);

// index = 3 - 2 = 1
index = index - 2;

// Prints decimalValues[1]
console.log(`Cell 1: ${decimalValues[index].toFixed(2)}`);
// Prints decimalValues[1 - 1] = decimalValues[0]
console.log(`Cell 0: ${decimalValues[index - 1].toFixed(2)}`);

console.log("\nDone.");
