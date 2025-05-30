import { writable } from "svelte/store";

export const clipboardHistory = writable<string[]>([]);

window.runtime.EventsOn("clipboardHistoryUpdated",(data: string[]) => {
    clipboardHistory.set(data)
})