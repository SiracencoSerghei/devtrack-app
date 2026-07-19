<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		type?: 'success' | 'danger' | 'info';
		children: Snippet;
	}

	let { type = 'info', children }: Props = $props();
</script>

<div class="alert alert-{type}">
	{@render children()}
</div>

<style>
	.alert {
		padding: 0.8rem 1rem;
		border-radius: 6px;
		font-weight: 500;
		font-size: 0.875rem;
		border: 1px solid transparent;
		margin-bottom: 1rem;
	}
	.alert-info { background: #eff6ff; color: #1e40af; border-color: #bfdbfe; }
	.alert-danger { background: #fef2f2; color: #991b1b; border-color: #fca5a5; }
	.alert-success { background: #f0fdf4; color: #166534; border-color: #bbf7d0; }
</style>