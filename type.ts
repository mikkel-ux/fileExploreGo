export type testTabType = {
  id: number;
  name: string;
  isActive: boolean;
};

export type TabType = {
  id: number;
  layout: "single" | "split";
  view: ViewType[];
  splitRatio: number;
  activeViewId: string;
};

export type ViewType = {
  id: string;
  title: string;
  image: string;
  history: string[];
  currentIndex: number;
};

export type FileDataType = {
  name: string;
  path: string;
  size: string;
  extension: string;
  created: string;
  modified: string;
  accessed: string;
  type: string;
  permissions: number;
  isHidden: boolean;
  isReadOnly: boolean;
  base64: string;
  firstFrame?: string;
};