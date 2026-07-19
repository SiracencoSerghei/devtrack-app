<script lang="ts">
	import { i18n } from '$lib/i18n/i18n.svelte.js';

	let isOpen = $state(false);

	function setLang(lang: 'it' | 'uk' | 'en') {
		i18n.setLang(lang);
		isOpen = false;
	}
</script>

<div class="lang-selector">
	<button class="lang-current" onclick={() => (isOpen = !isOpen)} type="button">
		{i18n.lang.toUpperCase()}
	</button>

	{#if isOpen}
		<div class="lang-dropdown">
			<button class={i18n.lang === 'it' ? 'active' : ''} onclick={() => setLang('it')} type="button">IT</button>
			<button class={i18n.lang === 'uk' ? 'active' : ''} onclick={() => setLang('uk')} type="button">UA</button>
			<button class={i18n.lang === 'en' ? 'active' : ''} onclick={() => setLang('en')} type="button">EN</button>
		</div>
	{/if}
</div>

<style>
	.lang-selector {
		position: relative;
		display: inline-block;
	}
	.lang-current {
		background: transparent;
		border: 1px solid #334155;
		padding: 0.4rem 0.75rem;
		color: #94a3b8;
		cursor: pointer;
		border-radius: 6px;
		font-size: 0.75rem;
		font-weight: 700;
		transition: all 0.2s ease;
	}
	.lang-current:hover {
		color: white;
		border-color: #475569;
	}
	.lang-dropdown {
		position: absolute;
		top: calc(100% + 6px);
		right: 0;
		background: #0f172a;
		border: 1px solid #334155;
		border-radius: 6px;
		min-width: 80px;
		z-index: 100;
		display: flex;
		flex-direction: column;
		box-shadow: 0 10px 25px rgba(0, 0, 0, 0.3);
		overflow: hidden;
	}
	.lang-dropdown button {
		width: 100%;
		padding: 0.5rem 0.75rem;
		background: transparent;
		border: none;
		color: #94a3b8;
		cursor: pointer;
		font-size: 0.75rem;
		text-align: left;
		transition: background 0.15s, color 0.15s;
	}
	.lang-dropdown button:hover {
		background: #1e293b;
		color: white;
	}
	.lang-dropdown button.active {
		background: var(--accent, #2563eb);
		color: white;
	}
</style>