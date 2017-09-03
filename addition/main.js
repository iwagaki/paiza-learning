process.stdin.resume();
process.stdin.setEncoding('utf8');

function partInt10(elem) {
    return parseInt(elem, 10);
}

process.stdin.on('data', function (chunk) {
    var intArrary = chunk.trim().split(' ').map(partInt10);
    var a = intArrary[0];
    var b = intArrary[1];
    console.log(a + b);
});