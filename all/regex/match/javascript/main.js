
let str = "This is a string in javascript";

let r1 = /stri../;
// or
// let r1 = new RegExp('stri..', '');

console.log(r1.test(str));  // true

r1 = /STRI../;
console.log(r1.test(str));  // false

r1 = /STRI../i;
// or
// r1 = new RegExp('STRI..', 'i');
console.log(r1.test(str));  // true
