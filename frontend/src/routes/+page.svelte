<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';
    import { untrack } from 'svelte'; // 🌟 Importiamo untrack per fermare il loop

    // Stati reattivi (Svelte 5 Runes)
    let statusKey = $state('checking');
    let messageKey = $state('checking');
    let users = $state([]);
    let userLoggedIn = $state(null);
    
    // Stati specifici per il modulo Driver
    let driverProfile = $state(null);
    let loadingDriver = $state(false);
    
    // Campi del Form Driver Profile
    let licenseNumber = $state('');
    let phone = $state('');
    let formError = $state('');
    let formSuccess = $state(false);

    // 1. Verifica lo stato del backend
    async function checkBackend() {
        try {
            const res = await fetch('http://localhost:8080/health');
            if (!res.ok) throw new Error('Offline');
            const data = await res.json();
            if (data.status === 'OK') {
                statusKey = 'ok';
                messageKey = 'msg_running';
            } else {
                statusKey = 'error';
            }
        } catch (e) {
            statusKey = 'offline';
            messageKey = 'offline_msg';
        }
    }
    
    // 2. Carica i dettagli del profilo Driver
    async function loadDriverProfile(userID) {
        // Se abbiamo già cercato o trovato il profilo, evitiamo chiamate doppie
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
                // 🌟 Gestione pulita del 404: Nessun driver a DB, l'utente vedrà il Form!
                driverProfile = null; 
            } else {
                console.error("Errore API Driver:", res.statusText);
            }
        } catch (err) {
            console.error("Errore nel caricamento del profilo driver", err);
        } finally {
            loadingDriver = false;
        }
    }

    // 3. Crea il profilo Driver
    async function handleCreateDriver(e) {
        e.preventDefault();
        formError = '';
        formSuccess = false;
        
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

            formSuccess = true;
            driverProfile = await res.json();
        } catch (err) {
            formError = "offline";
        }
    }

    // 4. Carica la lista utenti globale
    async function loadUsers() {
        try {
            const token = localStorage.getItem('token');
            const response = await fetch('http://localhost:8080/api/users', {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            if (response.ok) {
                users = await response.json();
            }
        } catch (err) {
            console.error("Errore caricamento utenti", err);
        }
    }

    // 🌟 L'effetto esegue i caricamenti iniziali in modo sicuro senza andare in loop
    $effect(() => {
        checkBackend();
        const token = localStorage.getItem('token');
        const savedUser = localStorage.getItem('user_data');

        if (token && savedUser) {
            try {
                userLoggedIn = JSON.parse(savedUser);
                // Usiamo untrack per evitare che le mutazioni di stato interne ri-attivino l'effect
                untrack(() => {
                    loadUsers();
                    loadDriverProfile(userLoggedIn.id);
                });
            } catch (e) {
                localStorage.removeItem('user_data');
            }
        }
    });
</script>

<div class="dashboard-header card">
    <div>
        <h1>{i18n.t('home.title')}</h1>
        <p class="system-status">
            {i18n.t('dashboard.system_infra')}: 
            <span class="status-indicator status-{statusKey}">
                {statusKey === 'ok' ? 'ONLINE' : 'OFFLINE'}
            </span>
        </p>
    </div>
</div>

{#if userLoggedIn}
    <div class="dashboard-grid">
        
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

        <div class="card fleet-panel">
            <h2>{i18n.t('dashboard.fleet_panel_title')}</h2>
            <p class="panel-desc">{i18n.t('home.user_list')}</p>
            
            {#if users.length === 0}
                <p class="empty-msg">{i18n.t('home.loading_or_empty')}</p>
            {:else}
                <div class="table-container">
                    <table class="erp-table">
                        <thead>
                            <tr>
                                <th>{i18n.t('dashboard.operator')}</th>
                                <th>{i18n.t('dashboard.email_contact')}</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each users as u}
                                <tr>
                                    <td><strong>{u.name}</strong></td>
                                    <td><code>{u.email}</code></td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            {/if}
        </div>

    </div>
{:else}
    <div class="card guest-box">
        <h2>{i18n.t('home.area_reserved')}</h2>
        <p>{i18n.t('home.guest_msg')}</p>
        <div class="auth-buttons">
            <a href="/login" class="btn btn-light">{i18n.t('dashboard.btn_guest_login')}</a>
            <a href="/signup" class="btn btn-primary">{i18n.t('dashboard.btn_guest_signup')}</a>
        </div>
    </div>
{/if}

<style>
    .dashboard-header { margin-bottom: 2rem; display: flex; justify-content: space-between; align-items: center; }
    .dashboard-header h1 { margin: 0; font-size: 1.75rem; color: var(--text-dark); }
    .system-status { margin: 0.25rem 0 0 0; font-size: 0.875rem; color: #64748b; }
    
    .status-indicator { font-weight: 700; padding: 0.15rem 0.4rem; border-radius: 4px; font-size: 0.75rem; }
    .status-ok { background: #d1fae5; color: #065f46; }
    .status-offline { background: #fee2e2; color: #991b1b; }

    .dashboard-grid { display: grid; grid-template-columns: 1fr; gap: 2rem; }
    @media (min-width: 968px) {
        .dashboard-grid { grid-template-columns: 1.2fr 0.8fr; }
    }

    h2 { margin-top: 0; font-size: 1.35rem; color: var(--text-dark); display: flex; align-items: center; gap: 0.5rem; }
    .panel-desc { color: #64748b; font-size: 0.9rem; margin-top: -0.5rem; margin-bottom: 1.5rem; }

    /* Modulo Driver UI */
    .driver-badge-box { display: flex; justify-content: space-between; align-items: center; background: #f8fafc; border: 1px solid var(--border); padding: 1.25rem; border-radius: 8px; }
    .profile-info p { margin: 0.25rem 0; }
    .status-box { text-align: right; }
    .status-tip { font-size: 0.75rem; color: #64748b; margin: 0.5rem 0 0 0; }

    .info-alert { background: #eff6ff; color: #1e40af; border: 1px solid #bfdbfe; padding: 0.75rem 1rem; border-radius: 6px; margin-bottom: 1.2rem; font-size: 0.9rem; }
    .erp-form { display: flex; flex-direction: column; gap: 1rem; }
    .form-row { display: flex; gap: 1rem; flex-wrap: wrap; }
    .form-row .form-group { flex: 1; min-width: 200px; }
    .form-group { display: flex; flex-direction: column; gap: 0.4rem; }
    .form-group label { font-size: 0.85rem; font-weight: 600; color: #475569; }
    .form-group input { padding: 0.55rem; border: 1px solid var(--border); border-radius: 6px; font-size: 0.95rem; }
    .form-group input:focus { border-color: var(--accent); outline: none; }

    /* Tabelle ERP */
    .table-container { border: 1px solid var(--border); border-radius: 8px; overflow: hidden; }
    .erp-table { width: 100%; border-collapse: collapse; text-align: left; font-size: 0.9rem; }
    .erp-table th { background: #f1f5f9; padding: 0.75rem; color: #475569; font-weight: 600; }
    .erp-table td { padding: 0.75rem; border-top: 1px solid var(--border); }
    
    /* Guest Box */
    .guest-box { text-align: center; padding: 4rem 2rem; background: var(--surface); border: 1px dashed #cbd5e1; }
    .auth-buttons { margin-top: 2rem; display: flex; justify-content: center; gap: 1rem; }
    .btn-light { background: #f1f5f9; color: #334155; border: 1px solid #cbd5e1; }
    .btn-light:hover { background: #e2e8f0; }

    .alert { padding: 0.75rem; border-radius: 6px; margin-bottom: 1rem; font-size: 0.9rem; }
    .alert-danger { background: #fee2e2; color: #991b1b; border: 1px solid #fca5a5; }
    .loader { color: var(--accent); font-weight: 600; font-style: italic; }
</style>