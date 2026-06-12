<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';
    import { untrack } from 'svelte';

    let users = $state([]);
    let isAuthorized = $state(false);
    let loading = $state(true);

    async function loadUsers() {
        try {
            const token = localStorage.getItem('token');
            const response = await fetch('http://localhost:8080/api/users', {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            if (response.ok) {
                users = await response.json();
            } else if (response.status === 403 || response.status === 401) {
                isAuthorized = false;
            }
        } catch (err) {
            console.error("Error loading users", err);
        } finally {
            loading = false;
        }
    }

    $effect(() => {
        // Leggiamo i dati in locale
        const savedUser = localStorage.getItem('user_data');
        if (savedUser) {
            try {
                const parsedUser = JSON.parse(savedUser);
                const roles = parsedUser.roles || [];
                
                // Controllo accessi per Admin e Dispatcher
                if (roles.includes('admin') || roles.includes('dispatcher')) {
                    isAuthorized = true;
                    // untrack evita che le dipendenze interne a loadUsers riattivino l'effetto
                    untrack(() => loadUsers());
                    return;
                }
            } catch (e) {
                console.error("i18n parsing error", e);
            }
        }
        
        isAuthorized = false;
        loading = false;
    });
</script>

{#if loading}
    <div class="loading-screen">
        <div class="loader">Loading...</div>
    </div>
{:else if !isAuthorized}
    <div class="card access-denied">
        <div class="icon">🚫</div>
        <h2>{i18n.t('home.access_denied_title')}</h2>
        <p>{i18n.t('home.access_denied_msg')}</p>
        <a href="/" class="btn btn-primary">{i18n.t('home.back_to_dashboard')}</a>
    </div>
{:else}
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
                            <th>{i18n.t('home.active_roles_table')}</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each users as u}
                            <tr>
                                <td><strong>{u.name}</strong></td>
                                <td><code>{u.email}</code></td>
                                <td>
                                    {#each u.roles || [] as r}
                                        <span class="badge-role-inline">{r}</span>
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
    .panel-desc { color: #64748b; font-size: 0.9rem; margin-bottom: 1.5rem; }
    
    .access-denied { max-width: 500px; margin: 4rem auto; text-align: center; padding: 3rem 2rem; border: 1px solid #fca5a5; background: #fff5f5; }
    .access-denied .icon { font-size: 3rem; margin-bottom: 1rem; }
    .access-denied h2 { color: #991b1b; display: flex; justify-content: center; border: none; }
    .access-denied p { color: #7f1d1d; margin-bottom: 2rem; font-size: 0.95rem; }

    .table-container { border: 1px solid var(--border); border-radius: 8px; overflow: hidden; }
    .erp-table { width: 100%; border-collapse: collapse; text-align: left; font-size: 0.9rem; }
    .erp-table th { background: #f1f5f9; padding: 0.75rem; color: #475569; font-weight: 600; }
    .erp-table td { padding: 0.75rem; border-top: 1px solid var(--border); }
    .loading-screen { text-align: center; margin-top: 5rem; }
    
    .badge-role-inline {
        background: #e2e8f0; color: #334155; font-size: 0.7rem; font-weight: bold;
        padding: 0.15rem 0.4rem; border-radius: 4px; margin-right: 0.25rem; text-transform: uppercase;
    }
</style>