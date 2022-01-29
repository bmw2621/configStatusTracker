<script>
    import StatusBar from "./StatusBar.svelte"
	import { statuses } from "../store"
</script>

<div style={`height: calc(${Object.keys($statuses).length} * 4.6rem); width: 100%;`}>
	{#each Object.keys($statuses) as status, i (i)}
		<StatusBar delayIn={i * 100} service={status} status={$statuses[status]} />
	{/each}
</div>