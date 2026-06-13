function mapWordWeights(words: string[], weights: number[]): string {
  const len = words.length;
  const ans = new Uint8Array(len);
  const codeA = 97;
  const codeZ = 122;

  for (let i = 0; i < len; i++) {
    const word = words[i];
    let s = 0;

    for (let j = 0; j < word.length; j++) {
      s += weights[word.charCodeAt(j) - codeA];
    }

    ans[i] = codeZ - (s % 26);
  }

  return String.fromCharCode(...ans);
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
