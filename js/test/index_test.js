'use strict';

import assert from 'assert';
import {
  Board
}
from 'index';
import _ from 'lodash';

describe('solve', () => {
  let fixtures = {
    "79538246": "x won.",
    "35497162193": "x won.",
    "61978543": "x won.",
    "254961323121": "x won.",
    "6134278187": "x won.",
    "4319581": "Foul : x won.",
    "9625663381": "Foul : x won.",
    "7975662": "Foul : x won.",
    "2368799597": "Foul : x won.",
    "18652368566": "Foul : x won.",
    "965715": "o won.",
    "38745796": "o won.",
    "371929": "o won.",
    "758698769": "o won.",
    "42683953": "o won.",
    "618843927": "Foul : o won.",
    "36535224": "Foul : o won.",
    "882973": "Foul : o won.",
    "653675681": "Foul : o won.",
    "9729934662": "Foul : o won.",
    "972651483927": "Draw game.",
    "142583697": "Draw game.",
    "5439126787": "Draw game.",
    "42198637563": "Draw game.",
    "657391482": "Draw game.",
  };

  _.map(fixtures, (v, k) => {
    it('works', () => {
      assert.equal(Board.solve(k), v, k);
    });
  });


})
