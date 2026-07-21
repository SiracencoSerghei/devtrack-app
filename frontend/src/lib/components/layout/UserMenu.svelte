<script lang="ts">
	import { authStore } from '$lib/auth/auth.svelte';
	import { i18n } from '$lib/i18n/i18n.svelte.js';

	function handleLogout() {
		authStore.clear();
		window.location.href = '/login';
	}
</script>

<div class="auth-links">
	{#if authStore.isAuthenticated}
		<div class="user-badge">
			<span>{i18n.t('auth.welcome') || 'Welcome'}, <strong>{authStore.user?.name}</strong></span>
		</div>
		<button onclick={handleLogout} class="btn-logout" type="button">
			{i18n.t('auth.logout') || 'Logout'}
		</button>
	{:else}
		<a href="/login" class="btn-login-nav">
			{i18n.t('auth.login') || 'Login'}
		</a>
		<a href="/signup" class="btn-signup-nav">
			{i18n.t('auth.signup') || 'Sign Up'}
		</a>
	{/if}
</div>

<style>
	.auth-links {
		display: flex;
		align-items: center;
		gap: 1rem;
	}
	.user-badge {
		background: #1e293b;
		padding: 0.4rem 0.8rem;
		border-radius: 6px;
		font-size: 0.875rem;
		color: #f1f5f9;
	}
	.btn-logout {
		background: #ef4444;
		color: white;
		border: none;
		padding: 0.4rem 0.8rem;
		border-radius: 6px;
		cursor: pointer;
		font-weight: 600;
		font-size: 0.875rem;
		transition: background 0.2s;
	}
	.btn-logout:hover {
		background: #dc2626;
	}
	.btn-login-nav {
		color: #94a3b8;
		text-decoration: none;
		font-weight: 600;
		font-size: 0.875rem;
		transition: color 0.2s;
	}
	.btn-login-nav:hover {
		color: white;
	}
	.btn-signup-nav {
		background: var(--accent, #2563eb);
		color: white;
		text-decoration: none;
		padding: 0.4rem 0.8rem;
		border-radius: 6px;
		font-weight: 600;
		font-size: 0.875rem;
		transition: background 0.2s;
	}
	.btn-signup-nav:hover {
		background: #1d4ed8;
	}
</style>