process.stdin.resume();
process.stdin.setEncoding('utf8');

function partInt10(elem) {
    return parseInt(elem, 10);
}

process.stdin.on('data', function (chunk) {
    var intArrary = chunk.trim().split(' ').map(partInt10);
    var N = intArrary[0];

    for (var i = 1; i <= N; i++) {
        var result = '';
        if (i % 3 === 0)
            result += 'Fizz ';
        if (i % 5 === 0)
            result += 'Buzz ';
        if (result === '')
            result += i.toString();

        console.log(result.trim());
    }
});
