import { createStore } from 'vuex'
import http from '../utils/http'

export default createStore({
  state: {
    token: "",
    user: "",
  },

  getters: {
  },

  mutations: {
  },

  actions: {
    login: async function(context, payload) {
      return await http.post('/api/login', payload)
    },

    listusers: async function() {
      return await http.get('/api/accounts')
    }
  },

  modules: {
  }
})
