<script lang="ts">
    import {fade} from "svelte/transition"
    let files:FileList;
    const getFileNames = ():string => Array.from(files).map(f => f.name).join(", ")
</script>

<div in:fade={{delay: 200, duration: 200}} out:fade={{duration: 200}}>
    <h2>Select Config</h2>
    {#if files?.length > 0}
    {#each files as file}
        <p>{file.name}</p>
    {/each}
    {/if}

    <div style="flex-direction:row; justify-content: center">
        <label for="file-upload" class="custom-upload">Browse</label>
        <input id="file-upload" accept=".yml,.yaml" type="file" multiple bind:files>
        <button 
            on:click={()=>{alert("Send files to API:\n" + getFileNames())}} 
            disabled='{!(files && files.length > 0)}'>Upload</button>
    </div>
</div>


<style>
    div {
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    input[type="file"]{
        display: none;
    }

    p {
        padding: 1rem 1.5rem;
        width: 100%;
        height: 3rem;
        border: 1px solid lightgray;
        margin: 0.2rem 0;
    }

    /* .custom-upload {
        border: 1px solid lightslategray;
        border-radius: 3px;
        display: inline-block;
        padding: 1rem;
        cursor: pointer;
    } */



    .custom-upload,
    button {
        border: 1px solid lightslategray;
        background: orangered;
        color: white;
        border-radius: 3px;
        padding: 1rem;
        margin-top: 1rem;
        margin-right: 1rem;
        width: 10rem;
        cursor: pointer;
    }

    button:disabled {
        background-color: lightcoral;
    }
</style>