import { writable } from "svelte/store";

export const clipboardHistory = writable<string[]>([]);
export const responseHistory = writable<string[]>([]);

window.runtime.EventsOn("clipboardHistoryUpdated", (data: string[]) => {
  clipboardHistory.set(data);
});

window.runtime.EventsOn("responseHistoryUpdated", (data: string[]) => {
  responseHistory.set(data);
});
