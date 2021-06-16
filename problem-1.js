//part 1: prints sum of all whole numbers 1 to n 

var n = /*replace w/ input value*/
var sum = (n * (n + 1)) / 2; 
console.log(sum);

//part 2: prints sum of all multiples of 3 & 5 to n

function threesAndFives(n, d) {
n = parseInt(n / d);
return n * ((n + 1) * parseInt(d) / 2);
}
function sum(n) {
n--;
return threesAndFives(n, 3) + threesAndFives(n, 5) - threesAndFives(n, 15);
}
var n = /*replace w/ input value*/
console.log(sum(n));
