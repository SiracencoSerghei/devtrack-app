<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		type?: 'default' | 'success' | 'warning' | 'danger';
		children: Snippet;
	}

	let { type = 'default', children }: Props = $props();
</script>

<span class="badge badge-{type}">
	{@render children()}
</span>

<style>
	.badge {
		display: inline-block;
		padding: 0.15rem 0.5rem;
		font-size: 0.7rem;
		font-weight: 700;
		border-radius: 4px;
		text-transform: uppercase;
		letter-spacing: 0.025em;
	}
	.badge-default { background: #e2e8f0; color: #334155; }
	.badge-success { background: #d1fae5; color: #065f46; }
	.badge-warning { background: #fef3c7; color: #92400e; }
	.badge-danger { background: #fee2e2; color: #991b1b; }
</style>