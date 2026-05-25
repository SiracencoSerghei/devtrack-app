<script>
    import { onMount } from 'svelte';
    
    let status = $state('Checking...');

    onMount(async () => {
        try {
            const res = await fetch('http://localhost:8080/health');
            if (!res.ok) throw new Error('Backend error');
            
            const data = await res.json();
            // This mutation will now trigger a UI update correctly!
            status = data.status || 'OK'; 
        } catch (e) {
            status = 'Server offline';
        }
    });
</script>

<h1>Home</h1>
<p>Stato del Backend Go: <strong>{status}</strong></p>