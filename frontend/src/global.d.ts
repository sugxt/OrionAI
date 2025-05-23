// src/global.d.ts
import type * as runtime from "../wailsjs/runtime/runtime";

declare global {
  interface Window {
    runtime: typeof runtime;
  }
}
