<script>
    let email = $state('');
    let password = $state('');
    let errorMessage = $state('');

    async function handleLogin(e) {
        e.preventDefault();
        errorMessage = '';

        try {
            const response = await fetch('http://localhost:8080/api/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || 'Credenziali non valide');
            }

            const data = await response.json();
            
            localStorage.setItem('access_token', data.access_token);
            localStorage.setItem('user_data', JSON.stringify(data.user));

            window.location.href = '/';

        } catch (err) {
            errorMessage = err.message;
        }
    }
</script>

<div class="auth-container">
    <h2>Accedi a DevTrack</h2>
    
    {#if errorMessage}
        <div class="alert alert-danger">{errorMessage}</div>
    {/if}

    <form onsubmit={handleLogin}>
        <div class="form-group">
            <label for="email">Indirizzo Email</label>
            <input type="email" id="email" bind:value={email} required placeholder="mario@rossi.it" />
        </div>

        <div class="form-group">
            <label for="password">Password</label>
            <input type="password" id="password" bind:value={password} required placeholder="••••••••" />
        </div>

        <button type="submit" class="btn-submit">Accedi</button>
    </form>
    
    <p class="switch-auth">Non hai un account? <a href="/signup">Registrati qui</a></p>
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
</style>