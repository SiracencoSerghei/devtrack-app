<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';
    import '../app.css';
    let { children } = $props();
    let user = $state(null);
    let langOpen = $state(false);

    function checkAuth() {
        const token = localStorage.getItem('token');
        const userData = localStorage.getItem('user_data');

        if (token && userData) {
            try {
                user = JSON.parse(userData);
            } catch {
                user = null;
            }
        } else {
            user = null;
        }
    }

    $effect(() => {
        checkAuth();
    });

    function logout() {
        localStorage.removeItem('token');
        localStorage.removeItem('user_data');
        user = null;
        window.location.href = '/login';
    }

    function setLang(lang) {
        i18n.setLang(lang);
        langOpen = false;
    }

    function toggleLang() {
        langOpen = !langOpen;
    }
</script>

<header class="main-header">
    <div class="header-container">
        <a href="/" class="logo">DevTrack <span>ERP</span></a>

        <nav class="nav-links">
            <a href="/">{i18n.t('nav.home')}</a>
            <a href="/about">{i18n.t('nav.about')}</a>
            <a href="/health">{i18n.t('nav.health')}</a>
        </nav>

        <div class="actions-wrapper">

            <div class="lang-selector">
                <button class="lang-current" onclick={toggleLang}>
                    {i18n.lang.toUpperCase()}
                </button>

                {#if langOpen}
                    <div class="lang-dropdown">
                        <button class={i18n.lang === 'it' ? 'active' : ''} onclick={() => setLang('it')}>IT</button>
                        <button class={i18n.lang === 'uk' ? 'active' : ''} onclick={() => setLang('uk')}>UA</button>
                        <button class={i18n.lang === 'en' ? 'active' : ''} onclick={() => setLang('en')}>EN</button>
                    </div>
                {/if}
            </div>

            <div class="auth-links">
                {#if user}
                    <div class="user-badge">
                        <span>
                            {i18n.t('auth.welcome') || 'Welcome'}, <strong>{user.name}</strong>
                        </span>
                    </div>

                    <button onclick={logout} class="btn btn-logout">
                        {i18n.t('auth.logout') || 'Logout'}
                    </button>
                {:else}
                    <a href="/login" class="btn btn-login-nav">
                        {i18n.t('auth.login') || 'Login'}
                    </a>
                    <a href="/signup" class="btn btn-primary">
                        {i18n.t('auth.signup') || 'Sign Up'}
                    </a>
                {/if}
            </div>

        </div>
    </div>
</header>

<main class="container">
    {@render children()}
</main>

<style>
    .main-header {
        background: var(--primary);
        color: white;
        border-bottom: 1px solid var(--primary-light);
    }
    .header-container {
        max-width: 1200px;
        margin: 0 auto;
        padding: 1rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .logo {
        font-size: 1.25rem;
        font-weight: 800;
        color: white;
        text-decoration: none;
    }
    .logo span {
        color: var(--accent);
    }
    .nav-links a {
        color: #94a3b8;
        margin-left: 1.5rem;
        text-decoration: none;
        font-weight: 500;
        transition: color 0.2s;
    }
    .nav-links a:hover {
        color: white;
    }

    .actions-wrapper {
        display: flex;
        align-items: center;
        gap: 1.5rem;
    }
    .lang-selector {
        position: relative;
        display: inline-block;
    }

    .lang-current {
        background: transparent;
        border: 1px solid #334155;
        padding: 0.25rem 0.5rem;
        color: #94a3b8;
        cursor: pointer;
        border-radius: 4px;
        font-size: 0.75rem;
        font-weight: bold;
        transition: all 0.2s;
    }

    .lang-current:hover {
        color: white;
        border-color: #475569;
    }

    .lang-dropdown {
        position: absolute;
        top: calc(100% + 6px);
        right: 0;

        background: #0f172a;
        border: 1px solid #334155;
        border-radius: 6px;

        min-width: 70px;
        z-index: 100;

        display: flex;
        flex-direction: column;

        box-shadow: 0 10px 25px rgba(0,0,0,0.3);
    }

    .lang-dropdown button {
        width: 100%;
        padding: 0.4rem 0.6rem;

        background: transparent;
        border: none;

        color: #94a3b8;
        cursor: pointer;
        font-size: 0.75rem;

        text-align: left;
        transition: background 0.15s, color 0.15s;
    }

    .lang-dropdown button:hover {
        background: #1e293b;
        color: white;
    }

    .lang-dropdown button.active {
        background: var(--accent);
        color: white;
    }

    /* AUTH */
    .auth-links {
        display: flex;
        align-items: center;
        gap: 1rem;
    }

    .user-badge {
        background: #1e293b;
        padding: 0.4rem 0.8rem;
        border-radius: 6px;
        font-size: 0.875rem;
    }

    .btn-logout {
        background: #ef4444;
        color: white;
        border: none;
        padding: 0.4rem 0.8rem;
        border-radius: 6px;
        cursor: pointer;
        font-weight: bold;
    }

    .btn-login-nav {
        color: white;
        text-decoration: none;
        font-weight: 600;
    }
</style>