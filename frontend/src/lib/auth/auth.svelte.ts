import type { User } from '../types/api';

class AuthStore {
	#user = $state<User | null>(null);
	#token = $state<string | null>(null);
	#initialized = $state(false);

	constructor() {
		if (typeof window !== 'undefined') {
			const savedUser = localStorage.getItem('user_data');
			const savedToken = localStorage.getItem('token');
			if (savedUser && savedToken) {
				try {
					this.#user = JSON.parse(savedUser);
					this.#token = savedToken;
				} catch {
					this.clear();
				}
			}
			this.#initialized = true;
		}
	}

	get user() {
		return this.#user;
	}
	get token() {
		return this.#token;
	}
	get isInitialized() {
		return this.#initialized;
	}
	get isAuthenticated() {
		return !!this.#user;
	}

	get isDriver() {
		return this.#user?.roles?.includes('driver') ?? false;
	}
	get isStaff() {
		return this.#user?.roles?.some((role) => role === 'admin' || role === 'dispatcher') ?? false;
	}

	setAuth(user: User, token: string) {
		this.#user = user;
		this.#token = token;
		if (typeof window !== 'undefined') {
			localStorage.setItem('user_data', JSON.stringify(user));
			localStorage.setItem('token', token);
		}
	}

	clear() {
		this.#user = null;
		this.#token = null;
		if (typeof window !== 'undefined') {
			localStorage.removeItem('user_data');
			localStorage.removeItem('token');
		}
	}
}

export const authStore = new AuthStore();
