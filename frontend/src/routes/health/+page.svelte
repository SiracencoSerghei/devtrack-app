<script>
    let status = $state('Loading...');
    let message = $state('');

    $effect(() => {
        async function fetchHealth() {
            try {
                const res = await fetch('http://localhost:8080/health');
                const json = await res.json();
                console.log('Risposta ricevuta:', json);

                status = json.status ?? 'OK';
                message = json.message ?? 'No message from the server';
                
            } catch (e) {
                status = 'offline';
                message = e.message;
            }
        }

        fetchHealth();
    });
</script>

<h1>Health</h1>

<p>Status: {status}</p>
<p>Message: {message}</p>
