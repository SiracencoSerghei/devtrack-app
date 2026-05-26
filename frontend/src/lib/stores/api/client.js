// src/lib/stores/api/client.js
import { apiStore } from '$lib/stores/api/index.js';

export async function apiFetch(url, options = {}) {
	try {
		const res = await fetch(url, options);

		if (!res.ok) throw new Error(`HTTP ${res.status}`);

		const data = await res.json();

		apiStore.setStatus('online');
		apiStore.setMessage('');

		return data;
	} catch (e) {
		apiStore.setStatus('offline');
		apiStore.setMessage(e.message);

		throw e;
	}
}
