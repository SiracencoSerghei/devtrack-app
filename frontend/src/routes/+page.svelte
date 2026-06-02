<script>

    let status = $state('Verifica in corso...');
    let message = $state('');
    let users = $state([]);
    let userLoggedIn = $state(null);

    async function checkBackend() {
        try {
            const res = await fetch('http://localhost:8080/health');
            if (!res.ok) throw new Error('Errore del server');
            const data = await res.json();
            status = data.status || 'OK';
            message = data.messaggio || '';
        } catch (e) {
            status = 'Server offline';
            message = e.message;
        }
    }

    async function loadUsers() {
        try {
            const token = localStorage.getItem('access_token');
            if (!token) {
                console.error("Token non trovato. Utente non autenticato.");
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

                    localStorage.removeItem('access_token');
                    localStorage.removeItem('user_data');
                    userLoggedIn = null;
                    alert("Sessione scaduta. Per favore, effettua nuovamente il login.");
                }
                throw new Error(`Errore del server: ${response.status}`);
            }

            const data = await response.json();
            users = data || []; 

        } catch (err) {
            console.error("Impossibile caricare gli utenti:", err);
        }
    }

    $effect(() => {

        checkBackend();

        const token = localStorage.getItem('access_token');
        const savedUser = localStorage.getItem('user_data');

        if (token && savedUser) {
            try {

                userLoggedIn = JSON.parse(savedUser);

                loadUsers();
            } catch (e) {
                console.error("Errore nel parsing dei dati utente", e);
                localStorage.removeItem('user_data');
            }
        }
    });

    function handleLogout() {
        localStorage.removeItem('access_token');
        localStorage.removeItem('user_data');
        userLoggedIn = null;
        users = [];

        window.location.href = '/login';
    }
</script>

<h1>Benvenuto su DevTrack</h1>

<section class="status">
    <p>Stato del server: <strong class="status-{status.toLowerCase().replace(' ', '-')}">{status}</strong></p>
    {#if message}<p>Messaggio: {message}</p>{/if}
</section>

<hr />

{#if userLoggedIn}
    <section class="dashboard">
        
        <p>Ecco la lista degli utenti registrati nel sistema:</p>
        
        {#if users.length === 0}
            <p class="empty-msg">Nessun utente trovato o caricamento in corso...</p>
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
        <h2>Area Riservata</h2>
        <p>Per vedere gli utenti registrati ed accedere alle funzionalità di DevTrack, devi avere un account.</p>
        <div class="auth-buttons">
            <a href="/login" class="btn btn-login">Accedi (Login)</a>
            <a href="/signup" class="btn btn-signup">Registrati (Sign up)</a>
        </div>
    </section>
{/if}

<style>
    section { margin: 2rem 0; font-family: sans-serif; }
   
    .status-ok { color: green; }
    .status-server-offline { color: red; }

    .btn-logout {
        background: #ef4444;
        color: white;
        border: none;
        padding: 0.4rem 0.8rem;
        border-radius: 4px;
        cursor: pointer;
        font-weight: bold;
    }
    .btn-logout:hover { background: #dc2626; }

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