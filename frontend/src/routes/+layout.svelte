<script>
    // 1. Accept the 'children' prop using the Svelte 5 $props rune
    let { children } = $props();
</script>

<nav>
    <a href="/">Home</a>
    <a href="/about">About</a>
</nav>

<main>
    {@render children()}
</main>

<style>
    nav { border-bottom: 1px solid #ccc; padding: 1rem; }
    nav a { margin-right: 1rem; text-decoration: none; color: blue; }
</style>
