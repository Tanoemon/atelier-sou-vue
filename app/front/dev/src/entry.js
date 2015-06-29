var Vue = require('Vue');

document.addEventListener("DOMContentLoaded", function() {

  var mainVue = require('./app/app.vue');
  var app = new Vue(mainVue).$mount('#app');

});
