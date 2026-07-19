export const API_BASE_URL = (import.meta.env.VITE_API_URL as string) ?? 'http://localhost:8080';

export const ENDPOINTS = {
	AUTH: {
		SIGNUP: '/api/auth/signup',
		LOGIN: '/api/auth/login',
		GET_ALL: '/api/auth/users'
	},
	FLEET: {
		DRIVERS: '/api/fleet/drivers'
	},
	CORE: {
		EMPLOYEE: '/api/core/employee'
	},
	LOGISTICS: {
		ORDERS: '/api/orders'
	},
	SYSTEM: {
		HEALTH: '/health'
	}
} as const;
