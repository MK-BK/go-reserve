import { createStore } from 'vuex'

import http from '../utils/http'

export default createStore({
	state: {
		token: '',
		user: {},
	},

	getters: {
	},

	mutations: {
	},

	actions: {
		// +++++++++++++++++ API USER
		login: async function(context, payload) {
			let that = this
			http.post('/api/login', payload).then(function(resp) {
				that.state.user = resp
			}).catch(function(err) {
				return err
			})
		},

		register: async function(context, payload) {
			return await http.post('/api/register', payload)
		},

		listUsers: async function() {
			return await http.get('/api/accounts')
		},


		// +++++++++++++++++ API SHOP
		listShop: async function() {
			return await http.get('/api/shops')
		},

		createShop: async function(context, payload) {
			return await http.post('/api/shop', payload)
		},

		// +++++++++++++++++ API commodity
		listCommodity: async function() {
			return await http.get('/api/commodity')
		},

		createCommodity: async function(context, payload) {
			return await http.post('/api/commodity', payload)
		},

		updateCommdity: async function(context, payload) {
			return await http.put(`/api/commodity/${payload.ID}`, payload)
		},


		// +++++++++++++++++ API Job
		listJob: async function() {
			return await http.get('/api/jobs')
		},

		createJob: async function(context, payload) {
			return await http.post('/api/job', payload)
		},
		

		// +++++++++++++++++ API Request
		listRequests: async function() {
			return await http.get('/api/requests')
		},

		updateRquest: async function(context, payload) {
			return await http.put(`/api/requests/${payload.ID}`, payload)
		},

		// +++++++++++++++++++++ API avatar
		uploadAvatar: async function(context, payload) {
			return await http.post(`/api/avatar/${payload.resource}/${payload.id}`, payload)
		},

		getAvatat: async function(context, payload) {
			return await http.get(`/api/avatar/${payload.resource}/${payload.id}`)
		},

		// +++++++++++++++++ API audit_log
		listAuditLogs: async function() {
			return await http.get('/api/audit_log')
		}
	},

	modules: {
	}
})
