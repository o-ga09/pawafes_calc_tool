import Vue from 'vue'
import VueMq from 'vue-mq'

Vue.use(VueMq, {
  breakpoints: {
    sp: 959,
    pc: Infinity
  },
  defaultBreakpoint: 'sp'
})