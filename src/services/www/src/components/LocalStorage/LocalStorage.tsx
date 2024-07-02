import { useState } from "react";

export function useStoredItem(
  key: string
): [string | undefined, (v?: string) => void] {
  const [storedItem, setStoredItemState] = useState(
    localStorage.getItem(key) || undefined
  );

  function setStoredItem(value?: string) {
    if (value != null) {
      localStorage.setItem(key, value);
      setStoredItemState(value);
    } else {
      localStorage.removeItem(key);
      setStoredItemState(undefined);
    }
  }

  return [storedItem, setStoredItem];
}
