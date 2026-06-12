<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';

    let email = $state('');
    let password = $state('');
    let errorKey = $state(''); 
    let customError = $state('');

    let showPassword = $state(false); 

    async function handleLogin(e) {
        e.preventDefault();
        errorKey = '';
        customError = '';

        try {
            const response = await fetch('http://localhost:8080/api/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password })
            });

            if (!response.ok) {
                const errorData = await response.json();
                if (response.status === 401 || errorData.error === 'Credenziali non valide') {
                    errorKey = 'invalid_creds';
                } else {
                    customError = errorData.error || response.statusText;
                }
                return;
            }

            const data = await response.json();
            localStorage.setItem('token', data.token);
            localStorage.setItem('user_data', JSON.stringify(data.user));
            window.location.href = '/';

        } catch (err) {
            errorKey = 'offline_msg'; 
        }
    }
</script>

<div class="auth-container">
    <h2>{i18n.t('login_page.title')}</h2>
    
    {#if errorKey}
        <div class="alert alert-danger">
            {errorKey === 'offline_msg' ? i18n.t('server.offline_msg') : i18n.t(`login_page.${errorKey}`)}
        </div>
    {:else if customError}
        <div class="alert alert-danger">{customError}</div>
    {/if}

    <form onsubmit={handleLogin}>
        <div class="form-group">
            <label for="email">{i18n.t('login_page.email_label')}</label>
            <input type="email" id="email" bind:value={email} required placeholder="mario@rossi.it" />
        </div>

        <div class="form-group">
            <label for="password">{i18n.t('login_page.password_label')}</label>
            
            <div class="password-wrapper">
                <input 
                    type={showPassword ? 'text' : 'password'} 
                    id="password" 
                    bind:value={password} 
                    required 
                    placeholder="••••••••" 
                />
                <button 
                    type="button" 
                    class="toggle-password" 
                    onclick={() => showPassword = !showPassword}
                    title={showPassword ? i18n.t('auth.hide_password') : i18n.t('auth.show_password')}
                >
                    {#if showPassword}
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88" />
                        </svg>
                    {:else}
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z" />
                            <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
                        </svg>
                    {/if}
                </button>
            </div>
        </div>

        <button type="submit" class="btn-submit">{i18n.t('login_page.btn_submit')}</button>
    </form>
    
    <p class="switch-auth">
        {i18n.t('login_page.no_account')} 
        <a href="/signup">{i18n.t('login_page.register_link')}</a>
    </p>
</div>

<style>
    .auth-container { max-width: 400px; margin: 3rem auto; padding: 2rem; border: 1px solid #e0e0e0; border-radius: 8px; font-family: sans-serif; box-shadow: 0 4px 6px rgba(0,0,0,0.05); }
    h2 { margin-top: 0; color: #333; text-align: center; }
    .form-group { margin-bottom: 1.2rem; display: flex; flex-direction: column; }
    label { margin-bottom: 0.4rem; font-weight: 500; color: #666; }
    input { padding: 0.6rem; border: 1px solid #ccc; border-radius: 4px; font-size: 1rem; }
    input:focus { border-color: #0076ff; outline: none; }
    .btn-submit { width: 100%; padding: 0.7rem; background: #0076ff; color: white; border: none; border-radius: 4px; font-size: 1rem; font-weight: bold; cursor: pointer; }
    .btn-submit:hover { background: #005bc5; }
    .alert { padding: 0.8rem; margin-bottom: 1rem; border-radius: 4px; font-weight: 500; background: #ffe3e3; color: #d60000; border: 1px solid #fcc; }
    .switch-auth { text-align: center; margin-top: 1.5rem; color: #666; }
    .switch-auth a { color: #0076ff; text-decoration: none; }
    .password-wrapper {
        position: relative;
        display: flex;
        align-items: center;
    }
    .password-wrapper input {
        width: 100%;
        padding-right: 2.5rem;
        box-sizing: border-box;
    }
    .toggle-password {
        position: absolute;
        right: 0.6rem;
        background: none;
        border: none;
        cursor: pointer;
        color: #666;
        padding: 0.2rem;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .toggle-password:hover {
        color: #0076ff;
    }
    .icon {
        width: 1.25rem;
        height: 1.25rem;
    }
</style>