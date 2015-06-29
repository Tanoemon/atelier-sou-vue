import S from './S';
var _ = require('lodash');
var V = require('velocity-animate');

export default class Base {
  constructor() {
    this.els = [];
    this.du = 1000;
    let elns = Array.prototype.slice.call(arguments);
    elns.forEach((eln) => {
      let el = S.gi(eln);
      this.els.push({
        name: eln,
        el: el,
        du: 1000
      });
    });
  }

  show(eln) {
    this._showIt(eln);
    this._hideOthers(eln);
  }

  _showIt(eln) {
    var elo = _.filter(this.els, {
      name: eln
    })[0];
    let o = S.css(elo.el, 'opacity');
    V(elo.el, 'stop');
    V(elo.el, {
      opacity: 1
    }, {
      duration: this.du - (o * this.du)
    });
  }

  _hideOthers(eln) {
    this.els.forEach((elo) => {
      if (elo.name === eln) {
        return;
      }
      let o = S.css(elo.el, 'opacity');
      V(elo.el, 'stop');
      V(elo.el, {
        opacity: 0
      }, {
        duration: o * this.du
      });
    });
  }
}
