// import assert from 'assert'
import { convertNamesList } from './renameNamesConverter'

const sets = [
  {
    list: {
      a: 'b'
    },
    expected: [['a', 'b']]
  },
  {
    list: {
      'a/c': 'b/c'
    },
    expected: [['a', 'b']]
  },
  {
    list: {
      'a/b/c': 'a/e/c'
    },
    expected: [['a/b', 'e']]
  },
  {
    list: {
      'aaa/ccc': 'bbb/ddd'
    },
    expected: [['aaa', 'bbb'], ['bbb/ccc', 'ddd']]
  }
]

describe.each(sets)('convertNamesList', (set) => {
  it(`Should be ${JSON.stringify(set.expected)} for basic ${JSON.stringify(set.list)}`, () => {
      expect(convertNamesList(set.list)).toStrictEqual(set.expected)
  })
})
