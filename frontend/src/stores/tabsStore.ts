//# allFunctionsCalledOnLoad

import { writable, get, derived } from "svelte/store";
import { getCurrentWindow } from "@tauri-apps/api/window";
import type { TabType, ViewType, FileDataType } from "../../type";
import { v4 as uuidv4 } from "uuid";

const appWindow = getCurrentWindow();
const firstTabUuid = uuidv4();

//=========================
// writable stores that are not exported
//=========================

const nextTabId = writable(2);

//=========================
// Window Management
//=========================

function close() {
  appWindow.close();
}

//=========================
// helper functions
//=========================

function createTab() {
  const id = get(nextTabId);
  const viewId = uuidv4();
  /* TODO change image to the right image path */
  /* TODO change the first history path to the right one for "Home" */
  const newView: ViewType = {
    id: viewId,
    title: `Tab ${id}`,
    image: "https://",
    history: ["/"],
    currentIndex: 0,
  };

  const newTab: TabType = {
    id,
    layout: "single",
    view: [structuredClone(newView)],
    splitRatio: 1,
    activeViewId: viewId,
  };
  nextTabId.update((n) => n + 1);
  activeTabId.set(id);
  return newTab;
}

function findIndex(tabs: TabType[], id: number): number {
  return tabs.findIndex((tab) => tab.id === id);
}

function findActiveView(
  tabs: TabType[],
  activeTabIdVal: number
): ViewType | undefined {
  const tab = tabs.find((tab) => tab.id === activeTabIdVal);
  return tab?.view.find((v) => v.id === tab.activeViewId);
}

function updateActiveView(
  tabs: TabType[],
  activeTabIdVal: number,
  updateFn: (view: ViewType) => ViewType
): TabType[] {
  const tabIndex = findIndex(tabs, activeTabIdVal);
  if (tabIndex === -1) return tabs;

  const tab = tabs[tabIndex];
  const viewIndex = tab.view.findIndex((v) => v.id === tab.activeViewId);
  if (viewIndex === -1) return tabs;

  const activeView = tab.view[viewIndex];
  const updatedView = updateFn(activeView);

  const updatedViews = [...tab.view];
  updatedViews[viewIndex] = updatedView;

  const updateTab: TabType = {
    ...tab,
    view: updatedViews,
  };

  const updatedTabs = [...tabs];
  updatedTabs[tabIndex] = updateTab;
  return updatedTabs;
}

//=========================
// exported stores
//=========================

export const selectedFile = writable<FileDataType | null>(null);
export const isDragging = writable(false);
export const activeTabId = writable<number>(1);
export const tabsStore = writable<TabType[]>([
  {
    id: 1,
    layout: "single",
    view: [
      {
        id: firstTabUuid,
        title: "view 1",
        image: "image1.png",
        history: ["/"],
        currentIndex: 0,
      },
    ],
    splitRatio: 1,
    activeViewId: firstTabUuid,
  },
]);

//=========================
// File Management
//==========================
export const secsectFile = (file: FileDataType) => {
  selectedFile.set(file);
};

export const removeSelectedFile = () => {
  selectedFile.set(null);
};

//=========================
// Tab Management
//=========================

export const addTab = () => {
  tabsStore.update((tabs) => {
    const newTab = createTab();
    return [...tabs, newTab];
  });
};

export const removeTab = (id: number) => {
  if (get(tabsStore).length === 1) {
    close();
  }

  const currentTabs = get(tabsStore);
  const tabIndex = currentTabs.findIndex((tab) => tab.id === id);
  if (tabIndex === -1) return;

  const newTabs = currentTabs.filter((tab) => tab.id !== id);
  tabsStore.set(newTabs);

  if (newTabs.length === 0) return;
  const nextActive = newTabs[Math.max(0, tabIndex - 1)] ?? newTabs[0];
  if (nextActive) {
    setActiveTab(nextActive.id);
  }
};

export const setActiveTab = (id: number) => {
  activeTabId.set(id);
};

//=========================
// View Management
//=========================

export const addView = (view: ViewType) => {
  const activeTabIdVal = get(activeTabId);
  tabsStore.update((tabs) => {
    const tabIndex = findIndex(tabs, activeTabIdVal);
    if (tabIndex === -1) return tabs;

    const tab = tabs[tabIndex];
    const updatedTab: TabType = {
      ...tab,
      view: [...tab.view, view],
      layout: "split",
      splitRatio: 0.5,
    };

    const updatedTabs = [...tabs];
    updatedTabs[tabIndex] = updatedTab;
    return updatedTabs;
  });
};

export const removeView = (tabId: number, viewId: string) => {
  tabsStore.update((tabs) => {
    const tabIndex = findIndex(tabs, tabId);
    const tab = tabs[tabIndex];
    const updatedView = tab.view.filter((view) => view.id !== viewId);
    const updatedTab: TabType = {
      ...tab,
      view: updatedView,
      layout: "single",
      splitRatio: 1,
      activeViewId: updatedView[0]?.id ?? "",
    };

    const updatedTabs = [...tabs];
    updatedTabs[tabIndex] = updatedTab;
    return updatedTabs;
  });
};

export const setActiveView = (viewId: string) => {
  tabsStore.update((tabs) => {
    const activeTabIdVal = get(activeTabId);
    const tabIndex = findIndex(tabs, activeTabIdVal);
    if (tabIndex === -1) return tabs;

    const updatedTab: TabType = {
      ...tabs[tabIndex],
      activeViewId: viewId,
    };

    const updatedTabs = [...tabs];
    updatedTabs[tabIndex] = updatedTab;
    return updatedTabs;
  });
};

/* TODO figer out how we will update the title of the view */
/* export const setViewTitle = (tabId: number, viewId: number, title: string) => {
  tabsStore.update((tabs) => {
    const tabIndex = findIndex(tabs, tabId);
    const viewIndex = tabs[tabIndex].view.findIndex(
      (view) => view.id === viewId
    );
    tabs[tabIndex].view[viewIndex].title = title;
    return tabs;
  });
}; */

//=========================
// History Management
//=========================

export const updateHistory = (path: string) => {
  tabsStore.update((tabs) => {
    const activeTabIdVal = get(activeTabId);
    return updateActiveView(tabs, activeTabIdVal, (activeView) => {
      const newHistory = [
        ...activeView.history.slice(0, activeView.currentIndex + 1),
        path,
      ];
      return {
        ...activeView,
        history: newHistory,
        currentIndex: newHistory.length - 1,
      };
    });
  });
};

export const goBack = () => {
  tabsStore.update((tabs) => {
    const activeTabIdVal = get(activeTabId);
    return updateActiveView(tabs, activeTabIdVal, (activeView) => {
      const newIndex = Math.max(activeView.currentIndex - 1, 0);
      return {
        ...activeView,
        currentIndex: newIndex,
      };
    });
  });
  return getCurrentPath();
};

export const goForward = () => {
  tabsStore.update((tabs) => {
    const activeTabIdVal = get(activeTabId);
    return updateActiveView(tabs, activeTabIdVal, (activeView) => ({
      ...activeView,
      currentIndex: Math.min(
        activeView.currentIndex + 1,
        activeView.history.length - 1
      ),
    }));
  });
  return getCurrentPath();
};

export const getCurrentPath = () => {
  let currentPath: string = "";
  const tabs = get(tabsStore);
  const activeTabIdVal = get(activeTabId);
  const activeView = findActiveView(tabs, activeTabIdVal);
  if (activeView) {
    currentPath = activeView.history[activeView.currentIndex];
  }
  return currentPath;
};

export const getHistory = () => {
  let history: string[] = [];
  const tabs = get(tabsStore);
  const activeTabIdVal = get(activeTabId);
  const activeView = findActiveView(tabs, activeTabIdVal);
  if (activeView) {
    history = activeView.history;
  }
  return history;
};

export const canGoBack = derived(
  [tabsStore, activeTabId],
  ([$tabStore, $activeTabId]) => {
    const tab = $tabStore.find((tab) => tab.id === $activeTabId);
    const activeView = tab?.view.find((v) => v.id === tab?.activeViewId);
    return activeView ? activeView.currentIndex > 0 : false;
  }
);

export const canGoForward = derived(
  [tabsStore, activeTabId],
  ([$tabStore, $activeTabId]) => {
    const tab = $tabStore.find((tab) => tab.id === $activeTabId);
    const activeView = tab?.view.find((v) => v.id === tab?.activeViewId);
    return activeView
      ? activeView.currentIndex < activeView.history.length - 1
      : false;
  }
);

export const activeTabIndex = derived(tabsStore, ($tabsStore) => {
  /* const activeTab = $tabsStore.find((tab) => tab.isActive); */
  const activeTabIdVal = get(activeTabId);
  const activeTab = $tabsStore.find((tab) => tab.id === activeTabIdVal);
  return activeTab ? $tabsStore.indexOf(activeTab) : -1;
});

export const activeViews = derived(
  [tabsStore, activeTabIndex],
  ([$tabsStore, $activeTabIndex]) => {
    if ($activeTabIndex === -1) return [];
    return $tabsStore[$activeTabIndex].view;
  }
);
