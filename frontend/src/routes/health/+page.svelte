<script lang="ts">
	import { systemAPI } from '$lib/api/system';
	import { i18n } from '$lib/i18n/i18n.svelte.js';
	import Card from '$lib/components/ui/Card.svelte';
	import Badge from '$lib/components/ui/Badge.svelte';

	let statusKey = $state<'loading' | 'ok' | 'offline'>('loading');

	async function fetchHealth() {
		try {
			const res = await systemAPI.checkHealth();
			if (res && res.status === 'ok') {
				statusKey = 'ok';
			} else {
				statusKey = 'offline';
			}
		} catch (e) {
			statusKey = 'offline';
		}
	}

	$effect(() => {
		fetchHealth();
	});
</script>

<Card>
	<h1>{i18n.t('health.title') || 'System Status'}</h1>
	<p>
		{i18n.t('health.status') || 'Status:'} 
		<Badge type={statusKey === 'ok' ? 'success' : statusKey === 'loading' ? 'default' : 'danger'}>
			{statusKey.toUpperCase()}
		</Badge>
	</p>
</Card>