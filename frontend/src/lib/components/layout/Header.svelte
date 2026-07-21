<script lang="ts">
	import { i18n } from '$lib/i18n/i18n.svelte.js';
	import LanguageSelector from './LanguageSelector.svelte';
	import UserMenu from './UserMenu.svelte';
</script>

<header class="main-header">
	<div class="header-container">
		<a href="/" class="logo">DevTrack <span>ERP</span></a>

		<nav class="nav-links">
			<a href="/">{i18n.t('nav.home')}</a>
			<a href="/orders">{i18n.t('nav.orders')}</a>
			<a href="/about">{i18n.t('nav.about')}</a>
			<a href="/health">{i18n.t('nav.health')}</a>
		</nav>

		<div class="actions-wrapper">
			<LanguageSelector />
			<UserMenu />
		</div>
	</div>
</header>

<style>
	.main-header {
		background: #0f172a;
		color: white;
		border-bottom: 1px solid #1e293b;
	}
	.header-container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 1rem;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.logo {
		font-size: 1.25rem;
		font-weight: 800;
		color: white;
		text-decoration: none;
		letter-spacing: -0.025em;
	}
	.logo span {
		color: var(--accent, #2563eb);
	}
	.nav-links a {
		color: #94a3b8;
		margin-left: 1.5rem;
		text-decoration: none;
		font-weight: 500;
		font-size: 0.875rem;
		transition: color 0.2s;
	}
	.nav-links a:hover {
		color: white;
	}
	.actions-wrapper {
		display: flex;
		align-items: center;
		gap: 1.5rem;
	}
</style>