<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';

    let statusKey = $state('checking');
    let userLoggedIn = $state(null);

    let hasDriverRole = $derived(userLoggedIn?.roles?.includes('driver') || false);
    let hasStaffRole = $derived(userLoggedIn?.roles?.includes('admin') || userLoggedIn?.roles?.includes('dispatcher') || false);
    
    async function checkBackend() {
        try {
            const res = await fetch('http://localhost:8080/health');
            if (res.ok) statusKey = 'ok';
            else statusKey = 'offline';
        } catch (e) {
            statusKey = 'offline';
        }
    }

    $effect(() => {
        checkBackend();
        const savedUser = localStorage.getItem('user_data');
        if (savedUser) {
            try { userLoggedIn = JSON.parse(savedUser); } catch (e) {}
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
    <div class="welcome-box card">
        <p>👋 {i18n.t('home.welcome')}, <strong>{userLoggedIn.name}</strong>! {i18n.t('home.active_roles')} 
            {#each userLoggedIn.roles || [] as role}
                <span class="role-tag">{role}</span>
            {/each}
        </p>
        <p class="sub-text">{i18n.t('home.select_department')}</p>
    </div>

    <div class="hub-grid">
        <!-- Card 1: Area Driver -->
        {#if hasDriverRole}
            <a href="/driver" class="hub-card">
                <div class="icon">🚚</div>
                <h3>{i18n.t('dashboard.driver_panel_title')}</h3>
                <p>{i18n.t('dashboard.driver_panel_desc_placeholder' || 'Accedi ai tuoi dati logistici')}</p>
            </a>
        {/if}

        <!-- Card 2: Area Admin/Dispatcher -->
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
            <a href="/login" class="btn btn-light">{i18n.t('dashboard.btn_guest_login')}</a>
            <a href="/signup" class="btn btn-primary">{i18n.t('dashboard.btn_guest_signup')}</a>
        </div>
    </div>
{/if}

<style>
    .dashboard-header { margin-bottom: 1.5rem; padding: 1.5rem; }
    .system-status { margin: 0.25rem 0 0 0; font-size: 0.875rem; color: #64748b; }
    .status-indicator { font-weight: 700; padding: 0.15rem 0.4rem; border-radius: 4px; font-size: 0.75rem; }
    .status-ok { background: #d1fae5; color: #065f46; }
    .status-offline { background: #fee2e2; color: #991b1b; }

    .welcome-box { margin-bottom: 2rem; padding: 1.25rem; background: #f1f5f9; border: 1px solid var(--border); }
    .welcome-box p { margin: 0.25rem 0; }
    .sub-text { font-size: 0.875rem; color: #64748b; }
    
    .role-tag {
        display: inline-block; background: var(--accent, #2563eb); color: white; padding: 0.15rem 0.5rem;
        font-size: 0.75rem; font-weight: bold; border-radius: 4px; margin-left: 0.35rem; text-transform: uppercase;
    }

    .hub-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr)); gap: 1.5rem; }
    .hub-card { 
        background: white; border: 1px solid var(--border); padding: 2rem; border-radius: 12px; 
        text-decoration: none; color: inherit; transition: all 0.2s ease-in-out; display: block;
    }
    .hub-card:hover { transform: translateY(-3px); border-color: var(--accent, #2563eb); box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.05); }
    .hub-card .icon { font-size: 2.5rem; margin-bottom: 1rem; }
    .hub-card h3 { margin: 0 0 0.5rem 0; font-size: 1.2rem; color: var(--text-dark); }
    .hub-card p { margin: 0; font-size: 0.9rem; color: #64748b; line-height: 1.4; }

    .guest-box { text-align: center; padding: 4rem 2rem; }
    .auth-buttons { margin-top: 2rem; display: flex; justify-content: center; gap: 1rem; }
    .btn-light { background: #f1f5f9; color: #334155; border: 1px solid #cbd5e1; text-decoration: none; padding: 0.5rem 1rem; border-radius: 6px; }
</style>