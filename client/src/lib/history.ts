export interface HistoryItem {
  id: string;
  longUrl: string;
  shortUrl: string;
  fullShortUrl: string;
  createdAt: number;
}

const STORAGE_KEY = "titly_history_v1";

export const getHistory = (): HistoryItem[] => {
  if (typeof window === "undefined") return [];
  try {
    const data = localStorage.getItem(STORAGE_KEY);
    return data ? JSON.parse(data) : [];
  } catch {
    return [];
  }
};

export const addHistoryItem = (item: Omit<HistoryItem, "id" | "createdAt">): HistoryItem[] => {
  if (typeof window === "undefined") return [];
  const current = getHistory();
  
  // Deduplicate if exact longUrl & shortUrl already exists
  const filtered = current.filter(h => h.shortUrl !== item.shortUrl);
  
  const newItem: HistoryItem = {
    ...item,
    id: Math.random().toString(36).substring(2, 9),
    createdAt: Date.now()
  };

  const updated = [newItem, ...filtered].slice(0, 25); // store up to 25 items
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(updated));
  } catch {
    // ignore quota errors
  }
  return updated;
};

export const removeHistoryItem = (id: string): HistoryItem[] => {
  if (typeof window === "undefined") return [];
  const current = getHistory();
  const updated = current.filter(item => item.id !== id);
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(updated));
  } catch {
    // ignore
  }
  return updated;
};

export const clearHistory = (): HistoryItem[] => {
  if (typeof window === "undefined") return [];
  try {
    localStorage.removeItem(STORAGE_KEY);
  } catch {
    // ignore
  }
  return [];
};
