export default class S {
  static on(el, ev, fn) {
    el.addEventListener(ev, fn);
  }

  static log(s) {
    console.log(s);
  }

  static gi(s) {
    return document.getElementById(s);
  }

  static qs(s) {
    return document.querySelector(s);
  }

  static qa(s) {
    return document.querySelectorAll(s);
  }

  static css(el, s) {
    return window.getComputedStyle(el).getPropertyValue(s);
  }
}
