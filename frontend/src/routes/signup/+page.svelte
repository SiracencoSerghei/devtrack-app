<script>
    import { onMount } from 'svelte';

    
    let name = $state('');
    let email = $state('');
    let password = $state('');
    let errorMessage = $state('');
    let successMessage = $state('');

    async function handleSignUp(e) {
        e.preventDefault();
        errorMessage = '';
        successMessage = '';

        try {
            
            const response = await fetch('http://localhost:8080/api/signup', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name, email, password })
            });

            if (!response.ok) {
                const text = await response.text();
                throw new Error(text || 'Errore durante la registrazione');
            }

            successMessage = 'Registrazione completata! Verrai reindirizzato al login...';
            
            
            setTimeout(() => {
                window.location.href = '/login';
            }, 2000);

        } catch (err) {
            errorMessage = err.message;
        }
    }
</script>

<div class="auth-container">
    <h2>Crea un nuovo account</h2>
    
    {#if errorMessage}
        <div class="alert alert-danger">{errorMessage}</div>
    {/if}
    {#if successMessage}
        <div class="alert alert-success">{successMessage}</div>
    {/if}

    <form onsubmit={handleSignUp}>
        <div class="form-group">
            <label for="name">Nome completo</label>
            <input type="text" id="name" bind:value={name} required placeholder="Es. Mario Rossi" />
        </div>

        <div class="form-group">
            <label for="email">Indirizzo Email</label>
            <input type="email" id="email" bind:value={email} required placeholder="mario@rossi.it" />
        </div>

        <div class="form-group">
            <label for="password">Password</label>
            <input type="password" id="password" bind:value={password} required placeholder="••••••••" />
        </div>

        <button type="submit" class="btn-submit">Registrati</button>
    </form>
    
    <p class="switch-auth">Hai già un account? <a href="/login">Accedi qui</a></p>
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
    .alert { padding: 0.8rem; margin-bottom: 1rem; border-radius: 4px; font-weight: 500; }
    .alert-danger { background: #ffe3e3; color: #d60000; border: 1px solid #fcc; }
    .alert-success { background: #e3ffe3; color: #008a00; border: 1px solid #cfc; }
    .switch-auth { text-align: center; margin-top: 1.5rem; color: #666; }
    .switch-auth a { color: #0076ff; text-decoration: none; }
</style>