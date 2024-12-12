const fs = require("fs");
const { argv } = require("process");
const path = argv[2] ?? "sample";
const str = fs.readFileSync(`${path}.txt`).toString().split("\r\n");
const matrix = [],
  zeros = [];
str.forEach((_, row) => {
  matrix.push(
    str[row].split("").map((el, idx) => {
      if (+el === 0) zeros.push([row, idx]);
      return +el;
    })
  );
});
const isInBound = (matrix, row, col) =>
  row >= 0 && row < matrix.length && col >= 0 && col < matrix[0].length;
function walk(matrix, row, col, lastVal, result) {
  if (!isInBound(matrix, row, col) || matrix[row][col] - lastVal != 1) return;
  if (matrix[row][col] == 9) result.push([row, col]);
  walk(matrix, row, col - 1, matrix[row][col], result);
  walk(matrix, row, col + 1, matrix[row][col], result);
  walk(matrix, row - 1, col, matrix[row][col], result);
  walk(matrix, row + 1, col, matrix[row][col], result);
}
let total = zeros.reduce((total, zero) => {
  const result = [];
  walk(matrix, zero[0], zero[1], -1, result);
  return total + result.length;
}, 0);
