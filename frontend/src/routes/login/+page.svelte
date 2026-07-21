<script lang="ts">
	import { authAPI } from '$lib/api/auth';
	import { authStore } from '$lib/auth/auth.svelte';
	import { i18n } from '$lib/i18n/i18n.svelte.js';
	import { APIClientError } from '$lib/api/client';
	import { goto } from '$app/navigation';

	let email = $state('');
	let password = $state('');
	let errorMsg = $state('');
	let loading = $state(false);
	let showPassword = $state(false);

	async function handleLogin(e: SubmitEvent) {
		e.preventDefault();
		if (loading) return;

		errorMsg = '';
		loading = true;

		try {
			const res = await authAPI.login({ email, password });
			authStore.setAuth(res.user, res.token);
			await goto('/');
		} catch (err) {
			if (err instanceof APIClientError) {
				errorMsg = err.message;
			} else {
				errorMsg = i18n.t('server.offline_msg') || 'Errore di connessione';
			}
		} finally {
			loading = false;
		}
	}
</script>

<div class="auth-container">
	<h2>{i18n.t('login_page.title')}</h2>
	
	{#if errorMsg}
		<div class="alert alert-danger">{errorMsg}</div>
	{/if}

	<form onsubmit={handleLogin}>
		<div class="form-group">
			<label for="email">{i18n.t('login_page.email_label')}</label>
			<input type="email" id="email" bind:value={email} disabled={loading} required placeholder="mario@rossi.it" />
		</div>

		<div class="form-group">
			<label for="password">{i18n.t('login_page.password_label')}</label>
			<div class="password-wrapper">
				<input 
					type={showPassword ? 'text' : 'password'} 
					id="password" 
					bind:value={password} 
					disabled={loading}
					required 
					placeholder="••••••••" 
				/>
				<button 
					type="button" 
					class="toggle-password" 
					onclick={() => (showPassword = !showPassword)}
					title={showPassword ? i18n.t('auth.hide_password') : i18n.t('auth.show_password')}
				>
					{#if showPassword}
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
							<path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88" />
						</svg>
					{:else}
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
							<path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z" />
							<path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
						</svg>
					{/if}
				</button>
			</div>
		</div>

		<button type="submit" class="btn-submit" disabled={loading}>
			{loading ? 'Loading...' : i18n.t('login_page.btn_submit')}
		</button>
	</form>
	
	<p class="switch-auth">
		{i18n.t('login_page.no_account')} 
		<a href="/signup">{i18n.t('login_page.register_link')}</a>
	</p>
</div>

<style>
	.auth-container { max-width: 400px; margin: 3rem auto; padding: 2rem; border: 1px solid #e2e8f0; border-radius: 8px; font-family: sans-serif; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05); background: white; }
	h2 { margin-top: 0; color: #1e293b; text-align: center; }
	.switch-auth { text-align: center; margin-top: 1.5rem; color: #64748b; font-size: 0.875rem; }
	.switch-auth a { color: #2563eb; text-decoration: none; font-weight: 500; }
	.password-wrapper { position: relative; display: flex; align-items: center; }
	.password-wrapper input { width: 100%; padding-right: 2.5rem; }
	.toggle-password { position: absolute; right: 0.6rem; background: none; border: none; cursor: pointer; color: #64748b; padding: 0.2rem; display: flex; align-items: center; }
	.toggle-password:hover { color: #2563eb; }
	.icon { width: 1.25rem; height: 1.25rem; }
</style>