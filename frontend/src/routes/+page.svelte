<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';

    let statusKey = $state('checking');
    let message = $state('');
    let messageKey = $state('');
    let users = $state([]);
    let userLoggedIn = $state(null);

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
                messageKey = ''; 
            }
        } catch (e) {
            statusKey = 'offline';
            messageKey = 'offline_msg';
        }
    }
    
    async function loadUsers() {
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                console.error(i18n.t('errors.no_token'));
                return;
            }

            const response = await fetch('http://localhost:8080/api/users', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            });

            if (!response.ok) {
                if (response.status === 401) {
                    localStorage.removeItem('token');
                    localStorage.removeItem('user_data');
                    userLoggedIn = null;
                    alert(i18n.t('errors.session_expired'));
                }
                throw new Error(`${i18n.t('errors.server_error')}: ${response.status}`);
            }

            const data = await response.json();
            users = data || []; 

        } catch (err) {
            console.error(i18n.t('errors.load_failed'), err);
        }
    }

    $effect(() => {
        checkBackend();

        const token = localStorage.getItem('token');
        const savedUser = localStorage.getItem('user_data');

        if (token && savedUser) {
            try {
                userLoggedIn = JSON.parse(savedUser);
                loadUsers();
            } catch (e) {
                console.error("Error parsing user data", e);
                localStorage.removeItem('user_data');
            }
        }
    });

    function handleLogout() {
        localStorage.removeItem('token');
        localStorage.removeItem('user_data');
        userLoggedIn = null;
        users = [];
        window.location.href = '/login';
    }
</script>

<h1>{i18n.t('home.title')}</h1>

<section class="status">
    <p>
        {i18n.t('server.status')} 
        <strong class="status-{statusKey}">
            {i18n.t(`server.${statusKey}`)}
        </strong>
    </p>
    
    {#if messageKey}
        <p>{i18n.t('server.message')} {i18n.t(`server.${messageKey}`)}</p>
    {/if}
</section>

<hr />

{#if userLoggedIn}
    <section class="dashboard">
        
        <p>{i18n.t('home.user_list')}</p>
        
        {#if users.length === 0}
            <p class="empty-msg">{i18n.t('home.loading_or_empty')}</p>
        {:else}
            <ul class="user-list">
                {#each users as u}
                    <li>
                        <span class="user-name">{u.name}</span> 
                        <span class="user-email">({u.email})</span>
                    </li>
                {/each}
            </ul>
        {/if}
    </section>
{:else}
    <section class="guest-box">
        <h2>{i18n.t('home.area_reserved')}</h2>
        <p>{i18n.t('home.guest_msg')}</p>
        <div class="auth-buttons">
            <a href="/login" class="btn btn-login">{i18n.t('auth.login')}</a>
            <a href="/signup" class="btn btn-signup">{i18n.t('auth.signup')}</a>
        </div>
    </section>
{/if}

<style>
    section { margin: 2rem 0; font-family: sans-serif; }

    .status-ok { color: green; }
    .status-offline, .status-error { color: red; }
    .status-checking { color: orange; }

    .guest-box {
        background: #f9f9f9;
        padding: 2rem;
        border-radius: 8px;
        border: 1px dashed #ccc;
        text-align: center;
    }
    .auth-buttons { margin-top: 1.5rem; }
    .btn {
        display: inline-block;
        padding: 0.6rem 1.2rem;
        margin: 0 0.5rem;
        text-decoration: none;
        font-weight: bold;
        border-radius: 4px;
    }
    .btn-login { background: #eee; color: #333; border: 1px solid #ccc; }
    .btn-signup { background: #0076ff; color: white; }
    .btn-signup:hover { background: #005bc5; }

    .user-list { list-style: none; padding: 0; }
    .user-list li { 
        padding: 0.6rem; 
        background: #f1f5f9; 
        margin-bottom: 0.5rem; 
        border-radius: 4px; 
        border-left: 4px solid #0076ff;
    }
    .user-name { font-weight: bold; color: #1e293b; }
    .user-email { color: #64748b; margin-left: 0.5rem; }
    .empty-msg { color: #666; font-style: italic; }
</style>