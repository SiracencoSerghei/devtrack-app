<script>
    let status = $state('Verifica in corso...');
    let message = $state('');
    
    let name = $state('');
    let email = $state('');
    let users = $state([]);
    let formMessage = $state('');

    async function checkBackend() {
        try {
            const res = await fetch('http://localhost:8080');
            if (!res.ok) throw new Error('Errore del server');
            const data = await res.json();
            status = data.status || 'OK';
            message = data.message || '';
        } catch (e) {
            status = 'Server offline';
            message = e.message;
        }
    }

    async function loadUsers() {
        try {
            const res = await fetch('http://localhost:8080/users');
            if (res.ok) {
                users = await res.json();
            }
        } catch (e) {
            console.error("Errore nel caricamento degli utenti:", e);
        }
    }

    async function createUser(event) {
        event.preventDefault();
        formMessage = '';
        try {
            const res = await fetch('http://localhost:8080/users', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name, email })
            });

            const data = await res.json();

            if (!res.ok) {
                if (data.error === "name and email are required") {
                    formMessage = "Il nome e l'email sono obbligatori.";
                } else if (data.error === "invalid email format") {
                    formMessage = "Formato email non valido.";
                } else if (data.error === "email already exists") {
                    formMessage = "Questa email è già registrata.";
                } else {
                    formMessage = `Errore: ${data.error}`;
                }
                return;
            }

            formMessage = "Utente creato con successo!";
            name = '';
            email = '';
            loadUsers();
        } catch (e) {
            formMessage = `Errore di connessione: ${e.message}`;
        }
    }

    $effect(() => {
        checkBackend();
        loadUsers();
    });
</script>

<h1>Benvenuto su DevTrack</h1>

<section class="status">
    <p>Stato del server: <strong>{status}</strong></p>
    {#if message}<p>Messaggio: {message}</p>{/if}
</section>

<hr />

<section>
    <h2>Crea un nuovo utente</h2>
    <form onsubmit={createUser}>
        <div>
            <label for="name">Nome:</label>
            <input type="text" id="name" bind:value={name} placeholder="Es. Mario Rossi" />
        </div>
        <div>
            <label for="email">Email:</label>
            <input type="email" id="email" bind:value={email} placeholder="mario.rossi@example.com" />
        </div>
        <button type="submit">Registra Utente</button>
    </form>
    
    {#if formMessage}
        <p class="notification">{formMessage}</p>
    {/if}
</section>

<hr />

<section>
    <h2>Lista Utenti registrati</h2>
    {#if users.length === 0}
        <p>Nessun utente trovato.</p>
    {:else}
        <ul>
            {#each users as u}
                <li><strong>{u.name}</strong> ({u.email})</li>
            {/each}
        </ul>
    {/if}
</section>

<style>
    section { margin: 1.5rem 0; }
    form div { margin-bottom: 0.5rem; }
    label { display: inline-block; width: 80px; }
    .notification { color: blue; font-weight: bold; }
</style>