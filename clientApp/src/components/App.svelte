<script>
	import Editor from "./Editor.svelte"
	import StatusBars from "./StatusBars.svelte"
	import ConfigSelector from "./ConfigSelector.svelte";
	import GetConfig from "./GetConfig.svelte"

	let tabToShow = "create"

	const changeTab = (tab) => tabToShow = tab
	
</script>

<header>
	<h1>Status Tracker</h1>
	<h2>Service Configuration Manager</h2>
</header>

<main>
	<StatusBars />
	<div id="config-btns">
		<button class={`${tabToShow === "create" && "selectedTab"}`} on:click={() => changeTab("create")}>Create Config</button>
		<button class={`${tabToShow === "load" && "selectedTab"}`} on:click={() => changeTab("load")}>Load From System</button>
		<button class={`${tabToShow === "get" && "selectedTab"}`} on:click={() => changeTab("get")}>Get Saved Config</button>
	</div>
	{#if tabToShow === "create"}
		<Editor />
	{:else if tabToShow === "load"}
		<ConfigSelector />
	{:else if tabToShow === "get"}
		<GetConfig />
	{/if}
	
	
</main>

<style>
	header {
		width: 100%;
		background: orangered;
		color: white;
		text-shadow: 3px 3px 8px black;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 20px;
	}
	main {
		width: 100%;
		text-align: center;
		padding: 1em;
		max-width: 800px;
		margin: 0 auto;
	}

	#config-btns {
		margin: 1rem 0;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		width: 100%;
	}

	button {
		border: 1px solid lightslategray;
		background: white;
		color: black;
		border-radius: 3px;
		padding: 1rem;
		margin-right: 1rem;
	}

	.selectedTab {
		background: orangered;
		color: white
	}
</style>