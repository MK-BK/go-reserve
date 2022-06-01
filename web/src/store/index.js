import { createStore } from 'vuex'

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
    listusers: async function(context) {
      const { axios } = getCurrentInstance()
      return await axios.get('/api/users')
    }
  },

  modules: {
  }
})
