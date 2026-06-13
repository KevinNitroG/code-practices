const indexedReversedChar = [
  "z",
  "y",
  "x",
  "w",
  "v",
  "u",
  "t",
  "s",
  "r",
  "q",
  "p",
  "o",
  "n",
  "m",
  "l",
  "k",
  "j",
  "i",
  "h",
  "g",
  "f",
  "e",
  "d",
  "c",
  "b",
  "a",
];

function mapWordWeights(words: string[], weights: number[]): string {
  const result = new Array(words.length);
  for (let i = 0; i < words.length; ++i) {
    let accumulate = 0;
    for (const char of words[i]) {
      accumulate += weights[char.charCodeAt(0) - 97];
    }
    result[i] = indexedReversedChar[accumulate % 26];
  }
  return result.join("");
}

console.log(
  mapWordWeights(
    ["abcd", "def", "xyz"],
    [
      5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3,
      7, 7, 2,
    ],
  ), // "rij"
);

console.log(
  mapWordWeights(
    ["a", "b", "c"],
    [
      1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
      1,
    ],
  ),
); // "yyy"

console.log(
  mapWordWeights(
    ["abcd"],
    [
      7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4,
      4, 5,
    ],
  ),
);
