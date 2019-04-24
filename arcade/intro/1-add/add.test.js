const add = require("./add");
const assert = require("assert");
const describe = require("mocha").describe;
const it = require("mocha").it;

describe("add", function() {
  it("should add the two input numbers together", function() {
    let tests = [
      {
        n1: 1,
        n2: 2,
        expect: 3
      },
      {
        n1: 0,
        n2: 1000,
        expect: 1000
      },
      {
        n1: 2,
        n2: -39,
        expect: -37
      },
      {
        n1: 99,
        n2: 100,
        expect: 199
      },
      {
        n1: -100,
        n2: 100,
        expect: 0
      },
      {
        n1: -1000,
        n2: -1000,
        expect: -2000r
      }
    ];
    for (let test of tests) {
      let { n1, n2, expect } = test;
      let out = add(n1, n2);
      let msg = `[FAIL] add{${n1}, ${n2}} == ${out} expect ${expect}`;
      assert.strictEqual(out, expect, msg);
    }
  });
});
