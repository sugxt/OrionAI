import { LoadConfig, SaveConfig } from "../../wailsjs/go/api/App.js";

export async function loadConfig() {
  const load = await LoadConfig();
  if (load) {
    return JSON.parse(load);
  }
  return {};
}

export async function saveConfig(config: { theme: string; username: string }) {
  const json = JSON.stringify(config);
  const save = SaveConfig(json);
  if (save) {
    return save;
  }
}
