<script lang="ts">
	import { systemAPI } from '$lib/api/system';
	import { authStore } from '$lib/auth/auth.svelte';
	import { i18n } from '$lib/i18n/i18n.svelte.js';

	let statusKey = $state<'checking' | 'ok' | 'offline'>('checking');

	// Обчислювальні руни Svelte 5 для безпечного керування доступом до інтерфейсу
	let hasDriverRole = $derived(authStore.user?.roles?.includes('driver') || false);
	let hasStaffRole = $derived(
		authStore.user?.roles?.some(role => role === 'admin' || role === 'dispatcher') || false
	);
	
	async function checkBackend() {
		try {
			const res = await systemAPI.checkHealth();
			if (res && res.status === 'ok') {
				statusKey = 'ok';
			} else {
				statusKey = 'offline';
			}
		} catch (e) {
			// М'яко гасимо помилку, щоб не ламати життєвий цикл компонента
			statusKey = 'offline';
		}
	}

	// Запускаємо перевірку ізольовано від інших реактивних контекстів
	$effect(() => {
		checkBackend();
	});
</script>

<div class="dashboard-header card">
	<div>
		<h1>{i18n.t('home.title')}</h1>
		<p class="system-status">
			{i18n.t('dashboard.system_infra')}: 
			{#if statusKey === 'checking'}
				<span class="status-indicator status-checking">CONNECTING...</span>
			{:else}
				<span class="status-indicator status-{statusKey}">
					{statusKey === 'ok' ? 'ONLINE' : 'OFFLINE'}
				</span>
			{/if}
		</p>
	</div>
</div>

{#if authStore.isAuthenticated}
	<div class="welcome-box card">
		<p>👋 {i18n.t('home.welcome')}, <strong>{authStore.user?.name}</strong>! {i18n.t('home.active_roles')} 
			{#each authStore.user?.roles || [] as role}
				<span class="role-tag">{role}</span>
			{/each}
		</p>
		<p class="sub-text">{i18n.t('home.select_department')}</p>
	</div>

	<div class="hub-grid">
		<!-- Модуль 1: Кабінет Водія -->
		{#if hasDriverRole}
			<a href="/driver" class="hub-card">
				<div class="icon">🚚</div>
				<h3>{i18n.t('dashboard.driver_panel_title')}</h3>
				<p>{i18n.t('dashboard.driver_panel_desc_placeholder' || 'Accedi ai tuoi dati logistici')}</p>
			</a>
		{/if}

		<!-- Модуль 2: Адмін-панель Флоту -->
		{#if hasStaffRole}
			<a href="/admin" class="hub-card">
				<div class="icon">🏢</div>
				<h3>{i18n.t('dashboard.fleet_panel_title')}</h3>
				<p>{i18n.t('dashboard.fleet_panel_desc_placeholder' || 'Monitora flotta e operatori')}</p>
			</a>
		{/if}
	</div>
{:else}
	<div class="card guest-box">
		<h2>{i18n.t('home.area_reserved')}</h2>
		<p>{i18n.t('home.guest_msg')}</p>
		<div class="auth-buttons">
			<a href="/login" class="btn btn-light">{i18n.t('dashboard.btn_guest_login') || 'Login'}</a>
			<a href="/signup" class="btn btn-primary">{i18n.t('dashboard.btn_guest_signup') || 'Register'}</a>
		</div>
	</div>
{/if}

<style>
	.card { background: white; border: 1px solid #e2e8f0; border-radius: 8px; padding: 1.5rem; }
	.dashboard-header { margin-bottom: 1.5rem; box-shadow: 0 1px 3px rgba(0,0,0,0.05); }
	.dashboard-header h1 { margin: 0; font-size: 1.75rem; color: #0f172a; }
	.system-status { margin: 0.5rem 0 0 0; font-size: 0.875rem; color: #64748b; font-weight: 500; }
	.status-indicator { font-weight: 700; padding: 0.2rem 0.5rem; border-radius: 4px; font-size: 0.75rem; letter-spacing: 0.05em; }
	.status-ok { background: #d1fae5; color: #065f46; }
	.status-offline { background: #fee2e2; color: #991b1b; }
	.status-checking { background: #e2e8f0; color: #475569; }

	.welcome-box { margin-bottom: 2rem; padding: 1.25rem; background: #f8fafc; border-color: #e2e8f0; }
	.welcome-box p { margin: 0.25rem 0; color: #334155; font-size: 0.95rem; }
	.sub-text { font-size: 0.875rem; color: #64748b; }
	
	.role-tag {
		display: inline-block; background: var(--accent, #2563eb); color: white; padding: 0.15rem 0.4rem;
		font-size: 0.7rem; font-weight: 700; border-radius: 4px; margin-left: 0.35rem; text-transform: uppercase; letter-spacing: 0.025em;
	}

	.hub-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr)); gap: 1.5rem; }
	.hub-card { 
		background: white; border: 1px solid #e2e8f0; padding: 2rem; border-radius: 12px; 
		text-decoration: none; color: inherit; transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1); display: block;
		box-shadow: 0 1px 2px rgba(0,0,0,0.02);
	}
	.hub-card:hover { transform: translateY(-4px); border-color: var(--accent, #2563eb); box-shadow: 0 12px 20px -3px rgba(0, 0, 0, 0.08); }
	.hub-card .icon { font-size: 2.5rem; margin-bottom: 1rem; }
	.hub-card h3 { margin: 0 0 0.5rem 0; font-size: 1.2rem; color: #0f172a; font-weight: 700; }
	.hub-card p { margin: 0; font-size: 0.875rem; color: #64748b; line-height: 1.5; }

	.guest-box { text-align: center; padding: 4rem 2rem; }
	.guest-box h2 { color: #0f172a; margin-top: 0; }
	.guest-box p { color: #64748b; margin-bottom: 2rem; }
	.auth-buttons { margin-top: 2rem; display: flex; justify-content: center; gap: 1rem; }
	.btn { display: inline-block; padding: 0.6rem 1.2rem; border-radius: 6px; font-weight: 600; text-decoration: none; font-size: 0.9rem; transition: all 0.2s; }
	.btn-light { background: #f1f5f9; color: #334155; border: 1px solid #cbd5e1; }
	.btn-light:hover { background: #e2e8f0; }
	.btn-primary { background: #2563eb; color: white; border: 1px solid transparent; }
	.btn-primary:hover { background: #1d4ed8; }
</style>