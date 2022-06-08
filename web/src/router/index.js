import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
	{
		path: '/',
		redirect: to => {
		return { path: '/login', query: { q: to.params.searchText } }
		},
	},
	{
		path: '/login',
		component: () => import(/* webpackChunkName: "about" */ '../views/login.vue')
	},
	{
		path: '/index',
		component: () => import(/* webpackChunkName: "about" */ '../views/index.vue'),
		children: [
		{
			path: '/users',
			component: () => import(/* webpackChunkName: "about" */ '../views/user.vue')
		},
		{
			path: '/shops',
			component: () => import(/* webpackChunkName: "about" */ '../views/shop.vue')
		},
		{
			path: '/requests',
			component: () => import(/* webpackChunkName: "about" */ '../views/request.vue')
		},
		{
			path: '/audit_logs',
			component: () => import(/* webpackChunkName: "about" */ '../views/audit_log.vue')
		},
		{
			path: '/setting',
			component: () => import(/* webpackChunkName: "about" */ '../views/setting.vue')
		},
		{
			path: '/job',
			component: () => import(/* webpackChunkName: "about" */ '../views/job.vue')
		},
		{
			path: '/commodity',
			component: () => import(/* webpackChunkName: "about" */ '../views/commodity.vue')
		}
		]
	}
]

const router = createRouter({
	history: createWebHashHistory(),
	routes
})

export default router
