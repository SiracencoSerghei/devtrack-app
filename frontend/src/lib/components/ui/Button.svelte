<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		type?: 'button' | 'submit' | 'reset';
		variant?: 'primary' | 'secondary' | 'danger' | 'light';
		disabled?: boolean;
		loading?: boolean;
		onclick?: () => void;
		children: Snippet;
	}

	let {
		type = 'button',
		variant = 'primary',
		disabled = false,
		loading = false,
		onclick,
		children
	}: Props = $props();
</script>

<button
	{type}
	class="btn btn-{variant}"
	disabled={disabled || loading}
	{onclick}
>
	{#if loading}
		<span class="btn-spinner"></span>
		<span>Loading...</span>
	{:else}
		{@render children()}
	{/if}
</button>

<style>
	.btn {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;
		padding: 0.6rem 1.2rem;
		border-radius: 6px;
		font-size: 0.9rem;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.2s;
		border: 1px solid transparent;
		box-sizing: border-box;
	}
	.btn:disabled {
		background: #cbd5e1 !important;
		color: #94a3b8 !important;
		border-color: transparent !important;
		cursor: not-allowed;
	}
	.btn-primary { background: #2563eb; color: white; }
	.btn-primary:hover:not(:disabled) { background: #1d4ed8; }

	.btn-secondary { background: #64748b; color: white; }
	.btn-secondary:hover:not(:disabled) { background: #475569; }

	.btn-danger { background: #ef4444; color: white; }
	.btn-danger:hover:not(:disabled) { background: #dc2626; }

	.btn-light { background: #f1f5f9; color: #334155; border: 1px solid #cbd5e1; }
	.btn-light:hover:not(:disabled) { background: #e2e8f0; }

	.btn-spinner {
		width: 1rem;
		height: 1rem;
		border: 2px solid currentColor;
		border-bottom-color: transparent;
		border-radius: 50%;
		display: inline-block;
		animation: rotation 1s linear infinite;
	}
	@keyframes rotation {
		0% { transform: rotate(0deg); }
		100% { transform: rotate(360deg); }
	}
</style>