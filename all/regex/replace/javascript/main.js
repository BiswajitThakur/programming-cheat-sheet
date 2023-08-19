
let str = "I hate you. I hate javascript.";

let reg = /hate/i;

let st1 = str.replace(reg, 'love')
console.log(st1);
// I love you. I hate javascript.

let st2 = str.replace(reg, (s)=>{
    return `!${s}`;
});
console.log(st2);
// I !hate you. I hate javascript.

reg = /hate/gi;

let st3 = str.replace(reg, 'love')
console.log(st3);
// I love you. I love javascript.

let st4 = str.replace(reg, (s)=>{
    return `!${s}`;
});
console.log(st4);
// I !hate you. I !hate javascript.
