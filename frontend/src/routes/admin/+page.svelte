<script lang="ts">
	import { authAPI } from '$lib/api/auth';
	import { authStore } from '$lib/auth/auth.svelte';
	import { i18n } from '$lib/i18n/i18n.svelte.js';
	import { APIClientError } from '$lib/api/client';
	import type { User } from '$lib/types/api';
	import { goto } from '$app/navigation';

	let users = $state<User[]>([]); 
	let loading = $state(true);
	let errorMsg = $state('');

	async function loadUsers() {
		try {
			users = await authAPI.getAll();
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

	$effect(() => {
		if (authStore.isInitialized) {
			if (authStore.isAuthenticated && authStore.isStaff) {
				loadUsers();
			} else {
				loading = false;
			}
		}
	});
</script>

{#if loading}
	<div class="loading-screen">
		<div class="loader">Loading...</div>
	</div>
{:else if !authStore.isAuthenticated || !authStore.isStaff}
	<div class="card access-denied">
		<div class="icon">🚫</div>
		<h2>{i18n.t('home.access_denied_title')}</h2>
		<p>{i18n.t('home.access_denied_msg')}</p>
		<button class="btn btn-primary" onclick={() => goto('/')} type="button">
			{i18n.t('home.back_to_dashboard')}
		</button>
	</div>
{:else}
	<div class="card fleet-panel">
		<h2>{i18n.t('dashboard.fleet_panel_title')}</h2>
		<p class="panel-desc">{i18n.t('home.user_list')}</p>
		
		{#if errorMsg}
			<div class="alert alert-danger">{errorMsg}</div>
		{:else if users.length === 0}
			<p class="empty-msg">{i18n.t('home.loading_or_empty')}</p>
		{:else}
			<div class="table-container">
				<table class="erp-table">
					<thead>
						<tr>
							<th>{i18n.t('dashboard.operator') || 'Користувач'}</th>
							<th>{i18n.t('dashboard.email_contact') || 'Email'}</th>
							<th>{i18n.t('home.active_roles_table') || 'Ролі'}</th>
						</tr>
					</thead>
					<tbody>
						{#each users as user}
							<tr>
								<td><strong>{user.name}</strong></td>
								<td><code>{user.email}</code></td>
								<td>
									{#each user.roles || [] as role}
										<span class="badge-role-inline">{role}</span>
									{/each}
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>
		{/if}
	</div>
{/if}

<style>
	.fleet-panel { max-width: 1000px; margin: 2rem auto; }
	.panel-desc { color: #64748b; font-size: 0.9rem; margin-bottom: 1.5rem; margin-top: 0.25rem; }
	.access-denied { max-width: 500px; margin: 4rem auto; text-align: center; padding: 3rem 2rem; border: 1px solid #fca5a5; background: #fff5f5; border-radius: 8px; }
	.access-denied .icon { font-size: 3rem; margin-bottom: 1rem; }
	.access-denied h2 { color: #991b1b; font-size: 1.5rem; margin-bottom: 0.5rem; border: none; padding: 0; }
	.access-denied p { color: #7f1d1d; margin-bottom: 2rem; font-size: 0.95rem; }
	.table-container { border: 1px solid #e2e8f0; border-radius: 8px; overflow: hidden; background: white; }
	.erp-table { width: 100%; border-collapse: collapse; text-align: left; font-size: 0.9rem; }
	.erp-table th { background: #f1f5f9; padding: 0.75rem; color: #475569; font-weight: 600; border-bottom: 1px solid #e2e8f0; }
	.erp-table td { padding: 0.75rem; border-top: 1px solid #e2e8f0; color: #1e293b; }
	.loading-screen { text-align: center; margin-top: 5rem; color: #64748b; font-weight: 500; }
	.empty-msg { color: #64748b; text-align: center; padding: 2rem; }
	.badge-role-inline { background: #e2e8f0; color: #334155; font-size: 0.7rem; font-weight: 700; padding: 0.15rem 0.4rem; border-radius: 4px; text-transform: uppercase; letter-spacing: 0.05em; margin-right: 0.25rem; }
</style>