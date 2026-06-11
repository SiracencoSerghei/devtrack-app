<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';
    import { untrack } from 'svelte';

    let userLoggedIn = $state(null);
    let driverProfile = $state(null);
    let loadingDriver = $state(false);
    
    let licenseNumber = $state('');
    let phone = $state('');
    let formError = $state('');
    let formSuccess = $state(false);

    async function loadDriverProfile(userID) {
        if (loadingDriver) return;
        loadingDriver = true;
        try {
            const token = localStorage.getItem('token');
            const res = await fetch(`http://localhost:8080/api/drivers?user_id=${userID}`, {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            if (res.ok) {
                driverProfile = await res.json();
            } else if (res.status === 404) {
                driverProfile = null;
            }
        } catch (err) {
            console.error("Errore profilo driver", err);
        } finally {
            loadingDriver = false;
        }
    }

    async function handleCreateDriver(e) {
        e.preventDefault();
        formError = '';
        try {
            const token = localStorage.getItem('token');
            const res = await fetch('http://localhost:8080/api/drivers', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({
                    user_id: userLoggedIn.id,
                    license_number: licenseNumber,
                    phone: phone
                })
            });

            if (!res.ok) {
                const data = await res.json();
                formError = data.error || "Errore";
                return;
            }
            driverProfile = await res.json();
        } catch (err) {
            formError = "offline";
        }
    }

    $effect(() => {
        const token = localStorage.getItem('token');
        const savedUser = localStorage.getItem('user_data');
        if (token && savedUser) {
            try {
                userLoggedIn = JSON.parse(savedUser);
                untrack(() => loadDriverProfile(userLoggedIn.id));
            } catch (e) {
                localStorage.removeItem('user_data');
            }
        }
    });
</script>

<div class="card driver-panel">
    <h2>{i18n.t('dashboard.driver_panel_title')}</h2>
    <p class="panel-desc">{i18n.t('dashboard.driver_panel_desc')}</p>
    
    {#if loadingDriver}
        <div class="loader">{i18n.t('dashboard.loading_profile')}</div>
    {:else if driverProfile}
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
                <div class="alert alert-danger">
                    {formError === 'offline' ? i18n.t('dashboard.form_conn_error') : formError}
                </div>
            {/if}

            <form onsubmit={handleCreateDriver} class="erp-form">
                <div class="form-row">
                    <div class="form-group">
                        <label for="license">{i18n.t('dashboard.license_label')}</label>
                        <input type="text" id="license" bind:value={licenseNumber} required placeholder="Es. U11223344K" />
                    </div>
                    <div class="form-group">
                        <label for="phone">{i18n.t('dashboard.phone_label')}</label>
                        <input type="text" id="phone" bind:value={phone} placeholder="Es. +39 333 1234567" />
                    </div>
                </div>
                <button type="submit" class="btn btn-primary">{i18n.t('dashboard.btn_activate')}</button>
            </form>
        </div>
    {/if}
</div>

<style>
    .driver-panel { max-width: 800px; margin: 2rem auto; }
    .panel-desc { color: #64748b; font-size: 0.9rem; margin-bottom: 1.5rem; }
    .driver-badge-box { display: flex; justify-content: space-between; align-items: center; background: #f8fafc; border: 1px solid var(--border); padding: 1.25rem; border-radius: 8px; }
    .status-box { text-align: right; }
    .status-tip { font-size: 0.75rem; color: #64748b; margin-top: 0.5rem; }
    .info-alert { background: #eff6ff; color: #1e40af; border: 1px solid #bfdbfe; padding: 0.75rem 1rem; border-radius: 6px; margin-bottom: 1.2rem; }
    .erp-form { display: flex; flex-direction: column; gap: 1rem; }
    .form-row { display: flex; gap: 1rem; flex-wrap: wrap; }
    .form-row .form-group { flex: 1; min-width: 200px; }
    .form-group { display: flex; flex-direction: column; gap: 0.4rem; }
    .form-group label { font-size: 0.85rem; font-weight: 600; color: #475569; }
    .form-group input { padding: 0.55rem; border: 1px solid var(--border); border-radius: 6px; }
    .alert-danger { background: #fee2e2; color: #991b1b; padding: 0.75rem; border-radius: 6px; }
</style>