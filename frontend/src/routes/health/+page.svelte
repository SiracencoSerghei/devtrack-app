<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';

    let statusKey = $state('loading');
    let message = $state('');

    $effect(() => {
        async function fetchHealth() {
            try {
                const res = await fetch('http://localhost:8080/health');
                const json = await res.json();
                
                if (json.status === 'OK') {
                    statusKey = 'ok';
                } else {
                    statusKey = 'offline';
                }
            } catch (e) {
                statusKey = 'offline';
                message = e.message;
            }
        }
        fetchHealth();
    });
</script>

<h1>{i18n.t('health.title')}</h1>

<p>{i18n.t('health.status')} <strong>{i18n.t(`health.${statusKey}`)}</strong></p>

{#if message}
    <p>{i18n.t('health.error_detail')} {message}</p>
{/if}