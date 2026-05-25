<script>
    import { onMount } from 'svelte';
    let data = "In attesa...";

    onMount(async () => {
        try {
            const response = await fetch('http://localhost:8080/health');
            const result = await response.json();
            data = JSON.stringify(result); // Dovrebbe stampare {"status":"ok"}
        } catch (err) {
            data = "Errore di connessione: " + err.message;
        }
    });
</script>

<h1>DevTrack</h1>
<p>Risposta dal backend: <strong>{data}</strong></p>