<script lang="ts">
	interface Props {
		type?: string;
		id: string;
		label?: string;
		value: string;
		placeholder?: string;
		required?: boolean;
		disabled?: boolean;
	}

	let {
		type = 'text',
		id,
		label,
		value = $bindable(),
		placeholder = '',
		required = false,
		disabled = false
	}: Props = $props();
</script>

<div class="form-group">
	{#if label}
		<label for={id}>{label}</label>
	{/if}
	<input
		{type}
		{id}
		bind:value
		{placeholder}
		{required}
		{disabled}
	/>
</div>

<style>
	.form-group {
		display: flex;
		flex-direction: column;
		gap: 0.4rem;
		margin-bottom: 1.2rem;
		width: 100%;
	}
	label {
		font-size: 0.85rem;
		font-weight: 600;
		color: #475569;
	}
	input {
		padding: 0.6rem;
		border: 1px solid #cbd5e1;
		border-radius: 6px;
		font-size: 0.95rem;
		color: #1e293b;
		background: white;
		transition: border-color 0.2s, box-shadow 0.2s;
		width: 100%;
		box-sizing: border-box;
	}
	input:focus {
		border-color: #2563eb;
		outline: none;
		box-shadow: 0 0 0 2px rgba(37, 99, 235, 0.1);
	}
	input:disabled {
		background: #f8fafc;
		color: #94a3b8;
		cursor: not-allowed;
	}
</style>