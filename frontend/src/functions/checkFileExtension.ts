import type { FileDataType } from "../../type";
export function isImage(fileData: FileDataType): boolean {
  const imageExtensions = ["jpg", "jpeg", "png", "gif", "bmp", "webp"];
  return imageExtensions.includes(fileData.extension.toLowerCase());
}
