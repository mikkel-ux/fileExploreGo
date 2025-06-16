import { writable, derived } from "svelte/store";

export const historyStore = writable({
  history: ["/"],
  currentIndex: 0,
});

export const addToHistory = (path: string) => {
  historyStore.update((state) => {
    const newHistory = [
      ...state.history.slice(0, state.currentIndex + 1),
      path,
    ];
    return {
      history: newHistory,
      currentIndex: newHistory.length - 1,
    };
  });
};

export const goBack = () => {
  historyStore.update((state) => {
    const newIndex = Math.max(state.currentIndex - 1, 0);
    return {
      ...state,
      currentIndex: newIndex,
    };
  });
  let currentPath: string = "";
  historyStore.subscribe((state) => {
    currentPath = state.history[state.currentIndex];
  });
  return currentPath;
};

export const goForward = () => {
  historyStore.update((state) => {
    const newIndex = Math.min(state.currentIndex + 1, state.history.length - 1);
    return {
      ...state,
      currentIndex: newIndex,
    };
  });
  let currentPath: string = "";
  historyStore.subscribe((state) => {
    currentPath = state.history[state.currentIndex];
  });
  return currentPath;
};

export const getCurrentPath = () => {
  let currentPath: string = "";
  historyStore.subscribe((state) => {
    currentPath = state.history[state.currentIndex];
  });
  return currentPath;
};

export const getHistory = () => {
  let history: string[] = [];
  historyStore.subscribe((state) => {
    history = state.history;
  });
  return history;
};

export const canGoBack = derived(historyStore, ($historyStore) => {
  return $historyStore.currentIndex > 0;
});

export const canGoForward = derived(historyStore, ($historyStore) => {
  return $historyStore.currentIndex < $historyStore.history.length - 1;
});
