<script lang="ts">
    import { AceEditor } from "svelte-ace";
	import "brace/mode/yaml";
	import "brace/theme/tomorrow";
    import {fade} from "svelte/transition"

	export let yamlText = ""

    const handleSave = async(): Promise<void> => {
        const resp = await fetch("./config", {method: "POST", body: yamlText})
        const respText = await resp.text()
        
        if(respText === "Success") {
            yamlText = ""
        }
    }
</script>

<div in:fade={{delay: 200, duration: 200}} out:fade={{duration: 200}}>
    <AceEditor
        on:paste={(obj) => yamlText = obj.detail}
        on:input={(obj) => yamlText = obj.detail}
        on:cut={(obj) => yamlText = obj.detail}
        width='100%'
        height='300px'
        lang="yaml"
        theme="tomorrow"
        options={{fontSize: "1rem"}}
        bind:value={yamlText} />
    
    <button on:click={handleSave}>Save</button>
</div>


<style>
    div {
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
    }

    button {
        border: 1px solid lightslategray;
        background: orangered;
        color: white;
        border-radius: 3px;
        padding: 1rem;
        margin-top: 1rem;
    }
</style>