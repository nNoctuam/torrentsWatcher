// import assert from 'assert'
import convertNamesList from "./renameNamesConverter";

class Set {
  list: [string, string][] = [];
  expected: [string, string][] = [];
}

const sets: Set[] = [
  {
    list: [["a", "b"]],
    expected: [["a", "b"]],
  },
  {
    list: [["a/c", "b/c"]],
    expected: [["a", "b"]],
  },
  {
    list: [["a/b/c", "a/e/c"]],
    expected: [["a/b", "e"]],
  },
  {
    list: [
      ["aaa/ccc", "bbb/ccc"],
      ["aaa/eee", "bbb/eee"],
    ],
    expected: [["aaa", "bbb"]],
  },
  {
    list: [
      ["aaa/ccc", "bbb/ddd"],
      ["aaa/eee", "bbb/eee"],
    ],
    expected: [
      ["aaa", "bbb"],
      ["bbb/ccc", "ddd"],
    ],
  },
  {
    list: [
      ["aaa/ccc", "bbb/ddd"],
      ["aaa/eee", "bbb/fff"],
    ],
    expected: [
      ["aaa", "bbb"],
      ["bbb/ccc", "ddd"],
      ["bbb/eee", "fff"],
    ],
  },
];

describe.each(sets)("convertNamesList", (set) => {
  it(`Should be ${JSON.stringify(set.expected)} for basic ${JSON.stringify(
    set.list
  )}`, () => {
    expect(convertNamesList(set.list)).toStrictEqual(set.expected);
  });
});
