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
		login: async function({payload}) {
			return await http.post('/api/login', payload)
		},

		register: async function({payload}) {
			return await http.post('/api/register', payload)
		},

		listUsers: async function() {
			return await http.get('/api/accounts')
		},

		listShop: async function() {
			return await http.get('/api/shops')
		},

		listCommodity: async function() {
			return await http.get('/api/commodity')
		},

		listJob: async function() {
			return await http.get('/api/jobs')
		},

		listAuditLogs: async function() {
			return await http.get('/api/audit_log')
		},

		listRequests: async function() {
			return await http.get('/api/requests')
		}
	},

	modules: {
	}
})
