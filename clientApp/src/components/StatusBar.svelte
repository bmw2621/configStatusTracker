<script>
    import { slide } from "svelte/transition"
    export let service
    export let status
    export let delayIn

    function fadeSlide(node, options) {
		const slideTrans = slide(node, options)
		return {
            ...options,
			css: t => `
				${slideTrans.css(t)}
				opacity: ${t};
			`
		};
	}
</script>

<div transition:fadeSlide="{{duration: 500, delay: delayIn}}" class="statusContainer">
    <span style={`${status.progress > 60 ? "color: white" : ""}`}>{`${service.toUpperCase()} : ${status.status}`}</span>
    <div class="statusBar" style={`width: ${status.progress}%`}></div>

</div>

<style>
    .statusContainer{
        width: 100%;
        position: relative;
        padding: 0.2rem;
        border: 1px solid lightslategray;
        margin-bottom:0.2rem;
        border-radius: 3px;
    }

    .statusBar {
        position: relative;
        height: 4rem;
        top: 0;
        left: 0;
        background: orangered;
        transition: width 500ms linear;
    }
    span {
        position: absolute;
        height: 100%;
        top: 1.8rem;
        z-index: 10;
        left: 50%;
        transform: translateX(-50%);
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
    }
</style>

