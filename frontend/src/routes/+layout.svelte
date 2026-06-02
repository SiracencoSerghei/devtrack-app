<script>

    let { children } = $props();

    let user = $state(null);

    function checkAuth() {
        const token = localStorage.getItem('access_token');
        const userData = localStorage.getItem('user_data');
        
        if (token && userData) {
            try {
                user = JSON.parse(userData);
            } catch (e) {
                console.error("Errore nel parsing dei dati utente nel layout:", e);
                user = null;
            }
        } else {
            user = null;
        }
    }


    $effect(() => {
        checkAuth();

        window.addEventListener('storage', checkAuth);
        return () => window.removeEventListener('storage', checkAuth);
    });

    function logout() {
        localStorage.removeItem('access_token');
        localStorage.removeItem('user_data');
        user = null;
        window.location.href = '/login';
    }
</script>

<nav>
    <div class="nav-links">
        <a href="/">Home</a>
        <a href="/about">About</a>
        <a href="/health">Health</a>
    </div>

    <div class="auth-links">
        {#if user}
            <div class="dashboard-header">
            <span class="welcome-msg">Ciao, Benvenuto a bordo! <strong>{user.name}</strong>!</span>
            <button onclick={logout} class="btn-logout">Esci</button>
            </div>
        {:else}
            <a href="/login" class="btn-login">Accedi</a>
            <a href="/signup" class="btn-signup">Registrati</a>
        {/if}
    </div>
</nav>

<main>
    {@render children()}
</main>

<style>
    nav { 
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-bottom: 1px solid #ccc; 
        padding: 1rem; 
        font-family: sans-serif;
    }
    nav a { 
        margin-right: 1rem; 
        text-decoration: none; 
        color: #0076ff; 
        font-weight: 500;
    }
    nav a:hover {
        text-decoration: underline;
    }
    .welcome-msg {
        margin-right: 1rem;
        color: #333;
    }
    .btn-logout {
        background: #ff3e00;
        color: white;
        border: none;
        padding: 0.4rem 0.8rem;
        border-radius: 4px;
        cursor: pointer;
        font-weight: bold;
    }
    .btn-logout:hover {
        background: #d63300;
    }
    .btn-signup {
        background: #0076ff;
        color: white !important;
        padding: 0.4rem 0.8rem;
        border-radius: 4px;
    }
    .btn-signup:hover {
        background: #005bc5;
        text-decoration: none !important;
    }

    .dashboard-header {
        /* display: flex; */
        justify-content: space-between;
        align-items: center;
    }
</style>