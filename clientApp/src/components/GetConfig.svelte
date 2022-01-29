<script lang="ts">
    let configIds : Promise<string[]> = fetch("./config").then(data => data.text()).then(data => JSON.parse(data))

    const convertIdToDate = (id : string): string => {
        return new Date(parseInt(id) * 1000).toLocaleDateString("en-US" ,{month: "short", day: "numeric", year: "numeric", hour: "numeric", minute: "numeric"})
    }

    const getConfig = async(id : string) : Promise<void> => {
        const data = await fetch(`./config/${id}`)
        const yamlFile = await data.text()
        alert(yamlFile)
    }
</script>

<div>
    {#await configIds then ids}
        {#each ids as id}
            <p>{`${id}.yaml`}</p>
            <button on:click={() => getConfig(id)}>{convertIdToDate(id)}</button>
        {/each}
    {/await}
</div>

<style>
    button {
		border: 1px solid lightslategray;
		background: white;
		color: black;
		border-radius: 3px;
		padding: 1rem;
		margin-right: 1rem;
	}
</style>