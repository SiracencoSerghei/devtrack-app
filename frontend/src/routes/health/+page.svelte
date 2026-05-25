<script>
    import { onMount } from 'svelte';
    
    // Runa per la reattività
    let data = $state("In attesa...");

    onMount(async () => {
        try {
            console.log("Tentativo fetch...");
            const response = await fetch('http://localhost:8080/health');
            const result = await response.json();
            console.log("Risposta ricevuta:", result);
            data = JSON.stringify(result);
        } catch (err) {
            console.error("Errore fetch:", err);
            data = "Errore: " + err.message;
        }
    });
</script>

<h1>DevTrack</h1>
<p>Risposta dal backend: <strong>{data}</strong></p>