<script>
    let status = $state('Caricamento...');
    let message = $state('');

    $effect(() => {
        async function fetchHealth() {
            try {
                const res = await fetch('http://localhost:8080/health');
                const json = await res.json();
                status = json.status ?? 'OK';
            } catch (e) {
                status = 'Non in linea (offline)';
                message = e.message;
            }
        }
        fetchHealth();
    });
</script>

<h1>Stato del Sistema</h1>
<p>Stato: <strong>{status}</strong></p>
{#if message}<p>Dettaglio errore: {message}</p>{/if}