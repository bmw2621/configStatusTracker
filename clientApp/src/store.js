import { writable, readable } from "svelte/store";
import Socket from "./socket";

export const statuses = writable({});

export const socket = readable(new Socket("ws://localhost:3000/ws"));
