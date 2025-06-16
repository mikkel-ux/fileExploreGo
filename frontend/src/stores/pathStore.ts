import { writable, derived } from "svelte/store";

export const pathStore = writable("/");

export const setPath = (path: string) => {
  pathStore.set(path);
};

export const getPath = () => {
  let currentPath: string = "";
  pathStore.subscribe((path) => {
    currentPath = path;
  })();
  return currentPath;
};