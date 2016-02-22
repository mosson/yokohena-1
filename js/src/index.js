'use strict';

import _ from 'lodash';

const players = {
  white: 0x1,
  black: 0x10,
  blank: 0
};

const marks = {
  [players.white]: 'o',
  [players.black]: 'x',
  [players.blank]: '-'
}

class Board {
  static read(str) {
    return _.map(str.split(''), (char) => {
      return parseInt(char, 10);
    });
  }

  static solve(str) {
    return new Board().solve(Board.read(str));
  }

  constructor() {
    this.mem = {};
    _.each(_.range(1, 10), (i) => {
      this.mem[i] = players.blank;
    });;
    this._result = 'Draw game.'
  }

  check() {
    return _.find([
      this.mem[1] & this.mem[2] & this.mem[3],
      this.mem[4] & this.mem[5] & this.mem[6],
      this.mem[7] & this.mem[8] & this.mem[9],
      this.mem[1] & this.mem[4] & this.mem[7],
      this.mem[2] & this.mem[5] & this.mem[8],
      this.mem[3] & this.mem[6] & this.mem[9],
      this.mem[1] & this.mem[5] & this.mem[9],
      this.mem[3] & this.mem[5] & this.mem[7],
    ], (m) => {
      return m > 0;
    });
  }

  description() {
    let table = [
      '\n',
      marks[this.mem[1]],
      marks[this.mem[2]],
      marks[this.mem[3]],
      '\n',
      marks[this.mem[4]],
      marks[this.mem[5]],
      marks[this.mem[6]],
      '\n',
      marks[this.mem[7]],
      marks[this.mem[8]],
      marks[this.mem[9]],
      '\n'
    ].join('\t');

    console.log(table);
  }

  place(n, i) {
    if (this.mem[n] !== players.blank) {
      this._result = `Foul : ${(i % 2 == 0 ? 'x' : 'o')} won.`;
      return true;
    }

    this.mem[n] = (i % 2 == 0 ? players.white : players.black);

    let c = this.check();
    if (c) {
      this._result = c;
      return true;
    }
  }

  solve(order) {
    for (let i = 0; i < order.length; i++) {
      let n = order[i];

      if (this.place(n, i)) break;

      if (_.all(this.mem, (v) => {
          return v > 0;
        })) {
        this._result = 'Draw game.';
        break;
      }
    };

    this.description();
    return this.result;
  }

  get result() {
    if (this._result === players.white) {
      return 'o won.';
    } else if (this._result === players.black) {
      return 'x won.';
    } else {
      return this._result;
    }
  }
}

export {
  Board
};
