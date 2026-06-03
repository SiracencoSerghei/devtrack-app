<script>
    import { i18n } from '$lib/i18n/i18n.svelte.js';
</script>

<h1>{i18n.t('about.title')}</h1>

<p>{i18n.t('about.description')}</p>

<p>{i18n.t('about.subtitle')}</p>

<style>
    p {
        line-height: 1.6;
        color: #333;
        font-family: sans-serif;
    }
</style>