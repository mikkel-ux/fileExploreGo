import { writable } from "svelte/store";
import type { testTabType } from "../../type";

export const nextTabId = writable(2);
export const activeTabId = writable<number>(1);
export const tabs = writable<testTabType[]>([
  {
    id: 1,
    name: "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
    isActive: true,
  },
]);
export const isDragging = writable(false);
