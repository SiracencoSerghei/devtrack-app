<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';
    
    let { children } = $props();
    let user = $state(null);
    let currentLang = $derived(i18n.lang);

    function checkAuth() {
        const token = localStorage.getItem('token');
        const userData = localStorage.getItem('user_data');
        if (token && userData) {
            try { user = JSON.parse(userData); } catch (e) { user = null; }
        } else { user = null; }
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
</script>

<nav>

    <div class="nav-links">
        <a href="/">{i18n.t('nav.home')}</a>
        <a href="/about">{i18n.t('nav.about')}</a>
        <a href="/health">{i18n.t('nav.health')}</a>
    </div>

    <div class="lang-selector">
        <button class={i18n.lang === 'it' ? 'active' : ''} onclick={() => i18n.lang = 'it'}>IT</button>
        <button class={i18n.lang === 'uk' ? 'active' : ''} onclick={() => i18n.lang = 'uk'}>UA</button>
        <button class={i18n.lang === 'en' ? 'active' : ''} onclick={() => i18n.lang = 'en'}>EN</button>
    </div>

    <div class="auth-links">
        {#if user}
            <span class="welcome-msg">{i18n.t('auth.welcome')}, <strong>{user.name}</strong>!</span>
            <button onclick={logout} class="btn-logout">{i18n.t('auth.logout')}</button>
        {:else}
            <a href="/login" class="btn-login">{i18n.t('auth.login') === 'auth.login' ? 'Login' : i18n.t('auth.login')}</a>
            <a href="/signup" class="btn-signup">{i18n.t('auth.signup') === 'auth.signup' ? 'Sign Up' : i18n.t('auth.signup')}</a>
        {/if}
    </div>
</nav>

<main>
    {@render children()}
</main>

<style>
    nav { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #ccc; padding: 1rem; font-family: sans-serif; }
    nav a { margin-right: 1rem; text-decoration: none; color: #0076ff; font-weight: 500; }
    
    .lang-selector button {
        background: none; border: 1px solid #ccc; padding: 0.2rem 0.5rem; margin: 0 0.2rem; cursor: pointer; border-radius: 4px; font-size: 0.8rem;
    }
    .lang-selector button.active { background: #0076ff; color: white; border-color: #0076ff; font-weight: bold; }
    
    .welcome-msg { margin-right: 1rem; color: #333; }
    .btn-logout { background: #ff3e00; color: white; border: none; padding: 0.4rem 0.8rem; border-radius: 4px; cursor: pointer; font-weight: bold; }
    .btn-signup { background: #0076ff; color: white !important; padding: 0.4rem 0.8rem; border-radius: 4px; }
</style>