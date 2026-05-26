<script>
    let status = $state('Checking...');
    let message = $state('');

    $effect(() => {
        async function checkBackend() {
            try {
                const res = await fetch('http://localhost:8080');
                if (!res.ok) throw new Error('Server error');

                const data = await res.json();

                status = data.status || 'OK';
                message = data.message || '';
            } catch (e) {
                status = 'Server offline';
                message = e.message;
            }
        }

        checkBackend();
    });
</script>

<h1>Home</h1>

<p>Response status: <strong>{status}</strong></p>

{#if message}
    <p>Message: {message}</p>
{/if}

