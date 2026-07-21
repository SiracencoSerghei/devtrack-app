<script lang="ts">
	import { fleetAPI } from '$lib/api/fleet';
	import { authStore } from '$lib/auth/auth.svelte';
	import { i18n } from '$lib/i18n/i18n.svelte.js';
	import { APIClientError } from '$lib/api/client';
	import type { DriverProfile } from '$lib/types/api';

	let driverProfile = $state<DriverProfile | null>(null);
	let loadingDriver = $state(true);
	
	let licenseNumber = $state('');
	let phone = $state('');
	let formError = $state('');
	let isSubmitting = $state(false);

	async function loadDriverProfile() {
		try {
			driverProfile = await fleetAPI.getProfile();
		} catch (err) {
			if (err instanceof APIClientError && err.status === 404) {
				driverProfile = null;
			} else {
				console.error("Errore nel caricamento del profilo driver", err);
			}
		} finally {
			loadingDriver = false;
		}
	}

	async function handleCreateDriver(e: SubmitEvent) {
		e.preventDefault();
		if (isSubmitting) return;

		formError = '';
		isSubmitting = true;

		try {
			driverProfile = await fleetAPI.createProfile({
				license_number: licenseNumber,
				phone: phone
			});
		} catch (err) {
			if (err instanceof APIClientError) {
				formError = err.message;
			} else {
				formError = i18n.t('dashboard.form_conn_error') || 'Errore di connessione';
			}
		} finally {
			isSubmitting = false;
		}
	}

	$effect(() => {
		if (authStore.isInitialized) {
			if (authStore.isAuthenticated && authStore.isDriver) {
				loadDriverProfile();
			} else {
				loadingDriver = false;
			}
		}
	});
</script>

{#if loadingDriver}
	<div class="loader-container">
		<div class="loader">{i18n.t('dashboard.loading_profile')}</div>
	</div>
{:else if !authStore.isAuthenticated || !authStore.isDriver}
	<div class="alert alert-danger max-width-container">
		<h2>🚫 {i18n.t('home.access_denied_title')}</h2>
		<p>{i18n.t('home.access_denied_msg')}</p>
		<a href="/" class="btn-link">{i18n.t('home.back_to_dashboard')}</a>
	</div>
{:else}
	<div class="card driver-panel">
		<h2>{i18n.t('dashboard.driver_panel_title')}</h2>
		<p class="panel-desc">{i18n.t('dashboard.driver_panel_desc')}</p>
		
		{#if driverProfile}
			<div class="driver-badge-box">
				<div class="profile-info">
					<p><strong>{i18n.t('dashboard.license_num')}</strong> <code>{driverProfile.license_number}</code></p>
					<p><strong>{i18n.t('dashboard.phone_num')}</strong> {driverProfile.phone || '—'}</p>
				</div>
				<div class="status-box">
					<span class="badge badge-{driverProfile.status.toLowerCase()}">
						{driverProfile.status}
					</span>
					<p class="status-tip">{i18n.t('dashboard.status_tip')}</p>
				</div>
			</div>
		{:else}
			<div class="activation-box">
				<div class="info-alert">
					<strong>{i18n.t('dashboard.activation_required')}</strong> {i18n.t('dashboard.activation_desc')}
				</div>
				
				{#if formError}
					<div class="alert alert-danger">{formError}</div>
				{/if}

				<form onsubmit={handleCreateDriver} class="erp-form">
					<div class="form-row">
						<div class="form-group">
							<label for="license">{i18n.t('dashboard.license_label')}</label>
							<input type="text" id="license" bind:value={licenseNumber} disabled={isSubmitting} required placeholder="Es. U11223344K" />
						</div>
						<div class="form-group">
							<label for="phone">{i18n.t('dashboard.phone_label')}</label>
							<input type="text" id="phone" bind:value={phone} disabled={isSubmitting} placeholder="Es. +39 333 1234567" />
						</div>
					</div>
					<button type="submit" class="btn btn-primary" disabled={isSubmitting}>
						{isSubmitting ? 'Inviando...' : i18n.t('dashboard.btn_activate')}
					</button>
				</form>
			</div>
		{/if}
	</div>
{/if}

<style>
	.max-width-container { max-width: 600px; margin: 4rem auto; text-align: center; padding: 2rem; border-radius: 8px; }
	.btn-link { display: inline-block; margin-top: 1rem; color: #2563eb; text-decoration: none; font-weight: 600; }
	.loader-container { display: flex; justify-content: center; padding: 4rem; }
	.loader { color: #64748b; font-size: 1rem; font-weight: 500; }
	.driver-panel { max-width: 800px; margin: 2rem auto; background: white; border: 1px solid #e2e8f0; padding: 2rem; border-radius: 8px; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05); font-family: sans-serif; }
	.panel-desc { color: #64748b; font-size: 0.9rem; margin-bottom: 1.5rem; margin-top: 0.25rem; }
	.driver-badge-box { display: flex; justify-content: space-between; align-items: center; background: #f8fafc; border: 1px solid #e2e8f0; padding: 1.25rem; border-radius: 8px; }
	.profile-info p { margin: 0.5rem 0; color: #334155; }
	.profile-info code { background: #e2e8f0; padding: 0.2rem 0.4rem; border-radius: 4px; font-weight: 600; color: #0f172a; }
	.status-box { text-align: right; }
	.badge { display: inline-block; padding: 0.25rem 0.75rem; border-radius: 9999px; font-size: 0.75rem; font-weight: 700; text-transform: uppercase; }
	.badge-available { background: #d1fae5; color: #065f46; }
	.badge-in_transit { background: #fef3c7; color: #92400e; }
	.badge-off_duty { background: #f1f5f9; color: #475569; }
	.status-tip { font-size: 0.75rem; color: #64748b; margin-top: 0.5rem; margin-bottom: 0; }
	.info-alert { background: #eff6ff; color: #1e40af; border: 1px solid #bfdbfe; padding: 0.75rem 1rem; border-radius: 6px; margin-bottom: 1.2rem; font-size: 0.9rem; }
	.erp-form { display: flex; flex-direction: column; gap: 1rem; }
	.form-row { display: flex; gap: 1rem; flex-wrap: wrap; }
	.form-row .form-group { flex: 1; min-width: 200px; }
	.form-group { display: flex; flex-direction: column; gap: 0.4rem; }
	.form-group label { font-size: 0.85rem; font-weight: 600; color: #475569; }
	.form-group input { padding: 0.55rem; border: 1px solid #cbd5e1; border-radius: 6px; font-size: 0.95rem; }
	.form-group input:focus { border-color: #2563eb; outline: none; }
	.btn-primary { background: #2563eb; color: white; border: none; padding: 0.6rem 1.2rem; border-radius: 6px; font-weight: 600; cursor: pointer; transition: background 0.2s; align-self: flex-start; }
	.btn-primary:hover:not(:disabled) { background: #1d4ed8; }
	.btn-primary:disabled { background: #94a3b8; cursor: not-allowed; }
	.alert-danger { background: #fef2f2; color: #991b1b; padding: 1rem; border: 1px solid #fca5a5; border-radius: 6px; margin-bottom: 1rem; }
</style>