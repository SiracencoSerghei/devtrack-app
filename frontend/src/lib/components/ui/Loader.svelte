<script lang="ts">
	interface Props {
		message?: string;
	}

	let { message = 'Loading...' }: Props = $props();
</script>

<div class="loader-container">
	<span class="spinner"></span>
	{#if message}
		<p>{message}</p>
	{/if}
</div>

<style>
	.loader-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 3rem;
		gap: 1rem;
	}
	.spinner {
		width: 2.5rem;
		height: 2.5rem;
		border: 4px solid #e2e8f0;
		border-bottom-color: #2563eb;
		border-radius: 50%;
		display: inline-block;
		animation: rotation 1s linear infinite;
	}
	p {
		color: #64748b;
		font-size: 0.9rem;
		font-weight: 500;
		margin: 0;
	}
	@keyframes rotation {
		0% { transform: rotate(0deg); }
		100% { transform: rotate(360deg); }
	}
</style>