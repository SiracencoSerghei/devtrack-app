import { API_BASE_URL } from '../constants/api';
import type { APIErrorResponse } from '../types/api';

export class APIClientError extends Error {
	constructor(
		public status: number,
		message: string
	) {
		super(message);
		this.name = 'APIClientError';
	}
}

async function request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
	const url = `${API_BASE_URL}${endpoint}`;
	const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;

	const headers = new Headers(options.headers);
	if (!headers.has('Content-Type') && !(options.body instanceof FormData)) {
		headers.set('Content-Type', 'application/json');
	}
	if (token) {
		headers.set('Authorization', `Bearer ${token}`);
	}

	const config: RequestInit = {
		...options,
		headers
	};

	try {
		const response = await fetch(url, config);

		if (response.status === 204) {
			return {} as T;
		}

		if (!response.ok) {
			let errorMessage = `Error: ${response.status}`;
			try {
				const errorData = await response.json();
				// Бекенд на Go повертає помилки у форматі {"error": "text"}
				errorMessage = errorData.error || errorMessage;
			} catch {
				// Якщо бекенд віддав не JSON
			}
			throw new APIClientError(response.status, errorMessage);
		}

		return (await response.json()) as T;
	} catch (error) {
		if (error instanceof APIClientError) throw error;
		// Якщо сервер взагалі вимкнений, повертаємо зрозумілу помилку
		throw new APIClientError(0, 'Server is offline. Please start your backend.');
	}
}

export const client = {
	get: <T>(endpoint: string, options?: RequestInit) =>
		request<T>(endpoint, { ...options, method: 'GET' }),
	post: <T>(endpoint: string, body: any, options?: RequestInit) =>
		request<T>(endpoint, { ...options, method: 'POST', body: JSON.stringify(body) }),
	put: <T>(endpoint: string, body: any, options?: RequestInit) =>
		request<T>(endpoint, { ...options, method: 'PUT', body: JSON.stringify(body) }),
	patch: <T>(endpoint: string, body: any, options?: RequestInit) =>
		request<T>(endpoint, { ...options, method: 'PATCH', body: JSON.stringify(body) }),
	delete: <T>(endpoint: string, options?: RequestInit) =>
		request<T>(endpoint, { ...options, method: 'DELETE' })
};
