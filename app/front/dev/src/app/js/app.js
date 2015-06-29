import B from './js/classes/Base'
let b

module.exports = {
  components: {
    "v-menu": require('./pages/menu/menu.vue'),
    "v-staff": require('./pages/staff/staff.vue')
  },
  methods: {
    muClick: () => {
      b.show('menu')
    },
    sfClick: () => {
      b.show('staff')
    },
    ctClick: () => {
      b.show('concept')
    },
    asClick: (e) => {
      console.log(e);
      b.show('access')
    }
  },
  ready: () => {
    b = new B('menu', 'staff', 'concept', 'access')
  }
};
