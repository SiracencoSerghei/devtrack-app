<script>
    import { onMount } from 'svelte';
    let status = 'Checking...';

    onMount(async () => {
        try {
            // Nota: in SvelteKit lato client, la porta è quella del backend
            const res = await fetch('http://localhost:8080/health');
            const data = await res.json();
            status = data.status;
        } catch (e) {
            status = 'Server offline';
        }
    });
</script>

<h1>Home</h1>
<p>Stato del Backend Go: <strong>{status}</strong></p>